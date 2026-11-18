package yini

import (
	"errors"
	"strings"
	"testing"
)

func TestDecodeString(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name  string
		token string
		want  string
	}{
		{"empty string", "''", ""},
		{"raw single quotes", `'C:\Depot\new'`, `C:\Depot\new`},
		{"raw double quotes", `"keep\ntext"`, `keep\ntext`},
		{"uppercase raw prefix", `R'keep\ttext'`, `keep\ttext`},
		{"lowercase raw prefix", `r"keep\ttext"`, `keep\ttext`},
		{"classic prefix", `C'first\nsecond'`, "first\nsecond"},
		{"lowercase classic prefix", `c'first\tsecond'`, "first\tsecond"},
		{"escaped prefix", `E"first\nsecond"`, "first\nsecond"},
		{"lowercase escaped prefix", `e'first\tsecond'`, "first\tsecond"},
		{"Unicode escapes", `E'\x41\u03A9\U0001F331'`, "AΩ🌱"},
		{"octal escapes", `E'\o1\o12\o101\o377'`, "\x01\nAÿ"},
		{"simple escapes", `E'\\\'\"\/\0\?\a\b\f\n\r\t\v'`, "\\'\"/\x00?\a\b\f\n\r\t\v"},
		{"escaped backtick", "E'\\`'", "`"},
		{"literal Unicode", "'moss 🌱'", "moss 🌱"},
		{"literal tab", "'left\tright'", "left\tright"},
		{"multiline LF", "\"\"\"first\nsecond\"\"\"", "first\nsecond"},
		{"multiline CRLF", "R\"\"\"first\r\nsecond\"\"\"", "first\nsecond"},
		{"multiline CR", "r\"\"\"first\rsecond\"\"\"", "first\nsecond"},
		{"escaped versus source newlines", "E\"\"\"first\r\n\\r\\nsecond\"\"\"", "first\n\r\nsecond"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			token := test.token

			// 2. Act.
			got, err := decodeString(token)

			// 3. Assert.
			if err != nil {
				t.Fatalf("decodeString(%q): %v", token, err)
			}
			if got != test.want {
				t.Errorf("decoded = %q, want %q", got, test.want)
			}
		})
	}
}

func TestDecodeStringErrors(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		token   string
		offset  int
		message string
	}{
		{"empty token", "", 0, "Empty string token"},
		{"unclosed quote", "'unfinished", 0, "closing quote"},
		{"unclosed triple quote", `"""unfinished`, 0, "closing triple quote"},
		{"unsupported escape", `E'bad\z'`, 5, "Unsupported escape"},
		{"Unicode before escape", `E'é\z'`, 4, "Unsupported escape"},
		{"incomplete escape", `E'last\'`, 6, "incomplete escape"},
		{"short Unicode escape", `E'\u12'`, 2, "needs 4 digits"},
		{"invalid hexadecimal escape", `E'\xG0'`, 2, "Invalid escape"},
		{"surrogate escape", `E'\uD800'`, 2, "Invalid escape"},
		{"out of range Unicode", `E'\U00110000'`, 2, "Invalid escape"},
		{"octal overflow", `E'\o400'`, 2, "Invalid escape"},
		{"invalid octal digit", `E'\o8'`, 2, "Invalid octal escape"},
		{"newline in single line string", "'a\nb'", 2, "control characters"},
		{"literal control character", "'é\x01'", 3, "control characters"},
		{"invalid UTF-8", "'a\xff'", 2, "invalid UTF-8"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			token := test.token
			var detail *stringDecodeError

			// 2. Act.
			got, err := decodeString(token)

			// 3. Assert.
			if !errors.As(err, &detail) {
				t.Fatalf("error = %v, want *stringDecodeError", err)
			}
			if detail.offset != test.offset || !strings.Contains(detail.Error(), test.message) {
				t.Errorf("error = %q at byte %d, want %q at byte %d", detail.Error(), detail.offset, test.message, test.offset)
			}
			if got != "" {
				t.Errorf("failed decode returned %q", got)
			}
		})
	}
}

func TestDecodeCommandString(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name string
		body string
		want string
	}{
		{"empty", "", ""},
		{"structural indentation", "  first\n    second\n", "first\n  second"},
		{"blank lines", " \n  first\n  \n", "\nfirst\n"},
		{"blank content", "  \n \n", "\n"},
		{"literal escapes", "  keep\\ntext\n", `keep\ntext`},
		{"tabs after indentation", "  first\n  \tsecond\n", "first\n\tsecond"},
	}
	for _, test := range tests {
		for _, ending := range []string{"\n", "\r\n", "\r"} {
			t.Run(test.name+"/"+strings.ReplaceAll(strings.ReplaceAll(ending, "\r", "CR"), "\n", "LF"), func(t *testing.T) {
				// 1. Assign.
				token := "command\"\"\"" + ending + strings.ReplaceAll(test.body, "\n", ending) + `"""`

				// 2. Act.
				got, err := decodeString(token)

				// 3. Assert.
				if err != nil {
					t.Fatalf("decodeString(%q): %v", token, err)
				}
				if got != test.want {
					t.Errorf("decoded command = %q, want %q", got, test.want)
				}
			})
		}
	}
}

func TestDecodeCommandStringErrors(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		token   string
		offset  int
		message string
	}{
		{"missing opening newline", `command"""text"""`, 10, "begin with a line break"},
		{"text after closing delimiter", "command\"\"\"\n  text\n\"\"\" tail", 18, "on its own line"},
		{"insufficient indentation", "command\"\"\"\n  é\n x\n\"\"\"", 17, "insufficient indentation"},
		{"CRLF indentation", "command\"\"\"\r\n  é\r\n x\r\n\"\"\"", 19, "insufficient indentation"},
		{"leading blank indentation", "command\"\"\"\n   \n  text\n\"\"\"", 13, "leading blank"},
		{"literal control character", "command\"\"\"\n  é\x01\n\"\"\"", 15, "control characters"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			token := test.token
			var detail *stringDecodeError

			// 2. Act.
			_, err := decodeString(token)

			// 3. Assert.
			if !errors.As(err, &detail) {
				t.Fatalf("error = %v, want *stringDecodeError", err)
			}
			if detail.offset != test.offset || !strings.Contains(detail.Error(), test.message) {
				t.Errorf("error = %q at byte %d, want %q at byte %d", detail.Error(), detail.offset, test.message, test.offset)
			}
		})
	}
}
