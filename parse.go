package yini

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/YINI-lang/yini-parser-go/internal/grammar/generated"
	"github.com/antlr4-go/antlr/v4"
)

// ParseString parses inline YINI content. A result with diagnostics is returned
// even when the input is invalid.
func ParseString(content string, options Options) (Result, error) {
	mode := options.Mode
	if mode == "" {
		mode = ModeLenient
	}
	if mode != ModeLenient && mode != ModeStrict {
		return Result{}, fmt.Errorf("invalid YINI mode %q", mode)
	}
	result := Result{Document: &Document{}}
	if !utf8.ValidString(content) {
		result.Diagnostics = append(result.Diagnostics, Diagnostic{
			Severity: SeverityError,
			Code:     "invalid-utf8",
			Message:  "YINI content must be valid UTF-8",
			Position: firstInvalidUTF8Position(content),
		})
		return result, &ParseError{Diagnostic: result.Diagnostics[0]}
	}
	bom := strings.HasPrefix(content, "\uFEFF")
	input := strings.TrimPrefix(content, "\uFEFF")
	// The checked-in grammar expects an end-of-line token after assignments and
	// headers. EOF is also a valid final line ending for YINI input.
	if !strings.HasSuffix(input, "\n") && !strings.HasSuffix(input, "\r") {
		input += "\n"
	}
	input, misplacedShebang := normalizeShebang(input)
	locator := newSourceLocator(input, bom)
	if misplacedShebang >= 0 {
		severity := SeverityWarning
		if mode == ModeStrict {
			severity = SeverityError
		}
		result.Diagnostics = append(result.Diagnostics, Diagnostic{
			Severity: severity,
			Code:     "misplaced-shebang",
			Message:  "A shebang must start at the beginning of the first line; this line is treated as a comment",
			Position: locator.index(misplacedShebang),
		})
	}
	listener := &syntaxListener{result: &result, locator: locator, mode: mode}
	lexer := generated.NewYiniLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := generated.NewYiniParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	tree := parser.Yini()
	builder := documentBuilder{result: &result, mode: mode, locator: locator}
	builder.build(tree)
	sort.SliceStable(result.Diagnostics, func(i, j int) bool {
		left, right := result.Diagnostics[i].Position, result.Diagnostics[j].Position
		if left.Line != right.Line {
			return left.Line < right.Line
		}
		return left.Column < right.Column
	})
	result.Diagnostics = uniqueDiagnostics(result.Diagnostics)
	for _, diagnostic := range result.Diagnostics {
		if diagnostic.Severity == SeverityError {
			return result, &ParseError{Diagnostic: diagnostic}
		}
	}
	return result, nil
}

// The grammar recognizes #! as a shebang at any position. Later occurrences
// are comments by the specification, so replace the ! without moving tokens.
func normalizeShebang(input string) (string, int) {
	if !strings.Contains(input, "#!") {
		return input, -1
	}
	lexer := generated.NewYiniLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	runes := []rune(input)
	first := -1
	for _, token := range lexer.GetAllTokens() {
		if token.GetTokenType() != generated.YiniLexerSHEBANG || token.GetStart() == 0 {
			continue
		}
		index := token.GetStart()
		if first < 0 {
			first = index
		}
		runes[index+1] = '#'
	}
	if first < 0 {
		return input, -1
	}
	return string(runes), first
}

// ParseFile reads and parses a UTF-8 YINI file. The filename does not select
// the parsing mode; a .strict.yini file parsed leniently gets a warning.
// Parse errors include the file path and can be inspected with errors.As.
func ParseFile(path string, options Options) (Result, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Result{}, err
	}
	result, err := ParseString(string(content), options)
	if strings.HasSuffix(strings.ToLower(path), ".strict.yini") && options.Mode != ModeStrict && (options.Mode == "" || options.Mode == ModeLenient) {
		result.Diagnostics = append([]Diagnostic{{
			Severity: SeverityWarning,
			Code:     "filename-mode",
			Message:  "The .strict.yini filename suggests strict mode, but lenient mode was selected",
			Position: Position{Line: 1, Column: 1},
		}}, result.Diagnostics...)
	}
	if err != nil {
		return result, fmt.Errorf("%s: %w", path, err)
	}
	return result, err
}

// ParseReader parses YINI content from a reader.
func ParseReader(reader io.Reader, options Options) (Result, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return Result{}, err
	}
	return ParseString(string(content), options)
}

// ParseForTooling returns diagnostics and any recoverable document content
// without returning a Go error for normal YINI syntax and validation issues.
func ParseForTooling(content string, options Options) (Result, error) {
	result, err := ParseString(content, options)
	if _, ok := err.(*ParseError); ok {
		return result, nil
	}
	return result, err
}

type syntaxListener struct {
	*antlr.DefaultErrorListener
	result  *Result
	locator *sourceLocator
	mode    Mode
}

func (l *syntaxListener) SyntaxError(recognizer antlr.Recognizer, offending any, line, column int, _ string, _ antlr.RecognitionException) {
	token, _ := offending.(antlr.Token)
	message := l.syntaxMessage(recognizer, token)
	position := Position{Line: line, Column: column + 1}
	if token != nil {
		position = l.locator.token(token)
	} else if lexer, ok := recognizer.(antlr.Lexer); ok {
		index := lexer.GetInputStream().Index()
		if index > 0 {
			index--
		}
		position = l.locator.index(index)
	}
	l.result.Diagnostics = append(l.result.Diagnostics, Diagnostic{
		Severity: SeverityError,
		Code:     "syntax",
		Message:  message,
		Position: position,
	})
}

func (l *syntaxListener) syntaxMessage(recognizer antlr.Recognizer, token antlr.Token) string {
	if token == nil {
		return "Unrecognized YINI text; check for an invalid character or an unclosed quote"
	}
	parser, ok := recognizer.(antlr.Parser)
	if !ok {
		return fmt.Sprintf("Unexpected %q in YINI content", token.GetText())
	}
	context := parser.GetParserRuleContext()
	tokenType := token.GetTokenType()
	collection := syntaxNearestCollection(context)
	switch {
	case tokenType == antlr.TokenEOF && collection == generated.YiniParserRULE_list_literal:
		return "Unclosed list; add ]"
	case tokenType == antlr.TokenEOF && collection == generated.YiniParserRULE_object_literal:
		return "Unclosed inline object; add }"
	case tokenType == antlr.TokenEOF && syntaxInRule(context, generated.YiniParserRULE_concat_tail):
		return "Expected a value after +"
	case syntaxInRule(context, generated.YiniParserRULE_object_member_separator):
		if l.mode == ModeStrict {
			return "Expected : after an inline object key"
		}
		return "Expected : or = after an inline object key"
	case tokenType == generated.YiniLexerCOMMA && collection == generated.YiniParserRULE_list_literal:
		return "Unexpected comma in list; add a value before it"
	case tokenType == generated.YiniLexerCC && collection == generated.YiniParserRULE_list_literal:
		return "Expected ] to close the list"
	case tokenType == generated.YiniLexerCB && collection == generated.YiniParserRULE_object_literal:
		return "Expected } to close the inline object"
	case tokenType == generated.YiniLexerFULL_LINE_COMMENT && syntaxInRule(context, generated.YiniParserRULE_eol):
		return "Semicolon comments must start on their own line"
	case tokenType == generated.YiniLexerPLUS && syntaxInRule(context, generated.YiniParserRULE_eol):
		return "Concatenation must start with a string and place + on the same line"
	}
	if root, ok := context.(*generated.YiniContext); ok && root.Terminal_stmt() != nil {
		return "Only comments and whitespace may follow /END"
	}
	if tokenType == antlr.TokenEOF {
		return "Unexpected end of YINI content"
	}
	return fmt.Sprintf("Unexpected %q in YINI content", token.GetText())
}

func syntaxInRule(context antlr.ParserRuleContext, rule int) bool {
	for context != nil {
		if context.GetRuleIndex() == rule {
			return true
		}
		parent, ok := context.GetParent().(antlr.ParserRuleContext)
		if !ok {
			break
		}
		context = parent
	}
	return false
}

func syntaxNearestCollection(context antlr.ParserRuleContext) int {
	for context != nil {
		switch context.GetRuleIndex() {
		case generated.YiniParserRULE_list_literal, generated.YiniParserRULE_object_literal:
			return context.GetRuleIndex()
		}
		parent, ok := context.GetParent().(antlr.ParserRuleContext)
		if !ok {
			break
		}
		context = parent
	}
	return -1
}

func uniqueDiagnostics(diagnostics []Diagnostic) []Diagnostic {
	unique := diagnostics[:0]
	for _, diagnostic := range diagnostics {
		if len(unique) > 0 {
			last := unique[len(unique)-1]
			if last.Code == diagnostic.Code && last.Position == diagnostic.Position && last.Message == diagnostic.Message {
				continue
			}
		}
		unique = append(unique, diagnostic)
	}
	return unique
}

func firstInvalidUTF8Position(content string) Position {
	position := Position{Line: 1, Column: 1}
	previousCR := false
	for len(content) > 0 {
		char, size := utf8.DecodeRuneInString(content)
		if char == utf8.RuneError && size == 1 {
			return position
		}
		if char == '\r' {
			position.Line++
			position.Column = 1
			previousCR = true
		} else if char == '\n' {
			if !previousCR {
				position.Line++
			}
			position.Column = 1
			previousCR = false
		} else {
			position.Column++
			previousCR = false
		}
		content = content[size:]
	}
	return position
}
