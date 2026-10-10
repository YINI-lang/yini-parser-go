package yini

import (
	"fmt"
	"strings"

	"github.com/YINI-lang/yini-parser-go/internal/grammar/generated"
	"github.com/antlr4-go/antlr/v4"
)

func (b *documentBuilder) parseValue(context generated.IValueContext) (any, bool) {
	switch {
	case context == nil:
		return nil, false
	case context.Concat_expression() != nil:
		return b.parseConcat(context.Concat_expression())
	case context.Scalar_value() != nil:
		return b.parseScalar(context.Scalar_value())
	case context.List_literal() != nil:
		return b.parseList(context.List_literal())
	case context.Object_literal() != nil:
		return b.parseObject(context.Object_literal())
	default:
		b.report(SeverityError, "value", "Invalid YINI value", b.position(context.GetStart()))
		return nil, false
	}
}

func (b *documentBuilder) parseScalar(context generated.IScalar_valueContext) (any, bool) {
	switch {
	case context.Null_literal() != nil:
		return nil, true
	case context.String_literal() != nil:
		return b.parseStringToken(context.String_literal().GetStart())
	case context.Number_literal() != nil:
		value, _, ok := b.parseNumberToken(context.Number_literal().GetStart())
		return value, ok
	case context.Boolean_literal() != nil:
		text := strings.ToLower(context.Boolean_literal().GetText())
		return text == "true" || text == "yes" || text == "on", true
	default:
		b.report(SeverityError, "value", "Invalid YINI scalar value", b.position(context.GetStart()))
		return nil, false
	}
}

func (b *documentBuilder) parseStringToken(token antlr.Token) (any, bool) {
	if token == nil {
		return nil, false
	}
	value, err := decodeString(token.GetText())
	if err != nil {
		position := b.position(token)
		if detail, ok := err.(*stringDecodeError); ok {
			position = positionAtOffset(position, token.GetText(), detail.offset)
		}
		b.report(SeverityError, "string", err.Error(), position)
		return nil, false
	}
	return value, true
}

func (b *documentBuilder) parseNumberToken(token antlr.Token) (any, string, bool) {
	if token == nil {
		return nil, "", false
	}
	value, canonical, err := parseNumber(token.GetText())
	if err != nil {
		b.report(SeverityError, "number", err.Error(), b.position(token))
		return nil, "", false
	}
	return value, canonical, true
}

func (b *documentBuilder) parseConcat(context generated.IConcat_expressionContext) (any, bool) {
	first := context.STRING()
	if first == nil {
		return nil, false
	}
	parsed, ok := b.parseStringToken(first.GetSymbol())
	if !ok {
		return nil, false
	}
	var result strings.Builder
	result.WriteString(parsed.(string))
	for _, tail := range context.AllConcat_tail() {
		operand := tail.Concat_operand()
		if operand == nil {
			return nil, false
		}
		switch {
		case operand.STRING() != nil:
			value, valid := b.parseStringToken(operand.STRING().GetSymbol())
			if !valid {
				return nil, false
			}
			result.WriteString(value.(string))
		case b.mode == ModeStrict:
			b.report(SeverityError, "concat-strict", "Strict mode allows only strings in concatenation", b.position(operand.GetStart()))
			return nil, false
		case operand.NUMBER() != nil:
			_, canonical, valid := b.parseNumberToken(operand.NUMBER().GetSymbol())
			if !valid {
				return nil, false
			}
			result.WriteString(canonical)
		case operand.BOOLEAN_TRUE() != nil:
			result.WriteString("true")
		case operand.BOOLEAN_FALSE() != nil:
			result.WriteString("false")
		case operand.NULL() != nil:
			result.WriteString("null")
		default:
			return nil, false
		}
	}
	return result.String(), true
}

func (b *documentBuilder) parseList(context generated.IList_literalContext) (any, bool) {
	values := make([]any, 0)
	elements := context.Elements()
	if elements == nil {
		return values, true
	}
	if b.mode == ModeStrict && endsWithComma(elements) {
		b.report(SeverityError, "trailing-comma", "Strict mode does not allow a trailing comma in a list", b.position(elements.GetStop()))
		return nil, false
	}
	for _, element := range elements.AllValue() {
		value, ok := b.parseValue(element)
		if !ok {
			return nil, false
		}
		values = append(values, value)
	}
	return values, true
}

func (b *documentBuilder) parseObject(context generated.IObject_literalContext) (any, bool) {
	values := make(map[string]any)
	members := context.Object_members()
	if members == nil {
		return values, true
	}
	if b.mode == ModeStrict && endsWithComma(members) {
		b.report(SeverityError, "trailing-comma", "Strict mode does not allow a trailing comma in an inline object", b.position(members.GetStop()))
		return nil, false
	}
	for _, member := range members.AllObject_member() {
		if member.KEY() == nil {
			return nil, false
		}
		position := b.position(member.KEY().GetSymbol())
		name, err := parseIdentifier(member.KEY().GetText())
		if err != nil {
			b.report(SeverityError, "identifier", err.Error(), position)
			return nil, false
		}
		separator := member.Object_member_separator()
		if b.mode == ModeStrict && separator != nil && separator.EQ() != nil {
			b.report(SeverityError, "object-separator", "Strict mode requires ':' between inline object keys and values", b.position(separator.EQ().GetSymbol()))
			return nil, false
		}
		if _, exists := values[name]; exists {
			severity := SeverityWarning
			if b.mode == ModeStrict {
				severity = SeverityError
			}
			b.report(severity, "duplicate-object-key", fmt.Sprintf("Inline object member %q is already defined; keeping the first value", name), position)
			continue
		}
		value, ok := b.parseValue(member.Value())
		if !ok {
			return nil, false
		}
		values[name] = value
	}
	return values, true
}

func endsWithComma(context antlr.ParserRuleContext) bool {
	stop := context.GetStop()
	return stop != nil && stop.GetTokenType() == generated.YiniParserCOMMA
}
