package integration_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestParseFileErrorIncludesPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.yini")
	if err := os.WriteFile(path, []byte("^ Settings\nport = 1 ; note\n/END"), 0600); err != nil {
		t.Fatal(err)
	}

	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			result, err := yini.ParseFile(path, yini.Options{Mode: mode})
			if err == nil || !strings.Contains(err.Error(), path) {
				t.Fatalf("ParseFile error = %v; want file path", err)
			}
			var parseErr *yini.ParseError
			if !errors.As(err, &parseErr) {
				t.Fatalf("ParseFile error = %T; want wrapped *yini.ParseError", err)
			}
			if parseErr.Diagnostic.Position != (yini.Position{Line: 2, Column: 10}) {
				t.Fatalf("ParseFile location = %#v; want 2:10 (diagnostics: %#v)", parseErr.Diagnostic.Position, result.Diagnostics)
			}
			if !result.HasErrors() {
				t.Fatal("ParseFile omitted error diagnostics")
			}
		})
	}

	validPath := filepath.Join(t.TempDir(), "valid.yini")
	if err := os.WriteFile(validPath, []byte("^ Settings\nport = 1\n/END"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := yini.ParseFile(validPath, yini.Options{Mode: yini.ModeStrict}); err != nil {
		t.Fatalf("valid file: %v", err)
	}
}

func TestSyntaxDiagnosticsExplainCommonMistakes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		mode    yini.Mode
		message string
		line    int
		column  int
	}{
		{"unclosed list lenient", "^ Store\nitems = [1, 2", yini.ModeLenient, "Unclosed list; add ]", 3, 1},
		{"unclosed list strict", "^ Store\nitems = [1, 2", yini.ModeStrict, "Unclosed list; add ]", 3, 1},
		{"missing object separator lenient", "^ Store\nitem = {name 'x'}\n", yini.ModeLenient, "Expected : or = after an inline object key", 2, 14},
		{"missing object separator strict", "^ Store\nitem = {name 'x'}\n/END", yini.ModeStrict, "Expected : after an inline object key", 2, 14},
		{"missing list value", "^ Store\nitems = [1,,2]\n", yini.ModeLenient, "Unexpected comma in list; add a value before it", 2, 12},
		{"inline semicolon", "^ Store\nitem = 1 ; note\n", yini.ModeLenient, "Semicolon comments must start on their own line", 2, 10},
		{"content after end", "^ Store\n/END\nitem = 1\n", yini.ModeStrict, "Only comments and whitespace may follow /END", 3, 1},
		{"missing concatenation value", "^ Store\nitem = 'x' +\n", yini.ModeLenient, "Expected a value after +", 3, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := yini.ParseString(test.input, yini.Options{Mode: test.mode})
			if err == nil {
				t.Fatal("expected parse error")
			}
			for _, diagnostic := range result.Diagnostics {
				if diagnostic.Code == "syntax" && diagnostic.Message == test.message {
					if diagnostic.Position != (yini.Position{Line: test.line, Column: test.column}) {
						t.Fatalf("diagnostic location = %#v; want %d:%d", diagnostic.Position, test.line, test.column)
					}
					return
				}
			}
			t.Fatalf("missing syntax message %q in %#v", test.message, result.Diagnostics)
		})
	}
}
