// Generated from d:/Sources/YINI-lang-WORK/yini-parser-typescript/grammar/v1.0.0-rc.2x/YiniParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link YiniParser}.
 */
public interface YiniParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link YiniParser#yini}.
	 * @param ctx the parse tree
	 */
	void enterYini(YiniParser.YiniContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#yini}.
	 * @param ctx the parse tree
	 */
	void exitYini(YiniParser.YiniContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#prolog}.
	 * @param ctx the parse tree
	 */
	void enterProlog(YiniParser.PrologContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#prolog}.
	 * @param ctx the parse tree
	 */
	void exitProlog(YiniParser.PrologContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#terminal}.
	 * @param ctx the parse tree
	 */
	void enterTerminal(YiniParser.TerminalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#terminal}.
	 * @param ctx the parse tree
	 */
	void exitTerminal(YiniParser.TerminalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(YiniParser.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(YiniParser.StmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#marker_stmt}.
	 * @param ctx the parse tree
	 */
	void enterMarker_stmt(YiniParser.Marker_stmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#marker_stmt}.
	 * @param ctx the parse tree
	 */
	void exitMarker_stmt(YiniParser.Marker_stmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#eol}.
	 * @param ctx the parse tree
	 */
	void enterEol(YiniParser.EolContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#eol}.
	 * @param ctx the parse tree
	 */
	void exitEol(YiniParser.EolContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(YiniParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(YiniParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#member}.
	 * @param ctx the parse tree
	 */
	void enterMember(YiniParser.MemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#member}.
	 * @param ctx the parse tree
	 */
	void exitMember(YiniParser.MemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#listAfterColon}.
	 * @param ctx the parse tree
	 */
	void enterListAfterColon(YiniParser.ListAfterColonContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#listAfterColon}.
	 * @param ctx the parse tree
	 */
	void exitListAfterColon(YiniParser.ListAfterColonContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(YiniParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(YiniParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#object_literal}.
	 * @param ctx the parse tree
	 */
	void enterObject_literal(YiniParser.Object_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#object_literal}.
	 * @param ctx the parse tree
	 */
	void exitObject_literal(YiniParser.Object_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#object_members}.
	 * @param ctx the parse tree
	 */
	void enterObject_members(YiniParser.Object_membersContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#object_members}.
	 * @param ctx the parse tree
	 */
	void exitObject_members(YiniParser.Object_membersContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#object_member}.
	 * @param ctx the parse tree
	 */
	void enterObject_member(YiniParser.Object_memberContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#object_member}.
	 * @param ctx the parse tree
	 */
	void exitObject_member(YiniParser.Object_memberContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#list_literal}.
	 * @param ctx the parse tree
	 */
	void enterList_literal(YiniParser.List_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#list_literal}.
	 * @param ctx the parse tree
	 */
	void exitList_literal(YiniParser.List_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#elements}.
	 * @param ctx the parse tree
	 */
	void enterElements(YiniParser.ElementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#elements}.
	 * @param ctx the parse tree
	 */
	void exitElements(YiniParser.ElementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#number_literal}.
	 * @param ctx the parse tree
	 */
	void enterNumber_literal(YiniParser.Number_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#number_literal}.
	 * @param ctx the parse tree
	 */
	void exitNumber_literal(YiniParser.Number_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#null_literal}.
	 * @param ctx the parse tree
	 */
	void enterNull_literal(YiniParser.Null_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#null_literal}.
	 * @param ctx the parse tree
	 */
	void exitNull_literal(YiniParser.Null_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#string_literal}.
	 * @param ctx the parse tree
	 */
	void enterString_literal(YiniParser.String_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#string_literal}.
	 * @param ctx the parse tree
	 */
	void exitString_literal(YiniParser.String_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#string_concat}.
	 * @param ctx the parse tree
	 */
	void enterString_concat(YiniParser.String_concatContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#string_concat}.
	 * @param ctx the parse tree
	 */
	void exitString_concat(YiniParser.String_concatContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#boolean_literal}.
	 * @param ctx the parse tree
	 */
	void enterBoolean_literal(YiniParser.Boolean_literalContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#boolean_literal}.
	 * @param ctx the parse tree
	 */
	void exitBoolean_literal(YiniParser.Boolean_literalContext ctx);
	/**
	 * Enter a parse tree produced by {@link YiniParser#bad_member}.
	 * @param ctx the parse tree
	 */
	void enterBad_member(YiniParser.Bad_memberContext ctx);
	/**
	 * Exit a parse tree produced by {@link YiniParser#bad_member}.
	 * @param ctx the parse tree
	 */
	void exitBad_member(YiniParser.Bad_memberContext ctx);
}