package integration_test

import (
	"reflect"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestSectionAndIdentifierRules(t *testing.T) {
	tests := []struct {
		name string
		text string
		code string
	}{
		{"skipped level", "^ Ledger\n^^^ Audit\nentry = 1", "section-depth"},
		{"too many markers", "^^^^^^^^^^ Archive\n", "section-header"},
		{"dotted key", "^ Ledger\nentry.part = 1\n", "identifier"},
		{"dotted section", "^ Ledger.Part\n", "section-header"},
		{"root collision", "Ledger = 1\n^ Ledger\n", "name-collision"},
		{"inline semicolon", "^ Ledger\nentry = 1 ; note\n", "syntax"},
		{"reserved directive", "@include 'other.yini'\n^ Ledger\n", "reserved-directive"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := yini.ParseString(test.text, yini.Options{})
			if err == nil || !hasDiagnostic(result, test.code) {
				t.Fatalf("want %s error, got %v and %#v", test.code, err, result.Diagnostics)
			}
		})
	}
	result, err := yini.ParseString("^`Road Map`// comment\n^^Milestones\nstep = 1", yini.Options{})
	if err != nil {
		t.Fatalf("valid backticked header: %v (%#v)", err, result.Diagnostics)
	}
	if got := result.Document.Values(); !reflect.DeepEqual(got, map[string]any{"Road Map": map[string]any{"Milestones": map[string]any{"step": int64(1)}}}) {
		t.Fatalf("unexpected values: %#v", got)
	}
}

func TestInlineObjectDuplicateAndStrictConcatenation(t *testing.T) {
	content := "^ Specimen\nitem = { id: 4, id: 9 }\n"
	lenient, err := yini.ParseString(content, yini.Options{})
	if err != nil || !hasDiagnostic(lenient, "duplicate-object-key") {
		t.Fatalf("lenient duplicate: %v, %#v", err, lenient.Diagnostics)
	}
	item := lenient.Document.Values()["Specimen"].(map[string]any)["item"]
	if item.(map[string]any)["id"] != int64(4) {
		t.Fatalf("first value was not kept: %#v", item)
	}
	strict, err := yini.ParseString(content+"/END", yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "duplicate-object-key") {
		t.Fatalf("strict duplicate: %v, %#v", err, strict.Diagnostics)
	}
	strict, err = yini.ParseString("^ Specimen\nlabel = 'v' + 4\n/END", yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "concat-strict") {
		t.Fatalf("strict concatenation: %v, %#v", err, strict.Diagnostics)
	}
}

func TestNumericFormsAndEscapes(t *testing.T) {
	content := strings.Join([]string{
		"^ Materials",
		"binary = 0b_1010_0101",
		"octal = 0o_17",
		"dozenal = 0z_E",
		"hexadecimal = 0x_FF",
		"explicitHex = HEX:10",
		"fraction = .5",
		"tiny = 5e-324",
		"classic = C'one\\n\\u263A\\o101'",
		"raw = 'C:\\Raw\\Path'",
		"",
	}, "\n")
	result, err := yini.ParseString(content, yini.Options{})
	if err != nil {
		t.Fatalf("ParseString: %v (%#v)", err, result.Diagnostics)
	}
	values := result.Document.Values()["Materials"].(map[string]any)
	for key, want := range map[string]any{
		"binary": int64(165), "octal": int64(15), "dozenal": int64(11),
		"hexadecimal": int64(255), "explicitHex": int64(16),
		"fraction": 0.5, "classic": "one\n☺A", "raw": `C:\Raw\Path`,
	} {
		if values[key] != want {
			t.Errorf("%s = %#v, want %#v", key, values[key], want)
		}
	}
	if values["tiny"].(float64) == 0 {
		t.Error("nonzero subnormal was lost")
	}
	for _, bad := range []string{"C'bad\\z'", "C'bad\\uD800'", "C'bad\\o400'"} {
		result, err := yini.ParseString("^ Materials\nvalue = "+bad, yini.Options{})
		if err == nil || !hasDiagnostic(result, "string") {
			t.Errorf("bad escape %s: %v, %#v", bad, err, result.Diagnostics)
		}
	}
}

func TestEmptyDocumentsAndLineEndings(t *testing.T) {
	lenient, err := yini.ParseString("; note only", yini.Options{})
	if err != nil || !hasDiagnostic(lenient, "empty-document") {
		t.Fatalf("lenient empty: %v, %#v", err, lenient.Diagnostics)
	}
	strict, err := yini.ParseString("; note only\r/END", yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "empty-document") {
		t.Fatalf("strict empty: %v, %#v", err, strict.Diagnostics)
	}
	result, err := yini.ParseString("\uFEFF^ Transit\r\nline = 1\rline = 2\r", yini.Options{})
	if err != nil {
		t.Fatalf("mixed line endings: %v (%#v)", err, result.Diagnostics)
	}
	for _, diagnostic := range result.Diagnostics {
		if diagnostic.Code == "duplicate-key" {
			if diagnostic.Position.Line != 3 || diagnostic.Position.Column != 1 {
				t.Fatalf("duplicate location = %#v", diagnostic.Position)
			}
			return
		}
	}
	t.Fatal("missing duplicate-key diagnostic")
}

func TestMisplacedShebang(t *testing.T) {
	content := "^ Journal\ntext = '#! stays text'\n#! later comment\nentry = 1\n"
	lenient, err := yini.ParseString(content, yini.Options{})
	if err != nil {
		t.Fatalf("lenient misplaced shebang: %v (%#v)", err, lenient.Diagnostics)
	}
	if got := lenient.Document.Values()["Journal"].(map[string]any)["entry"]; got != int64(1) {
		t.Fatalf("content after comment = %#v", got)
	}
	if len(lenient.Diagnostics) != 1 || lenient.Diagnostics[0].Code != "misplaced-shebang" || lenient.Diagnostics[0].Position.Line != 3 || lenient.Diagnostics[0].Severity != yini.SeverityWarning {
		t.Fatalf("lenient diagnostics = %#v", lenient.Diagnostics)
	}
	strict, err := yini.ParseString(content+"/END", yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "misplaced-shebang") {
		t.Fatalf("strict misplaced shebang: %v (%#v)", err, strict.Diagnostics)
	}
	leading, err := yini.ParseString(" #! leading space\n^ Journal\n", yini.Options{})
	if err != nil || !hasDiagnostic(leading, "misplaced-shebang") {
		t.Fatalf("leading-space shebang: %v (%#v)", err, leading.Diagnostics)
	}
}

func hasDiagnostic(result yini.Result, code string) bool {
	for _, diagnostic := range result.Diagnostics {
		if diagnostic.Code == code {
			return true
		}
	}
	return false
}
