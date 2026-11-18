package yini

import "testing"

func TestSourceLocatorIndex(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name  string
		input string
		bom   bool
		index int
		want  Position
	}{
		{"empty input", "", false, 0, Position{Line: 1, Column: 1}},
		{"negative offset", "abc", false, -1, Position{Line: 1, Column: 1}},
		{"EOF", "abc", false, 3, Position{Line: 1, Column: 4}},
		{"beyond EOF", "abc", false, 20, Position{Line: 1, Column: 4}},
		{"Unicode rune offset", "é🌱x", false, 2, Position{Line: 1, Column: 3}},
		{"Unicode EOF", "é🌱x", false, 3, Position{Line: 1, Column: 4}},
		{"tab column", "\tx", false, 1, Position{Line: 1, Column: 2}},
		{"LF", "a\nb", false, 2, Position{Line: 2, Column: 1}},
		{"CR", "a\rb", false, 2, Position{Line: 2, Column: 1}},
		{"CRLF", "a\r\nb", false, 3, Position{Line: 2, Column: 1}},
		{"consecutive CR", "a\r\rb", false, 3, Position{Line: 3, Column: 1}},
		{"mixed endings", "a\r\nb\rc\nd", false, 7, Position{Line: 4, Column: 1}},
		{"trailing newline", "a\r\n", false, 3, Position{Line: 2, Column: 1}},
		{"BOM first character", "abc", true, 0, Position{Line: 1, Column: 2}},
		{"BOM first line EOF", "é", true, 1, Position{Line: 1, Column: 3}},
		{"BOM next line", "a\nb", true, 2, Position{Line: 2, Column: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			locator := newSourceLocator(test.input, test.bom)

			// 2. Act.
			got := locator.index(test.index)

			// 3. Assert.
			if got != test.want {
				t.Errorf("position = %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestSourceLocatorNilToken(t *testing.T) {
	// 1. Assign.
	locator := newSourceLocator("text", true)

	// 2. Act.
	got := locator.token(nil)

	// 3. Assert.
	if got != (Position{Line: 1, Column: 1}) {
		t.Errorf("nil token position = %#v, want 1:1", got)
	}
}

func TestPositionAtOffset(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name   string
		text   string
		offset int
		want   Position
	}{
		{"start", "abc", 0, Position{Line: 4, Column: 7}},
		{"ASCII", "abc", 2, Position{Line: 4, Column: 9}},
		{"Unicode byte offset", "é🌱x", 6, Position{Line: 4, Column: 9}},
		{"EOF", "é🌱x", 7, Position{Line: 4, Column: 10}},
		{"beyond EOF", "é🌱x", 100, Position{Line: 4, Column: 10}},
		{"LF", "é\nx", 3, Position{Line: 5, Column: 1}},
		{"CR", "é\rx", 3, Position{Line: 5, Column: 1}},
		{"CRLF", "é\r\nx", 4, Position{Line: 5, Column: 1}},
		{"Unicode after newline", "é\r\n🌱x", 8, Position{Line: 5, Column: 2}},
		{"mixed endings", "a\r\nb\rc\nd", 7, Position{Line: 7, Column: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			start := Position{Line: 4, Column: 7}

			// 2. Act.
			got := positionAtOffset(start, test.text, test.offset)

			// 3. Assert.
			if got != test.want {
				t.Errorf("position = %#v, want %#v", got, test.want)
			}
		})
	}
}
