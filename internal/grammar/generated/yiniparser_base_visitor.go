// Code generated from grammar/v1.0.0-rc.6/YiniParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // YiniParser
import "github.com/antlr4-go/antlr/v4"

type BaseYiniParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseYiniParserVisitor) VisitYini(ctx *YiniContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitProlog(ctx *PrologContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitTerminal_stmt(ctx *Terminal_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitTerminal_trivia(ctx *Terminal_triviaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitFull_line_comment_stmt(ctx *Full_line_comment_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitDisabled_line_stmt(ctx *Disabled_line_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitInvalid_section_stmt(ctx *Invalid_section_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitMeta_stmt(ctx *Meta_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitYini_directive(ctx *Yini_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitYini_mode_declaration(ctx *Yini_mode_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitEol(ctx *EolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitScalar_value(ctx *Scalar_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitConcat_expression(ctx *Concat_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitConcat_tail(ctx *Concat_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitConcat_operand(ctx *Concat_operandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitObject_literal(ctx *Object_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitObject_members(ctx *Object_membersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitObject_member(ctx *Object_memberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitObject_member_separator(ctx *Object_member_separatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitList_literal(ctx *List_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitNumber_literal(ctx *Number_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitNull_literal(ctx *Null_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitBad_meta_text(ctx *Bad_meta_textContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYiniParserVisitor) VisitBad_member(ctx *Bad_memberContext) interface{} {
	return v.VisitChildren(ctx)
}
