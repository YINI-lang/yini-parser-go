// Code generated from grammar/v1.0.0-rc.7/YiniParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // YiniParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type YiniParser struct {
	*antlr.BaseParser
}

var YiniParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yiniparserParserInit() {
	staticData := &YiniParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "'{}'", "'[]'", "", "",
		"'^'", "'\\u00A7'", "'>'", "'<'", "'='", "','", "':'", "'['", "']'",
		"'{'", "'}'", "'+'", "'-'", "'$'", "'%'", "'@'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "SHEBANG", "YINI_TOKEN", "INCLUDE_TOKEN", "DEPRECATED_TOKEN", "TERMINAL_TOKEN",
		"SECTION_HEAD", "INVALID_SECTION_HEAD", "BOOLEAN_FALSE", "BOOLEAN_TRUE",
		"NULL", "EMPTY_OBJECT", "EMPTY_LIST", "STRING", "NUMBER", "CARET", "SS",
		"GT", "LT", "EQ", "COMMA", "COLON", "OB", "CB", "OC", "CC", "PLUS",
		"MINUS", "DOLLAR", "PC", "AT", "SEMICOLON", "NL", "WS", "BLOCK_COMMENT",
		"DISABLED_LINE", "FULL_LINE_COMMENT", "INLINE_COMMENT", "KEY", "IDENT_INVALID",
		"REST", "META_INVALID",
	}
	staticData.RuleNames = []string{
		"yini", "prolog", "terminal_stmt", "terminal_trivia", "stmt", "full_line_comment_stmt",
		"disabled_line_stmt", "invalid_section_stmt", "meta_stmt", "directive",
		"yini_directive", "yini_mode_declaration", "annotation", "eol", "line_end",
		"assignment", "identifier", "member", "value", "scalar_value", "concat_expression",
		"concat_tail", "concat_operand", "string_literal", "object_literal",
		"object_members", "object_member", "object_member_separator", "list_literal",
		"elements", "multiline_trivia", "number_literal", "null_literal", "boolean_literal",
		"bad_meta_text", "bad_member",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 340, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 3, 0,
		74, 8, 0, 1, 0, 5, 0, 77, 8, 0, 10, 0, 12, 0, 80, 9, 0, 1, 0, 3, 0, 83,
		8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1,
		1, 4, 1, 95, 8, 1, 11, 1, 12, 1, 96, 3, 1, 99, 8, 1, 1, 2, 1, 2, 5, 2,
		103, 8, 2, 10, 2, 12, 2, 106, 9, 2, 1, 3, 1, 3, 3, 3, 110, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 120, 8, 4, 1, 5, 1, 5,
		3, 5, 124, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 135, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 140, 8, 9, 1, 9, 3, 9, 143, 8,
		9, 1, 10, 1, 10, 3, 10, 147, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 4, 13, 157, 8, 13, 11, 13, 12, 13, 158, 1, 14, 1,
		14, 3, 14, 163, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 3, 17, 173, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 179, 8, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 185, 8, 19, 1, 20, 1, 20, 4, 20, 189,
		8, 20, 11, 20, 12, 20, 190, 1, 21, 1, 21, 5, 21, 195, 8, 21, 10, 21, 12,
		21, 198, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		207, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 213, 8, 24, 10, 24, 12,
		24, 216, 9, 24, 1, 24, 3, 24, 219, 8, 24, 1, 24, 5, 24, 222, 8, 24, 10,
		24, 12, 24, 225, 9, 24, 1, 24, 1, 24, 3, 24, 229, 8, 24, 1, 25, 1, 25,
		5, 25, 233, 8, 25, 10, 25, 12, 25, 236, 9, 25, 1, 25, 1, 25, 5, 25, 240,
		8, 25, 10, 25, 12, 25, 243, 9, 25, 1, 25, 5, 25, 246, 8, 25, 10, 25, 12,
		25, 249, 9, 25, 1, 25, 5, 25, 252, 8, 25, 10, 25, 12, 25, 255, 9, 25, 1,
		25, 3, 25, 258, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28,
		1, 28, 5, 28, 268, 8, 28, 10, 28, 12, 28, 271, 9, 28, 1, 28, 3, 28, 274,
		8, 28, 1, 28, 5, 28, 277, 8, 28, 10, 28, 12, 28, 280, 9, 28, 1, 28, 1,
		28, 3, 28, 284, 8, 28, 1, 29, 1, 29, 5, 29, 288, 8, 29, 10, 29, 12, 29,
		291, 9, 29, 1, 29, 1, 29, 5, 29, 295, 8, 29, 10, 29, 12, 29, 298, 9, 29,
		1, 29, 5, 29, 301, 8, 29, 10, 29, 12, 29, 304, 9, 29, 1, 29, 5, 29, 307,
		8, 29, 10, 29, 12, 29, 310, 9, 29, 1, 29, 3, 29, 313, 8, 29, 1, 30, 1,
		30, 1, 31, 3, 31, 318, 8, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 330, 8, 35, 1, 35, 1, 35, 1, 35, 3,
		35, 335, 8, 35, 1, 35, 3, 35, 338, 8, 35, 1, 35, 0, 0, 36, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 5, 2, 0, 8,
		10, 38, 38, 2, 0, 19, 19, 21, 21, 2, 0, 32, 32, 35, 36, 1, 0, 26, 27, 1,
		0, 8, 9, 362, 0, 73, 1, 0, 0, 0, 2, 98, 1, 0, 0, 0, 4, 100, 1, 0, 0, 0,
		6, 109, 1, 0, 0, 0, 8, 119, 1, 0, 0, 0, 10, 121, 1, 0, 0, 0, 12, 125, 1,
		0, 0, 0, 14, 127, 1, 0, 0, 0, 16, 134, 1, 0, 0, 0, 18, 142, 1, 0, 0, 0,
		20, 144, 1, 0, 0, 0, 22, 150, 1, 0, 0, 0, 24, 152, 1, 0, 0, 0, 26, 156,
		1, 0, 0, 0, 28, 162, 1, 0, 0, 0, 30, 164, 1, 0, 0, 0, 32, 167, 1, 0, 0,
		0, 34, 169, 1, 0, 0, 0, 36, 178, 1, 0, 0, 0, 38, 184, 1, 0, 0, 0, 40, 186,
		1, 0, 0, 0, 42, 192, 1, 0, 0, 0, 44, 206, 1, 0, 0, 0, 46, 208, 1, 0, 0,
		0, 48, 228, 1, 0, 0, 0, 50, 230, 1, 0, 0, 0, 52, 259, 1, 0, 0, 0, 54, 263,
		1, 0, 0, 0, 56, 283, 1, 0, 0, 0, 58, 285, 1, 0, 0, 0, 60, 314, 1, 0, 0,
		0, 62, 317, 1, 0, 0, 0, 64, 321, 1, 0, 0, 0, 66, 323, 1, 0, 0, 0, 68, 325,
		1, 0, 0, 0, 70, 329, 1, 0, 0, 0, 72, 74, 3, 2, 1, 0, 73, 72, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 78, 1, 0, 0, 0, 75, 77, 3, 8, 4, 0, 76, 75, 1,
		0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79,
		82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 83, 3, 4, 2, 0, 82, 81, 1, 0, 0,
		0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 5, 0, 0, 1, 85, 1, 1,
		0, 0, 0, 86, 90, 5, 1, 0, 0, 87, 89, 3, 26, 13, 0, 88, 87, 1, 0, 0, 0,
		89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 99, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 3, 26, 13, 0, 94, 93, 1, 0, 0, 0,
		95, 96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1,
		0, 0, 0, 98, 86, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 99, 3, 1, 0, 0, 0, 100,
		104, 5, 5, 0, 0, 101, 103, 3, 6, 3, 0, 102, 101, 1, 0, 0, 0, 103, 106,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 5, 1, 0, 0,
		0, 106, 104, 1, 0, 0, 0, 107, 110, 3, 26, 13, 0, 108, 110, 3, 10, 5, 0,
		109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 7, 1, 0, 0, 0, 111, 120,
		3, 26, 13, 0, 112, 120, 3, 10, 5, 0, 113, 120, 3, 12, 6, 0, 114, 120, 5,
		6, 0, 0, 115, 120, 3, 14, 7, 0, 116, 120, 3, 30, 15, 0, 117, 120, 3, 16,
		8, 0, 118, 120, 3, 70, 35, 0, 119, 111, 1, 0, 0, 0, 119, 112, 1, 0, 0,
		0, 119, 113, 1, 0, 0, 0, 119, 114, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 119,
		116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 9, 1,
		0, 0, 0, 121, 123, 5, 36, 0, 0, 122, 124, 3, 26, 13, 0, 123, 122, 1, 0,
		0, 0, 123, 124, 1, 0, 0, 0, 124, 11, 1, 0, 0, 0, 125, 126, 5, 35, 0, 0,
		126, 13, 1, 0, 0, 0, 127, 128, 5, 7, 0, 0, 128, 15, 1, 0, 0, 0, 129, 135,
		3, 18, 9, 0, 130, 135, 3, 24, 12, 0, 131, 132, 3, 68, 34, 0, 132, 133,
		3, 28, 14, 0, 133, 135, 1, 0, 0, 0, 134, 129, 1, 0, 0, 0, 134, 130, 1,
		0, 0, 0, 134, 131, 1, 0, 0, 0, 135, 17, 1, 0, 0, 0, 136, 143, 3, 20, 10,
		0, 137, 139, 5, 3, 0, 0, 138, 140, 3, 46, 23, 0, 139, 138, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 3, 28, 14, 0, 142,
		136, 1, 0, 0, 0, 142, 137, 1, 0, 0, 0, 143, 19, 1, 0, 0, 0, 144, 146, 5,
		2, 0, 0, 145, 147, 3, 22, 11, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 3, 28, 14, 0, 149, 21, 1, 0, 0, 0,
		150, 151, 5, 38, 0, 0, 151, 23, 1, 0, 0, 0, 152, 153, 5, 4, 0, 0, 153,
		154, 3, 28, 14, 0, 154, 25, 1, 0, 0, 0, 155, 157, 5, 32, 0, 0, 156, 155,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 27, 1, 0, 0, 0, 160, 163, 3, 26, 13, 0, 161, 163, 5, 0, 0, 1,
		162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 29, 1, 0, 0, 0, 164, 165,
		3, 34, 17, 0, 165, 166, 3, 28, 14, 0, 166, 31, 1, 0, 0, 0, 167, 168, 7,
		0, 0, 0, 168, 33, 1, 0, 0, 0, 169, 170, 3, 32, 16, 0, 170, 172, 5, 19,
		0, 0, 171, 173, 3, 36, 18, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0,
		0, 173, 35, 1, 0, 0, 0, 174, 179, 3, 40, 20, 0, 175, 179, 3, 38, 19, 0,
		176, 179, 3, 56, 28, 0, 177, 179, 3, 48, 24, 0, 178, 174, 1, 0, 0, 0, 178,
		175, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179, 37, 1,
		0, 0, 0, 180, 185, 3, 64, 32, 0, 181, 185, 3, 46, 23, 0, 182, 185, 3, 62,
		31, 0, 183, 185, 3, 66, 33, 0, 184, 180, 1, 0, 0, 0, 184, 181, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 39, 1, 0, 0, 0, 186,
		188, 5, 13, 0, 0, 187, 189, 3, 42, 21, 0, 188, 187, 1, 0, 0, 0, 189, 190,
		1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 41, 1, 0,
		0, 0, 192, 196, 5, 26, 0, 0, 193, 195, 3, 60, 30, 0, 194, 193, 1, 0, 0,
		0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 3, 44, 22, 0, 200, 43,
		1, 0, 0, 0, 201, 207, 5, 13, 0, 0, 202, 207, 3, 62, 31, 0, 203, 207, 5,
		9, 0, 0, 204, 207, 5, 8, 0, 0, 205, 207, 5, 10, 0, 0, 206, 201, 1, 0, 0,
		0, 206, 202, 1, 0, 0, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		205, 1, 0, 0, 0, 207, 45, 1, 0, 0, 0, 208, 209, 5, 13, 0, 0, 209, 47, 1,
		0, 0, 0, 210, 214, 5, 24, 0, 0, 211, 213, 3, 60, 30, 0, 212, 211, 1, 0,
		0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 219, 3, 50, 25, 0, 218,
		217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 223, 1, 0, 0, 0, 220, 222,
		3, 60, 30, 0, 221, 220, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1,
		0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0, 0, 225, 223, 1, 0, 0,
		0, 226, 229, 5, 25, 0, 0, 227, 229, 5, 11, 0, 0, 228, 210, 1, 0, 0, 0,
		228, 227, 1, 0, 0, 0, 229, 49, 1, 0, 0, 0, 230, 247, 3, 52, 26, 0, 231,
		233, 3, 60, 30, 0, 232, 231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 237, 241, 5, 20, 0, 0, 238, 240, 3, 60, 30, 0, 239, 238, 1, 0, 0,
		0, 240, 243, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		244, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 246, 3, 52, 26, 0, 245, 234,
		1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0,
		0, 0, 248, 257, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 252, 3, 60, 30,
		0, 251, 250, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 258,
		5, 20, 0, 0, 257, 253, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 51, 1, 0,
		0, 0, 259, 260, 3, 32, 16, 0, 260, 261, 3, 54, 27, 0, 261, 262, 3, 36,
		18, 0, 262, 53, 1, 0, 0, 0, 263, 264, 7, 1, 0, 0, 264, 55, 1, 0, 0, 0,
		265, 269, 5, 22, 0, 0, 266, 268, 3, 60, 30, 0, 267, 266, 1, 0, 0, 0, 268,
		271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 273,
		1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 274, 3, 58, 29, 0, 273, 272, 1,
		0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 278, 1, 0, 0, 0, 275, 277, 3, 60, 30,
		0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278,
		279, 1, 0, 0, 0, 279, 281, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 284,
		5, 23, 0, 0, 282, 284, 5, 12, 0, 0, 283, 265, 1, 0, 0, 0, 283, 282, 1,
		0, 0, 0, 284, 57, 1, 0, 0, 0, 285, 302, 3, 36, 18, 0, 286, 288, 3, 60,
		30, 0, 287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0,
		289, 290, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292,
		296, 5, 20, 0, 0, 293, 295, 3, 60, 30, 0, 294, 293, 1, 0, 0, 0, 295, 298,
		1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 299, 1, 0,
		0, 0, 298, 296, 1, 0, 0, 0, 299, 301, 3, 36, 18, 0, 300, 289, 1, 0, 0,
		0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		312, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 307, 3, 60, 30, 0, 306, 305,
		1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0,
		0, 0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 313, 5, 20, 0, 0,
		312, 308, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 59, 1, 0, 0, 0, 314, 315,
		7, 2, 0, 0, 315, 61, 1, 0, 0, 0, 316, 318, 7, 3, 0, 0, 317, 316, 1, 0,
		0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 5, 14, 0, 0,
		320, 63, 1, 0, 0, 0, 321, 322, 5, 10, 0, 0, 322, 65, 1, 0, 0, 0, 323, 324,
		7, 4, 0, 0, 324, 67, 1, 0, 0, 0, 325, 326, 5, 41, 0, 0, 326, 69, 1, 0,
		0, 0, 327, 330, 5, 40, 0, 0, 328, 330, 3, 36, 18, 0, 329, 327, 1, 0, 0,
		0, 329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331,
		334, 5, 19, 0, 0, 332, 335, 3, 36, 18, 0, 333, 335, 5, 40, 0, 0, 334, 332,
		1, 0, 0, 0, 334, 333, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 338, 3, 26,
		13, 0, 337, 336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 71, 1, 0, 0, 0,
		44, 73, 78, 82, 90, 96, 98, 104, 109, 119, 123, 134, 139, 142, 146, 158,
		162, 172, 178, 184, 190, 196, 206, 214, 218, 223, 228, 234, 241, 247, 253,
		257, 269, 273, 278, 283, 289, 296, 302, 308, 312, 317, 329, 334, 337,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// YiniParserInit initializes any static state used to implement YiniParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYiniParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YiniParserInit() {
	staticData := &YiniParserParserStaticData
	staticData.once.Do(yiniparserParserInit)
}

// NewYiniParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYiniParser(input antlr.TokenStream) *YiniParser {
	YiniParserInit()
	this := new(YiniParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &YiniParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "YiniParser.g4"

	return this
}

// YiniParser tokens.
const (
	YiniParserEOF                  = antlr.TokenEOF
	YiniParserSHEBANG              = 1
	YiniParserYINI_TOKEN           = 2
	YiniParserINCLUDE_TOKEN        = 3
	YiniParserDEPRECATED_TOKEN     = 4
	YiniParserTERMINAL_TOKEN       = 5
	YiniParserSECTION_HEAD         = 6
	YiniParserINVALID_SECTION_HEAD = 7
	YiniParserBOOLEAN_FALSE        = 8
	YiniParserBOOLEAN_TRUE         = 9
	YiniParserNULL                 = 10
	YiniParserEMPTY_OBJECT         = 11
	YiniParserEMPTY_LIST           = 12
	YiniParserSTRING               = 13
	YiniParserNUMBER               = 14
	YiniParserCARET                = 15
	YiniParserSS                   = 16
	YiniParserGT                   = 17
	YiniParserLT                   = 18
	YiniParserEQ                   = 19
	YiniParserCOMMA                = 20
	YiniParserCOLON                = 21
	YiniParserOB                   = 22
	YiniParserCB                   = 23
	YiniParserOC                   = 24
	YiniParserCC                   = 25
	YiniParserPLUS                 = 26
	YiniParserMINUS                = 27
	YiniParserDOLLAR               = 28
	YiniParserPC                   = 29
	YiniParserAT                   = 30
	YiniParserSEMICOLON            = 31
	YiniParserNL                   = 32
	YiniParserWS                   = 33
	YiniParserBLOCK_COMMENT        = 34
	YiniParserDISABLED_LINE        = 35
	YiniParserFULL_LINE_COMMENT    = 36
	YiniParserINLINE_COMMENT       = 37
	YiniParserKEY                  = 38
	YiniParserIDENT_INVALID        = 39
	YiniParserREST                 = 40
	YiniParserMETA_INVALID         = 41
)

// YiniParser rules.
const (
	YiniParserRULE_yini                    = 0
	YiniParserRULE_prolog                  = 1
	YiniParserRULE_terminal_stmt           = 2
	YiniParserRULE_terminal_trivia         = 3
	YiniParserRULE_stmt                    = 4
	YiniParserRULE_full_line_comment_stmt  = 5
	YiniParserRULE_disabled_line_stmt      = 6
	YiniParserRULE_invalid_section_stmt    = 7
	YiniParserRULE_meta_stmt               = 8
	YiniParserRULE_directive               = 9
	YiniParserRULE_yini_directive          = 10
	YiniParserRULE_yini_mode_declaration   = 11
	YiniParserRULE_annotation              = 12
	YiniParserRULE_eol                     = 13
	YiniParserRULE_line_end                = 14
	YiniParserRULE_assignment              = 15
	YiniParserRULE_identifier              = 16
	YiniParserRULE_member                  = 17
	YiniParserRULE_value                   = 18
	YiniParserRULE_scalar_value            = 19
	YiniParserRULE_concat_expression       = 20
	YiniParserRULE_concat_tail             = 21
	YiniParserRULE_concat_operand          = 22
	YiniParserRULE_string_literal          = 23
	YiniParserRULE_object_literal          = 24
	YiniParserRULE_object_members          = 25
	YiniParserRULE_object_member           = 26
	YiniParserRULE_object_member_separator = 27
	YiniParserRULE_list_literal            = 28
	YiniParserRULE_elements                = 29
	YiniParserRULE_multiline_trivia        = 30
	YiniParserRULE_number_literal          = 31
	YiniParserRULE_null_literal            = 32
	YiniParserRULE_boolean_literal         = 33
	YiniParserRULE_bad_meta_text           = 34
	YiniParserRULE_bad_member              = 35
)

// IYiniContext is an interface to support dynamic dispatch.
type IYiniContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Prolog() IPrologContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Terminal_stmt() ITerminal_stmtContext

	// IsYiniContext differentiates from other interfaces.
	IsYiniContext()
}

type YiniContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYiniContext() *YiniContext {
	var p = new(YiniContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini
	return p
}

func InitEmptyYiniContext(p *YiniContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini
}

func (*YiniContext) IsYiniContext() {}

func NewYiniContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YiniContext {
	var p = new(YiniContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_yini

	return p
}

func (s *YiniContext) GetParser() antlr.Parser { return s.parser }

func (s *YiniContext) EOF() antlr.TerminalNode {
	return s.GetToken(YiniParserEOF, 0)
}

func (s *YiniContext) Prolog() IPrologContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrologContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrologContext)
}

func (s *YiniContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *YiniContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *YiniContext) Terminal_stmt() ITerminal_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminal_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminal_stmtContext)
}

func (s *YiniContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YiniContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YiniContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitYini(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Yini() (localctx IYiniContext) {
	localctx = NewYiniContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YiniParserRULE_yini)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(72)
			p.Prolog()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3681009827804) != 0 {
		{
			p.SetState(75)
			p.Stmt()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserTERMINAL_TOKEN {
		{
			p.SetState(81)
			p.Terminal_stmt()
		}

	}
	{
		p.SetState(84)
		p.Match(YiniParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrologContext is an interface to support dynamic dispatch.
type IPrologContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHEBANG() antlr.TerminalNode
	AllEol() []IEolContext
	Eol(i int) IEolContext

	// IsPrologContext differentiates from other interfaces.
	IsPrologContext()
}

type PrologContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologContext() *PrologContext {
	var p = new(PrologContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_prolog
	return p
}

func InitEmptyPrologContext(p *PrologContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_prolog
}

func (*PrologContext) IsPrologContext() {}

func NewPrologContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologContext {
	var p = new(PrologContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_prolog

	return p
}

func (s *PrologContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologContext) SHEBANG() antlr.TerminalNode {
	return s.GetToken(YiniParserSHEBANG, 0)
}

func (s *PrologContext) AllEol() []IEolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEolContext); ok {
			len++
		}
	}

	tst := make([]IEolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEolContext); ok {
			tst[i] = t.(IEolContext)
			i++
		}
	}

	return tst
}

func (s *PrologContext) Eol(i int) IEolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PrologContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrologContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrologContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitProlog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Prolog() (localctx IPrologContext) {
	localctx = NewPrologContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YiniParserRULE_prolog)
	var _alt int

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserSHEBANG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(YiniParserSHEBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(87)
					p.Eol()
				}

			}
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case YiniParserNL:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(93)
					p.Eol()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminal_stmtContext is an interface to support dynamic dispatch.
type ITerminal_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TERMINAL_TOKEN() antlr.TerminalNode
	AllTerminal_trivia() []ITerminal_triviaContext
	Terminal_trivia(i int) ITerminal_triviaContext

	// IsTerminal_stmtContext differentiates from other interfaces.
	IsTerminal_stmtContext()
}

type Terminal_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminal_stmtContext() *Terminal_stmtContext {
	var p = new(Terminal_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_terminal_stmt
	return p
}

func InitEmptyTerminal_stmtContext(p *Terminal_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_terminal_stmt
}

func (*Terminal_stmtContext) IsTerminal_stmtContext() {}

func NewTerminal_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Terminal_stmtContext {
	var p = new(Terminal_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_terminal_stmt

	return p
}

func (s *Terminal_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Terminal_stmtContext) TERMINAL_TOKEN() antlr.TerminalNode {
	return s.GetToken(YiniParserTERMINAL_TOKEN, 0)
}

func (s *Terminal_stmtContext) AllTerminal_trivia() []ITerminal_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminal_triviaContext); ok {
			len++
		}
	}

	tst := make([]ITerminal_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminal_triviaContext); ok {
			tst[i] = t.(ITerminal_triviaContext)
			i++
		}
	}

	return tst
}

func (s *Terminal_stmtContext) Terminal_trivia(i int) ITerminal_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminal_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminal_triviaContext)
}

func (s *Terminal_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Terminal_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Terminal_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitTerminal_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Terminal_stmt() (localctx ITerminal_stmtContext) {
	localctx = NewTerminal_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YiniParserRULE_terminal_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(YiniParserTERMINAL_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YiniParserNL || _la == YiniParserFULL_LINE_COMMENT {
		{
			p.SetState(101)
			p.Terminal_trivia()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminal_triviaContext is an interface to support dynamic dispatch.
type ITerminal_triviaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Eol() IEolContext
	Full_line_comment_stmt() IFull_line_comment_stmtContext

	// IsTerminal_triviaContext differentiates from other interfaces.
	IsTerminal_triviaContext()
}

type Terminal_triviaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminal_triviaContext() *Terminal_triviaContext {
	var p = new(Terminal_triviaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_terminal_trivia
	return p
}

func InitEmptyTerminal_triviaContext(p *Terminal_triviaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_terminal_trivia
}

func (*Terminal_triviaContext) IsTerminal_triviaContext() {}

func NewTerminal_triviaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Terminal_triviaContext {
	var p = new(Terminal_triviaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_terminal_trivia

	return p
}

func (s *Terminal_triviaContext) GetParser() antlr.Parser { return s.parser }

func (s *Terminal_triviaContext) Eol() IEolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Terminal_triviaContext) Full_line_comment_stmt() IFull_line_comment_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFull_line_comment_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFull_line_comment_stmtContext)
}

func (s *Terminal_triviaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Terminal_triviaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Terminal_triviaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitTerminal_trivia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Terminal_trivia() (localctx ITerminal_triviaContext) {
	localctx = NewTerminal_triviaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YiniParserRULE_terminal_trivia)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Eol()
		}

	case YiniParserFULL_LINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Full_line_comment_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Eol() IEolContext
	Full_line_comment_stmt() IFull_line_comment_stmtContext
	Disabled_line_stmt() IDisabled_line_stmtContext
	SECTION_HEAD() antlr.TerminalNode
	Invalid_section_stmt() IInvalid_section_stmtContext
	Assignment() IAssignmentContext
	Meta_stmt() IMeta_stmtContext
	Bad_member() IBad_memberContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Eol() IEolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *StmtContext) Full_line_comment_stmt() IFull_line_comment_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFull_line_comment_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFull_line_comment_stmtContext)
}

func (s *StmtContext) Disabled_line_stmt() IDisabled_line_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisabled_line_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisabled_line_stmtContext)
}

func (s *StmtContext) SECTION_HEAD() antlr.TerminalNode {
	return s.GetToken(YiniParserSECTION_HEAD, 0)
}

func (s *StmtContext) Invalid_section_stmt() IInvalid_section_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvalid_section_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvalid_section_stmtContext)
}

func (s *StmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtContext) Meta_stmt() IMeta_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_stmtContext)
}

func (s *StmtContext) Bad_member() IBad_memberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBad_memberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBad_memberContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YiniParserRULE_stmt)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Eol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Full_line_comment_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Disabled_line_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Match(YiniParserSECTION_HEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Invalid_section_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Assignment()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.Meta_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(118)
			p.Bad_member()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFull_line_comment_stmtContext is an interface to support dynamic dispatch.
type IFull_line_comment_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FULL_LINE_COMMENT() antlr.TerminalNode
	Eol() IEolContext

	// IsFull_line_comment_stmtContext differentiates from other interfaces.
	IsFull_line_comment_stmtContext()
}

type Full_line_comment_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFull_line_comment_stmtContext() *Full_line_comment_stmtContext {
	var p = new(Full_line_comment_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_full_line_comment_stmt
	return p
}

func InitEmptyFull_line_comment_stmtContext(p *Full_line_comment_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_full_line_comment_stmt
}

func (*Full_line_comment_stmtContext) IsFull_line_comment_stmtContext() {}

func NewFull_line_comment_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Full_line_comment_stmtContext {
	var p = new(Full_line_comment_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_full_line_comment_stmt

	return p
}

func (s *Full_line_comment_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Full_line_comment_stmtContext) FULL_LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(YiniParserFULL_LINE_COMMENT, 0)
}

func (s *Full_line_comment_stmtContext) Eol() IEolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Full_line_comment_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Full_line_comment_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Full_line_comment_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitFull_line_comment_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Full_line_comment_stmt() (localctx IFull_line_comment_stmtContext) {
	localctx = NewFull_line_comment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, YiniParserRULE_full_line_comment_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(YiniParserFULL_LINE_COMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(122)
			p.Eol()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDisabled_line_stmtContext is an interface to support dynamic dispatch.
type IDisabled_line_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DISABLED_LINE() antlr.TerminalNode

	// IsDisabled_line_stmtContext differentiates from other interfaces.
	IsDisabled_line_stmtContext()
}

type Disabled_line_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisabled_line_stmtContext() *Disabled_line_stmtContext {
	var p = new(Disabled_line_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_disabled_line_stmt
	return p
}

func InitEmptyDisabled_line_stmtContext(p *Disabled_line_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_disabled_line_stmt
}

func (*Disabled_line_stmtContext) IsDisabled_line_stmtContext() {}

func NewDisabled_line_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Disabled_line_stmtContext {
	var p = new(Disabled_line_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_disabled_line_stmt

	return p
}

func (s *Disabled_line_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Disabled_line_stmtContext) DISABLED_LINE() antlr.TerminalNode {
	return s.GetToken(YiniParserDISABLED_LINE, 0)
}

func (s *Disabled_line_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Disabled_line_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Disabled_line_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitDisabled_line_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Disabled_line_stmt() (localctx IDisabled_line_stmtContext) {
	localctx = NewDisabled_line_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YiniParserRULE_disabled_line_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(YiniParserDISABLED_LINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInvalid_section_stmtContext is an interface to support dynamic dispatch.
type IInvalid_section_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INVALID_SECTION_HEAD() antlr.TerminalNode

	// IsInvalid_section_stmtContext differentiates from other interfaces.
	IsInvalid_section_stmtContext()
}

type Invalid_section_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvalid_section_stmtContext() *Invalid_section_stmtContext {
	var p = new(Invalid_section_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_invalid_section_stmt
	return p
}

func InitEmptyInvalid_section_stmtContext(p *Invalid_section_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_invalid_section_stmt
}

func (*Invalid_section_stmtContext) IsInvalid_section_stmtContext() {}

func NewInvalid_section_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Invalid_section_stmtContext {
	var p = new(Invalid_section_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_invalid_section_stmt

	return p
}

func (s *Invalid_section_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Invalid_section_stmtContext) INVALID_SECTION_HEAD() antlr.TerminalNode {
	return s.GetToken(YiniParserINVALID_SECTION_HEAD, 0)
}

func (s *Invalid_section_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Invalid_section_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Invalid_section_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitInvalid_section_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Invalid_section_stmt() (localctx IInvalid_section_stmtContext) {
	localctx = NewInvalid_section_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YiniParserRULE_invalid_section_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(YiniParserINVALID_SECTION_HEAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMeta_stmtContext is an interface to support dynamic dispatch.
type IMeta_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directive() IDirectiveContext
	Annotation() IAnnotationContext
	Bad_meta_text() IBad_meta_textContext
	Line_end() ILine_endContext

	// IsMeta_stmtContext differentiates from other interfaces.
	IsMeta_stmtContext()
}

type Meta_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_stmtContext() *Meta_stmtContext {
	var p = new(Meta_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_meta_stmt
	return p
}

func InitEmptyMeta_stmtContext(p *Meta_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_meta_stmt
}

func (*Meta_stmtContext) IsMeta_stmtContext() {}

func NewMeta_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_stmtContext {
	var p = new(Meta_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_meta_stmt

	return p
}

func (s *Meta_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_stmtContext) Directive() IDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *Meta_stmtContext) Annotation() IAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *Meta_stmtContext) Bad_meta_text() IBad_meta_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBad_meta_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBad_meta_textContext)
}

func (s *Meta_stmtContext) Line_end() ILine_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_endContext)
}

func (s *Meta_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitMeta_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Meta_stmt() (localctx IMeta_stmtContext) {
	localctx = NewMeta_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, YiniParserRULE_meta_stmt)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserYINI_TOKEN, YiniParserINCLUDE_TOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Directive()
		}

	case YiniParserDEPRECATED_TOKEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Annotation()
		}

	case YiniParserMETA_INVALID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Bad_meta_text()
		}
		{
			p.SetState(132)
			p.Line_end()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Yini_directive() IYini_directiveContext
	INCLUDE_TOKEN() antlr.TerminalNode
	Line_end() ILine_endContext
	String_literal() IString_literalContext

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_directive
	return p
}

func InitEmptyDirectiveContext(p *DirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_directive
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Yini_directive() IYini_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYini_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYini_directiveContext)
}

func (s *DirectiveContext) INCLUDE_TOKEN() antlr.TerminalNode {
	return s.GetToken(YiniParserINCLUDE_TOKEN, 0)
}

func (s *DirectiveContext) Line_end() ILine_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_endContext)
}

func (s *DirectiveContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YiniParserRULE_directive)
	var _la int

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserYINI_TOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Yini_directive()
		}

	case YiniParserINCLUDE_TOKEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(YiniParserINCLUDE_TOKEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YiniParserSTRING {
			{
				p.SetState(138)
				p.String_literal()
			}

		}
		{
			p.SetState(141)
			p.Line_end()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IYini_directiveContext is an interface to support dynamic dispatch.
type IYini_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	YINI_TOKEN() antlr.TerminalNode
	Line_end() ILine_endContext
	Yini_mode_declaration() IYini_mode_declarationContext

	// IsYini_directiveContext differentiates from other interfaces.
	IsYini_directiveContext()
}

type Yini_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYini_directiveContext() *Yini_directiveContext {
	var p = new(Yini_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini_directive
	return p
}

func InitEmptyYini_directiveContext(p *Yini_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini_directive
}

func (*Yini_directiveContext) IsYini_directiveContext() {}

func NewYini_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yini_directiveContext {
	var p = new(Yini_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_yini_directive

	return p
}

func (s *Yini_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Yini_directiveContext) YINI_TOKEN() antlr.TerminalNode {
	return s.GetToken(YiniParserYINI_TOKEN, 0)
}

func (s *Yini_directiveContext) Line_end() ILine_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_endContext)
}

func (s *Yini_directiveContext) Yini_mode_declaration() IYini_mode_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYini_mode_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYini_mode_declarationContext)
}

func (s *Yini_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yini_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yini_directiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitYini_directive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Yini_directive() (localctx IYini_directiveContext) {
	localctx = NewYini_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, YiniParserRULE_yini_directive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(YiniParserYINI_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserKEY {
		{
			p.SetState(145)
			p.Yini_mode_declaration()
		}

	}
	{
		p.SetState(148)
		p.Line_end()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IYini_mode_declarationContext is an interface to support dynamic dispatch.
type IYini_mode_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEY() antlr.TerminalNode

	// IsYini_mode_declarationContext differentiates from other interfaces.
	IsYini_mode_declarationContext()
}

type Yini_mode_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYini_mode_declarationContext() *Yini_mode_declarationContext {
	var p = new(Yini_mode_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini_mode_declaration
	return p
}

func InitEmptyYini_mode_declarationContext(p *Yini_mode_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_yini_mode_declaration
}

func (*Yini_mode_declarationContext) IsYini_mode_declarationContext() {}

func NewYini_mode_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yini_mode_declarationContext {
	var p = new(Yini_mode_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_yini_mode_declaration

	return p
}

func (s *Yini_mode_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Yini_mode_declarationContext) KEY() antlr.TerminalNode {
	return s.GetToken(YiniParserKEY, 0)
}

func (s *Yini_mode_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yini_mode_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yini_mode_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitYini_mode_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Yini_mode_declaration() (localctx IYini_mode_declarationContext) {
	localctx = NewYini_mode_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YiniParserRULE_yini_mode_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(YiniParserKEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEPRECATED_TOKEN() antlr.TerminalNode
	Line_end() ILine_endContext

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_annotation
	return p
}

func InitEmptyAnnotationContext(p *AnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_annotation
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) DEPRECATED_TOKEN() antlr.TerminalNode {
	return s.GetToken(YiniParserDEPRECATED_TOKEN, 0)
}

func (s *AnnotationContext) Line_end() ILine_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_endContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, YiniParserRULE_annotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(YiniParserDEPRECATED_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Line_end()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEolContext is an interface to support dynamic dispatch.
type IEolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsEolContext differentiates from other interfaces.
	IsEolContext()
}

type EolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEolContext() *EolContext {
	var p = new(EolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_eol
	return p
}

func InitEmptyEolContext(p *EolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_eol
}

func (*EolContext) IsEolContext() {}

func NewEolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EolContext {
	var p = new(EolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_eol

	return p
}

func (s *EolContext) GetParser() antlr.Parser { return s.parser }

func (s *EolContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *EolContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
}

func (s *EolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitEol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Eol() (localctx IEolContext) {
	localctx = NewEolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, YiniParserRULE_eol)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(155)
				p.Match(YiniParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_endContext is an interface to support dynamic dispatch.
type ILine_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Eol() IEolContext
	EOF() antlr.TerminalNode

	// IsLine_endContext differentiates from other interfaces.
	IsLine_endContext()
}

type Line_endContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_endContext() *Line_endContext {
	var p = new(Line_endContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_line_end
	return p
}

func InitEmptyLine_endContext(p *Line_endContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_line_end
}

func (*Line_endContext) IsLine_endContext() {}

func NewLine_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_endContext {
	var p = new(Line_endContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_line_end

	return p
}

func (s *Line_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_endContext) Eol() IEolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Line_endContext) EOF() antlr.TerminalNode {
	return s.GetToken(YiniParserEOF, 0)
}

func (s *Line_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitLine_end(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Line_end() (localctx ILine_endContext) {
	localctx = NewLine_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, YiniParserRULE_line_end)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Eol()
		}

	case YiniParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(YiniParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Member() IMemberContext
	Line_end() ILine_endContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Member() IMemberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *AssignmentContext) Line_end() ILine_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_endContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, YiniParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Member()
	}
	{
		p.SetState(165)
		p.Line_end()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEY() antlr.TerminalNode
	BOOLEAN_TRUE() antlr.TerminalNode
	BOOLEAN_FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) KEY() antlr.TerminalNode {
	return s.GetToken(YiniParserKEY, 0)
}

func (s *IdentifierContext) BOOLEAN_TRUE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_TRUE, 0)
}

func (s *IdentifierContext) BOOLEAN_FALSE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_FALSE, 0)
}

func (s *IdentifierContext) NULL() antlr.TerminalNode {
	return s.GetToken(YiniParserNULL, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, YiniParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877908736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	EQ() antlr.TerminalNode
	Value() IValueContext

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_member
	return p
}

func InitEmptyMemberContext(p *MemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_member
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberContext) EQ() antlr.TerminalNode {
	return s.GetToken(YiniParserEQ, 0)
}

func (s *MemberContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, YiniParserRULE_member)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Identifier()
	}
	{
		p.SetState(170)
		p.Match(YiniParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&222330624) != 0 {
		{
			p.SetState(171)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Concat_expression() IConcat_expressionContext
	Scalar_value() IScalar_valueContext
	List_literal() IList_literalContext
	Object_literal() IObject_literalContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Concat_expression() IConcat_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcat_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcat_expressionContext)
}

func (s *ValueContext) Scalar_value() IScalar_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_valueContext)
}

func (s *ValueContext) List_literal() IList_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_literalContext)
}

func (s *ValueContext) Object_literal() IObject_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_literalContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, YiniParserRULE_value)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Concat_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Scalar_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.List_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(177)
			p.Object_literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalar_valueContext is an interface to support dynamic dispatch.
type IScalar_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Null_literal() INull_literalContext
	String_literal() IString_literalContext
	Number_literal() INumber_literalContext
	Boolean_literal() IBoolean_literalContext

	// IsScalar_valueContext differentiates from other interfaces.
	IsScalar_valueContext()
}

type Scalar_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalar_valueContext() *Scalar_valueContext {
	var p = new(Scalar_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_scalar_value
	return p
}

func InitEmptyScalar_valueContext(p *Scalar_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_scalar_value
}

func (*Scalar_valueContext) IsScalar_valueContext() {}

func NewScalar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scalar_valueContext {
	var p = new(Scalar_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_scalar_value

	return p
}

func (s *Scalar_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Scalar_valueContext) Null_literal() INull_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INull_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INull_literalContext)
}

func (s *Scalar_valueContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Scalar_valueContext) Number_literal() INumber_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_literalContext)
}

func (s *Scalar_valueContext) Boolean_literal() IBoolean_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_literalContext)
}

func (s *Scalar_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scalar_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scalar_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitScalar_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Scalar_value() (localctx IScalar_valueContext) {
	localctx = NewScalar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, YiniParserRULE_scalar_value)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNULL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Null_literal()
		}

	case YiniParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.String_literal()
		}

	case YiniParserNUMBER, YiniParserPLUS, YiniParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Number_literal()
		}

	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Boolean_literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConcat_expressionContext is an interface to support dynamic dispatch.
type IConcat_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	AllConcat_tail() []IConcat_tailContext
	Concat_tail(i int) IConcat_tailContext

	// IsConcat_expressionContext differentiates from other interfaces.
	IsConcat_expressionContext()
}

type Concat_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcat_expressionContext() *Concat_expressionContext {
	var p = new(Concat_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_expression
	return p
}

func InitEmptyConcat_expressionContext(p *Concat_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_expression
}

func (*Concat_expressionContext) IsConcat_expressionContext() {}

func NewConcat_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Concat_expressionContext {
	var p = new(Concat_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_concat_expression

	return p
}

func (s *Concat_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Concat_expressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(YiniParserSTRING, 0)
}

func (s *Concat_expressionContext) AllConcat_tail() []IConcat_tailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConcat_tailContext); ok {
			len++
		}
	}

	tst := make([]IConcat_tailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConcat_tailContext); ok {
			tst[i] = t.(IConcat_tailContext)
			i++
		}
	}

	return tst
}

func (s *Concat_expressionContext) Concat_tail(i int) IConcat_tailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcat_tailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcat_tailContext)
}

func (s *Concat_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Concat_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Concat_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitConcat_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Concat_expression() (localctx IConcat_expressionContext) {
	localctx = NewConcat_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, YiniParserRULE_concat_expression)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(YiniParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(187)
				p.Concat_tail()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConcat_tailContext is an interface to support dynamic dispatch.
type IConcat_tailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	Concat_operand() IConcat_operandContext
	AllMultiline_trivia() []IMultiline_triviaContext
	Multiline_trivia(i int) IMultiline_triviaContext

	// IsConcat_tailContext differentiates from other interfaces.
	IsConcat_tailContext()
}

type Concat_tailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcat_tailContext() *Concat_tailContext {
	var p = new(Concat_tailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_tail
	return p
}

func InitEmptyConcat_tailContext(p *Concat_tailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_tail
}

func (*Concat_tailContext) IsConcat_tailContext() {}

func NewConcat_tailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Concat_tailContext {
	var p = new(Concat_tailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_concat_tail

	return p
}

func (s *Concat_tailContext) GetParser() antlr.Parser { return s.parser }

func (s *Concat_tailContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YiniParserPLUS, 0)
}

func (s *Concat_tailContext) Concat_operand() IConcat_operandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcat_operandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcat_operandContext)
}

func (s *Concat_tailContext) AllMultiline_trivia() []IMultiline_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			len++
		}
	}

	tst := make([]IMultiline_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiline_triviaContext); ok {
			tst[i] = t.(IMultiline_triviaContext)
			i++
		}
	}

	return tst
}

func (s *Concat_tailContext) Multiline_trivia(i int) IMultiline_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiline_triviaContext)
}

func (s *Concat_tailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Concat_tailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Concat_tailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitConcat_tail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Concat_tail() (localctx IConcat_tailContext) {
	localctx = NewConcat_tailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, YiniParserRULE_concat_tail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(YiniParserPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
		{
			p.SetState(193)
			p.Multiline_trivia()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Concat_operand()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConcat_operandContext is an interface to support dynamic dispatch.
type IConcat_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Number_literal() INumber_literalContext
	BOOLEAN_TRUE() antlr.TerminalNode
	BOOLEAN_FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsConcat_operandContext differentiates from other interfaces.
	IsConcat_operandContext()
}

type Concat_operandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcat_operandContext() *Concat_operandContext {
	var p = new(Concat_operandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_operand
	return p
}

func InitEmptyConcat_operandContext(p *Concat_operandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_concat_operand
}

func (*Concat_operandContext) IsConcat_operandContext() {}

func NewConcat_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Concat_operandContext {
	var p = new(Concat_operandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_concat_operand

	return p
}

func (s *Concat_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Concat_operandContext) STRING() antlr.TerminalNode {
	return s.GetToken(YiniParserSTRING, 0)
}

func (s *Concat_operandContext) Number_literal() INumber_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_literalContext)
}

func (s *Concat_operandContext) BOOLEAN_TRUE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_TRUE, 0)
}

func (s *Concat_operandContext) BOOLEAN_FALSE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_FALSE, 0)
}

func (s *Concat_operandContext) NULL() antlr.TerminalNode {
	return s.GetToken(YiniParserNULL, 0)
}

func (s *Concat_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Concat_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Concat_operandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitConcat_operand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Concat_operand() (localctx IConcat_operandContext) {
	localctx = NewConcat_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, YiniParserRULE_concat_operand)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Match(YiniParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserNUMBER, YiniParserPLUS, YiniParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Number_literal()
		}

	case YiniParserBOOLEAN_TRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(YiniParserBOOLEAN_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserBOOLEAN_FALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(YiniParserBOOLEAN_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserNULL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)
			p.Match(YiniParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_string_literal
	return p
}

func InitEmptyString_literalContext(p *String_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_string_literal
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING() antlr.TerminalNode {
	return s.GetToken(YiniParserSTRING, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitString_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, YiniParserRULE_string_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(YiniParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_literalContext is an interface to support dynamic dispatch.
type IObject_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OC() antlr.TerminalNode
	CC() antlr.TerminalNode
	AllMultiline_trivia() []IMultiline_triviaContext
	Multiline_trivia(i int) IMultiline_triviaContext
	Object_members() IObject_membersContext
	EMPTY_OBJECT() antlr.TerminalNode

	// IsObject_literalContext differentiates from other interfaces.
	IsObject_literalContext()
}

type Object_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_literalContext() *Object_literalContext {
	var p = new(Object_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_literal
	return p
}

func InitEmptyObject_literalContext(p *Object_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_literal
}

func (*Object_literalContext) IsObject_literalContext() {}

func NewObject_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_literalContext {
	var p = new(Object_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_object_literal

	return p
}

func (s *Object_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_literalContext) OC() antlr.TerminalNode {
	return s.GetToken(YiniParserOC, 0)
}

func (s *Object_literalContext) CC() antlr.TerminalNode {
	return s.GetToken(YiniParserCC, 0)
}

func (s *Object_literalContext) AllMultiline_trivia() []IMultiline_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			len++
		}
	}

	tst := make([]IMultiline_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiline_triviaContext); ok {
			tst[i] = t.(IMultiline_triviaContext)
			i++
		}
	}

	return tst
}

func (s *Object_literalContext) Multiline_trivia(i int) IMultiline_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiline_triviaContext)
}

func (s *Object_literalContext) Object_members() IObject_membersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_membersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_membersContext)
}

func (s *Object_literalContext) EMPTY_OBJECT() antlr.TerminalNode {
	return s.GetToken(YiniParserEMPTY_OBJECT, 0)
}

func (s *Object_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitObject_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Object_literal() (localctx IObject_literalContext) {
	localctx = NewObject_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, YiniParserRULE_object_literal)
	var _la int

	var _alt int

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserOC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(YiniParserOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(211)
					p.Multiline_trivia()
				}

			}
			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877908736) != 0 {
			{
				p.SetState(217)
				p.Object_members()
			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
			{
				p.SetState(220)
				p.Multiline_trivia()
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Match(YiniParserCC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserEMPTY_OBJECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(YiniParserEMPTY_OBJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_membersContext is an interface to support dynamic dispatch.
type IObject_membersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllObject_member() []IObject_memberContext
	Object_member(i int) IObject_memberContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllMultiline_trivia() []IMultiline_triviaContext
	Multiline_trivia(i int) IMultiline_triviaContext

	// IsObject_membersContext differentiates from other interfaces.
	IsObject_membersContext()
}

type Object_membersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_membersContext() *Object_membersContext {
	var p = new(Object_membersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_members
	return p
}

func InitEmptyObject_membersContext(p *Object_membersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_members
}

func (*Object_membersContext) IsObject_membersContext() {}

func NewObject_membersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_membersContext {
	var p = new(Object_membersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_object_members

	return p
}

func (s *Object_membersContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_membersContext) AllObject_member() []IObject_memberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_memberContext); ok {
			len++
		}
	}

	tst := make([]IObject_memberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_memberContext); ok {
			tst[i] = t.(IObject_memberContext)
			i++
		}
	}

	return tst
}

func (s *Object_membersContext) Object_member(i int) IObject_memberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_memberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_memberContext)
}

func (s *Object_membersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YiniParserCOMMA)
}

func (s *Object_membersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserCOMMA, i)
}

func (s *Object_membersContext) AllMultiline_trivia() []IMultiline_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			len++
		}
	}

	tst := make([]IMultiline_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiline_triviaContext); ok {
			tst[i] = t.(IMultiline_triviaContext)
			i++
		}
	}

	return tst
}

func (s *Object_membersContext) Multiline_trivia(i int) IMultiline_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiline_triviaContext)
}

func (s *Object_membersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_membersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_membersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitObject_members(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Object_members() (localctx IObject_membersContext) {
	localctx = NewObject_membersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, YiniParserRULE_object_members)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Object_member()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
				{
					p.SetState(231)
					p.Multiline_trivia()
				}

				p.SetState(236)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(237)
				p.Match(YiniParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
				{
					p.SetState(238)
					p.Multiline_trivia()
				}

				p.SetState(243)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(244)
				p.Object_member()
			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
			{
				p.SetState(250)
				p.Multiline_trivia()
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(256)
			p.Match(YiniParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_memberContext is an interface to support dynamic dispatch.
type IObject_memberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Object_member_separator() IObject_member_separatorContext
	Value() IValueContext

	// IsObject_memberContext differentiates from other interfaces.
	IsObject_memberContext()
}

type Object_memberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_memberContext() *Object_memberContext {
	var p = new(Object_memberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_member
	return p
}

func InitEmptyObject_memberContext(p *Object_memberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_member
}

func (*Object_memberContext) IsObject_memberContext() {}

func NewObject_memberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_memberContext {
	var p = new(Object_memberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_object_member

	return p
}

func (s *Object_memberContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_memberContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Object_memberContext) Object_member_separator() IObject_member_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_member_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_member_separatorContext)
}

func (s *Object_memberContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Object_memberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_memberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_memberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitObject_member(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Object_member() (localctx IObject_memberContext) {
	localctx = NewObject_memberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, YiniParserRULE_object_member)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Identifier()
	}
	{
		p.SetState(260)
		p.Object_member_separator()
	}
	{
		p.SetState(261)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_member_separatorContext is an interface to support dynamic dispatch.
type IObject_member_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	EQ() antlr.TerminalNode

	// IsObject_member_separatorContext differentiates from other interfaces.
	IsObject_member_separatorContext()
}

type Object_member_separatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_member_separatorContext() *Object_member_separatorContext {
	var p = new(Object_member_separatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_member_separator
	return p
}

func InitEmptyObject_member_separatorContext(p *Object_member_separatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_object_member_separator
}

func (*Object_member_separatorContext) IsObject_member_separatorContext() {}

func NewObject_member_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_member_separatorContext {
	var p = new(Object_member_separatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_object_member_separator

	return p
}

func (s *Object_member_separatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_member_separatorContext) COLON() antlr.TerminalNode {
	return s.GetToken(YiniParserCOLON, 0)
}

func (s *Object_member_separatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(YiniParserEQ, 0)
}

func (s *Object_member_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_member_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_member_separatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitObject_member_separator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Object_member_separator() (localctx IObject_member_separatorContext) {
	localctx = NewObject_member_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, YiniParserRULE_object_member_separator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !(_la == YiniParserEQ || _la == YiniParserCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_literalContext is an interface to support dynamic dispatch.
type IList_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OB() antlr.TerminalNode
	CB() antlr.TerminalNode
	AllMultiline_trivia() []IMultiline_triviaContext
	Multiline_trivia(i int) IMultiline_triviaContext
	Elements() IElementsContext
	EMPTY_LIST() antlr.TerminalNode

	// IsList_literalContext differentiates from other interfaces.
	IsList_literalContext()
}

type List_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_literalContext() *List_literalContext {
	var p = new(List_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_list_literal
	return p
}

func InitEmptyList_literalContext(p *List_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_list_literal
}

func (*List_literalContext) IsList_literalContext() {}

func NewList_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_literalContext {
	var p = new(List_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_list_literal

	return p
}

func (s *List_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *List_literalContext) OB() antlr.TerminalNode {
	return s.GetToken(YiniParserOB, 0)
}

func (s *List_literalContext) CB() antlr.TerminalNode {
	return s.GetToken(YiniParserCB, 0)
}

func (s *List_literalContext) AllMultiline_trivia() []IMultiline_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			len++
		}
	}

	tst := make([]IMultiline_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiline_triviaContext); ok {
			tst[i] = t.(IMultiline_triviaContext)
			i++
		}
	}

	return tst
}

func (s *List_literalContext) Multiline_trivia(i int) IMultiline_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiline_triviaContext)
}

func (s *List_literalContext) Elements() IElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *List_literalContext) EMPTY_LIST() antlr.TerminalNode {
	return s.GetToken(YiniParserEMPTY_LIST, 0)
}

func (s *List_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitList_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) List_literal() (localctx IList_literalContext) {
	localctx = NewList_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, YiniParserRULE_list_literal)
	var _la int

	var _alt int

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserOB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Match(YiniParserOB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(266)
					p.Multiline_trivia()
				}

			}
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&222330624) != 0 {
			{
				p.SetState(272)
				p.Elements()
			}

		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
			{
				p.SetState(275)
				p.Multiline_trivia()
			}

			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(281)
			p.Match(YiniParserCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserEMPTY_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(YiniParserEMPTY_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementsContext is an interface to support dynamic dispatch.
type IElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllMultiline_trivia() []IMultiline_triviaContext
	Multiline_trivia(i int) IMultiline_triviaContext

	// IsElementsContext differentiates from other interfaces.
	IsElementsContext()
}

type ElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementsContext() *ElementsContext {
	var p = new(ElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_elements
	return p
}

func InitEmptyElementsContext(p *ElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_elements
}

func (*ElementsContext) IsElementsContext() {}

func NewElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementsContext {
	var p = new(ElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_elements

	return p
}

func (s *ElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementsContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ElementsContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ElementsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YiniParserCOMMA)
}

func (s *ElementsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserCOMMA, i)
}

func (s *ElementsContext) AllMultiline_trivia() []IMultiline_triviaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			len++
		}
	}

	tst := make([]IMultiline_triviaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiline_triviaContext); ok {
			tst[i] = t.(IMultiline_triviaContext)
			i++
		}
	}

	return tst
}

func (s *ElementsContext) Multiline_trivia(i int) IMultiline_triviaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiline_triviaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiline_triviaContext)
}

func (s *ElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Elements() (localctx IElementsContext) {
	localctx = NewElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, YiniParserRULE_elements)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Value()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(289)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
				{
					p.SetState(286)
					p.Multiline_trivia()
				}

				p.SetState(291)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(292)
				p.Match(YiniParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
				{
					p.SetState(293)
					p.Multiline_trivia()
				}

				p.SetState(298)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(299)
				p.Value()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0 {
			{
				p.SetState(305)
				p.Multiline_trivia()
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(311)
			p.Match(YiniParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiline_triviaContext is an interface to support dynamic dispatch.
type IMultiline_triviaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NL() antlr.TerminalNode
	DISABLED_LINE() antlr.TerminalNode
	FULL_LINE_COMMENT() antlr.TerminalNode

	// IsMultiline_triviaContext differentiates from other interfaces.
	IsMultiline_triviaContext()
}

type Multiline_triviaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiline_triviaContext() *Multiline_triviaContext {
	var p = new(Multiline_triviaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_multiline_trivia
	return p
}

func InitEmptyMultiline_triviaContext(p *Multiline_triviaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_multiline_trivia
}

func (*Multiline_triviaContext) IsMultiline_triviaContext() {}

func NewMultiline_triviaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multiline_triviaContext {
	var p = new(Multiline_triviaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_multiline_trivia

	return p
}

func (s *Multiline_triviaContext) GetParser() antlr.Parser { return s.parser }

func (s *Multiline_triviaContext) NL() antlr.TerminalNode {
	return s.GetToken(YiniParserNL, 0)
}

func (s *Multiline_triviaContext) DISABLED_LINE() antlr.TerminalNode {
	return s.GetToken(YiniParserDISABLED_LINE, 0)
}

func (s *Multiline_triviaContext) FULL_LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(YiniParserFULL_LINE_COMMENT, 0)
}

func (s *Multiline_triviaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multiline_triviaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multiline_triviaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitMultiline_trivia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Multiline_trivia() (localctx IMultiline_triviaContext) {
	localctx = NewMultiline_triviaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, YiniParserRULE_multiline_trivia)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107374182400) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumber_literalContext is an interface to support dynamic dispatch.
type INumber_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsNumber_literalContext differentiates from other interfaces.
	IsNumber_literalContext()
}

type Number_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_literalContext() *Number_literalContext {
	var p = new(Number_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_number_literal
	return p
}

func InitEmptyNumber_literalContext(p *Number_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_number_literal
}

func (*Number_literalContext) IsNumber_literalContext() {}

func NewNumber_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_literalContext {
	var p = new(Number_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_number_literal

	return p
}

func (s *Number_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(YiniParserNUMBER, 0)
}

func (s *Number_literalContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YiniParserPLUS, 0)
}

func (s *Number_literalContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YiniParserMINUS, 0)
}

func (s *Number_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitNumber_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Number_literal() (localctx INumber_literalContext) {
	localctx = NewNumber_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, YiniParserRULE_number_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserPLUS || _la == YiniParserMINUS {
		{
			p.SetState(316)
			_la = p.GetTokenStream().LA(1)

			if !(_la == YiniParserPLUS || _la == YiniParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(319)
		p.Match(YiniParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INull_literalContext is an interface to support dynamic dispatch.
type INull_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NULL() antlr.TerminalNode

	// IsNull_literalContext differentiates from other interfaces.
	IsNull_literalContext()
}

type Null_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNull_literalContext() *Null_literalContext {
	var p = new(Null_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_null_literal
	return p
}

func InitEmptyNull_literalContext(p *Null_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_null_literal
}

func (*Null_literalContext) IsNull_literalContext() {}

func NewNull_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Null_literalContext {
	var p = new(Null_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_null_literal

	return p
}

func (s *Null_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Null_literalContext) NULL() antlr.TerminalNode {
	return s.GetToken(YiniParserNULL, 0)
}

func (s *Null_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Null_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Null_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitNull_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Null_literal() (localctx INull_literalContext) {
	localctx = NewNull_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, YiniParserRULE_null_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(YiniParserNULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_literalContext is an interface to support dynamic dispatch.
type IBoolean_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN_TRUE() antlr.TerminalNode
	BOOLEAN_FALSE() antlr.TerminalNode

	// IsBoolean_literalContext differentiates from other interfaces.
	IsBoolean_literalContext()
}

type Boolean_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_literalContext() *Boolean_literalContext {
	var p = new(Boolean_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_boolean_literal
	return p
}

func InitEmptyBoolean_literalContext(p *Boolean_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_boolean_literal
}

func (*Boolean_literalContext) IsBoolean_literalContext() {}

func NewBoolean_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_literalContext {
	var p = new(Boolean_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_boolean_literal

	return p
}

func (s *Boolean_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_literalContext) BOOLEAN_TRUE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_TRUE, 0)
}

func (s *Boolean_literalContext) BOOLEAN_FALSE() antlr.TerminalNode {
	return s.GetToken(YiniParserBOOLEAN_FALSE, 0)
}

func (s *Boolean_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitBoolean_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Boolean_literal() (localctx IBoolean_literalContext) {
	localctx = NewBoolean_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, YiniParserRULE_boolean_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		_la = p.GetTokenStream().LA(1)

		if !(_la == YiniParserBOOLEAN_FALSE || _la == YiniParserBOOLEAN_TRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBad_meta_textContext is an interface to support dynamic dispatch.
type IBad_meta_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	META_INVALID() antlr.TerminalNode

	// IsBad_meta_textContext differentiates from other interfaces.
	IsBad_meta_textContext()
}

type Bad_meta_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_meta_textContext() *Bad_meta_textContext {
	var p = new(Bad_meta_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_bad_meta_text
	return p
}

func InitEmptyBad_meta_textContext(p *Bad_meta_textContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_bad_meta_text
}

func (*Bad_meta_textContext) IsBad_meta_textContext() {}

func NewBad_meta_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_meta_textContext {
	var p = new(Bad_meta_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_bad_meta_text

	return p
}

func (s *Bad_meta_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_meta_textContext) META_INVALID() antlr.TerminalNode {
	return s.GetToken(YiniParserMETA_INVALID, 0)
}

func (s *Bad_meta_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_meta_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_meta_textContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitBad_meta_text(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Bad_meta_text() (localctx IBad_meta_textContext) {
	localctx = NewBad_meta_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, YiniParserRULE_bad_meta_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(YiniParserMETA_INVALID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBad_memberContext is an interface to support dynamic dispatch.
type IBad_memberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllREST() []antlr.TerminalNode
	REST(i int) antlr.TerminalNode
	Eol() IEolContext

	// IsBad_memberContext differentiates from other interfaces.
	IsBad_memberContext()
}

type Bad_memberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_memberContext() *Bad_memberContext {
	var p = new(Bad_memberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_bad_member
	return p
}

func InitEmptyBad_memberContext(p *Bad_memberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YiniParserRULE_bad_member
}

func (*Bad_memberContext) IsBad_memberContext() {}

func NewBad_memberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_memberContext {
	var p = new(Bad_memberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YiniParserRULE_bad_member

	return p
}

func (s *Bad_memberContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_memberContext) EQ() antlr.TerminalNode {
	return s.GetToken(YiniParserEQ, 0)
}

func (s *Bad_memberContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *Bad_memberContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Bad_memberContext) AllREST() []antlr.TerminalNode {
	return s.GetTokens(YiniParserREST)
}

func (s *Bad_memberContext) REST(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserREST, i)
}

func (s *Bad_memberContext) Eol() IEolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Bad_memberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_memberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_memberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YiniParserVisitor:
		return t.VisitBad_member(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YiniParser) Bad_member() (localctx IBad_memberContext) {
	localctx = NewBad_memberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, YiniParserRULE_bad_member)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case YiniParserREST:
		{
			p.SetState(327)
			p.Match(YiniParserREST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE, YiniParserNULL, YiniParserEMPTY_OBJECT, YiniParserEMPTY_LIST, YiniParserSTRING, YiniParserNUMBER, YiniParserOB, YiniParserOC, YiniParserPLUS, YiniParserMINUS:
		{
			p.SetState(328)
			p.Value()
		}

	case YiniParserEQ:

	default:
	}
	{
		p.SetState(331)
		p.Match(YiniParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE, YiniParserNULL, YiniParserEMPTY_OBJECT, YiniParserEMPTY_LIST, YiniParserSTRING, YiniParserNUMBER, YiniParserOB, YiniParserOC, YiniParserPLUS, YiniParserMINUS:
		{
			p.SetState(332)
			p.Value()
		}

	case YiniParserREST:
		{
			p.SetState(333)
			p.Match(YiniParserREST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(336)
			p.Eol()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
