package yini

import "github.com/antlr4-go/antlr/v4"

// sourceLocator maps ANTLR rune offsets to positions in the original input.
// ANTLR's line counter does not advance for a lone CR, but YINI treats it as
// a line ending.
type sourceLocator struct {
	positions []Position
}

func newSourceLocator(input string, bom bool) *sourceLocator {
	characters := []rune(input)
	positions := make([]Position, len(characters)+1)
	position := Position{Line: 1, Column: 1}
	if bom {
		position.Column++
	}
	previousCR := false
	for index, char := range characters {
		positions[index] = position
		switch char {
		case '\r':
			position.Line++
			position.Column = 1
			previousCR = true
		case '\n':
			if !previousCR {
				position.Line++
			}
			position.Column = 1
			previousCR = false
		default:
			position.Column++
			previousCR = false
		}
	}
	positions[len(characters)] = position
	return &sourceLocator{positions: positions}
}

func (l *sourceLocator) token(token antlr.Token) Position {
	if token == nil {
		return Position{Line: 1, Column: 1}
	}
	return l.index(token.GetStart())
}

func (l *sourceLocator) index(index int) Position {
	if index < 0 {
		index = 0
	}
	if index >= len(l.positions) {
		index = len(l.positions) - 1
	}
	return l.positions[index]
}
