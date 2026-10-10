package integration_test

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestParseFilesAndValues(t *testing.T) {
	tests := []struct {
		name string
		file string
		mode yini.Mode
		want map[string]any
	}{
		{
			name: "lenient",
			file: "workshop.yini",
			want: map[string]any{
				"label":      "north wing",
				"retryCount": int64(3),
				"Workshop": map[string]any{
					"active":     true,
					"tags":       []any{"wood", "steel"},
					"details":    map[string]any{"lead": "Mira", "floor": int64(2)},
					"Calibrator": map[string]any{"margin": 0.125},
				},
			},
		},
		{
			name: "strict",
			file: "observatory.strict.yini",
			mode: yini.ModeStrict,
			want: map[string]any{
				"Observatory": map[string]any{
					"title":       "Night survey",
					"count":       int64(17),
					"ready":       true,
					"coordinates": []any{12.5, -4.25},
					"Instrument":  map[string]any{"model": "Orion"},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := yini.ParseFile(filepath.Join("testdata", test.file), yini.Options{Mode: test.mode})
			if err != nil {
				t.Fatalf("ParseFile: %v (diagnostics: %#v)", err, result.Diagnostics)
			}
			if got := result.Document.Values(); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Values() = %#v, want %#v", got, test.want)
			}
			values := result.Document.Values()
			values["changed"] = true
			if _, exists := result.Document.Values()["changed"]; exists {
				t.Fatal("Values() returned document-owned map")
			}
			for _, section := range values {
				if nested, ok := section.(map[string]any); ok {
					nested["changed"] = true
				}
			}
			if !reflect.DeepEqual(result.Document.Values(), test.want) {
				t.Fatal("Values() returned document-owned nested map")
			}
		})
	}
}

func TestModesAndDiagnostics(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		mode     yini.Mode
		code     string
		severity yini.Severity
	}{
		{"strict terminator", "^ Ship\nname = 'Dawn'", yini.ModeStrict, "missing-end", yini.SeverityError},
		{"strict orphan", "name = 'Dawn'\n^ Ship\n/END", yini.ModeStrict, "orphan-member", yini.SeverityError},
		{"strict extra root", "^ Ship\n^ Dock\n/END", yini.ModeStrict, "top-level-section", yini.SeverityError},
		{"declared strict", "@yini strict\n^ Ship\n", yini.ModeLenient, "mode-mismatch", yini.SeverityError},
		{"declared lenient", "@yini lenient\n^ Ship\n/END", yini.ModeStrict, "mode-mismatch", yini.SeverityWarning},
		{"strict trailing comma", "^ Ship\nports = [1, 2,]\n/END", yini.ModeStrict, "trailing-comma", yini.SeverityError},
		{"strict object separator", "^ Ship\nport = { number = 7 }\n/END", yini.ModeStrict, "object-separator", yini.SeverityError},
		{"strict empty value", "^ Ship\nname =\n/END", yini.ModeStrict, "empty-value", yini.SeverityError},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := yini.ParseString(test.text, yini.Options{Mode: test.mode})
			if test.severity == yini.SeverityError && err == nil {
				t.Fatalf("expected parse error, got diagnostics %#v", result.Diagnostics)
			}
			if test.severity == yini.SeverityWarning && err != nil {
				t.Fatalf("unexpected parse error: %v", err)
			}
			for _, diagnostic := range result.Diagnostics {
				if diagnostic.Code == test.code && diagnostic.Severity == test.severity {
					return
				}
			}
			t.Fatalf("missing %s diagnostic %q: %#v", test.severity, test.code, result.Diagnostics)
		})
	}
}

func TestDuplicateRecoveryAndLocations(t *testing.T) {
	text := "\uFEFF#!/usr/bin/env yini\n^ Fleet\nname = 'First'\nname = 'Second'\n^^ Vessel\nid = 1\n^^ Vessel\nid = 2\n"
	result, err := yini.ParseString(text, yini.Options{})
	if err != nil {
		t.Fatalf("ParseString: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	want := map[string]any{"Fleet": map[string]any{"name": "First", "Vessel": map[string]any{"id": int64(1)}}}
	if got := result.Document.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values() = %#v, want %#v", got, want)
	}
	for _, check := range []struct {
		code string
		line int
	}{
		{"duplicate-key", 4}, {"duplicate-section", 7},
	} {
		found := false
		for _, diagnostic := range result.Diagnostics {
			if diagnostic.Code == check.code && diagnostic.Position.Line == check.line && diagnostic.Position.Column == 1 {
				found = true
			}
		}
		if !found {
			t.Fatalf("missing %s at %d:1: %#v", check.code, check.line, result.Diagnostics)
		}
	}
}

func TestNumbersAndStringConcatenation(t *testing.T) {
	text := "^ Bench\nmin = -9223372036854775808\nmax = 9223372036854775807\nhex = hex:FF\nduo = 0zX\nbin = %101\nexp = 3e4\nnegativeZero = -0.0\nlabel = 'count=' + 0x10\n"
	result, err := yini.ParseString(text, yini.Options{})
	if err != nil {
		t.Fatalf("ParseString: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	values := result.Document.Values()["Bench"].(map[string]any)
	for key, want := range map[string]any{
		"min": int64(math.MinInt64), "max": int64(math.MaxInt64),
		"hex": int64(255), "duo": int64(10), "bin": int64(5),
		"exp": float64(30000), "label": "count=16",
	} {
		if values[key] != want {
			t.Errorf("%s = %#v, want %#v", key, values[key], want)
		}
	}
	if !math.Signbit(values["negativeZero"].(float64)) {
		t.Error("negative floating-point zero lost its sign")
	}
	for _, literal := range []string{"9223372036854775808", "-9223372036854775809", "1e400", "1e-400"} {
		result, err := yini.ParseString("^ Bench\nvalue = "+literal+"\n", yini.Options{})
		if err == nil || !result.HasErrors() {
			t.Errorf("%s should fail: %#v", literal, result.Diagnostics)
		}
	}
}

func TestEOFAndTooling(t *testing.T) {
	result, err := yini.ParseString("^ Last\nvalue = C'line\\nend'", yini.Options{})
	if err != nil {
		t.Fatalf("final assignment at EOF: %v", err)
	}
	if got := result.Document.Values()["Last"].(map[string]any)["value"]; got != "line\nend" {
		t.Fatalf("decoded string = %#v", got)
	}
	_, err = yini.ParseString("^ Last\nvalue = 1\n/END", yini.Options{Mode: yini.ModeStrict})
	if err != nil {
		t.Fatalf("/END at EOF: %v", err)
	}
	result, err = yini.ParseForTooling("^ Last\nvalue = [1,,2]", yini.Options{})
	if err != nil || !result.HasErrors() {
		t.Fatalf("tooling parse = %#v, %v", result.Diagnostics, err)
	}
	_, err = yini.ParseString("^ Last\nvalue = [1,,2]", yini.Options{})
	var parseErr *yini.ParseError
	if !errors.As(err, &parseErr) || !strings.Contains(parseErr.Error(), "line") {
		t.Fatalf("expected positioned ParseError, got %v", err)
	}
}

func TestStrictFilenameDoesNotSelectMode(t *testing.T) {
	path := filepath.Join(t.TempDir(), "schedule.strict.yini")
	if err := os.WriteFile(path, []byte("orphan = 1"), 0600); err != nil {
		t.Fatal(err)
	}
	result, err := yini.ParseFile(path, yini.Options{})
	if err != nil || !hasDiagnostic(result, "filename-mode") {
		t.Fatalf("lenient file: %v (%#v)", err, result.Diagnostics)
	}
	if result.Document.Values()["orphan"] != int64(1) {
		t.Fatalf("filename changed parsing mode: %#v", result.Document.Values())
	}
	strict, err := yini.ParseFile(path, yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "orphan-member") || hasDiagnostic(strict, "filename-mode") {
		t.Fatalf("strict file: %v (%#v)", err, strict.Diagnostics)
	}
}

func TestReaderAndInvalidUTF8(t *testing.T) {
	result, err := yini.ParseReader(strings.NewReader("^ Passage\nopen = yes"), yini.Options{})
	if err != nil || result.Document.Values()["Passage"].(map[string]any)["open"] != true {
		t.Fatalf("ParseReader: %v (%#v)", err, result.Diagnostics)
	}
	result, err = yini.ParseString("^ Passage\rentry = \xff", yini.Options{})
	if err == nil || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "invalid-utf8" || result.Diagnostics[0].Position.Line != 2 {
		t.Fatalf("invalid UTF-8: %v (%#v)", err, result.Diagnostics)
	}
}
