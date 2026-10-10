// Generated from d:/Sources/YINI-lang-WORK/yini-parser-typescript/grammar/v1.0.0-rc.2x/YiniParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class YiniParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		YINI_MARKER=1, SECTION_HEAD=2, TERMINAL_TOKEN=3, SS=4, EUR=5, CARET=6, 
		GT=7, LT=8, EQ=9, HASH=10, COMMA=11, COLON=12, OB=13, CB=14, OC=15, CC=16, 
		PLUS=17, DOLLAR=18, PC=19, AT=20, BOOLEAN_FALSE=21, BOOLEAN_TRUE=22, NULL=23, 
		EMPTY_OBJECT=24, EMPTY_LIST=25, SHEBANG=26, NUMBER=27, KEY=28, IDENT=29, 
		IDENT_BACKTICKED=30, STRING=31, TRIPLE_QUOTED_STRING=32, SINGLE_OR_DOUBLE=33, 
		R_AND_C_STRING=34, HYPER_STRING=35, ESC_SEQ=36, ESC_SEQ_BASE=37, NL=38, 
		SINGLE_NL=39, WS=40, BLOCK_COMMENT=41, COMMENT=42, LINE_COMMENT=43, INLINE_COMMENT=44, 
		IDENT_INVALID=45, REST=46;
	public static final int
		RULE_yini = 0, RULE_prolog = 1, RULE_terminal_stmt = 2, RULE_stmt = 3, 
		RULE_marker_stmt = 4, RULE_eol = 5, RULE_assignment = 6, RULE_member = 7, 
		RULE_colon_list_decl = 8, RULE_value = 9, RULE_object_literal = 10, RULE_object_members = 11, 
		RULE_object_member = 12, RULE_list_literal = 13, RULE_elements = 14, RULE_number_literal = 15, 
		RULE_null_literal = 16, RULE_string_literal = 17, RULE_string_concat = 18, 
		RULE_boolean_literal = 19, RULE_bad_member = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"yini", "prolog", "terminal_stmt", "stmt", "marker_stmt", "eol", "assignment", 
			"member", "colon_list_decl", "value", "object_literal", "object_members", 
			"object_member", "list_literal", "elements", "number_literal", "null_literal", 
			"string_literal", "string_concat", "boolean_literal", "bad_member"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'\\u00A7'", "'\\u20AC'", "'^'", "'>'", "'<'", 
			"'='", "'#'", "','", "':'", "'['", "']'", "'{'", "'}'", "'+'", "'$'", 
			"'%'", "'@'", null, null, null, "'{}'", "'[]'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "YINI_MARKER", "SECTION_HEAD", "TERMINAL_TOKEN", "SS", "EUR", "CARET", 
			"GT", "LT", "EQ", "HASH", "COMMA", "COLON", "OB", "CB", "OC", "CC", "PLUS", 
			"DOLLAR", "PC", "AT", "BOOLEAN_FALSE", "BOOLEAN_TRUE", "NULL", "EMPTY_OBJECT", 
			"EMPTY_LIST", "SHEBANG", "NUMBER", "KEY", "IDENT", "IDENT_BACKTICKED", 
			"STRING", "TRIPLE_QUOTED_STRING", "SINGLE_OR_DOUBLE", "R_AND_C_STRING", 
			"HYPER_STRING", "ESC_SEQ", "ESC_SEQ_BASE", "NL", "SINGLE_NL", "WS", "BLOCK_COMMENT", 
			"COMMENT", "LINE_COMMENT", "INLINE_COMMENT", "IDENT_INVALID", "REST"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "YiniParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public YiniParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class YiniContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(YiniParser.EOF, 0); }
		public PrologContext prolog() {
			return getRuleContext(PrologContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public Terminal_stmtContext terminal_stmt() {
			return getRuleContext(Terminal_stmtContext.class,0);
		}
		public YiniContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_yini; }
	}

	public final YiniContext yini() throws RecognitionException {
		YiniContext _localctx = new YiniContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_yini);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(42);
				prolog();
				}
				break;
			}
			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 89337934946822L) != 0)) {
				{
				{
				setState(45);
				stmt();
				}
				}
				setState(50);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TERMINAL_TOKEN) {
				{
				setState(51);
				terminal_stmt();
				}
			}

			setState(54);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrologContext extends ParserRuleContext {
		public TerminalNode SHEBANG() { return getToken(YiniParser.SHEBANG, 0); }
		public List<EolContext> eol() {
			return getRuleContexts(EolContext.class);
		}
		public EolContext eol(int i) {
			return getRuleContext(EolContext.class,i);
		}
		public PrologContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prolog; }
	}

	public final PrologContext prolog() throws RecognitionException {
		PrologContext _localctx = new PrologContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_prolog);
		try {
			int _alt;
			setState(68);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SHEBANG:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				match(SHEBANG);
				setState(60);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(57);
						eol();
						}
						} 
					}
					setState(62);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				}
				}
				break;
			case NL:
			case INLINE_COMMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(64); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(63);
						eol();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(66); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Terminal_stmtContext extends ParserRuleContext {
		public TerminalNode TERMINAL_TOKEN() { return getToken(YiniParser.TERMINAL_TOKEN, 0); }
		public EolContext eol() {
			return getRuleContext(EolContext.class,0);
		}
		public TerminalNode INLINE_COMMENT() { return getToken(YiniParser.INLINE_COMMENT, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public Terminal_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_terminal_stmt; }
	}

	public final Terminal_stmtContext terminal_stmt() throws RecognitionException {
		Terminal_stmtContext _localctx = new Terminal_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_terminal_stmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(TERMINAL_TOKEN);
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(71);
				eol();
				}
				break;
			case 2:
				{
				setState(72);
				match(INLINE_COMMENT);
				}
				break;
			}
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NL) {
				{
				{
				setState(75);
				match(NL);
				}
				}
				setState(80);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public EolContext eol() {
			return getRuleContext(EolContext.class,0);
		}
		public TerminalNode SECTION_HEAD() { return getToken(YiniParser.SECTION_HEAD, 0); }
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public Colon_list_declContext colon_list_decl() {
			return getRuleContext(Colon_list_declContext.class,0);
		}
		public Marker_stmtContext marker_stmt() {
			return getRuleContext(Marker_stmtContext.class,0);
		}
		public Bad_memberContext bad_member() {
			return getRuleContext(Bad_memberContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_stmt);
		try {
			setState(87);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				eol();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				match(SECTION_HEAD);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				assignment();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(84);
				colon_list_decl();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(85);
				marker_stmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(86);
				bad_member();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Marker_stmtContext extends ParserRuleContext {
		public TerminalNode YINI_MARKER() { return getToken(YiniParser.YINI_MARKER, 0); }
		public EolContext eol() {
			return getRuleContext(EolContext.class,0);
		}
		public Marker_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_marker_stmt; }
	}

	public final Marker_stmtContext marker_stmt() throws RecognitionException {
		Marker_stmtContext _localctx = new Marker_stmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_marker_stmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(YINI_MARKER);
			setState(90);
			eol();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EolContext extends ParserRuleContext {
		public TerminalNode INLINE_COMMENT() { return getToken(YiniParser.INLINE_COMMENT, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public EolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eol; }
	}

	public final EolContext eol() throws RecognitionException {
		EolContext _localctx = new EolContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_eol);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INLINE_COMMENT) {
				{
				setState(92);
				match(INLINE_COMMENT);
				}
			}

			setState(96); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(95);
					match(NL);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(98); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public MemberContext member() {
			return getRuleContext(MemberContext.class,0);
		}
		public EolContext eol() {
			return getRuleContext(EolContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			member();
			setState(101);
			eol();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MemberContext extends ParserRuleContext {
		public TerminalNode KEY() { return getToken(YiniParser.KEY, 0); }
		public TerminalNode EQ() { return getToken(YiniParser.EQ, 0); }
		public List<TerminalNode> WS() { return getTokens(YiniParser.WS); }
		public TerminalNode WS(int i) {
			return getToken(YiniParser.WS, i);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public MemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_member; }
	}

	public final MemberContext member() throws RecognitionException {
		MemberContext _localctx = new MemberContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_member);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			match(KEY);
			setState(105);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WS) {
				{
				setState(104);
				match(WS);
				}
			}

			setState(107);
			match(EQ);
			setState(109);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WS) {
				{
				setState(108);
				match(WS);
				}
			}

			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2346754048L) != 0)) {
				{
				setState(111);
				value();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Colon_list_declContext extends ParserRuleContext {
		public TerminalNode KEY() { return getToken(YiniParser.KEY, 0); }
		public TerminalNode COLON() { return getToken(YiniParser.COLON, 0); }
		public ElementsContext elements() {
			return getRuleContext(ElementsContext.class,0);
		}
		public List<EolContext> eol() {
			return getRuleContexts(EolContext.class);
		}
		public EolContext eol(int i) {
			return getRuleContext(EolContext.class,i);
		}
		public List<TerminalNode> WS() { return getTokens(YiniParser.WS); }
		public TerminalNode WS(int i) {
			return getToken(YiniParser.WS, i);
		}
		public Colon_list_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_colon_list_decl; }
	}

	public final Colon_list_declContext colon_list_decl() throws RecognitionException {
		Colon_list_declContext _localctx = new Colon_list_declContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_colon_list_decl);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(KEY);
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WS) {
				{
				setState(115);
				match(WS);
				}
			}

			setState(118);
			match(COLON);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 18966575579136L) != 0)) {
				{
				setState(125);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case NL:
				case INLINE_COMMENT:
					{
					setState(119);
					eol();
					}
					break;
				case WS:
					{
					setState(121); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(120);
							match(WS);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(123); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			elements();
			setState(139);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					setState(137);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case NL:
					case INLINE_COMMENT:
						{
						setState(131);
						eol();
						}
						break;
					case WS:
						{
						setState(133); 
						_errHandler.sync(this);
						_alt = 1;
						do {
							switch (_alt) {
							case 1:
								{
								{
								setState(132);
								match(WS);
								}
								}
								break;
							default:
								throw new NoViableAltException(this);
							}
							setState(135); 
							_errHandler.sync(this);
							_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
						} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					} 
				}
				setState(141);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			}
			setState(142);
			eol();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ParserRuleContext {
		public Null_literalContext null_literal() {
			return getRuleContext(Null_literalContext.class,0);
		}
		public String_literalContext string_literal() {
			return getRuleContext(String_literalContext.class,0);
		}
		public Number_literalContext number_literal() {
			return getRuleContext(Number_literalContext.class,0);
		}
		public Boolean_literalContext boolean_literal() {
			return getRuleContext(Boolean_literalContext.class,0);
		}
		public List_literalContext list_literal() {
			return getRuleContext(List_literalContext.class,0);
		}
		public Object_literalContext object_literal() {
			return getRuleContext(Object_literalContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_value);
		try {
			setState(150);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NULL:
				enterOuterAlt(_localctx, 1);
				{
				setState(144);
				null_literal();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(145);
				string_literal();
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(146);
				number_literal();
				}
				break;
			case BOOLEAN_FALSE:
			case BOOLEAN_TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(147);
				boolean_literal();
				}
				break;
			case OB:
			case EMPTY_LIST:
				enterOuterAlt(_localctx, 5);
				{
				setState(148);
				list_literal();
				}
				break;
			case OC:
			case EMPTY_OBJECT:
				enterOuterAlt(_localctx, 6);
				{
				setState(149);
				object_literal();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Object_literalContext extends ParserRuleContext {
		public TerminalNode OC() { return getToken(YiniParser.OC, 0); }
		public TerminalNode CC() { return getToken(YiniParser.CC, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public Object_membersContext object_members() {
			return getRuleContext(Object_membersContext.class,0);
		}
		public TerminalNode EMPTY_OBJECT() { return getToken(YiniParser.EMPTY_OBJECT, 0); }
		public Object_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_literal; }
	}

	public final Object_literalContext object_literal() throws RecognitionException {
		Object_literalContext _localctx = new Object_literalContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_object_literal);
		int _la;
		try {
			int _alt;
			setState(182);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OC:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				match(OC);
				setState(156);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(153);
						match(NL);
						}
						} 
					}
					setState(158);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
				}
				setState(160);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==KEY) {
					{
					setState(159);
					object_members();
					}
				}

				setState(165);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NL) {
					{
					{
					setState(162);
					match(NL);
					}
					}
					setState(167);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(168);
				match(CC);
				setState(172);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(169);
						match(NL);
						}
						} 
					}
					setState(174);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
				}
				}
				break;
			case EMPTY_OBJECT:
				enterOuterAlt(_localctx, 2);
				{
				setState(175);
				match(EMPTY_OBJECT);
				setState(179);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(176);
						match(NL);
						}
						} 
					}
					setState(181);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Object_membersContext extends ParserRuleContext {
		public List<Object_memberContext> object_member() {
			return getRuleContexts(Object_memberContext.class);
		}
		public Object_memberContext object_member(int i) {
			return getRuleContext(Object_memberContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(YiniParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(YiniParser.COMMA, i);
		}
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public Object_membersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_members; }
	}

	public final Object_membersContext object_members() throws RecognitionException {
		Object_membersContext _localctx = new Object_membersContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_object_members);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			object_member();
			setState(195);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(185);
					match(COMMA);
					setState(189);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NL) {
						{
						{
						setState(186);
						match(NL);
						}
						}
						setState(191);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(192);
					object_member();
					}
					} 
				}
				setState(197);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			}
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(198);
				match(COMMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Object_memberContext extends ParserRuleContext {
		public TerminalNode KEY() { return getToken(YiniParser.KEY, 0); }
		public TerminalNode COLON() { return getToken(YiniParser.COLON, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public TerminalNode WS() { return getToken(YiniParser.WS, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public Object_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object_member; }
	}

	public final Object_memberContext object_member() throws RecognitionException {
		Object_memberContext _localctx = new Object_memberContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_object_member);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(KEY);
			setState(203);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WS) {
				{
				setState(202);
				match(WS);
				}
			}

			setState(205);
			match(COLON);
			setState(209);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NL) {
				{
				{
				setState(206);
				match(NL);
				}
				}
				setState(211);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(212);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_literalContext extends ParserRuleContext {
		public TerminalNode OB() { return getToken(YiniParser.OB, 0); }
		public TerminalNode CB() { return getToken(YiniParser.CB, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public ElementsContext elements() {
			return getRuleContext(ElementsContext.class,0);
		}
		public TerminalNode EMPTY_LIST() { return getToken(YiniParser.EMPTY_LIST, 0); }
		public List_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_literal; }
	}

	public final List_literalContext list_literal() throws RecognitionException {
		List_literalContext _localctx = new List_literalContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_list_literal);
		int _la;
		try {
			int _alt;
			setState(244);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OB:
				enterOuterAlt(_localctx, 1);
				{
				setState(214);
				match(OB);
				setState(218);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(215);
						match(NL);
						}
						} 
					}
					setState(220);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
				}
				setState(222);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2346754048L) != 0)) {
					{
					setState(221);
					elements();
					}
				}

				setState(227);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NL) {
					{
					{
					setState(224);
					match(NL);
					}
					}
					setState(229);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(230);
				match(CB);
				setState(234);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(231);
						match(NL);
						}
						} 
					}
					setState(236);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
				}
				}
				break;
			case EMPTY_LIST:
				enterOuterAlt(_localctx, 2);
				{
				setState(237);
				match(EMPTY_LIST);
				setState(241);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(238);
						match(NL);
						}
						} 
					}
					setState(243);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElementsContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(YiniParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(YiniParser.COMMA, i);
		}
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public ElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elements; }
	}

	public final ElementsContext elements() throws RecognitionException {
		ElementsContext _localctx = new ElementsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_elements);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			value();
			setState(263);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(250);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NL) {
						{
						{
						setState(247);
						match(NL);
						}
						}
						setState(252);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(253);
					match(COMMA);
					setState(257);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NL) {
						{
						{
						setState(254);
						match(NL);
						}
						}
						setState(259);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(260);
					value();
					}
					} 
				}
				setState(265);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			}
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(266);
				match(COMMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Number_literalContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(YiniParser.NUMBER, 0); }
		public Number_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number_literal; }
	}

	public final Number_literalContext number_literal() throws RecognitionException {
		Number_literalContext _localctx = new Number_literalContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_number_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			match(NUMBER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Null_literalContext extends ParserRuleContext {
		public TerminalNode NULL() { return getToken(YiniParser.NULL, 0); }
		public Null_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_null_literal; }
	}

	public final Null_literalContext null_literal() throws RecognitionException {
		Null_literalContext _localctx = new Null_literalContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_null_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271);
			match(NULL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class String_literalContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(YiniParser.STRING, 0); }
		public List<String_concatContext> string_concat() {
			return getRuleContexts(String_concatContext.class);
		}
		public String_concatContext string_concat(int i) {
			return getRuleContext(String_concatContext.class,i);
		}
		public String_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_literal; }
	}

	public final String_literalContext string_literal() throws RecognitionException {
		String_literalContext _localctx = new String_literalContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_string_literal);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(273);
			match(STRING);
			setState(277);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(274);
					string_concat();
					}
					} 
				}
				setState(279);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class String_concatContext extends ParserRuleContext {
		public TerminalNode PLUS() { return getToken(YiniParser.PLUS, 0); }
		public TerminalNode STRING() { return getToken(YiniParser.STRING, 0); }
		public List<TerminalNode> NL() { return getTokens(YiniParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(YiniParser.NL, i);
		}
		public String_concatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_concat; }
	}

	public final String_concatContext string_concat() throws RecognitionException {
		String_concatContext _localctx = new String_concatContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_string_concat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NL) {
				{
				{
				setState(280);
				match(NL);
				}
				}
				setState(285);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(286);
			match(PLUS);
			setState(290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NL) {
				{
				{
				setState(287);
				match(NL);
				}
				}
				setState(292);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(293);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Boolean_literalContext extends ParserRuleContext {
		public TerminalNode BOOLEAN_TRUE() { return getToken(YiniParser.BOOLEAN_TRUE, 0); }
		public TerminalNode BOOLEAN_FALSE() { return getToken(YiniParser.BOOLEAN_FALSE, 0); }
		public Boolean_literalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolean_literal; }
	}

	public final Boolean_literalContext boolean_literal() throws RecognitionException {
		Boolean_literalContext _localctx = new Boolean_literalContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_boolean_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			_la = _input.LA(1);
			if ( !(_la==BOOLEAN_FALSE || _la==BOOLEAN_TRUE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Bad_memberContext extends ParserRuleContext {
		public TerminalNode EQ() { return getToken(YiniParser.EQ, 0); }
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> REST() { return getTokens(YiniParser.REST); }
		public TerminalNode REST(int i) {
			return getToken(YiniParser.REST, i);
		}
		public List<TerminalNode> WS() { return getTokens(YiniParser.WS); }
		public TerminalNode WS(int i) {
			return getToken(YiniParser.WS, i);
		}
		public EolContext eol() {
			return getRuleContext(EolContext.class,0);
		}
		public Bad_memberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bad_member; }
	}

	public final Bad_memberContext bad_member() throws RecognitionException {
		Bad_memberContext _localctx = new Bad_memberContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_bad_member);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,46,_ctx) ) {
			case 1:
				{
				setState(297);
				match(WS);
				}
				break;
			}
			setState(302);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case REST:
				{
				setState(300);
				match(REST);
				}
				break;
			case OB:
			case OC:
			case BOOLEAN_FALSE:
			case BOOLEAN_TRUE:
			case NULL:
			case EMPTY_OBJECT:
			case EMPTY_LIST:
			case NUMBER:
			case STRING:
				{
				setState(301);
				value();
				}
				break;
			case EQ:
			case WS:
				break;
			default:
				break;
			}
			setState(305);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WS) {
				{
				setState(304);
				match(WS);
				}
			}

			setState(307);
			match(EQ);
			setState(310);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OB:
			case OC:
			case BOOLEAN_FALSE:
			case BOOLEAN_TRUE:
			case NULL:
			case EMPTY_OBJECT:
			case EMPTY_LIST:
			case NUMBER:
			case STRING:
				{
				setState(308);
				value();
				}
				break;
			case REST:
				{
				setState(309);
				match(REST);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(313);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
			case 1:
				{
				setState(312);
				eol();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001.\u013c\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001\u0000\u0003\u0000"+
		",\b\u0000\u0001\u0000\u0005\u0000/\b\u0000\n\u0000\f\u00002\t\u0000\u0001"+
		"\u0000\u0003\u00005\b\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0005\u0001;\b\u0001\n\u0001\f\u0001>\t\u0001\u0001\u0001\u0004"+
		"\u0001A\b\u0001\u000b\u0001\f\u0001B\u0003\u0001E\b\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002J\b\u0002\u0001\u0002\u0005\u0002"+
		"M\b\u0002\n\u0002\f\u0002P\t\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003X\b\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0003\u0005^\b\u0005\u0001\u0005"+
		"\u0004\u0005a\b\u0005\u000b\u0005\f\u0005b\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0007\u0001\u0007\u0003\u0007j\b\u0007\u0001\u0007\u0001"+
		"\u0007\u0003\u0007n\b\u0007\u0001\u0007\u0003\u0007q\b\u0007\u0001\b\u0001"+
		"\b\u0003\bu\b\b\u0001\b\u0001\b\u0001\b\u0004\bz\b\b\u000b\b\f\b{\u0005"+
		"\b~\b\b\n\b\f\b\u0081\t\b\u0001\b\u0001\b\u0001\b\u0004\b\u0086\b\b\u000b"+
		"\b\f\b\u0087\u0005\b\u008a\b\b\n\b\f\b\u008d\t\b\u0001\b\u0001\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u0097\b\t\u0001\n\u0001"+
		"\n\u0005\n\u009b\b\n\n\n\f\n\u009e\t\n\u0001\n\u0003\n\u00a1\b\n\u0001"+
		"\n\u0005\n\u00a4\b\n\n\n\f\n\u00a7\t\n\u0001\n\u0001\n\u0005\n\u00ab\b"+
		"\n\n\n\f\n\u00ae\t\n\u0001\n\u0001\n\u0005\n\u00b2\b\n\n\n\f\n\u00b5\t"+
		"\n\u0003\n\u00b7\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00bc"+
		"\b\u000b\n\u000b\f\u000b\u00bf\t\u000b\u0001\u000b\u0005\u000b\u00c2\b"+
		"\u000b\n\u000b\f\u000b\u00c5\t\u000b\u0001\u000b\u0003\u000b\u00c8\b\u000b"+
		"\u0001\f\u0001\f\u0003\f\u00cc\b\f\u0001\f\u0001\f\u0005\f\u00d0\b\f\n"+
		"\f\f\f\u00d3\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0005\r\u00d9\b\r\n\r"+
		"\f\r\u00dc\t\r\u0001\r\u0003\r\u00df\b\r\u0001\r\u0005\r\u00e2\b\r\n\r"+
		"\f\r\u00e5\t\r\u0001\r\u0001\r\u0005\r\u00e9\b\r\n\r\f\r\u00ec\t\r\u0001"+
		"\r\u0001\r\u0005\r\u00f0\b\r\n\r\f\r\u00f3\t\r\u0003\r\u00f5\b\r\u0001"+
		"\u000e\u0001\u000e\u0005\u000e\u00f9\b\u000e\n\u000e\f\u000e\u00fc\t\u000e"+
		"\u0001\u000e\u0001\u000e\u0005\u000e\u0100\b\u000e\n\u000e\f\u000e\u0103"+
		"\t\u000e\u0001\u000e\u0005\u000e\u0106\b\u000e\n\u000e\f\u000e\u0109\t"+
		"\u000e\u0001\u000e\u0003\u000e\u010c\b\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0005\u0011\u0114\b\u0011\n"+
		"\u0011\f\u0011\u0117\t\u0011\u0001\u0012\u0005\u0012\u011a\b\u0012\n\u0012"+
		"\f\u0012\u011d\t\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u0121\b\u0012"+
		"\n\u0012\f\u0012\u0124\t\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0014\u0003\u0014\u012b\b\u0014\u0001\u0014\u0001\u0014\u0003"+
		"\u0014\u012f\b\u0014\u0001\u0014\u0003\u0014\u0132\b\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0003\u0014\u0137\b\u0014\u0001\u0014\u0003\u0014"+
		"\u013a\b\u0014\u0001\u0014\u0000\u0000\u0015\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(\u0000"+
		"\u0001\u0001\u0000\u0015\u0016\u0163\u0000+\u0001\u0000\u0000\u0000\u0002"+
		"D\u0001\u0000\u0000\u0000\u0004F\u0001\u0000\u0000\u0000\u0006W\u0001"+
		"\u0000\u0000\u0000\bY\u0001\u0000\u0000\u0000\n]\u0001\u0000\u0000\u0000"+
		"\fd\u0001\u0000\u0000\u0000\u000eg\u0001\u0000\u0000\u0000\u0010r\u0001"+
		"\u0000\u0000\u0000\u0012\u0096\u0001\u0000\u0000\u0000\u0014\u00b6\u0001"+
		"\u0000\u0000\u0000\u0016\u00b8\u0001\u0000\u0000\u0000\u0018\u00c9\u0001"+
		"\u0000\u0000\u0000\u001a\u00f4\u0001\u0000\u0000\u0000\u001c\u00f6\u0001"+
		"\u0000\u0000\u0000\u001e\u010d\u0001\u0000\u0000\u0000 \u010f\u0001\u0000"+
		"\u0000\u0000\"\u0111\u0001\u0000\u0000\u0000$\u011b\u0001\u0000\u0000"+
		"\u0000&\u0127\u0001\u0000\u0000\u0000(\u012a\u0001\u0000\u0000\u0000*"+
		",\u0003\u0002\u0001\u0000+*\u0001\u0000\u0000\u0000+,\u0001\u0000\u0000"+
		"\u0000,0\u0001\u0000\u0000\u0000-/\u0003\u0006\u0003\u0000.-\u0001\u0000"+
		"\u0000\u0000/2\u0001\u0000\u0000\u00000.\u0001\u0000\u0000\u000001\u0001"+
		"\u0000\u0000\u000014\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u0000"+
		"35\u0003\u0004\u0002\u000043\u0001\u0000\u0000\u000045\u0001\u0000\u0000"+
		"\u000056\u0001\u0000\u0000\u000067\u0005\u0000\u0000\u00017\u0001\u0001"+
		"\u0000\u0000\u00008<\u0005\u001a\u0000\u00009;\u0003\n\u0005\u0000:9\u0001"+
		"\u0000\u0000\u0000;>\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000"+
		"<=\u0001\u0000\u0000\u0000=E\u0001\u0000\u0000\u0000><\u0001\u0000\u0000"+
		"\u0000?A\u0003\n\u0005\u0000@?\u0001\u0000\u0000\u0000AB\u0001\u0000\u0000"+
		"\u0000B@\u0001\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000CE\u0001\u0000"+
		"\u0000\u0000D8\u0001\u0000\u0000\u0000D@\u0001\u0000\u0000\u0000E\u0003"+
		"\u0001\u0000\u0000\u0000FI\u0005\u0003\u0000\u0000GJ\u0003\n\u0005\u0000"+
		"HJ\u0005,\u0000\u0000IG\u0001\u0000\u0000\u0000IH\u0001\u0000\u0000\u0000"+
		"IJ\u0001\u0000\u0000\u0000JN\u0001\u0000\u0000\u0000KM\u0005&\u0000\u0000"+
		"LK\u0001\u0000\u0000\u0000MP\u0001\u0000\u0000\u0000NL\u0001\u0000\u0000"+
		"\u0000NO\u0001\u0000\u0000\u0000O\u0005\u0001\u0000\u0000\u0000PN\u0001"+
		"\u0000\u0000\u0000QX\u0003\n\u0005\u0000RX\u0005\u0002\u0000\u0000SX\u0003"+
		"\f\u0006\u0000TX\u0003\u0010\b\u0000UX\u0003\b\u0004\u0000VX\u0003(\u0014"+
		"\u0000WQ\u0001\u0000\u0000\u0000WR\u0001\u0000\u0000\u0000WS\u0001\u0000"+
		"\u0000\u0000WT\u0001\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000WV\u0001"+
		"\u0000\u0000\u0000X\u0007\u0001\u0000\u0000\u0000YZ\u0005\u0001\u0000"+
		"\u0000Z[\u0003\n\u0005\u0000[\t\u0001\u0000\u0000\u0000\\^\u0005,\u0000"+
		"\u0000]\\\u0001\u0000\u0000\u0000]^\u0001\u0000\u0000\u0000^`\u0001\u0000"+
		"\u0000\u0000_a\u0005&\u0000\u0000`_\u0001\u0000\u0000\u0000ab\u0001\u0000"+
		"\u0000\u0000b`\u0001\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000c\u000b"+
		"\u0001\u0000\u0000\u0000de\u0003\u000e\u0007\u0000ef\u0003\n\u0005\u0000"+
		"f\r\u0001\u0000\u0000\u0000gi\u0005\u001c\u0000\u0000hj\u0005(\u0000\u0000"+
		"ih\u0001\u0000\u0000\u0000ij\u0001\u0000\u0000\u0000jk\u0001\u0000\u0000"+
		"\u0000km\u0005\t\u0000\u0000ln\u0005(\u0000\u0000ml\u0001\u0000\u0000"+
		"\u0000mn\u0001\u0000\u0000\u0000np\u0001\u0000\u0000\u0000oq\u0003\u0012"+
		"\t\u0000po\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000q\u000f\u0001"+
		"\u0000\u0000\u0000rt\u0005\u001c\u0000\u0000su\u0005(\u0000\u0000ts\u0001"+
		"\u0000\u0000\u0000tu\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000"+
		"v\u007f\u0005\f\u0000\u0000w~\u0003\n\u0005\u0000xz\u0005(\u0000\u0000"+
		"yx\u0001\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000{y\u0001\u0000\u0000"+
		"\u0000{|\u0001\u0000\u0000\u0000|~\u0001\u0000\u0000\u0000}w\u0001\u0000"+
		"\u0000\u0000}y\u0001\u0000\u0000\u0000~\u0081\u0001\u0000\u0000\u0000"+
		"\u007f}\u0001\u0000\u0000\u0000\u007f\u0080\u0001\u0000\u0000\u0000\u0080"+
		"\u0082\u0001\u0000\u0000\u0000\u0081\u007f\u0001\u0000\u0000\u0000\u0082"+
		"\u008b\u0003\u001c\u000e\u0000\u0083\u008a\u0003\n\u0005\u0000\u0084\u0086"+
		"\u0005(\u0000\u0000\u0085\u0084\u0001\u0000\u0000\u0000\u0086\u0087\u0001"+
		"\u0000\u0000\u0000\u0087\u0085\u0001\u0000\u0000\u0000\u0087\u0088\u0001"+
		"\u0000\u0000\u0000\u0088\u008a\u0001\u0000\u0000\u0000\u0089\u0083\u0001"+
		"\u0000\u0000\u0000\u0089\u0085\u0001\u0000\u0000\u0000\u008a\u008d\u0001"+
		"\u0000\u0000\u0000\u008b\u0089\u0001\u0000\u0000\u0000\u008b\u008c\u0001"+
		"\u0000\u0000\u0000\u008c\u008e\u0001\u0000\u0000\u0000\u008d\u008b\u0001"+
		"\u0000\u0000\u0000\u008e\u008f\u0003\n\u0005\u0000\u008f\u0011\u0001\u0000"+
		"\u0000\u0000\u0090\u0097\u0003 \u0010\u0000\u0091\u0097\u0003\"\u0011"+
		"\u0000\u0092\u0097\u0003\u001e\u000f\u0000\u0093\u0097\u0003&\u0013\u0000"+
		"\u0094\u0097\u0003\u001a\r\u0000\u0095\u0097\u0003\u0014\n\u0000\u0096"+
		"\u0090\u0001\u0000\u0000\u0000\u0096\u0091\u0001\u0000\u0000\u0000\u0096"+
		"\u0092\u0001\u0000\u0000\u0000\u0096\u0093\u0001\u0000\u0000\u0000\u0096"+
		"\u0094\u0001\u0000\u0000\u0000\u0096\u0095\u0001\u0000\u0000\u0000\u0097"+
		"\u0013\u0001\u0000\u0000\u0000\u0098\u009c\u0005\u000f\u0000\u0000\u0099"+
		"\u009b\u0005&\u0000\u0000\u009a\u0099\u0001\u0000\u0000\u0000\u009b\u009e"+
		"\u0001\u0000\u0000\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c\u009d"+
		"\u0001\u0000\u0000\u0000\u009d\u00a0\u0001\u0000\u0000\u0000\u009e\u009c"+
		"\u0001\u0000\u0000\u0000\u009f\u00a1\u0003\u0016\u000b\u0000\u00a0\u009f"+
		"\u0001\u0000\u0000\u0000\u00a0\u00a1\u0001\u0000\u0000\u0000\u00a1\u00a5"+
		"\u0001\u0000\u0000\u0000\u00a2\u00a4\u0005&\u0000\u0000\u00a3\u00a2\u0001"+
		"\u0000\u0000\u0000\u00a4\u00a7\u0001\u0000\u0000\u0000\u00a5\u00a3\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a6\u0001\u0000\u0000\u0000\u00a6\u00a8\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a5\u0001\u0000\u0000\u0000\u00a8\u00ac\u0005"+
		"\u0010\u0000\u0000\u00a9\u00ab\u0005&\u0000\u0000\u00aa\u00a9\u0001\u0000"+
		"\u0000\u0000\u00ab\u00ae\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000"+
		"\u0000\u0000\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad\u00b7\u0001\u0000"+
		"\u0000\u0000\u00ae\u00ac\u0001\u0000\u0000\u0000\u00af\u00b3\u0005\u0018"+
		"\u0000\u0000\u00b0\u00b2\u0005&\u0000\u0000\u00b1\u00b0\u0001\u0000\u0000"+
		"\u0000\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001\u0000\u0000"+
		"\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b7\u0001\u0000\u0000"+
		"\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u0098\u0001\u0000\u0000"+
		"\u0000\u00b6\u00af\u0001\u0000\u0000\u0000\u00b7\u0015\u0001\u0000\u0000"+
		"\u0000\u00b8\u00c3\u0003\u0018\f\u0000\u00b9\u00bd\u0005\u000b\u0000\u0000"+
		"\u00ba\u00bc\u0005&\u0000\u0000\u00bb\u00ba\u0001\u0000\u0000\u0000\u00bc"+
		"\u00bf\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00bd"+
		"\u00be\u0001\u0000\u0000\u0000\u00be\u00c0\u0001\u0000\u0000\u0000\u00bf"+
		"\u00bd\u0001\u0000\u0000\u0000\u00c0\u00c2\u0003\u0018\f\u0000\u00c1\u00b9"+
		"\u0001\u0000\u0000\u0000\u00c2\u00c5\u0001\u0000\u0000\u0000\u00c3\u00c1"+
		"\u0001\u0000\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4\u00c7"+
		"\u0001\u0000\u0000\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c6\u00c8"+
		"\u0005\u000b\u0000\u0000\u00c7\u00c6\u0001\u0000\u0000\u0000\u00c7\u00c8"+
		"\u0001\u0000\u0000\u0000\u00c8\u0017\u0001\u0000\u0000\u0000\u00c9\u00cb"+
		"\u0005\u001c\u0000\u0000\u00ca\u00cc\u0005(\u0000\u0000\u00cb\u00ca\u0001"+
		"\u0000\u0000\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc\u00cd\u0001"+
		"\u0000\u0000\u0000\u00cd\u00d1\u0005\f\u0000\u0000\u00ce\u00d0\u0005&"+
		"\u0000\u0000\u00cf\u00ce\u0001\u0000\u0000\u0000\u00d0\u00d3\u0001\u0000"+
		"\u0000\u0000\u00d1\u00cf\u0001\u0000\u0000\u0000\u00d1\u00d2\u0001\u0000"+
		"\u0000\u0000\u00d2\u00d4\u0001\u0000\u0000\u0000\u00d3\u00d1\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0003\u0012\t\u0000\u00d5\u0019\u0001\u0000\u0000"+
		"\u0000\u00d6\u00da\u0005\r\u0000\u0000\u00d7\u00d9\u0005&\u0000\u0000"+
		"\u00d8\u00d7\u0001\u0000\u0000\u0000\u00d9\u00dc\u0001\u0000\u0000\u0000"+
		"\u00da\u00d8\u0001\u0000\u0000\u0000\u00da\u00db\u0001\u0000\u0000\u0000"+
		"\u00db\u00de\u0001\u0000\u0000\u0000\u00dc\u00da\u0001\u0000\u0000\u0000"+
		"\u00dd\u00df\u0003\u001c\u000e\u0000\u00de\u00dd\u0001\u0000\u0000\u0000"+
		"\u00de\u00df\u0001\u0000\u0000\u0000\u00df\u00e3\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e2\u0005&\u0000\u0000\u00e1\u00e0\u0001\u0000\u0000\u0000\u00e2"+
		"\u00e5\u0001\u0000\u0000\u0000\u00e3\u00e1\u0001\u0000\u0000\u0000\u00e3"+
		"\u00e4\u0001\u0000\u0000\u0000\u00e4\u00e6\u0001\u0000\u0000\u0000\u00e5"+
		"\u00e3\u0001\u0000\u0000\u0000\u00e6\u00ea\u0005\u000e\u0000\u0000\u00e7"+
		"\u00e9\u0005&\u0000\u0000\u00e8\u00e7\u0001\u0000\u0000\u0000\u00e9\u00ec"+
		"\u0001\u0000\u0000\u0000\u00ea\u00e8\u0001\u0000\u0000\u0000\u00ea\u00eb"+
		"\u0001\u0000\u0000\u0000\u00eb\u00f5\u0001\u0000\u0000\u0000\u00ec\u00ea"+
		"\u0001\u0000\u0000\u0000\u00ed\u00f1\u0005\u0019\u0000\u0000\u00ee\u00f0"+
		"\u0005&\u0000\u0000\u00ef\u00ee\u0001\u0000\u0000\u0000\u00f0\u00f3\u0001"+
		"\u0000\u0000\u0000\u00f1\u00ef\u0001\u0000\u0000\u0000\u00f1\u00f2\u0001"+
		"\u0000\u0000\u0000\u00f2\u00f5\u0001\u0000\u0000\u0000\u00f3\u00f1\u0001"+
		"\u0000\u0000\u0000\u00f4\u00d6\u0001\u0000\u0000\u0000\u00f4\u00ed\u0001"+
		"\u0000\u0000\u0000\u00f5\u001b\u0001\u0000\u0000\u0000\u00f6\u0107\u0003"+
		"\u0012\t\u0000\u00f7\u00f9\u0005&\u0000\u0000\u00f8\u00f7\u0001\u0000"+
		"\u0000\u0000\u00f9\u00fc\u0001\u0000\u0000\u0000\u00fa\u00f8\u0001\u0000"+
		"\u0000\u0000\u00fa\u00fb\u0001\u0000\u0000\u0000\u00fb\u00fd\u0001\u0000"+
		"\u0000\u0000\u00fc\u00fa\u0001\u0000\u0000\u0000\u00fd\u0101\u0005\u000b"+
		"\u0000\u0000\u00fe\u0100\u0005&\u0000\u0000\u00ff\u00fe\u0001\u0000\u0000"+
		"\u0000\u0100\u0103\u0001\u0000\u0000\u0000\u0101\u00ff\u0001\u0000\u0000"+
		"\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102\u0104\u0001\u0000\u0000"+
		"\u0000\u0103\u0101\u0001\u0000\u0000\u0000\u0104\u0106\u0003\u0012\t\u0000"+
		"\u0105\u00fa\u0001\u0000\u0000\u0000\u0106\u0109\u0001\u0000\u0000\u0000"+
		"\u0107\u0105\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000\u0000"+
		"\u0108\u010b\u0001\u0000\u0000\u0000\u0109\u0107\u0001\u0000\u0000\u0000"+
		"\u010a\u010c\u0005\u000b\u0000\u0000\u010b\u010a\u0001\u0000\u0000\u0000"+
		"\u010b\u010c\u0001\u0000\u0000\u0000\u010c\u001d\u0001\u0000\u0000\u0000"+
		"\u010d\u010e\u0005\u001b\u0000\u0000\u010e\u001f\u0001\u0000\u0000\u0000"+
		"\u010f\u0110\u0005\u0017\u0000\u0000\u0110!\u0001\u0000\u0000\u0000\u0111"+
		"\u0115\u0005\u001f\u0000\u0000\u0112\u0114\u0003$\u0012\u0000\u0113\u0112"+
		"\u0001\u0000\u0000\u0000\u0114\u0117\u0001\u0000\u0000\u0000\u0115\u0113"+
		"\u0001\u0000\u0000\u0000\u0115\u0116\u0001\u0000\u0000\u0000\u0116#\u0001"+
		"\u0000\u0000\u0000\u0117\u0115\u0001\u0000\u0000\u0000\u0118\u011a\u0005"+
		"&\u0000\u0000\u0119\u0118\u0001\u0000\u0000\u0000\u011a\u011d\u0001\u0000"+
		"\u0000\u0000\u011b\u0119\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000"+
		"\u0000\u0000\u011c\u011e\u0001\u0000\u0000\u0000\u011d\u011b\u0001\u0000"+
		"\u0000\u0000\u011e\u0122\u0005\u0011\u0000\u0000\u011f\u0121\u0005&\u0000"+
		"\u0000\u0120\u011f\u0001\u0000\u0000\u0000\u0121\u0124\u0001\u0000\u0000"+
		"\u0000\u0122\u0120\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000\u0000"+
		"\u0000\u0123\u0125\u0001\u0000\u0000\u0000\u0124\u0122\u0001\u0000\u0000"+
		"\u0000\u0125\u0126\u0005\u001f\u0000\u0000\u0126%\u0001\u0000\u0000\u0000"+
		"\u0127\u0128\u0007\u0000\u0000\u0000\u0128\'\u0001\u0000\u0000\u0000\u0129"+
		"\u012b\u0005(\u0000\u0000\u012a\u0129\u0001\u0000\u0000\u0000\u012a\u012b"+
		"\u0001\u0000\u0000\u0000\u012b\u012e\u0001\u0000\u0000\u0000\u012c\u012f"+
		"\u0005.\u0000\u0000\u012d\u012f\u0003\u0012\t\u0000\u012e\u012c\u0001"+
		"\u0000\u0000\u0000\u012e\u012d\u0001\u0000\u0000\u0000\u012e\u012f\u0001"+
		"\u0000\u0000\u0000\u012f\u0131\u0001\u0000\u0000\u0000\u0130\u0132\u0005"+
		"(\u0000\u0000\u0131\u0130\u0001\u0000\u0000\u0000\u0131\u0132\u0001\u0000"+
		"\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0136\u0005\t\u0000"+
		"\u0000\u0134\u0137\u0003\u0012\t\u0000\u0135\u0137\u0005.\u0000\u0000"+
		"\u0136\u0134\u0001\u0000\u0000\u0000\u0136\u0135\u0001\u0000\u0000\u0000"+
		"\u0137\u0139\u0001\u0000\u0000\u0000\u0138\u013a\u0003\n\u0005\u0000\u0139"+
		"\u0138\u0001\u0000\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000\u013a"+
		")\u0001\u0000\u0000\u00003+04<BDINW]bimpt{}\u007f\u0087\u0089\u008b\u0096"+
		"\u009c\u00a0\u00a5\u00ac\u00b3\u00b6\u00bd\u00c3\u00c7\u00cb\u00d1\u00da"+
		"\u00de\u00e3\u00ea\u00f1\u00f4\u00fa\u0101\u0107\u010b\u0115\u011b\u0122"+
		"\u012a\u012e\u0131\u0136\u0139";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}