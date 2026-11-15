package yini

import (
	"strings"

	"github.com/YINI-lang/yini-parser-go/internal/grammar/generated"
	"github.com/antlr4-go/antlr/v4"
)

type delimiterState struct {
	kind int
	key  bool
}

// objectKeyLexer resolves the grammar's hex: numeric prefix in object key
// positions, where {hex:10} means the key "hex" with the decimal value 10.
type objectKeyLexer struct {
	*generated.YiniLexer
	delimiters []delimiterState
}

func (l *objectKeyLexer) NextToken() antlr.Token {
	token := l.YiniLexer.NextToken()
	depth := len(l.delimiters)
	if depth > 0 && l.delimiters[depth-1].key && token.GetTokenType() == generated.YiniLexerNUMBER && strings.HasPrefix(strings.ToLower(token.GetText()), "hex:") {
		key := l.GetTokenFactory().Create(token.GetSource(), generated.YiniLexerKEY,
			token.GetText()[:3], token.GetChannel(), token.GetStart(), token.GetStart()+2,
			token.GetLine(), token.GetColumn())
		// Re-lex from the colon so fractions, signs, and base prefixes in the
		// value follow the ordinary lexer rules. Reset also clears cached EOF.
		input := l.GetInputStream()
		l.YiniLexer.Reset()
		input.Seek(token.GetStart() + 3)
		simulator := l.Interpreter.(*antlr.LexerATNSimulator)
		simulator.Line = token.GetLine()
		simulator.CharPositionInLine = token.GetColumn() + 3
		token = key
	}
	switch token.GetTokenType() {
	case generated.YiniLexerNL, generated.YiniLexerFULL_LINE_COMMENT, generated.YiniLexerDISABLED_LINE:
		return token
	case generated.YiniLexerOC, generated.YiniLexerOB:
		l.delimiters = append(l.delimiters, delimiterState{
			kind: token.GetTokenType(),
			key:  token.GetTokenType() == generated.YiniLexerOC,
		})
	case generated.YiniLexerCC, generated.YiniLexerCB:
		if depth > 0 {
			l.delimiters = l.delimiters[:depth-1]
		}
	case generated.YiniLexerCOMMA:
		if depth > 0 {
			l.delimiters[depth-1].key = l.delimiters[depth-1].kind == generated.YiniLexerOC
		}
	default:
		if depth > 0 {
			l.delimiters[depth-1].key = false
		}
	}
	return token
}
