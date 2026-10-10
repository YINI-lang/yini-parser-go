// Code generated from grammar/v1.0.0-rc.6/YiniParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // YiniParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by YiniParser.
type YiniParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by YiniParser#yini.
	VisitYini(ctx *YiniContext) interface{}

	// Visit a parse tree produced by YiniParser#prolog.
	VisitProlog(ctx *PrologContext) interface{}

	// Visit a parse tree produced by YiniParser#terminal_stmt.
	VisitTerminal_stmt(ctx *Terminal_stmtContext) interface{}

	// Visit a parse tree produced by YiniParser#terminal_trivia.
	VisitTerminal_trivia(ctx *Terminal_triviaContext) interface{}

	// Visit a parse tree produced by YiniParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by YiniParser#full_line_comment_stmt.
	VisitFull_line_comment_stmt(ctx *Full_line_comment_stmtContext) interface{}

	// Visit a parse tree produced by YiniParser#disabled_line_stmt.
	VisitDisabled_line_stmt(ctx *Disabled_line_stmtContext) interface{}

	// Visit a parse tree produced by YiniParser#invalid_section_stmt.
	VisitInvalid_section_stmt(ctx *Invalid_section_stmtContext) interface{}

	// Visit a parse tree produced by YiniParser#meta_stmt.
	VisitMeta_stmt(ctx *Meta_stmtContext) interface{}

	// Visit a parse tree produced by YiniParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by YiniParser#yini_directive.
	VisitYini_directive(ctx *Yini_directiveContext) interface{}

	// Visit a parse tree produced by YiniParser#yini_mode_declaration.
	VisitYini_mode_declaration(ctx *Yini_mode_declarationContext) interface{}

	// Visit a parse tree produced by YiniParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by YiniParser#eol.
	VisitEol(ctx *EolContext) interface{}

	// Visit a parse tree produced by YiniParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by YiniParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by YiniParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by YiniParser#scalar_value.
	VisitScalar_value(ctx *Scalar_valueContext) interface{}

	// Visit a parse tree produced by YiniParser#concat_expression.
	VisitConcat_expression(ctx *Concat_expressionContext) interface{}

	// Visit a parse tree produced by YiniParser#concat_tail.
	VisitConcat_tail(ctx *Concat_tailContext) interface{}

	// Visit a parse tree produced by YiniParser#concat_operand.
	VisitConcat_operand(ctx *Concat_operandContext) interface{}

	// Visit a parse tree produced by YiniParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#object_literal.
	VisitObject_literal(ctx *Object_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#object_members.
	VisitObject_members(ctx *Object_membersContext) interface{}

	// Visit a parse tree produced by YiniParser#object_member.
	VisitObject_member(ctx *Object_memberContext) interface{}

	// Visit a parse tree produced by YiniParser#object_member_separator.
	VisitObject_member_separator(ctx *Object_member_separatorContext) interface{}

	// Visit a parse tree produced by YiniParser#list_literal.
	VisitList_literal(ctx *List_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by YiniParser#number_literal.
	VisitNumber_literal(ctx *Number_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#null_literal.
	VisitNull_literal(ctx *Null_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by YiniParser#bad_meta_text.
	VisitBad_meta_text(ctx *Bad_meta_textContext) interface{}

	// Visit a parse tree produced by YiniParser#bad_member.
	VisitBad_member(ctx *Bad_memberContext) interface{}
}
