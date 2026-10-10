package yini

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/YINI-lang/yini-parser-go/internal/grammar/generated"
	"github.com/antlr4-go/antlr/v4"
)

type documentBuilder struct {
	result       *Result
	mode         Mode
	locator      *sourceLocator
	stack        []*Section
	ignoredLevel int
	seenContent  bool
	seenMarker   bool
}

func (b *documentBuilder) build(tree generated.IYiniContext) {
	for _, statement := range tree.AllStmt() {
		switch {
		case statement.SECTION_HEAD() != nil:
			b.addSection(statement.SECTION_HEAD().GetSymbol())
		case statement.Invalid_section_stmt() != nil:
			b.report(SeverityError, "section-header", "Invalid section header", b.position(statement.GetStart()))
		case statement.Assignment() != nil:
			b.addMember(statement.Assignment().Member())
		case statement.Meta_stmt() != nil:
			b.handleMeta(statement.Meta_stmt())
		case statement.Bad_member() != nil:
			b.report(SeverityError, "member", "Invalid member; expected a key, '=', and a value", b.position(statement.GetStart()))
		}
	}
	if !b.seenContent {
		severity := SeverityWarning
		if b.mode == ModeStrict {
			severity = SeverityError
		}
		b.report(severity, "empty-document", "YINI document contains no sections or members", Position{1, 1})
	}
	if b.mode == ModeStrict {
		if len(b.result.Document.Sections) != 1 {
			b.report(SeverityError, "top-level-section", "Strict mode requires exactly one top-level section", Position{1, 1})
		}
		if tree.Terminal_stmt() == nil {
			b.report(SeverityError, "missing-end", "Strict mode requires /END at the end of the document", b.position(tree.GetStop()))
		}
	}
}

func (b *documentBuilder) handleMeta(meta generated.IMeta_stmtContext) {
	if directive := meta.Directive(); directive != nil {
		if marker := directive.Yini_directive(); marker != nil {
			position := b.position(marker.GetStart())
			if b.seenContent || b.seenMarker {
				b.report(SeverityError, "marker-position", "@yini must appear once, before sections and members", position)
			}
			b.seenMarker = true
			if declaration := marker.Yini_mode_declaration(); declaration != nil && declaration.KEY() != nil {
				declared := strings.ToLower(declaration.KEY().GetText())
				switch declared {
				case string(ModeStrict):
					if b.mode == ModeLenient {
						b.report(SeverityError, "mode-mismatch", "Document declares strict mode, but the parser is in lenient mode", position)
					}
				case string(ModeLenient):
					if b.mode == ModeStrict {
						b.report(SeverityWarning, "mode-mismatch", "Document declares lenient mode, but the parser is in strict mode", position)
					}
				default:
					b.report(SeverityError, "mode-declaration", "@yini mode must be strict or lenient", position)
				}
			}
		} else {
			b.report(SeverityError, "reserved-directive", "This directive is reserved and is not supported by YINI RC.6", b.position(directive.GetStart()))
		}
		return
	}
	b.report(SeverityError, "reserved-directive", "This annotation or directive is not supported by YINI RC.6", b.position(meta.GetStart()))
}

func (b *documentBuilder) addSection(token antlr.Token) {
	level, name, err := parseSectionHeader(token.GetText())
	position := b.position(token)
	if err != nil {
		b.report(SeverityError, "section-header", err.Error(), position)
		return
	}
	if b.ignoredLevel > 0 {
		if level > b.ignoredLevel {
			return
		}
		b.ignoredLevel = 0
	}
	b.seenContent = true
	if level > 255 {
		b.report(SeverityError, "section-depth", "Section depth cannot exceed 255", position)
		b.ignoredLevel = level
		return
	}
	if level > len(b.stack)+1 {
		b.report(SeverityError, "section-depth", "Section nesting cannot skip a level", position)
		b.ignoredLevel = level
		return
	}
	if level <= len(b.stack) {
		b.stack = b.stack[:level-1]
	}
	if b.mode == ModeStrict && level == 1 && len(b.result.Document.Sections) > 0 {
		b.report(SeverityError, "top-level-section", "Strict mode allows only one top-level section", position)
		b.ignoredLevel = level
		return
	}
	var members []*Member
	var siblings []*Section
	if level == 1 {
		members = b.result.Document.Members
		siblings = b.result.Document.Sections
	} else {
		parent := b.stack[level-2]
		members, siblings = parent.Members, parent.Sections
	}
	for _, member := range members {
		if member.Name == name {
			b.report(SeverityError, "name-collision", fmt.Sprintf("Section %q conflicts with a member at the same level", name), position)
			b.ignoredLevel = level
			return
		}
	}
	for _, sibling := range siblings {
		if sibling.Name == name {
			severity := SeverityWarning
			if b.mode == ModeStrict {
				severity = SeverityError
			}
			b.report(severity, "duplicate-section", fmt.Sprintf("Section %q is already defined; keeping the first section", name), position)
			b.ignoredLevel = level
			return
		}
	}
	section := &Section{Name: name, Level: level, Position: position}
	if level == 1 {
		b.result.Document.Sections = append(b.result.Document.Sections, section)
	} else {
		parent := b.stack[level-2]
		parent.Sections = append(parent.Sections, section)
	}
	b.stack = append(b.stack, section)
}

func (b *documentBuilder) addMember(member generated.IMemberContext) {
	if member == nil || member.KEY() == nil {
		return
	}
	if b.ignoredLevel > 0 {
		return
	}
	b.seenContent = true
	keyToken := member.KEY().GetSymbol()
	position := b.position(keyToken)
	name, err := parseIdentifier(keyToken.GetText())
	if err != nil {
		b.report(SeverityError, "identifier", err.Error(), position)
		return
	}
	if b.mode == ModeStrict && len(b.stack) == 0 {
		b.report(SeverityError, "orphan-member", fmt.Sprintf("Member %q must be inside the top-level section in strict mode", name), position)
		return
	}
	var members *[]*Member
	var sections []*Section
	if len(b.stack) == 0 {
		members = &b.result.Document.Members
		sections = b.result.Document.Sections
	} else {
		current := b.stack[len(b.stack)-1]
		members = &current.Members
		sections = current.Sections
	}
	for _, previous := range *members {
		if previous.Name == name {
			severity := SeverityWarning
			if b.mode == ModeStrict {
				severity = SeverityError
			}
			b.report(severity, "duplicate-key", fmt.Sprintf("Member %q is already defined; keeping the first value", name), position)
			return
		}
	}
	for _, section := range sections {
		if section.Name == name {
			b.report(SeverityError, "name-collision", fmt.Sprintf("Member %q conflicts with a section at the same level", name), position)
			return
		}
	}
	var value any
	if member.Value() == nil {
		if b.mode == ModeStrict {
			b.report(SeverityError, "empty-value", fmt.Sprintf("Member %q needs a value; write null explicitly if intended", name), b.position(member.EQ().GetSymbol()))
			return
		}
	} else {
		var valid bool
		value, valid = b.parseValue(member.Value())
		if !valid {
			return
		}
	}
	*members = append(*members, &Member{Name: name, Value: value, Position: position})
}

func (b *documentBuilder) report(severity Severity, code, message string, position Position) {
	b.result.Diagnostics = append(b.result.Diagnostics, Diagnostic{
		Severity: severity,
		Code:     code,
		Message:  message,
		Position: position,
	})
}

func (b *documentBuilder) position(token antlr.Token) Position {
	return b.locator.token(token)
}

func positionAtOffset(start Position, text string, offset int) Position {
	if offset > len(text) {
		offset = len(text)
	}
	for index := 0; index < offset; {
		char, size := utf8.DecodeRuneInString(text[index:])
		if char == '\r' {
			start.Line++
			start.Column = 1
			if index+size < offset && text[index+size] == '\n' {
				size++
			}
		} else if char == '\n' {
			start.Line++
			start.Column = 1
		} else {
			start.Column++
		}
		index += size
	}
	return start
}

func parseSectionHeader(text string) (int, string, error) {
	text = strings.TrimRight(text, "\r\n")
	characters := []rune(text)
	if len(characters) < 2 {
		return 0, "", fmt.Errorf("section header needs a name")
	}
	marker := characters[0]
	index := 1
	level := 1
	if index < len(characters) && characters[index] >= '0' && characters[index] <= '9' {
		start := index
		for index < len(characters) && characters[index] >= '0' && characters[index] <= '9' {
			index++
		}
		parsed, err := strconv.Atoi(string(characters[start:index]))
		if err != nil {
			return 0, "", fmt.Errorf("section depth is too large")
		}
		level = parsed
	} else {
		for index < len(characters) && (characters[index] == marker || characters[index] == '_') {
			if characters[index] == marker {
				level++
			}
			index++
		}
		if level > 9 {
			return 0, "", fmt.Errorf("repeated section markers support only levels 1 through 9; use numeric shorthand")
		}
	}
	for index < len(characters) && (characters[index] == ' ' || characters[index] == '\t') {
		index++
	}
	if index >= len(characters) {
		return 0, "", fmt.Errorf("section header needs a name")
	}
	start := index
	if characters[index] == '`' {
		index++
		for index < len(characters) && characters[index] != '`' {
			index++
		}
		if index >= len(characters) {
			return 0, "", fmt.Errorf("section name is missing a closing backtick")
		}
		index++
	} else {
		for index < len(characters) && characters[index] != ' ' && characters[index] != '\t' && characters[index] != '#' {
			if characters[index] == '/' && index+1 < len(characters) && characters[index+1] == '/' {
				break
			}
			index++
		}
	}
	name, err := parseIdentifier(string(characters[start:index]))
	return level, name, err
}

func parseIdentifier(text string) (string, error) {
	if strings.HasPrefix(text, "`") && strings.HasSuffix(text, "`") && len(text) >= 2 {
		name := text[1 : len(text)-1]
		for _, char := range name {
			if char < 0x20 || (char >= 0x7F && char <= 0x9F) {
				return "", fmt.Errorf("backticked identifier contains a control character")
			}
		}
		return name, nil
	}
	for index, char := range text {
		valid := char == '_' || char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
		if index > 0 {
			valid = valid || char >= '0' && char <= '9'
		}
		if !valid {
			return "", fmt.Errorf("identifier %q is invalid; use backticks for spaces or special characters", text)
		}
	}
	if text == "" {
		return "", fmt.Errorf("identifier cannot be empty unless written as backticks")
	}
	return text, nil
}
