package yini

import (
	"strings"
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name string
		text string
		want string
	}{
		{"plain", "station7", "station7"},
		{"underscore", "_station", "_station"},
		{"Boolean spelling", "TRUE", "TRUE"},
		{"null spelling", "null", "null"},
		{"quoted punctuation", "`ui.theme`", "ui.theme"},
		{"quoted Unicode", "`café 🌱`", "café 🌱"},
		{"empty quoted name", "``", ""},
		{"escaped backtick", "`bay\\`name`", "bay`name"},
		{"escaped backslash", "`bay\\\\name`", "bay\\name"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			text := test.text

			// 2. Act.
			got, err := parseIdentifier(text)

			// 3. Assert.
			if err != nil {
				t.Fatalf("parseIdentifier(%q): %v", text, err)
			}
			if got != test.want {
				t.Errorf("name = %q, want %q", got, test.want)
			}
		})
	}
}

func TestParseIdentifierErrors(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		text    string
		message string
	}{
		{"empty", "", "cannot be empty"},
		{"leading digit", "7station", "use backticks"},
		{"unquoted punctuation", "ui.theme", "use backticks"},
		{"unquoted Unicode", "café", "use backticks"},
		{"unsupported escape", "`bay\\q`", "only escaped backticks and backslashes"},
		{"trailing backslash", "`bay\\`", "only escaped backticks and backslashes"},
		{"ASCII control", "`bay\tname`", "control character"},
		{"Unicode control", "`bay\u0085name`", "control character"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			text := test.text

			// 2. Act.
			_, err := parseIdentifier(text)

			// 3. Assert.
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Errorf("error = %v, want a message containing %q", err, test.message)
			}
		})
	}
}

func TestParseSectionHeader(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name  string
		text  string
		level int
		want  string
	}{
		{"root", "^ Depot\n", 1, "Depot"},
		{"repeated markers", "^^^ Depot\n", 3, "Depot"},
		{"marker separators", "^_^_^ Depot\n", 3, "Depot"},
		{"maximum repeated level", "^^^^^^^^^ Depot\n", 9, "Depot"},
		{"numeric shorthand", "^12 Depot\n", 12, "Depot"},
		{"maximum section depth", "^255 Depot\n", 255, "Depot"},
		{"section sign", "§§ Depot\n", 2, "Depot"},
		{"greater than markers", ">> Depot\n", 2, "Depot"},
		{"less than markers", "<< Depot\n", 2, "Depot"},
		{"adjacent quoted name", "^`Loading bay`\n", 1, "Loading bay"},
		{"escaped name", "^^`bay\\`name`\n", 2, "bay`name"},
		{"comments", "^ /* before */ Depot /* after */\r\n", 1, "Depot"},
		{"hash comment", "^ Depot# note", 1, "Depot"},
		{"slash comment", "^ Depot// note", 1, "Depot"},
		{"EOF", "^ Depot", 1, "Depot"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			text := test.text

			// 2. Act.
			level, name, err := parseSectionHeader(text)

			// 3. Assert.
			if err != nil {
				t.Fatalf("parseSectionHeader(%q): %v", text, err)
			}
			if level != test.level || name != test.want {
				t.Errorf("header = (%d, %q), want (%d, %q)", level, name, test.level, test.want)
			}
		})
	}
}

func TestParseSectionHeaderErrors(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		text    string
		message string
	}{
		{"missing name", "^", "needs a name"},
		{"whitespace name", "^ \t\n", "needs a name"},
		{"comment without name", "^ /* note */\n", "needs a name"},
		{"too many repeated markers", "^^^^^^^^^^ Depot\n", "numeric shorthand"},
		{"numeric depth overflow", "^999999999999999999999999999999999999999 Depot\n", "depth is too large"},
		{"unclosed comment", "^ /* unfinished", "unclosed block comment"},
		{"multiline comment before name", "^ /* first\nsecond */ Depot\n", "same line"},
		{"unclosed name", "^ `Loading bay\n", "closing backtick"},
		{"invalid name", "^ Depot.Part\n", "use backticks"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			text := test.text

			// 2. Act.
			_, _, err := parseSectionHeader(text)

			// 3. Assert.
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Errorf("error = %v, want a message containing %q", err, test.message)
			}
		})
	}
}
