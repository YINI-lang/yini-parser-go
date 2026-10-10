// Code generated from grammar/v1.0.0-rc.6/YiniParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"'{'", "'}'", "'+'", "'$'", "'%'", "'@'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "SHEBANG", "YINI_TOKEN", "INCLUDE_TOKEN", "DEPRECATED_TOKEN", "TERMINAL_TOKEN",
		"SECTION_HEAD", "INVALID_SECTION_HEAD", "BOOLEAN_FALSE", "BOOLEAN_TRUE",
		"NULL", "EMPTY_OBJECT", "EMPTY_LIST", "STRING", "NUMBER", "CARET", "SS",
		"GT", "LT", "EQ", "COMMA", "COLON", "OB", "CB", "OC", "CC", "PLUS",
		"DOLLAR", "PC", "AT", "SEMICOLON", "NL", "WS", "BLOCK_COMMENT", "DISABLED_LINE",
		"FULL_LINE_COMMENT", "INLINE_COMMENT", "KEY", "IDENT_INVALID", "REST",
		"META_INVALID",
	}
	staticData.RuleNames = []string{
		"yini", "prolog", "terminal_stmt", "terminal_trivia", "stmt", "full_line_comment_stmt",
		"disabled_line_stmt", "invalid_section_stmt", "meta_stmt", "directive",
		"yini_directive", "yini_mode_declaration", "annotation", "eol", "assignment",
		"member", "value", "scalar_value", "concat_expression", "concat_tail",
		"concat_operand", "string_literal", "object_literal", "object_members",
		"object_member", "object_member_separator", "list_literal", "elements",
		"number_literal", "null_literal", "boolean_literal", "bad_meta_text",
		"bad_member",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 324, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 3, 0, 68, 8, 0, 1, 0, 5, 0, 71, 8, 0, 10, 0, 12,
		0, 74, 9, 0, 1, 0, 3, 0, 77, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 83, 8,
		1, 10, 1, 12, 1, 86, 9, 1, 1, 1, 4, 1, 89, 8, 1, 11, 1, 12, 1, 90, 3, 1,
		93, 8, 1, 1, 2, 1, 2, 5, 2, 97, 8, 2, 10, 2, 12, 2, 100, 9, 2, 1, 3, 1,
		3, 3, 3, 104, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 114, 8, 4, 1, 5, 1, 5, 3, 5, 118, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 129, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 134, 8,
		9, 1, 9, 3, 9, 137, 8, 9, 1, 10, 1, 10, 3, 10, 141, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 151, 8, 13, 11, 13, 12,
		13, 152, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 161, 8, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 167, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		3, 17, 173, 8, 17, 1, 18, 1, 18, 4, 18, 177, 8, 18, 11, 18, 12, 18, 178,
		1, 19, 1, 19, 5, 19, 183, 8, 19, 10, 19, 12, 19, 186, 9, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 196, 8, 22, 10, 22,
		12, 22, 199, 9, 22, 1, 22, 3, 22, 202, 8, 22, 1, 22, 5, 22, 205, 8, 22,
		10, 22, 12, 22, 208, 9, 22, 1, 22, 1, 22, 5, 22, 212, 8, 22, 10, 22, 12,
		22, 215, 9, 22, 1, 22, 1, 22, 5, 22, 219, 8, 22, 10, 22, 12, 22, 222, 9,
		22, 3, 22, 224, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 229, 8, 23, 10, 23,
		12, 23, 232, 9, 23, 1, 23, 5, 23, 235, 8, 23, 10, 23, 12, 23, 238, 9, 23,
		1, 23, 3, 23, 241, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		26, 1, 26, 5, 26, 251, 8, 26, 10, 26, 12, 26, 254, 9, 26, 1, 26, 3, 26,
		257, 8, 26, 1, 26, 5, 26, 260, 8, 26, 10, 26, 12, 26, 263, 9, 26, 1, 26,
		1, 26, 5, 26, 267, 8, 26, 10, 26, 12, 26, 270, 9, 26, 1, 26, 1, 26, 5,
		26, 274, 8, 26, 10, 26, 12, 26, 277, 9, 26, 3, 26, 279, 8, 26, 1, 27, 1,
		27, 5, 27, 283, 8, 27, 10, 27, 12, 27, 286, 9, 27, 1, 27, 1, 27, 5, 27,
		290, 8, 27, 10, 27, 12, 27, 293, 9, 27, 1, 27, 5, 27, 296, 8, 27, 10, 27,
		12, 27, 299, 9, 27, 1, 27, 3, 27, 302, 8, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 3, 32, 314, 8, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 319, 8, 32, 1, 32, 3, 32, 322, 8, 32, 1, 32, 0, 0, 33,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 3, 2, 0, 8,
		10, 13, 14, 2, 0, 19, 19, 21, 21, 1, 0, 8, 9, 344, 0, 67, 1, 0, 0, 0, 2,
		92, 1, 0, 0, 0, 4, 94, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0, 8, 113, 1, 0, 0,
		0, 10, 115, 1, 0, 0, 0, 12, 119, 1, 0, 0, 0, 14, 121, 1, 0, 0, 0, 16, 128,
		1, 0, 0, 0, 18, 136, 1, 0, 0, 0, 20, 138, 1, 0, 0, 0, 22, 144, 1, 0, 0,
		0, 24, 146, 1, 0, 0, 0, 26, 150, 1, 0, 0, 0, 28, 154, 1, 0, 0, 0, 30, 157,
		1, 0, 0, 0, 32, 166, 1, 0, 0, 0, 34, 172, 1, 0, 0, 0, 36, 174, 1, 0, 0,
		0, 38, 180, 1, 0, 0, 0, 40, 189, 1, 0, 0, 0, 42, 191, 1, 0, 0, 0, 44, 223,
		1, 0, 0, 0, 46, 225, 1, 0, 0, 0, 48, 242, 1, 0, 0, 0, 50, 246, 1, 0, 0,
		0, 52, 278, 1, 0, 0, 0, 54, 280, 1, 0, 0, 0, 56, 303, 1, 0, 0, 0, 58, 305,
		1, 0, 0, 0, 60, 307, 1, 0, 0, 0, 62, 309, 1, 0, 0, 0, 64, 313, 1, 0, 0,
		0, 66, 68, 3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 72,
		1, 0, 0, 0, 69, 71, 3, 8, 4, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0,
		72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1,
		0, 0, 0, 75, 77, 3, 4, 2, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77,
		78, 1, 0, 0, 0, 78, 79, 5, 0, 0, 1, 79, 1, 1, 0, 0, 0, 80, 84, 5, 1, 0,
		0, 81, 83, 3, 26, 13, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 93, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		87, 89, 3, 26, 13, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 80, 1, 0, 0, 0, 92,
		88, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0, 94, 98, 5, 5, 0, 0, 95, 97, 3, 6, 3,
		0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 5, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 104, 3, 26, 13,
		0, 102, 104, 3, 10, 5, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104,
		7, 1, 0, 0, 0, 105, 114, 3, 26, 13, 0, 106, 114, 3, 10, 5, 0, 107, 114,
		3, 12, 6, 0, 108, 114, 5, 6, 0, 0, 109, 114, 3, 14, 7, 0, 110, 114, 3,
		28, 14, 0, 111, 114, 3, 16, 8, 0, 112, 114, 3, 64, 32, 0, 113, 105, 1,
		0, 0, 0, 113, 106, 1, 0, 0, 0, 113, 107, 1, 0, 0, 0, 113, 108, 1, 0, 0,
		0, 113, 109, 1, 0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		112, 1, 0, 0, 0, 114, 9, 1, 0, 0, 0, 115, 117, 5, 35, 0, 0, 116, 118, 3,
		26, 13, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 11, 1, 0, 0,
		0, 119, 120, 5, 34, 0, 0, 120, 13, 1, 0, 0, 0, 121, 122, 5, 7, 0, 0, 122,
		15, 1, 0, 0, 0, 123, 129, 3, 18, 9, 0, 124, 129, 3, 24, 12, 0, 125, 126,
		3, 62, 31, 0, 126, 127, 3, 26, 13, 0, 127, 129, 1, 0, 0, 0, 128, 123, 1,
		0, 0, 0, 128, 124, 1, 0, 0, 0, 128, 125, 1, 0, 0, 0, 129, 17, 1, 0, 0,
		0, 130, 137, 3, 20, 10, 0, 131, 133, 5, 3, 0, 0, 132, 134, 3, 42, 21, 0,
		133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		137, 3, 26, 13, 0, 136, 130, 1, 0, 0, 0, 136, 131, 1, 0, 0, 0, 137, 19,
		1, 0, 0, 0, 138, 140, 5, 2, 0, 0, 139, 141, 3, 22, 11, 0, 140, 139, 1,
		0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 3, 26, 13,
		0, 143, 21, 1, 0, 0, 0, 144, 145, 5, 37, 0, 0, 145, 23, 1, 0, 0, 0, 146,
		147, 5, 4, 0, 0, 147, 148, 3, 26, 13, 0, 148, 25, 1, 0, 0, 0, 149, 151,
		5, 31, 0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 152, 153, 1, 0, 0, 0, 153, 27, 1, 0, 0, 0, 154, 155, 3, 30, 15, 0,
		155, 156, 3, 26, 13, 0, 156, 29, 1, 0, 0, 0, 157, 158, 5, 37, 0, 0, 158,
		160, 5, 19, 0, 0, 159, 161, 3, 32, 16, 0, 160, 159, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 31, 1, 0, 0, 0, 162, 167, 3, 36, 18, 0, 163, 167, 3, 34,
		17, 0, 164, 167, 3, 52, 26, 0, 165, 167, 3, 44, 22, 0, 166, 162, 1, 0,
		0, 0, 166, 163, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 165, 1, 0, 0, 0,
		167, 33, 1, 0, 0, 0, 168, 173, 3, 58, 29, 0, 169, 173, 3, 42, 21, 0, 170,
		173, 3, 56, 28, 0, 171, 173, 3, 60, 30, 0, 172, 168, 1, 0, 0, 0, 172, 169,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 35, 1, 0,
		0, 0, 174, 176, 5, 13, 0, 0, 175, 177, 3, 38, 19, 0, 176, 175, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		37, 1, 0, 0, 0, 180, 184, 5, 26, 0, 0, 181, 183, 5, 31, 0, 0, 182, 181,
		1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0,
		0, 0, 185, 187, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 188, 3, 40, 20,
		0, 188, 39, 1, 0, 0, 0, 189, 190, 7, 0, 0, 0, 190, 41, 1, 0, 0, 0, 191,
		192, 5, 13, 0, 0, 192, 43, 1, 0, 0, 0, 193, 197, 5, 24, 0, 0, 194, 196,
		5, 31, 0, 0, 195, 194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		200, 202, 3, 46, 23, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202,
		206, 1, 0, 0, 0, 203, 205, 5, 31, 0, 0, 204, 203, 1, 0, 0, 0, 205, 208,
		1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0,
		0, 0, 208, 206, 1, 0, 0, 0, 209, 213, 5, 25, 0, 0, 210, 212, 5, 31, 0,
		0, 211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 224, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 220,
		5, 11, 0, 0, 217, 219, 5, 31, 0, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1,
		0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 224, 1, 0, 0,
		0, 222, 220, 1, 0, 0, 0, 223, 193, 1, 0, 0, 0, 223, 216, 1, 0, 0, 0, 224,
		45, 1, 0, 0, 0, 225, 236, 3, 48, 24, 0, 226, 230, 5, 20, 0, 0, 227, 229,
		5, 31, 0, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0,
		0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		233, 235, 3, 48, 24, 0, 234, 226, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236,
		234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236,
		1, 0, 0, 0, 239, 241, 5, 20, 0, 0, 240, 239, 1, 0, 0, 0, 240, 241, 1, 0,
		0, 0, 241, 47, 1, 0, 0, 0, 242, 243, 5, 37, 0, 0, 243, 244, 3, 50, 25,
		0, 244, 245, 3, 32, 16, 0, 245, 49, 1, 0, 0, 0, 246, 247, 7, 1, 0, 0, 247,
		51, 1, 0, 0, 0, 248, 252, 5, 22, 0, 0, 249, 251, 5, 31, 0, 0, 250, 249,
		1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0,
		0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 257, 3, 54, 27,
		0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 261, 1, 0, 0, 0, 258,
		260, 5, 31, 0, 0, 259, 258, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259,
		1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 261, 1, 0,
		0, 0, 264, 268, 5, 23, 0, 0, 265, 267, 5, 31, 0, 0, 266, 265, 1, 0, 0,
		0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269,
		279, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 275, 5, 12, 0, 0, 272, 274,
		5, 31, 0, 0, 273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 275, 276, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0,
		278, 248, 1, 0, 0, 0, 278, 271, 1, 0, 0, 0, 279, 53, 1, 0, 0, 0, 280, 297,
		3, 32, 16, 0, 281, 283, 5, 31, 0, 0, 282, 281, 1, 0, 0, 0, 283, 286, 1,
		0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 1, 0, 0,
		0, 286, 284, 1, 0, 0, 0, 287, 291, 5, 20, 0, 0, 288, 290, 5, 31, 0, 0,
		289, 288, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 296,
		3, 32, 16, 0, 295, 284, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1,
		0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0,
		0, 300, 302, 5, 20, 0, 0, 301, 300, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302,
		55, 1, 0, 0, 0, 303, 304, 5, 14, 0, 0, 304, 57, 1, 0, 0, 0, 305, 306, 5,
		10, 0, 0, 306, 59, 1, 0, 0, 0, 307, 308, 7, 2, 0, 0, 308, 61, 1, 0, 0,
		0, 309, 310, 5, 40, 0, 0, 310, 63, 1, 0, 0, 0, 311, 314, 5, 39, 0, 0, 312,
		314, 3, 32, 16, 0, 313, 311, 1, 0, 0, 0, 313, 312, 1, 0, 0, 0, 313, 314,
		1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 318, 5, 19, 0, 0, 316, 319, 3, 32,
		16, 0, 317, 319, 5, 39, 0, 0, 318, 316, 1, 0, 0, 0, 318, 317, 1, 0, 0,
		0, 319, 321, 1, 0, 0, 0, 320, 322, 3, 26, 13, 0, 321, 320, 1, 0, 0, 0,
		321, 322, 1, 0, 0, 0, 322, 65, 1, 0, 0, 0, 42, 67, 72, 76, 84, 90, 92,
		98, 103, 113, 117, 128, 133, 136, 140, 152, 160, 166, 172, 178, 184, 197,
		201, 206, 213, 220, 223, 230, 236, 240, 252, 256, 261, 268, 275, 278, 284,
		291, 297, 301, 313, 318, 321,
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
	YiniParserDOLLAR               = 27
	YiniParserPC                   = 28
	YiniParserAT                   = 29
	YiniParserSEMICOLON            = 30
	YiniParserNL                   = 31
	YiniParserWS                   = 32
	YiniParserBLOCK_COMMENT        = 33
	YiniParserDISABLED_LINE        = 34
	YiniParserFULL_LINE_COMMENT    = 35
	YiniParserINLINE_COMMENT       = 36
	YiniParserKEY                  = 37
	YiniParserIDENT_INVALID        = 38
	YiniParserREST                 = 39
	YiniParserMETA_INVALID         = 40
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
	YiniParserRULE_assignment              = 14
	YiniParserRULE_member                  = 15
	YiniParserRULE_value                   = 16
	YiniParserRULE_scalar_value            = 17
	YiniParserRULE_concat_expression       = 18
	YiniParserRULE_concat_tail             = 19
	YiniParserRULE_concat_operand          = 20
	YiniParserRULE_string_literal          = 21
	YiniParserRULE_object_literal          = 22
	YiniParserRULE_object_members          = 23
	YiniParserRULE_object_member           = 24
	YiniParserRULE_object_member_separator = 25
	YiniParserRULE_list_literal            = 26
	YiniParserRULE_elements                = 27
	YiniParserRULE_number_literal          = 28
	YiniParserRULE_null_literal            = 29
	YiniParserRULE_boolean_literal         = 30
	YiniParserRULE_bad_meta_text           = 31
	YiniParserRULE_bad_member              = 32
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(66)
			p.Prolog()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1840415014876) != 0 {
		{
			p.SetState(69)
			p.Stmt()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserTERMINAL_TOKEN {
		{
			p.SetState(75)
			p.Terminal_stmt()
		}

	}
	{
		p.SetState(78)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserSHEBANG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(YiniParserSHEBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(84)
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
					p.SetState(81)
					p.Eol()
				}

			}
			p.SetState(86)
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
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(87)
					p.Eol()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(90)
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
		p.SetState(94)
		p.Match(YiniParserTERMINAL_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YiniParserNL || _la == YiniParserFULL_LINE_COMMENT {
		{
			p.SetState(95)
			p.Terminal_trivia()
		}

		p.SetState(100)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Eol()
		}

	case YiniParserFULL_LINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Eol()
		}

	case YiniParserFULL_LINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Full_line_comment_stmt()
		}

	case YiniParserDISABLED_LINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.Disabled_line_stmt()
		}

	case YiniParserSECTION_HEAD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.Match(YiniParserSECTION_HEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserINVALID_SECTION_HEAD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.Invalid_section_stmt()
		}

	case YiniParserKEY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.Assignment()
		}

	case YiniParserYINI_TOKEN, YiniParserINCLUDE_TOKEN, YiniParserDEPRECATED_TOKEN, YiniParserMETA_INVALID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.Meta_stmt()
		}

	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE, YiniParserNULL, YiniParserEMPTY_OBJECT, YiniParserEMPTY_LIST, YiniParserSTRING, YiniParserNUMBER, YiniParserEQ, YiniParserOB, YiniParserOC, YiniParserREST:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(112)
			p.Bad_member()
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
		p.SetState(115)
		p.Match(YiniParserFULL_LINE_COMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
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
		p.SetState(119)
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
		p.SetState(121)
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
	Eol() IEolContext

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

func (s *Meta_stmtContext) Eol() IEolContext {
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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserYINI_TOKEN, YiniParserINCLUDE_TOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Directive()
		}

	case YiniParserDEPRECATED_TOKEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Annotation()
		}

	case YiniParserMETA_INVALID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Bad_meta_text()
		}
		{
			p.SetState(126)
			p.Eol()
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
	Eol() IEolContext
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

func (s *DirectiveContext) Eol() IEolContext {
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

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserYINI_TOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Yini_directive()
		}

	case YiniParserINCLUDE_TOKEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Match(YiniParserINCLUDE_TOKEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YiniParserSTRING {
			{
				p.SetState(132)
				p.String_literal()
			}

		}
		{
			p.SetState(135)
			p.Eol()
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
	Eol() IEolContext
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

func (s *Yini_directiveContext) Eol() IEolContext {
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
		p.SetState(138)
		p.Match(YiniParserYINI_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserKEY {
		{
			p.SetState(139)
			p.Yini_mode_declaration()
		}

	}
	{
		p.SetState(142)
		p.Eol()
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
		p.SetState(144)
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
	Eol() IEolContext

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

func (s *AnnotationContext) Eol() IEolContext {
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
		p.SetState(146)
		p.Match(YiniParserDEPRECATED_TOKEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Eol()
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(149)
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

		p.SetState(152)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Member() IMemberContext
	Eol() IEolContext

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

func (s *AssignmentContext) Eol() IEolContext {
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
	p.EnterRule(localctx, 28, YiniParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Member()
	}
	{
		p.SetState(155)
		p.Eol()
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
	KEY() antlr.TerminalNode
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

func (s *MemberContext) KEY() antlr.TerminalNode {
	return s.GetToken(YiniParserKEY, 0)
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
	p.EnterRule(localctx, 30, YiniParserRULE_member)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(YiniParserKEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(YiniParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21004032) != 0 {
		{
			p.SetState(159)
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
	p.EnterRule(localctx, 32, YiniParserRULE_value)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Concat_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Scalar_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.List_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(165)
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
	p.EnterRule(localctx, 34, YiniParserRULE_scalar_value)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserNULL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Null_literal()
		}

	case YiniParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.String_literal()
		}

	case YiniParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.Number_literal()
		}

	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)
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
	p.EnterRule(localctx, 36, YiniParserRULE_concat_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(YiniParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == YiniParserPLUS {
		{
			p.SetState(175)
			p.Concat_tail()
		}

		p.SetState(178)
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

// IConcat_tailContext is an interface to support dynamic dispatch.
type IConcat_tailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	Concat_operand() IConcat_operandContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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

func (s *Concat_tailContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *Concat_tailContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
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
	p.EnterRule(localctx, 38, YiniParserRULE_concat_tail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(YiniParserPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YiniParserNL {
		{
			p.SetState(181)
			p.Match(YiniParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
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
	NUMBER() antlr.TerminalNode
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

func (s *Concat_operandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(YiniParserNUMBER, 0)
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
	p.EnterRule(localctx, 40, YiniParserRULE_concat_operand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26368) != 0) {
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
	p.EnterRule(localctx, 42, YiniParserRULE_string_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
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
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
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

func (s *Object_literalContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *Object_literalContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
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
	p.EnterRule(localctx, 44, YiniParserRULE_object_literal)
	var _la int

	var _alt int

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserOC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(YiniParserOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(194)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YiniParserKEY {
			{
				p.SetState(200)
				p.Object_members()
			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == YiniParserNL {
			{
				p.SetState(203)
				p.Match(YiniParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(209)
			p.Match(YiniParserCC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(210)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case YiniParserEMPTY_OBJECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(YiniParserEMPTY_OBJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(217)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
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
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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

func (s *Object_membersContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *Object_membersContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
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
	p.EnterRule(localctx, 46, YiniParserRULE_object_members)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Object_member()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(226)
				p.Match(YiniParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == YiniParserNL {
				{
					p.SetState(227)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(232)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(233)
				p.Object_member()
			}

		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserCOMMA {
		{
			p.SetState(239)
			p.Match(YiniParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IObject_memberContext is an interface to support dynamic dispatch.
type IObject_memberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEY() antlr.TerminalNode
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

func (s *Object_memberContext) KEY() antlr.TerminalNode {
	return s.GetToken(YiniParserKEY, 0)
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
	p.EnterRule(localctx, 48, YiniParserRULE_object_member)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(YiniParserKEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Object_member_separator()
	}
	{
		p.SetState(244)
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
	p.EnterRule(localctx, 50, YiniParserRULE_object_member_separator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
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
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
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

func (s *List_literalContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *List_literalContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
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
	p.EnterRule(localctx, 52, YiniParserRULE_list_literal)
	var _la int

	var _alt int

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserOB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(YiniParserOB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(249)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21004032) != 0 {
			{
				p.SetState(255)
				p.Elements()
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == YiniParserNL {
			{
				p.SetState(258)
				p.Match(YiniParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(264)
			p.Match(YiniParserCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(265)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case YiniParserEMPTY_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(YiniParserEMPTY_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(272)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
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
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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

func (s *ElementsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(YiniParserNL)
}

func (s *ElementsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(YiniParserNL, i)
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
	p.EnterRule(localctx, 54, YiniParserRULE_elements)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Value()
	}
	p.SetState(297)
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
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == YiniParserNL {
				{
					p.SetState(281)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(286)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(287)
				p.Match(YiniParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == YiniParserNL {
				{
					p.SetState(288)
					p.Match(YiniParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(293)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(294)
				p.Value()
			}

		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YiniParserCOMMA {
		{
			p.SetState(300)
			p.Match(YiniParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 56, YiniParserRULE_number_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
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
	p.EnterRule(localctx, 58, YiniParserRULE_null_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
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
	p.EnterRule(localctx, 60, YiniParserRULE_boolean_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
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
	p.EnterRule(localctx, 62, YiniParserRULE_bad_meta_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
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
	p.EnterRule(localctx, 64, YiniParserRULE_bad_member)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case YiniParserREST:
		{
			p.SetState(311)
			p.Match(YiniParserREST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE, YiniParserNULL, YiniParserEMPTY_OBJECT, YiniParserEMPTY_LIST, YiniParserSTRING, YiniParserNUMBER, YiniParserOB, YiniParserOC:
		{
			p.SetState(312)
			p.Value()
		}

	case YiniParserEQ:

	default:
	}
	{
		p.SetState(315)
		p.Match(YiniParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YiniParserBOOLEAN_FALSE, YiniParserBOOLEAN_TRUE, YiniParserNULL, YiniParserEMPTY_OBJECT, YiniParserEMPTY_LIST, YiniParserSTRING, YiniParserNUMBER, YiniParserOB, YiniParserOC:
		{
			p.SetState(316)
			p.Value()
		}

	case YiniParserREST:
		{
			p.SetState(317)
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
	p.SetState(321)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(320)
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
