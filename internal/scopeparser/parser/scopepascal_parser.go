// Code generated from /home/tomas/development/tomas303/projects/palsp/internal/scopeparser/scopepascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // scopepascal

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

type scopepascalParser struct {
	*antlr.BaseParser
}

var ScopepascalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func scopepascalParserInit() {
	staticData := &ScopepascalParserStaticData
	staticData.LiteralNames = []string{
		"", "'begin'", "'end'", "'function'", "'procedure'", "'var'", "'const'",
		"'out'", "'.'", "'..'", "';'", "','", "':'", "'='", "'<>'", "'<'", "'<='",
		"'>'", "'>='", "'in'", "'*'", "'/'", "'div'", "'mod'", "'and'", "'or'",
		"'not'", "'+'", "'-'", "'shr'", "'shl'", "'['", "']'", "'('", "')'",
		"':='", "'label'", "'type'", "'resourcestring'", "'program'", "'unit'",
		"'interface'", "'implementation'", "'initialization'", "'finalization'",
		"'uses'", "'forward'", "'class'", "'PRIVATE'", "'PROTECTED'", "'PUBLIC'",
		"'PUBLISHED'", "", "", "", "'VIRTUAL'", "'OVERRIDE'", "'REINTRODUCE'",
		"'OVERLOAD'", "'INLINE'", "'STDCALL'", "'CDECL'", "'STATIC'", "'PACKED'",
		"'RECORD'", "'PROPERTY'", "'DEFAULT'", "'abstract'", "'read'", "'write'",
		"'index'", "'array'", "'of'", "'^'", "'@'", "'set'", "'file'", "'case'",
		"", "", "", "", "", "", "", "", "", "", "'\\uFEFF'",
	}
	staticData.SymbolicNames = []string{
		"", "BEGIN", "END", "FUNCTION", "PROCEDURE", "VAR", "CONST", "OUT",
		"DOT", "DOUBLEDOT", "SEMI", "COMMA", "COLON", "EQUAL", "NOT_EQUAL",
		"LT", "LE", "GT", "GE", "IN", "STAR", "SLASH", "DIV", "MOD", "AND",
		"OR", "NOT", "PLUS", "MINUS", "SHR", "SHL", "LBRACK", "RBRACK", "LPAREN",
		"RPAREN", "ASSIGN", "LABEL", "TYPE", "RESOURCESTRING", "PROGRAM", "UNIT",
		"INTERFACE", "IMPLEMENTATION", "INITIALIZATION", "FINALIZATION", "USES",
		"FORWARD", "CLASS", "PRIVATE", "PROTECTED", "PUBLIC", "PUBLISHED", "STRICTPRIVATE",
		"STRICTPROTECTED", "OFOBJECT", "VIRTUAL", "OVERRIDE", "REINTRODUCE",
		"OVERLOAD", "INLINE", "STDCALL", "CDECL", "STATIC", "PACKED", "RECORD",
		"PROPERTY", "DEFAULT", "ABSTRACT", "READ", "WRITE", "INDEX", "ARRAY",
		"OF", "DEREFERENCE", "AT", "SET", "FILE", "CASE", "IDENT", "HEX_LITERAL",
		"STRING_LITERAL", "STRING_CROSSHATCH_LITERAL", "NUM_INT", "NUM_REAL",
		"WS", "COMMENT_1", "COMMENT_2", "COMMENT_3", "UTF8BOM",
	}
	staticData.RuleNames = []string{
		"source", "program", "unit", "interfaceSection", "implementationSection",
		"initializationSection", "finalizationSection", "interfaceBlock", "implementationBlock",
		"unitList", "labelDeclaration", "constSection", "resourceSection", "typeSection",
		"typeBlock", "type", "accessSpecifier", "procedureOrFunctionModifiers",
		"classForwardDeclaration", "classType", "recordType", "propertyDeclaration",
		"propertyReadDeclaration", "propertyWriteDeclaration", "propertyDefaultValueDeclaration",
		"propertyIndexDeclaration", "propertyIndexParameters", "propertyIndexParametersList",
		"arrayType", "pointerType", "setType", "fileType", "scalarType", "subrangeType",
		"blockDeclaration", "functionDeclaration", "procedureDeclaration", "functionOrProcedureDeclaration",
		"functionOrProcedure", "blockStatement", "recordVariantDeclaration",
		"recordVariant", "statementError", "statement", "identifier", "identifierList",
		"expressionList", "paramsDeclaration", "paramsDeclarationSection", "paramSpecifier",
		"varSection", "varDeclaration", "inlinedVarDeclaration", "relationaloperator",
		"additiveoperator", "multiplicativeoperator", "operator", "string",
		"number", "expression", "term", "functionExpression", "procedureExpression",
		"functionDesignator", "errorExpression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 88, 723, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 3, 0, 133, 8, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 141, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 146, 8, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 158, 8,
		2, 1, 2, 3, 2, 161, 8, 2, 1, 2, 3, 2, 164, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 3, 3, 172, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 178, 8, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 194, 8, 7, 10, 7, 12, 7, 197, 9, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 5, 8, 205, 8, 8, 10, 8, 12, 8, 208, 9, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 218, 8, 10, 10, 10, 12, 10,
		221, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 229, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 235, 8, 11, 11, 11, 12, 11, 236, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 245, 8, 12, 11, 12, 12, 12,
		246, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14, 257,
		8, 14, 11, 14, 12, 14, 258, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 271, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		5, 17, 277, 8, 17, 10, 17, 12, 17, 280, 9, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 290, 8, 19, 1, 19, 1, 19, 3, 19,
		294, 8, 19, 1, 19, 3, 19, 297, 8, 19, 1, 19, 5, 19, 300, 8, 19, 10, 19,
		12, 19, 303, 9, 19, 1, 19, 1, 19, 1, 19, 5, 19, 308, 8, 19, 10, 19, 12,
		19, 311, 9, 19, 1, 19, 1, 19, 1, 20, 3, 20, 316, 8, 20, 1, 20, 1, 20, 3,
		20, 320, 8, 20, 1, 20, 1, 20, 1, 20, 5, 20, 325, 8, 20, 10, 20, 12, 20,
		328, 9, 20, 1, 20, 5, 20, 331, 8, 20, 10, 20, 12, 20, 334, 9, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 3, 21, 341, 8, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 346, 8, 21, 1, 21, 3, 21, 349, 8, 21, 1, 21, 3, 21, 352, 8, 21, 1,
		21, 3, 21, 355, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 360, 8, 21, 3, 21, 362,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 388, 8, 27, 10, 27, 12, 27, 391, 9,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 400, 8, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 410, 8,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 419, 8, 28,
		3, 28, 421, 8, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 3, 31, 434, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 448, 8,
		34, 1, 34, 1, 34, 3, 34, 452, 8, 34, 1, 35, 3, 35, 455, 8, 35, 1, 35, 1,
		35, 1, 35, 3, 35, 460, 8, 35, 1, 35, 1, 35, 1, 35, 3, 35, 465, 8, 35, 1,
		35, 1, 35, 1, 35, 1, 36, 3, 36, 471, 8, 36, 1, 36, 1, 36, 1, 36, 3, 36,
		476, 8, 36, 1, 36, 3, 36, 479, 8, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		3, 37, 486, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 491, 8, 38, 10, 38, 12,
		38, 494, 9, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38,
		503, 8, 38, 1, 39, 1, 39, 1, 39, 5, 39, 508, 8, 39, 10, 39, 12, 39, 511,
		9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 519, 8, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 5, 40, 525, 8, 40, 10, 40, 12, 40, 528, 9, 40,
		1, 40, 3, 40, 531, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 4, 41, 537, 8, 41,
		11, 41, 12, 41, 538, 1, 41, 1, 41, 1, 42, 4, 42, 544, 8, 42, 11, 42, 12,
		42, 545, 1, 43, 1, 43, 1, 43, 3, 43, 551, 8, 43, 1, 44, 1, 44, 1, 44, 5,
		44, 556, 8, 44, 10, 44, 12, 44, 559, 9, 44, 1, 45, 1, 45, 1, 45, 5, 45,
		564, 8, 45, 10, 45, 12, 45, 567, 9, 45, 1, 46, 1, 46, 1, 46, 5, 46, 572,
		8, 46, 10, 46, 12, 46, 575, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 581,
		8, 47, 10, 47, 12, 47, 584, 9, 47, 1, 47, 1, 47, 1, 48, 3, 48, 589, 8,
		48, 1, 48, 1, 48, 1, 48, 3, 48, 594, 8, 48, 1, 48, 1, 48, 3, 48, 598, 8,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 607, 8, 48,
		1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 4, 50, 615, 8, 50, 11, 50, 12,
		50, 616, 1, 51, 1, 51, 1, 51, 3, 51, 622, 8, 51, 1, 51, 1, 51, 3, 51, 626,
		8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 632, 8, 52, 1, 52, 1, 52, 3,
		52, 636, 8, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 56, 3, 56, 649, 8, 56, 1, 57, 4, 57, 652, 8, 57, 11, 57,
		12, 57, 653, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 662, 8, 59,
		10, 59, 12, 59, 665, 9, 59, 1, 59, 1, 59, 3, 59, 669, 8, 59, 1, 60, 3,
		60, 672, 8, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 60, 3, 60, 683, 8, 60, 1, 61, 1, 61, 3, 61, 687, 8, 61, 1, 61, 1, 61,
		3, 61, 691, 8, 61, 1, 61, 1, 61, 1, 62, 1, 62, 3, 62, 697, 8, 62, 1, 62,
		3, 62, 700, 8, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 3,
		63, 709, 8, 63, 1, 63, 1, 63, 3, 63, 713, 8, 63, 1, 63, 3, 63, 716, 8,
		63, 1, 64, 4, 64, 719, 8, 64, 11, 64, 12, 64, 720, 1, 64, 0, 0, 65, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108,
		110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 0, 10, 1, 0, 48, 53,
		2, 0, 55, 62, 67, 67, 2, 0, 1, 2, 10, 10, 1, 0, 5, 7, 1, 0, 13, 19, 2,
		0, 25, 25, 27, 28, 2, 0, 20, 24, 29, 30, 1, 0, 80, 81, 1, 0, 82, 83, 3,
		0, 1, 2, 10, 11, 34, 34, 774, 0, 132, 1, 0, 0, 0, 2, 134, 1, 0, 0, 0, 4,
		152, 1, 0, 0, 0, 6, 169, 1, 0, 0, 0, 8, 175, 1, 0, 0, 0, 10, 181, 1, 0,
		0, 0, 12, 184, 1, 0, 0, 0, 14, 195, 1, 0, 0, 0, 16, 206, 1, 0, 0, 0, 18,
		209, 1, 0, 0, 0, 20, 213, 1, 0, 0, 0, 22, 224, 1, 0, 0, 0, 24, 238, 1,
		0, 0, 0, 26, 248, 1, 0, 0, 0, 28, 256, 1, 0, 0, 0, 30, 270, 1, 0, 0, 0,
		32, 272, 1, 0, 0, 0, 34, 278, 1, 0, 0, 0, 36, 281, 1, 0, 0, 0, 38, 284,
		1, 0, 0, 0, 40, 315, 1, 0, 0, 0, 42, 361, 1, 0, 0, 0, 44, 363, 1, 0, 0,
		0, 46, 366, 1, 0, 0, 0, 48, 369, 1, 0, 0, 0, 50, 372, 1, 0, 0, 0, 52, 375,
		1, 0, 0, 0, 54, 379, 1, 0, 0, 0, 56, 420, 1, 0, 0, 0, 58, 422, 1, 0, 0,
		0, 60, 425, 1, 0, 0, 0, 62, 433, 1, 0, 0, 0, 64, 435, 1, 0, 0, 0, 66, 439,
		1, 0, 0, 0, 68, 451, 1, 0, 0, 0, 70, 454, 1, 0, 0, 0, 72, 470, 1, 0, 0,
		0, 74, 485, 1, 0, 0, 0, 76, 502, 1, 0, 0, 0, 78, 504, 1, 0, 0, 0, 80, 514,
		1, 0, 0, 0, 82, 532, 1, 0, 0, 0, 84, 543, 1, 0, 0, 0, 86, 550, 1, 0, 0,
		0, 88, 552, 1, 0, 0, 0, 90, 560, 1, 0, 0, 0, 92, 568, 1, 0, 0, 0, 94, 576,
		1, 0, 0, 0, 96, 606, 1, 0, 0, 0, 98, 608, 1, 0, 0, 0, 100, 610, 1, 0, 0,
		0, 102, 618, 1, 0, 0, 0, 104, 627, 1, 0, 0, 0, 106, 639, 1, 0, 0, 0, 108,
		641, 1, 0, 0, 0, 110, 643, 1, 0, 0, 0, 112, 648, 1, 0, 0, 0, 114, 651,
		1, 0, 0, 0, 116, 655, 1, 0, 0, 0, 118, 668, 1, 0, 0, 0, 120, 682, 1, 0,
		0, 0, 122, 684, 1, 0, 0, 0, 124, 694, 1, 0, 0, 0, 126, 703, 1, 0, 0, 0,
		128, 718, 1, 0, 0, 0, 130, 133, 3, 2, 1, 0, 131, 133, 3, 4, 2, 0, 132,
		130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 1, 1, 0, 0, 0, 134, 135, 5,
		39, 0, 0, 135, 140, 3, 88, 44, 0, 136, 137, 5, 33, 0, 0, 137, 138, 3, 90,
		45, 0, 138, 139, 5, 34, 0, 0, 139, 141, 1, 0, 0, 0, 140, 136, 1, 0, 0,
		0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 145, 5, 10, 0, 0, 143,
		144, 5, 45, 0, 0, 144, 146, 3, 90, 45, 0, 145, 143, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 3, 16, 8, 0, 148, 149, 3, 78,
		39, 0, 149, 150, 5, 8, 0, 0, 150, 151, 5, 0, 0, 1, 151, 3, 1, 0, 0, 0,
		152, 153, 5, 40, 0, 0, 153, 154, 3, 88, 44, 0, 154, 155, 5, 10, 0, 0, 155,
		157, 3, 6, 3, 0, 156, 158, 3, 8, 4, 0, 157, 156, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 161, 3, 10, 5, 0, 160, 159, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 164, 3, 12, 6, 0,
		163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		166, 5, 2, 0, 0, 166, 167, 5, 8, 0, 0, 167, 168, 5, 0, 0, 1, 168, 5, 1,
		0, 0, 0, 169, 171, 5, 41, 0, 0, 170, 172, 3, 18, 9, 0, 171, 170, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 3, 14, 7, 0,
		174, 7, 1, 0, 0, 0, 175, 177, 5, 42, 0, 0, 176, 178, 3, 18, 9, 0, 177,
		176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180,
		3, 16, 8, 0, 180, 9, 1, 0, 0, 0, 181, 182, 5, 43, 0, 0, 182, 183, 3, 78,
		39, 0, 183, 11, 1, 0, 0, 0, 184, 185, 5, 44, 0, 0, 185, 186, 3, 78, 39,
		0, 186, 13, 1, 0, 0, 0, 187, 194, 3, 26, 13, 0, 188, 194, 3, 20, 10, 0,
		189, 194, 3, 22, 11, 0, 190, 194, 3, 24, 12, 0, 191, 194, 3, 100, 50, 0,
		192, 194, 3, 74, 37, 0, 193, 187, 1, 0, 0, 0, 193, 188, 1, 0, 0, 0, 193,
		189, 1, 0, 0, 0, 193, 190, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 192,
		1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 15, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 205, 3, 20, 10, 0,
		199, 205, 3, 22, 11, 0, 200, 205, 3, 24, 12, 0, 201, 205, 3, 26, 13, 0,
		202, 205, 3, 100, 50, 0, 203, 205, 3, 76, 38, 0, 204, 198, 1, 0, 0, 0,
		204, 199, 1, 0, 0, 0, 204, 200, 1, 0, 0, 0, 204, 201, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204,
		1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 17, 1, 0, 0, 0, 208, 206, 1, 0,
		0, 0, 209, 210, 5, 45, 0, 0, 210, 211, 3, 90, 45, 0, 211, 212, 5, 10, 0,
		0, 212, 19, 1, 0, 0, 0, 213, 214, 5, 36, 0, 0, 214, 219, 3, 116, 58, 0,
		215, 216, 5, 11, 0, 0, 216, 218, 3, 116, 58, 0, 217, 215, 1, 0, 0, 0, 218,
		221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222,
		1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 5, 10, 0, 0, 223, 21, 1, 0,
		0, 0, 224, 234, 5, 6, 0, 0, 225, 228, 3, 90, 45, 0, 226, 227, 5, 12, 0,
		0, 227, 229, 3, 88, 44, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 230, 1, 0, 0, 0, 230, 231, 5, 13, 0, 0, 231, 232, 3, 118, 59, 0, 232,
		233, 5, 10, 0, 0, 233, 235, 1, 0, 0, 0, 234, 225, 1, 0, 0, 0, 235, 236,
		1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 23, 1, 0,
		0, 0, 238, 244, 5, 38, 0, 0, 239, 240, 3, 88, 44, 0, 240, 241, 5, 13, 0,
		0, 241, 242, 3, 118, 59, 0, 242, 243, 5, 10, 0, 0, 243, 245, 1, 0, 0, 0,
		244, 239, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246,
		247, 1, 0, 0, 0, 247, 25, 1, 0, 0, 0, 248, 249, 5, 37, 0, 0, 249, 250,
		3, 28, 14, 0, 250, 27, 1, 0, 0, 0, 251, 252, 3, 88, 44, 0, 252, 253, 5,
		13, 0, 0, 253, 254, 3, 30, 15, 0, 254, 255, 5, 10, 0, 0, 255, 257, 1, 0,
		0, 0, 256, 251, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0,
		258, 259, 1, 0, 0, 0, 259, 29, 1, 0, 0, 0, 260, 271, 3, 38, 19, 0, 261,
		271, 3, 40, 20, 0, 262, 271, 3, 56, 28, 0, 263, 271, 3, 58, 29, 0, 264,
		271, 3, 60, 30, 0, 265, 271, 3, 62, 31, 0, 266, 271, 3, 64, 32, 0, 267,
		271, 3, 66, 33, 0, 268, 271, 3, 70, 35, 0, 269, 271, 3, 72, 36, 0, 270,
		260, 1, 0, 0, 0, 270, 261, 1, 0, 0, 0, 270, 262, 1, 0, 0, 0, 270, 263,
		1, 0, 0, 0, 270, 264, 1, 0, 0, 0, 270, 265, 1, 0, 0, 0, 270, 266, 1, 0,
		0, 0, 270, 267, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 269, 1, 0, 0, 0,
		271, 31, 1, 0, 0, 0, 272, 273, 7, 0, 0, 0, 273, 33, 1, 0, 0, 0, 274, 275,
		5, 10, 0, 0, 275, 277, 7, 1, 0, 0, 276, 274, 1, 0, 0, 0, 277, 280, 1, 0,
		0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 35, 1, 0, 0, 0,
		280, 278, 1, 0, 0, 0, 281, 282, 5, 47, 0, 0, 282, 283, 5, 10, 0, 0, 283,
		37, 1, 0, 0, 0, 284, 293, 5, 47, 0, 0, 285, 286, 5, 33, 0, 0, 286, 289,
		3, 88, 44, 0, 287, 288, 5, 11, 0, 0, 288, 290, 3, 90, 45, 0, 289, 287,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 5, 34,
		0, 0, 292, 294, 1, 0, 0, 0, 293, 285, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0,
		294, 296, 1, 0, 0, 0, 295, 297, 5, 67, 0, 0, 296, 295, 1, 0, 0, 0, 296,
		297, 1, 0, 0, 0, 297, 301, 1, 0, 0, 0, 298, 300, 3, 68, 34, 0, 299, 298,
		1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0,
		0, 0, 302, 309, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 3, 32, 16,
		0, 305, 306, 3, 68, 34, 0, 306, 308, 1, 0, 0, 0, 307, 304, 1, 0, 0, 0,
		308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		312, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 5, 2, 0, 0, 313, 39, 1,
		0, 0, 0, 314, 316, 5, 63, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0,
		0, 316, 317, 1, 0, 0, 0, 317, 319, 5, 64, 0, 0, 318, 320, 3, 68, 34, 0,
		319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 326, 1, 0, 0, 0, 321,
		322, 3, 32, 16, 0, 322, 323, 3, 68, 34, 0, 323, 325, 1, 0, 0, 0, 324, 321,
		1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0,
		0, 0, 327, 332, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 329, 331, 3, 80, 40,
		0, 330, 329, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332,
		333, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 336,
		5, 2, 0, 0, 336, 41, 1, 0, 0, 0, 337, 338, 5, 65, 0, 0, 338, 340, 3, 88,
		44, 0, 339, 341, 3, 52, 26, 0, 340, 339, 1, 0, 0, 0, 340, 341, 1, 0, 0,
		0, 341, 342, 1, 0, 0, 0, 342, 343, 5, 12, 0, 0, 343, 345, 3, 88, 44, 0,
		344, 346, 3, 44, 22, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346,
		348, 1, 0, 0, 0, 347, 349, 3, 46, 23, 0, 348, 347, 1, 0, 0, 0, 348, 349,
		1, 0, 0, 0, 349, 351, 1, 0, 0, 0, 350, 352, 3, 48, 24, 0, 351, 350, 1,
		0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 354, 1, 0, 0, 0, 353, 355, 3, 50, 25,
		0, 354, 353, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 362, 1, 0, 0, 0, 356,
		357, 5, 65, 0, 0, 357, 359, 3, 88, 44, 0, 358, 360, 3, 48, 24, 0, 359,
		358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 362, 1, 0, 0, 0, 361, 337,
		1, 0, 0, 0, 361, 356, 1, 0, 0, 0, 362, 43, 1, 0, 0, 0, 363, 364, 5, 68,
		0, 0, 364, 365, 3, 88, 44, 0, 365, 45, 1, 0, 0, 0, 366, 367, 5, 69, 0,
		0, 367, 368, 3, 88, 44, 0, 368, 47, 1, 0, 0, 0, 369, 370, 5, 66, 0, 0,
		370, 371, 3, 118, 59, 0, 371, 49, 1, 0, 0, 0, 372, 373, 5, 70, 0, 0, 373,
		374, 3, 116, 58, 0, 374, 51, 1, 0, 0, 0, 375, 376, 5, 31, 0, 0, 376, 377,
		3, 54, 27, 0, 377, 378, 5, 32, 0, 0, 378, 53, 1, 0, 0, 0, 379, 380, 3,
		90, 45, 0, 380, 381, 5, 12, 0, 0, 381, 389, 3, 118, 59, 0, 382, 383, 5,
		10, 0, 0, 383, 384, 3, 90, 45, 0, 384, 385, 5, 12, 0, 0, 385, 386, 3, 118,
		59, 0, 386, 388, 1, 0, 0, 0, 387, 382, 1, 0, 0, 0, 388, 391, 1, 0, 0, 0,
		389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 55, 1, 0, 0, 0, 391, 389,
		1, 0, 0, 0, 392, 393, 5, 71, 0, 0, 393, 394, 5, 31, 0, 0, 394, 395, 3,
		90, 45, 0, 395, 396, 5, 32, 0, 0, 396, 399, 5, 72, 0, 0, 397, 400, 3, 30,
		15, 0, 398, 400, 3, 88, 44, 0, 399, 397, 1, 0, 0, 0, 399, 398, 1, 0, 0,
		0, 400, 421, 1, 0, 0, 0, 401, 402, 5, 71, 0, 0, 402, 403, 5, 31, 0, 0,
		403, 404, 3, 118, 59, 0, 404, 405, 5, 9, 0, 0, 405, 406, 3, 118, 59, 0,
		406, 409, 5, 72, 0, 0, 407, 410, 3, 30, 15, 0, 408, 410, 3, 88, 44, 0,
		409, 407, 1, 0, 0, 0, 409, 408, 1, 0, 0, 0, 410, 421, 1, 0, 0, 0, 411,
		412, 5, 71, 0, 0, 412, 413, 5, 72, 0, 0, 413, 421, 5, 6, 0, 0, 414, 415,
		5, 71, 0, 0, 415, 418, 5, 72, 0, 0, 416, 419, 3, 30, 15, 0, 417, 419, 3,
		88, 44, 0, 418, 416, 1, 0, 0, 0, 418, 417, 1, 0, 0, 0, 419, 421, 1, 0,
		0, 0, 420, 392, 1, 0, 0, 0, 420, 401, 1, 0, 0, 0, 420, 411, 1, 0, 0, 0,
		420, 414, 1, 0, 0, 0, 421, 57, 1, 0, 0, 0, 422, 423, 5, 73, 0, 0, 423,
		424, 3, 88, 44, 0, 424, 59, 1, 0, 0, 0, 425, 426, 5, 75, 0, 0, 426, 427,
		5, 72, 0, 0, 427, 428, 3, 88, 44, 0, 428, 61, 1, 0, 0, 0, 429, 430, 5,
		76, 0, 0, 430, 431, 5, 72, 0, 0, 431, 434, 3, 88, 44, 0, 432, 434, 5, 76,
		0, 0, 433, 429, 1, 0, 0, 0, 433, 432, 1, 0, 0, 0, 434, 63, 1, 0, 0, 0,
		435, 436, 5, 33, 0, 0, 436, 437, 3, 92, 46, 0, 437, 438, 5, 34, 0, 0, 438,
		65, 1, 0, 0, 0, 439, 440, 3, 118, 59, 0, 440, 441, 5, 9, 0, 0, 441, 442,
		3, 118, 59, 0, 442, 67, 1, 0, 0, 0, 443, 452, 3, 26, 13, 0, 444, 452, 3,
		22, 11, 0, 445, 447, 3, 102, 51, 0, 446, 448, 5, 10, 0, 0, 447, 446, 1,
		0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 452, 1, 0, 0, 0, 449, 452, 3, 74, 37,
		0, 450, 452, 3, 42, 21, 0, 451, 443, 1, 0, 0, 0, 451, 444, 1, 0, 0, 0,
		451, 445, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 450, 1, 0, 0, 0, 452,
		69, 1, 0, 0, 0, 453, 455, 5, 47, 0, 0, 454, 453, 1, 0, 0, 0, 454, 455,
		1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457, 5, 3, 0, 0, 457, 459, 5, 78,
		0, 0, 458, 460, 3, 94, 47, 0, 459, 458, 1, 0, 0, 0, 459, 460, 1, 0, 0,
		0, 460, 461, 1, 0, 0, 0, 461, 462, 5, 12, 0, 0, 462, 464, 5, 78, 0, 0,
		463, 465, 5, 54, 0, 0, 464, 463, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465,
		466, 1, 0, 0, 0, 466, 467, 3, 34, 17, 0, 467, 468, 5, 10, 0, 0, 468, 71,
		1, 0, 0, 0, 469, 471, 5, 47, 0, 0, 470, 469, 1, 0, 0, 0, 470, 471, 1, 0,
		0, 0, 471, 472, 1, 0, 0, 0, 472, 473, 5, 4, 0, 0, 473, 475, 5, 78, 0, 0,
		474, 476, 3, 94, 47, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476,
		478, 1, 0, 0, 0, 477, 479, 5, 54, 0, 0, 478, 477, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 481, 3, 34, 17, 0, 481, 482, 5,
		10, 0, 0, 482, 73, 1, 0, 0, 0, 483, 486, 3, 70, 35, 0, 484, 486, 3, 72,
		36, 0, 485, 483, 1, 0, 0, 0, 485, 484, 1, 0, 0, 0, 486, 75, 1, 0, 0, 0,
		487, 492, 3, 74, 37, 0, 488, 491, 3, 76, 38, 0, 489, 491, 3, 100, 50, 0,
		490, 488, 1, 0, 0, 0, 490, 489, 1, 0, 0, 0, 491, 494, 1, 0, 0, 0, 492,
		490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 495, 1, 0, 0, 0, 494, 492,
		1, 0, 0, 0, 495, 496, 3, 78, 39, 0, 496, 497, 5, 10, 0, 0, 497, 503, 1,
		0, 0, 0, 498, 499, 3, 74, 37, 0, 499, 500, 5, 46, 0, 0, 500, 501, 5, 10,
		0, 0, 501, 503, 1, 0, 0, 0, 502, 487, 1, 0, 0, 0, 502, 498, 1, 0, 0, 0,
		503, 77, 1, 0, 0, 0, 504, 509, 5, 1, 0, 0, 505, 508, 3, 86, 43, 0, 506,
		508, 3, 84, 42, 0, 507, 505, 1, 0, 0, 0, 507, 506, 1, 0, 0, 0, 508, 511,
		1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 512, 1, 0,
		0, 0, 511, 509, 1, 0, 0, 0, 512, 513, 5, 2, 0, 0, 513, 79, 1, 0, 0, 0,
		514, 515, 5, 77, 0, 0, 515, 518, 3, 88, 44, 0, 516, 517, 5, 12, 0, 0, 517,
		519, 3, 88, 44, 0, 518, 516, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 520,
		1, 0, 0, 0, 520, 521, 5, 72, 0, 0, 521, 526, 3, 82, 41, 0, 522, 523, 5,
		10, 0, 0, 523, 525, 3, 82, 41, 0, 524, 522, 1, 0, 0, 0, 525, 528, 1, 0,
		0, 0, 526, 524, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 530, 1, 0, 0, 0,
		528, 526, 1, 0, 0, 0, 529, 531, 5, 10, 0, 0, 530, 529, 1, 0, 0, 0, 530,
		531, 1, 0, 0, 0, 531, 81, 1, 0, 0, 0, 532, 533, 3, 90, 45, 0, 533, 534,
		5, 12, 0, 0, 534, 536, 5, 33, 0, 0, 535, 537, 3, 102, 51, 0, 536, 535,
		1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0, 538, 539, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 541, 5, 34, 0, 0, 541, 83, 1, 0, 0, 0,
		542, 544, 8, 2, 0, 0, 543, 542, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545,
		543, 1, 0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 85, 1, 0, 0, 0, 547, 551, 3,
		100, 50, 0, 548, 551, 3, 104, 52, 0, 549, 551, 3, 84, 42, 0, 550, 547,
		1, 0, 0, 0, 550, 548, 1, 0, 0, 0, 550, 549, 1, 0, 0, 0, 551, 87, 1, 0,
		0, 0, 552, 557, 5, 78, 0, 0, 553, 554, 5, 8, 0, 0, 554, 556, 5, 78, 0,
		0, 555, 553, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557,
		558, 1, 0, 0, 0, 558, 89, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 560, 565, 3,
		88, 44, 0, 561, 562, 5, 11, 0, 0, 562, 564, 3, 88, 44, 0, 563, 561, 1,
		0, 0, 0, 564, 567, 1, 0, 0, 0, 565, 563, 1, 0, 0, 0, 565, 566, 1, 0, 0,
		0, 566, 91, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0, 568, 573, 3, 118, 59, 0,
		569, 570, 5, 11, 0, 0, 570, 572, 3, 118, 59, 0, 571, 569, 1, 0, 0, 0, 572,
		575, 1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 93, 1,
		0, 0, 0, 575, 573, 1, 0, 0, 0, 576, 577, 5, 33, 0, 0, 577, 582, 3, 96,
		48, 0, 578, 579, 5, 10, 0, 0, 579, 581, 3, 96, 48, 0, 580, 578, 1, 0, 0,
		0, 581, 584, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583,
		585, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 585, 586, 5, 34, 0, 0, 586, 95,
		1, 0, 0, 0, 587, 589, 3, 98, 49, 0, 588, 587, 1, 0, 0, 0, 588, 589, 1,
		0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 593, 3, 90, 45, 0, 591, 592, 5, 12,
		0, 0, 592, 594, 3, 88, 44, 0, 593, 591, 1, 0, 0, 0, 593, 594, 1, 0, 0,
		0, 594, 597, 1, 0, 0, 0, 595, 596, 5, 13, 0, 0, 596, 598, 3, 118, 59, 0,
		597, 595, 1, 0, 0, 0, 597, 598, 1, 0, 0, 0, 598, 607, 1, 0, 0, 0, 599,
		600, 5, 3, 0, 0, 600, 601, 3, 94, 47, 0, 601, 602, 5, 12, 0, 0, 602, 603,
		3, 88, 44, 0, 603, 607, 1, 0, 0, 0, 604, 605, 5, 4, 0, 0, 605, 607, 3,
		94, 47, 0, 606, 588, 1, 0, 0, 0, 606, 599, 1, 0, 0, 0, 606, 604, 1, 0,
		0, 0, 607, 97, 1, 0, 0, 0, 608, 609, 7, 3, 0, 0, 609, 99, 1, 0, 0, 0, 610,
		614, 5, 5, 0, 0, 611, 612, 3, 102, 51, 0, 612, 613, 5, 10, 0, 0, 613, 615,
		1, 0, 0, 0, 614, 611, 1, 0, 0, 0, 615, 616, 1, 0, 0, 0, 616, 614, 1, 0,
		0, 0, 616, 617, 1, 0, 0, 0, 617, 101, 1, 0, 0, 0, 618, 621, 3, 90, 45,
		0, 619, 620, 5, 12, 0, 0, 620, 622, 3, 88, 44, 0, 621, 619, 1, 0, 0, 0,
		621, 622, 1, 0, 0, 0, 622, 625, 1, 0, 0, 0, 623, 624, 5, 13, 0, 0, 624,
		626, 3, 118, 59, 0, 625, 623, 1, 0, 0, 0, 625, 626, 1, 0, 0, 0, 626, 103,
		1, 0, 0, 0, 627, 628, 5, 5, 0, 0, 628, 631, 3, 90, 45, 0, 629, 630, 5,
		12, 0, 0, 630, 632, 3, 88, 44, 0, 631, 629, 1, 0, 0, 0, 631, 632, 1, 0,
		0, 0, 632, 635, 1, 0, 0, 0, 633, 634, 5, 35, 0, 0, 634, 636, 3, 118, 59,
		0, 635, 633, 1, 0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 637, 1, 0, 0, 0, 637,
		638, 5, 10, 0, 0, 638, 105, 1, 0, 0, 0, 639, 640, 7, 4, 0, 0, 640, 107,
		1, 0, 0, 0, 641, 642, 7, 5, 0, 0, 642, 109, 1, 0, 0, 0, 643, 644, 7, 6,
		0, 0, 644, 111, 1, 0, 0, 0, 645, 649, 3, 106, 53, 0, 646, 649, 3, 108,
		54, 0, 647, 649, 3, 110, 55, 0, 648, 645, 1, 0, 0, 0, 648, 646, 1, 0, 0,
		0, 648, 647, 1, 0, 0, 0, 649, 113, 1, 0, 0, 0, 650, 652, 7, 7, 0, 0, 651,
		650, 1, 0, 0, 0, 652, 653, 1, 0, 0, 0, 653, 651, 1, 0, 0, 0, 653, 654,
		1, 0, 0, 0, 654, 115, 1, 0, 0, 0, 655, 656, 7, 8, 0, 0, 656, 117, 1, 0,
		0, 0, 657, 663, 3, 120, 60, 0, 658, 659, 3, 112, 56, 0, 659, 660, 3, 120,
		60, 0, 660, 662, 1, 0, 0, 0, 661, 658, 1, 0, 0, 0, 662, 665, 1, 0, 0, 0,
		663, 661, 1, 0, 0, 0, 663, 664, 1, 0, 0, 0, 664, 669, 1, 0, 0, 0, 665,
		663, 1, 0, 0, 0, 666, 669, 3, 122, 61, 0, 667, 669, 3, 128, 64, 0, 668,
		657, 1, 0, 0, 0, 668, 666, 1, 0, 0, 0, 668, 667, 1, 0, 0, 0, 669, 119,
		1, 0, 0, 0, 670, 672, 5, 74, 0, 0, 671, 670, 1, 0, 0, 0, 671, 672, 1, 0,
		0, 0, 672, 673, 1, 0, 0, 0, 673, 683, 5, 78, 0, 0, 674, 683, 3, 116, 58,
		0, 675, 683, 3, 114, 57, 0, 676, 683, 3, 126, 63, 0, 677, 678, 5, 33, 0,
		0, 678, 679, 3, 118, 59, 0, 679, 680, 5, 34, 0, 0, 680, 683, 1, 0, 0, 0,
		681, 683, 3, 128, 64, 0, 682, 671, 1, 0, 0, 0, 682, 674, 1, 0, 0, 0, 682,
		675, 1, 0, 0, 0, 682, 676, 1, 0, 0, 0, 682, 677, 1, 0, 0, 0, 682, 681,
		1, 0, 0, 0, 683, 121, 1, 0, 0, 0, 684, 686, 5, 3, 0, 0, 685, 687, 3, 94,
		47, 0, 686, 685, 1, 0, 0, 0, 686, 687, 1, 0, 0, 0, 687, 688, 1, 0, 0, 0,
		688, 690, 5, 12, 0, 0, 689, 691, 3, 100, 50, 0, 690, 689, 1, 0, 0, 0, 690,
		691, 1, 0, 0, 0, 691, 692, 1, 0, 0, 0, 692, 693, 3, 78, 39, 0, 693, 123,
		1, 0, 0, 0, 694, 696, 5, 4, 0, 0, 695, 697, 3, 94, 47, 0, 696, 695, 1,
		0, 0, 0, 696, 697, 1, 0, 0, 0, 697, 699, 1, 0, 0, 0, 698, 700, 3, 100,
		50, 0, 699, 698, 1, 0, 0, 0, 699, 700, 1, 0, 0, 0, 700, 701, 1, 0, 0, 0,
		701, 702, 3, 78, 39, 0, 702, 125, 1, 0, 0, 0, 703, 708, 3, 88, 44, 0, 704,
		705, 5, 15, 0, 0, 705, 706, 3, 88, 44, 0, 706, 707, 5, 17, 0, 0, 707, 709,
		1, 0, 0, 0, 708, 704, 1, 0, 0, 0, 708, 709, 1, 0, 0, 0, 709, 715, 1, 0,
		0, 0, 710, 712, 5, 33, 0, 0, 711, 713, 3, 92, 46, 0, 712, 711, 1, 0, 0,
		0, 712, 713, 1, 0, 0, 0, 713, 714, 1, 0, 0, 0, 714, 716, 5, 34, 0, 0, 715,
		710, 1, 0, 0, 0, 715, 716, 1, 0, 0, 0, 716, 127, 1, 0, 0, 0, 717, 719,
		8, 9, 0, 0, 718, 717, 1, 0, 0, 0, 719, 720, 1, 0, 0, 0, 720, 718, 1, 0,
		0, 0, 720, 721, 1, 0, 0, 0, 721, 129, 1, 0, 0, 0, 88, 132, 140, 145, 157,
		160, 163, 171, 177, 193, 195, 204, 206, 219, 228, 236, 246, 258, 270, 278,
		289, 293, 296, 301, 309, 315, 319, 326, 332, 340, 345, 348, 351, 354, 359,
		361, 389, 399, 409, 418, 420, 433, 447, 451, 454, 459, 464, 470, 475, 478,
		485, 490, 492, 502, 507, 509, 518, 526, 530, 538, 545, 550, 557, 565, 573,
		582, 588, 593, 597, 606, 616, 621, 625, 631, 635, 648, 653, 663, 668, 671,
		682, 686, 690, 696, 699, 708, 712, 715, 720,
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

// scopepascalParserInit initializes any static state used to implement scopepascalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewscopepascalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ScopepascalParserInit() {
	staticData := &ScopepascalParserStaticData
	staticData.once.Do(scopepascalParserInit)
}

// NewscopepascalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewscopepascalParser(input antlr.TokenStream) *scopepascalParser {
	ScopepascalParserInit()
	this := new(scopepascalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ScopepascalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "scopepascal.g4"

	return this
}

// scopepascalParser tokens.
const (
	scopepascalParserEOF                       = antlr.TokenEOF
	scopepascalParserBEGIN                     = 1
	scopepascalParserEND                       = 2
	scopepascalParserFUNCTION                  = 3
	scopepascalParserPROCEDURE                 = 4
	scopepascalParserVAR                       = 5
	scopepascalParserCONST                     = 6
	scopepascalParserOUT                       = 7
	scopepascalParserDOT                       = 8
	scopepascalParserDOUBLEDOT                 = 9
	scopepascalParserSEMI                      = 10
	scopepascalParserCOMMA                     = 11
	scopepascalParserCOLON                     = 12
	scopepascalParserEQUAL                     = 13
	scopepascalParserNOT_EQUAL                 = 14
	scopepascalParserLT                        = 15
	scopepascalParserLE                        = 16
	scopepascalParserGT                        = 17
	scopepascalParserGE                        = 18
	scopepascalParserIN                        = 19
	scopepascalParserSTAR                      = 20
	scopepascalParserSLASH                     = 21
	scopepascalParserDIV                       = 22
	scopepascalParserMOD                       = 23
	scopepascalParserAND                       = 24
	scopepascalParserOR                        = 25
	scopepascalParserNOT                       = 26
	scopepascalParserPLUS                      = 27
	scopepascalParserMINUS                     = 28
	scopepascalParserSHR                       = 29
	scopepascalParserSHL                       = 30
	scopepascalParserLBRACK                    = 31
	scopepascalParserRBRACK                    = 32
	scopepascalParserLPAREN                    = 33
	scopepascalParserRPAREN                    = 34
	scopepascalParserASSIGN                    = 35
	scopepascalParserLABEL                     = 36
	scopepascalParserTYPE                      = 37
	scopepascalParserRESOURCESTRING            = 38
	scopepascalParserPROGRAM                   = 39
	scopepascalParserUNIT                      = 40
	scopepascalParserINTERFACE                 = 41
	scopepascalParserIMPLEMENTATION            = 42
	scopepascalParserINITIALIZATION            = 43
	scopepascalParserFINALIZATION              = 44
	scopepascalParserUSES                      = 45
	scopepascalParserFORWARD                   = 46
	scopepascalParserCLASS                     = 47
	scopepascalParserPRIVATE                   = 48
	scopepascalParserPROTECTED                 = 49
	scopepascalParserPUBLIC                    = 50
	scopepascalParserPUBLISHED                 = 51
	scopepascalParserSTRICTPRIVATE             = 52
	scopepascalParserSTRICTPROTECTED           = 53
	scopepascalParserOFOBJECT                  = 54
	scopepascalParserVIRTUAL                   = 55
	scopepascalParserOVERRIDE                  = 56
	scopepascalParserREINTRODUCE               = 57
	scopepascalParserOVERLOAD                  = 58
	scopepascalParserINLINE                    = 59
	scopepascalParserSTDCALL                   = 60
	scopepascalParserCDECL                     = 61
	scopepascalParserSTATIC                    = 62
	scopepascalParserPACKED                    = 63
	scopepascalParserRECORD                    = 64
	scopepascalParserPROPERTY                  = 65
	scopepascalParserDEFAULT                   = 66
	scopepascalParserABSTRACT                  = 67
	scopepascalParserREAD                      = 68
	scopepascalParserWRITE                     = 69
	scopepascalParserINDEX                     = 70
	scopepascalParserARRAY                     = 71
	scopepascalParserOF                        = 72
	scopepascalParserDEREFERENCE               = 73
	scopepascalParserAT                        = 74
	scopepascalParserSET                       = 75
	scopepascalParserFILE                      = 76
	scopepascalParserCASE                      = 77
	scopepascalParserIDENT                     = 78
	scopepascalParserHEX_LITERAL               = 79
	scopepascalParserSTRING_LITERAL            = 80
	scopepascalParserSTRING_CROSSHATCH_LITERAL = 81
	scopepascalParserNUM_INT                   = 82
	scopepascalParserNUM_REAL                  = 83
	scopepascalParserWS                        = 84
	scopepascalParserCOMMENT_1                 = 85
	scopepascalParserCOMMENT_2                 = 86
	scopepascalParserCOMMENT_3                 = 87
	scopepascalParserUTF8BOM                   = 88
)

// scopepascalParser rules.
const (
	scopepascalParserRULE_source                          = 0
	scopepascalParserRULE_program                         = 1
	scopepascalParserRULE_unit                            = 2
	scopepascalParserRULE_interfaceSection                = 3
	scopepascalParserRULE_implementationSection           = 4
	scopepascalParserRULE_initializationSection           = 5
	scopepascalParserRULE_finalizationSection             = 6
	scopepascalParserRULE_interfaceBlock                  = 7
	scopepascalParserRULE_implementationBlock             = 8
	scopepascalParserRULE_unitList                        = 9
	scopepascalParserRULE_labelDeclaration                = 10
	scopepascalParserRULE_constSection                    = 11
	scopepascalParserRULE_resourceSection                 = 12
	scopepascalParserRULE_typeSection                     = 13
	scopepascalParserRULE_typeBlock                       = 14
	scopepascalParserRULE_type                            = 15
	scopepascalParserRULE_accessSpecifier                 = 16
	scopepascalParserRULE_procedureOrFunctionModifiers    = 17
	scopepascalParserRULE_classForwardDeclaration         = 18
	scopepascalParserRULE_classType                       = 19
	scopepascalParserRULE_recordType                      = 20
	scopepascalParserRULE_propertyDeclaration             = 21
	scopepascalParserRULE_propertyReadDeclaration         = 22
	scopepascalParserRULE_propertyWriteDeclaration        = 23
	scopepascalParserRULE_propertyDefaultValueDeclaration = 24
	scopepascalParserRULE_propertyIndexDeclaration        = 25
	scopepascalParserRULE_propertyIndexParameters         = 26
	scopepascalParserRULE_propertyIndexParametersList     = 27
	scopepascalParserRULE_arrayType                       = 28
	scopepascalParserRULE_pointerType                     = 29
	scopepascalParserRULE_setType                         = 30
	scopepascalParserRULE_fileType                        = 31
	scopepascalParserRULE_scalarType                      = 32
	scopepascalParserRULE_subrangeType                    = 33
	scopepascalParserRULE_blockDeclaration                = 34
	scopepascalParserRULE_functionDeclaration             = 35
	scopepascalParserRULE_procedureDeclaration            = 36
	scopepascalParserRULE_functionOrProcedureDeclaration  = 37
	scopepascalParserRULE_functionOrProcedure             = 38
	scopepascalParserRULE_blockStatement                  = 39
	scopepascalParserRULE_recordVariantDeclaration        = 40
	scopepascalParserRULE_recordVariant                   = 41
	scopepascalParserRULE_statementError                  = 42
	scopepascalParserRULE_statement                       = 43
	scopepascalParserRULE_identifier                      = 44
	scopepascalParserRULE_identifierList                  = 45
	scopepascalParserRULE_expressionList                  = 46
	scopepascalParserRULE_paramsDeclaration               = 47
	scopepascalParserRULE_paramsDeclarationSection        = 48
	scopepascalParserRULE_paramSpecifier                  = 49
	scopepascalParserRULE_varSection                      = 50
	scopepascalParserRULE_varDeclaration                  = 51
	scopepascalParserRULE_inlinedVarDeclaration           = 52
	scopepascalParserRULE_relationaloperator              = 53
	scopepascalParserRULE_additiveoperator                = 54
	scopepascalParserRULE_multiplicativeoperator          = 55
	scopepascalParserRULE_operator                        = 56
	scopepascalParserRULE_string                          = 57
	scopepascalParserRULE_number                          = 58
	scopepascalParserRULE_expression                      = 59
	scopepascalParserRULE_term                            = 60
	scopepascalParserRULE_functionExpression              = 61
	scopepascalParserRULE_procedureExpression             = 62
	scopepascalParserRULE_functionDesignator              = 63
	scopepascalParserRULE_errorExpression                 = 64
)

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Program() IProgramContext
	Unit() IUnitContext

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_source
	return p
}

func InitEmptySourceContext(p *SourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_source
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *SourceContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitSource(s)
	}
}

func (s *SourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, scopepascalParserRULE_source)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case scopepascalParserPROGRAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Program()
		}

	case scopepascalParserUNIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Unit()
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

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROGRAM() antlr.TerminalNode
	Identifier() IIdentifierContext
	SEMI() antlr.TerminalNode
	ImplementationBlock() IImplementationBlockContext
	BlockStatement() IBlockStatementContext
	DOT() antlr.TerminalNode
	EOF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllIdentifierList() []IIdentifierListContext
	IdentifierList(i int) IIdentifierListContext
	RPAREN() antlr.TerminalNode
	USES() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROGRAM, 0)
}

func (s *ProgramContext) Identifier() IIdentifierContext {
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

func (s *ProgramContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *ProgramContext) ImplementationBlock() IImplementationBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementationBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementationBlockContext)
}

func (s *ProgramContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ProgramContext) DOT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDOT, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEOF, 0)
}

func (s *ProgramContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *ProgramContext) AllIdentifierList() []IIdentifierListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierListContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierListContext); ok {
			tst[i] = t.(IIdentifierListContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) IdentifierList(i int) IIdentifierListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
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

	return t.(IIdentifierListContext)
}

func (s *ProgramContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *ProgramContext) USES() antlr.TerminalNode {
	return s.GetToken(scopepascalParserUSES, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, scopepascalParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(scopepascalParserPROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Identifier()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(136)
			p.Match(scopepascalParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.IdentifierList()
		}
		{
			p.SetState(138)
			p.Match(scopepascalParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(142)
		p.Match(scopepascalParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserUSES {
		{
			p.SetState(143)
			p.Match(scopepascalParserUSES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.IdentifierList()
		}

	}
	{
		p.SetState(147)
		p.ImplementationBlock()
	}
	{
		p.SetState(148)
		p.BlockStatement()
	}
	{
		p.SetState(149)
		p.Match(scopepascalParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(scopepascalParserEOF)
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

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNIT() antlr.TerminalNode
	Identifier() IIdentifierContext
	SEMI() antlr.TerminalNode
	InterfaceSection() IInterfaceSectionContext
	END() antlr.TerminalNode
	DOT() antlr.TerminalNode
	EOF() antlr.TerminalNode
	ImplementationSection() IImplementationSectionContext
	InitializationSection() IInitializationSectionContext
	FinalizationSection() IFinalizationSectionContext

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) UNIT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserUNIT, 0)
}

func (s *UnitContext) Identifier() IIdentifierContext {
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

func (s *UnitContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *UnitContext) InterfaceSection() IInterfaceSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceSectionContext)
}

func (s *UnitContext) END() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, 0)
}

func (s *UnitContext) DOT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDOT, 0)
}

func (s *UnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEOF, 0)
}

func (s *UnitContext) ImplementationSection() IImplementationSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementationSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementationSectionContext)
}

func (s *UnitContext) InitializationSection() IInitializationSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializationSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializationSectionContext)
}

func (s *UnitContext) FinalizationSection() IFinalizationSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinalizationSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFinalizationSectionContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, scopepascalParserRULE_unit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(scopepascalParserUNIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Identifier()
	}
	{
		p.SetState(154)
		p.Match(scopepascalParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.InterfaceSection()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserIMPLEMENTATION {
		{
			p.SetState(156)
			p.ImplementationSection()
		}

	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserINITIALIZATION {
		{
			p.SetState(159)
			p.InitializationSection()
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserFINALIZATION {
		{
			p.SetState(162)
			p.FinalizationSection()
		}

	}
	{
		p.SetState(165)
		p.Match(scopepascalParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(scopepascalParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(scopepascalParserEOF)
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

// IInterfaceSectionContext is an interface to support dynamic dispatch.
type IInterfaceSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	InterfaceBlock() IInterfaceBlockContext
	UnitList() IUnitListContext

	// IsInterfaceSectionContext differentiates from other interfaces.
	IsInterfaceSectionContext()
}

type InterfaceSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceSectionContext() *InterfaceSectionContext {
	var p = new(InterfaceSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_interfaceSection
	return p
}

func InitEmptyInterfaceSectionContext(p *InterfaceSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_interfaceSection
}

func (*InterfaceSectionContext) IsInterfaceSectionContext() {}

func NewInterfaceSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceSectionContext {
	var p = new(InterfaceSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_interfaceSection

	return p
}

func (s *InterfaceSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceSectionContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserINTERFACE, 0)
}

func (s *InterfaceSectionContext) InterfaceBlock() IInterfaceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceBlockContext)
}

func (s *InterfaceSectionContext) UnitList() IUnitListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitListContext)
}

func (s *InterfaceSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterInterfaceSection(s)
	}
}

func (s *InterfaceSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitInterfaceSection(s)
	}
}

func (s *InterfaceSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitInterfaceSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) InterfaceSection() (localctx IInterfaceSectionContext) {
	localctx = NewInterfaceSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, scopepascalParserRULE_interfaceSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(scopepascalParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserUSES {
		{
			p.SetState(170)
			p.UnitList()
		}

	}
	{
		p.SetState(173)
		p.InterfaceBlock()
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

// IImplementationSectionContext is an interface to support dynamic dispatch.
type IImplementationSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPLEMENTATION() antlr.TerminalNode
	ImplementationBlock() IImplementationBlockContext
	UnitList() IUnitListContext

	// IsImplementationSectionContext differentiates from other interfaces.
	IsImplementationSectionContext()
}

type ImplementationSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementationSectionContext() *ImplementationSectionContext {
	var p = new(ImplementationSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_implementationSection
	return p
}

func InitEmptyImplementationSectionContext(p *ImplementationSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_implementationSection
}

func (*ImplementationSectionContext) IsImplementationSectionContext() {}

func NewImplementationSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementationSectionContext {
	var p = new(ImplementationSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_implementationSection

	return p
}

func (s *ImplementationSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementationSectionContext) IMPLEMENTATION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserIMPLEMENTATION, 0)
}

func (s *ImplementationSectionContext) ImplementationBlock() IImplementationBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementationBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementationBlockContext)
}

func (s *ImplementationSectionContext) UnitList() IUnitListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitListContext)
}

func (s *ImplementationSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementationSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementationSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterImplementationSection(s)
	}
}

func (s *ImplementationSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitImplementationSection(s)
	}
}

func (s *ImplementationSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitImplementationSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ImplementationSection() (localctx IImplementationSectionContext) {
	localctx = NewImplementationSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, scopepascalParserRULE_implementationSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(scopepascalParserIMPLEMENTATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserUSES {
		{
			p.SetState(176)
			p.UnitList()
		}

	}
	{
		p.SetState(179)
		p.ImplementationBlock()
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

// IInitializationSectionContext is an interface to support dynamic dispatch.
type IInitializationSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INITIALIZATION() antlr.TerminalNode
	BlockStatement() IBlockStatementContext

	// IsInitializationSectionContext differentiates from other interfaces.
	IsInitializationSectionContext()
}

type InitializationSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializationSectionContext() *InitializationSectionContext {
	var p = new(InitializationSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_initializationSection
	return p
}

func InitEmptyInitializationSectionContext(p *InitializationSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_initializationSection
}

func (*InitializationSectionContext) IsInitializationSectionContext() {}

func NewInitializationSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializationSectionContext {
	var p = new(InitializationSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_initializationSection

	return p
}

func (s *InitializationSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializationSectionContext) INITIALIZATION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserINITIALIZATION, 0)
}

func (s *InitializationSectionContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *InitializationSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializationSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializationSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterInitializationSection(s)
	}
}

func (s *InitializationSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitInitializationSection(s)
	}
}

func (s *InitializationSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitInitializationSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) InitializationSection() (localctx IInitializationSectionContext) {
	localctx = NewInitializationSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, scopepascalParserRULE_initializationSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(scopepascalParserINITIALIZATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.BlockStatement()
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

// IFinalizationSectionContext is an interface to support dynamic dispatch.
type IFinalizationSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FINALIZATION() antlr.TerminalNode
	BlockStatement() IBlockStatementContext

	// IsFinalizationSectionContext differentiates from other interfaces.
	IsFinalizationSectionContext()
}

type FinalizationSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinalizationSectionContext() *FinalizationSectionContext {
	var p = new(FinalizationSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_finalizationSection
	return p
}

func InitEmptyFinalizationSectionContext(p *FinalizationSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_finalizationSection
}

func (*FinalizationSectionContext) IsFinalizationSectionContext() {}

func NewFinalizationSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinalizationSectionContext {
	var p = new(FinalizationSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_finalizationSection

	return p
}

func (s *FinalizationSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *FinalizationSectionContext) FINALIZATION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFINALIZATION, 0)
}

func (s *FinalizationSectionContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FinalizationSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinalizationSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinalizationSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFinalizationSection(s)
	}
}

func (s *FinalizationSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFinalizationSection(s)
	}
}

func (s *FinalizationSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFinalizationSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FinalizationSection() (localctx IFinalizationSectionContext) {
	localctx = NewFinalizationSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, scopepascalParserRULE_finalizationSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(scopepascalParserFINALIZATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.BlockStatement()
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

// IInterfaceBlockContext is an interface to support dynamic dispatch.
type IInterfaceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeSection() []ITypeSectionContext
	TypeSection(i int) ITypeSectionContext
	AllLabelDeclaration() []ILabelDeclarationContext
	LabelDeclaration(i int) ILabelDeclarationContext
	AllConstSection() []IConstSectionContext
	ConstSection(i int) IConstSectionContext
	AllResourceSection() []IResourceSectionContext
	ResourceSection(i int) IResourceSectionContext
	AllVarSection() []IVarSectionContext
	VarSection(i int) IVarSectionContext
	AllFunctionOrProcedureDeclaration() []IFunctionOrProcedureDeclarationContext
	FunctionOrProcedureDeclaration(i int) IFunctionOrProcedureDeclarationContext

	// IsInterfaceBlockContext differentiates from other interfaces.
	IsInterfaceBlockContext()
}

type InterfaceBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceBlockContext() *InterfaceBlockContext {
	var p = new(InterfaceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_interfaceBlock
	return p
}

func InitEmptyInterfaceBlockContext(p *InterfaceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_interfaceBlock
}

func (*InterfaceBlockContext) IsInterfaceBlockContext() {}

func NewInterfaceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceBlockContext {
	var p = new(InterfaceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_interfaceBlock

	return p
}

func (s *InterfaceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceBlockContext) AllTypeSection() []ITypeSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSectionContext); ok {
			len++
		}
	}

	tst := make([]ITypeSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSectionContext); ok {
			tst[i] = t.(ITypeSectionContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) TypeSection(i int) ITypeSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSectionContext); ok {
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

	return t.(ITypeSectionContext)
}

func (s *InterfaceBlockContext) AllLabelDeclaration() []ILabelDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ILabelDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelDeclarationContext); ok {
			tst[i] = t.(ILabelDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) LabelDeclaration(i int) ILabelDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelDeclarationContext); ok {
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

	return t.(ILabelDeclarationContext)
}

func (s *InterfaceBlockContext) AllConstSection() []IConstSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstSectionContext); ok {
			len++
		}
	}

	tst := make([]IConstSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstSectionContext); ok {
			tst[i] = t.(IConstSectionContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) ConstSection(i int) IConstSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstSectionContext); ok {
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

	return t.(IConstSectionContext)
}

func (s *InterfaceBlockContext) AllResourceSection() []IResourceSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourceSectionContext); ok {
			len++
		}
	}

	tst := make([]IResourceSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourceSectionContext); ok {
			tst[i] = t.(IResourceSectionContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) ResourceSection(i int) IResourceSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceSectionContext); ok {
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

	return t.(IResourceSectionContext)
}

func (s *InterfaceBlockContext) AllVarSection() []IVarSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarSectionContext); ok {
			len++
		}
	}

	tst := make([]IVarSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarSectionContext); ok {
			tst[i] = t.(IVarSectionContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) VarSection(i int) IVarSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
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

	return t.(IVarSectionContext)
}

func (s *InterfaceBlockContext) AllFunctionOrProcedureDeclaration() []IFunctionOrProcedureDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionOrProcedureDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionOrProcedureDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionOrProcedureDeclarationContext); ok {
			tst[i] = t.(IFunctionOrProcedureDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceBlockContext) FunctionOrProcedureDeclaration(i int) IFunctionOrProcedureDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionOrProcedureDeclarationContext); ok {
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

	return t.(IFunctionOrProcedureDeclarationContext)
}

func (s *InterfaceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterInterfaceBlock(s)
	}
}

func (s *InterfaceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitInterfaceBlock(s)
	}
}

func (s *InterfaceBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitInterfaceBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) InterfaceBlock() (localctx IInterfaceBlockContext) {
	localctx = NewInterfaceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, scopepascalParserRULE_interfaceBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141218524692600) != 0 {
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case scopepascalParserTYPE:
			{
				p.SetState(187)
				p.TypeSection()
			}

		case scopepascalParserLABEL:
			{
				p.SetState(188)
				p.LabelDeclaration()
			}

		case scopepascalParserCONST:
			{
				p.SetState(189)
				p.ConstSection()
			}

		case scopepascalParserRESOURCESTRING:
			{
				p.SetState(190)
				p.ResourceSection()
			}

		case scopepascalParserVAR:
			{
				p.SetState(191)
				p.VarSection()
			}

		case scopepascalParserFUNCTION, scopepascalParserPROCEDURE, scopepascalParserCLASS:
			{
				p.SetState(192)
				p.FunctionOrProcedureDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(197)
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

// IImplementationBlockContext is an interface to support dynamic dispatch.
type IImplementationBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabelDeclaration() []ILabelDeclarationContext
	LabelDeclaration(i int) ILabelDeclarationContext
	AllConstSection() []IConstSectionContext
	ConstSection(i int) IConstSectionContext
	AllResourceSection() []IResourceSectionContext
	ResourceSection(i int) IResourceSectionContext
	AllTypeSection() []ITypeSectionContext
	TypeSection(i int) ITypeSectionContext
	AllVarSection() []IVarSectionContext
	VarSection(i int) IVarSectionContext
	AllFunctionOrProcedure() []IFunctionOrProcedureContext
	FunctionOrProcedure(i int) IFunctionOrProcedureContext

	// IsImplementationBlockContext differentiates from other interfaces.
	IsImplementationBlockContext()
}

type ImplementationBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementationBlockContext() *ImplementationBlockContext {
	var p = new(ImplementationBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_implementationBlock
	return p
}

func InitEmptyImplementationBlockContext(p *ImplementationBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_implementationBlock
}

func (*ImplementationBlockContext) IsImplementationBlockContext() {}

func NewImplementationBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementationBlockContext {
	var p = new(ImplementationBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_implementationBlock

	return p
}

func (s *ImplementationBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementationBlockContext) AllLabelDeclaration() []ILabelDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ILabelDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelDeclarationContext); ok {
			tst[i] = t.(ILabelDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) LabelDeclaration(i int) ILabelDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelDeclarationContext); ok {
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

	return t.(ILabelDeclarationContext)
}

func (s *ImplementationBlockContext) AllConstSection() []IConstSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstSectionContext); ok {
			len++
		}
	}

	tst := make([]IConstSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstSectionContext); ok {
			tst[i] = t.(IConstSectionContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) ConstSection(i int) IConstSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstSectionContext); ok {
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

	return t.(IConstSectionContext)
}

func (s *ImplementationBlockContext) AllResourceSection() []IResourceSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourceSectionContext); ok {
			len++
		}
	}

	tst := make([]IResourceSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourceSectionContext); ok {
			tst[i] = t.(IResourceSectionContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) ResourceSection(i int) IResourceSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceSectionContext); ok {
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

	return t.(IResourceSectionContext)
}

func (s *ImplementationBlockContext) AllTypeSection() []ITypeSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSectionContext); ok {
			len++
		}
	}

	tst := make([]ITypeSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSectionContext); ok {
			tst[i] = t.(ITypeSectionContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) TypeSection(i int) ITypeSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSectionContext); ok {
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

	return t.(ITypeSectionContext)
}

func (s *ImplementationBlockContext) AllVarSection() []IVarSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarSectionContext); ok {
			len++
		}
	}

	tst := make([]IVarSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarSectionContext); ok {
			tst[i] = t.(IVarSectionContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) VarSection(i int) IVarSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
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

	return t.(IVarSectionContext)
}

func (s *ImplementationBlockContext) AllFunctionOrProcedure() []IFunctionOrProcedureContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionOrProcedureContext); ok {
			len++
		}
	}

	tst := make([]IFunctionOrProcedureContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionOrProcedureContext); ok {
			tst[i] = t.(IFunctionOrProcedureContext)
			i++
		}
	}

	return tst
}

func (s *ImplementationBlockContext) FunctionOrProcedure(i int) IFunctionOrProcedureContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionOrProcedureContext); ok {
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

	return t.(IFunctionOrProcedureContext)
}

func (s *ImplementationBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementationBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementationBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterImplementationBlock(s)
	}
}

func (s *ImplementationBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitImplementationBlock(s)
	}
}

func (s *ImplementationBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitImplementationBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ImplementationBlock() (localctx IImplementationBlockContext) {
	localctx = NewImplementationBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, scopepascalParserRULE_implementationBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141218524692600) != 0 {
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case scopepascalParserLABEL:
			{
				p.SetState(198)
				p.LabelDeclaration()
			}

		case scopepascalParserCONST:
			{
				p.SetState(199)
				p.ConstSection()
			}

		case scopepascalParserRESOURCESTRING:
			{
				p.SetState(200)
				p.ResourceSection()
			}

		case scopepascalParserTYPE:
			{
				p.SetState(201)
				p.TypeSection()
			}

		case scopepascalParserVAR:
			{
				p.SetState(202)
				p.VarSection()
			}

		case scopepascalParserFUNCTION, scopepascalParserPROCEDURE, scopepascalParserCLASS:
			{
				p.SetState(203)
				p.FunctionOrProcedure()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(208)
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

// IUnitListContext is an interface to support dynamic dispatch.
type IUnitListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USES() antlr.TerminalNode
	IdentifierList() IIdentifierListContext
	SEMI() antlr.TerminalNode

	// IsUnitListContext differentiates from other interfaces.
	IsUnitListContext()
}

type UnitListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitListContext() *UnitListContext {
	var p = new(UnitListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_unitList
	return p
}

func InitEmptyUnitListContext(p *UnitListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_unitList
}

func (*UnitListContext) IsUnitListContext() {}

func NewUnitListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitListContext {
	var p = new(UnitListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_unitList

	return p
}

func (s *UnitListContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitListContext) USES() antlr.TerminalNode {
	return s.GetToken(scopepascalParserUSES, 0)
}

func (s *UnitListContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *UnitListContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *UnitListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterUnitList(s)
	}
}

func (s *UnitListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitUnitList(s)
	}
}

func (s *UnitListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitUnitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) UnitList() (localctx IUnitListContext) {
	localctx = NewUnitListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, scopepascalParserRULE_unitList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(scopepascalParserUSES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.IdentifierList()
	}
	{
		p.SetState(211)
		p.Match(scopepascalParserSEMI)
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

// ILabelDeclarationContext is an interface to support dynamic dispatch.
type ILabelDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LABEL() antlr.TerminalNode
	AllNumber() []INumberContext
	Number(i int) INumberContext
	SEMI() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLabelDeclarationContext differentiates from other interfaces.
	IsLabelDeclarationContext()
}

type LabelDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelDeclarationContext() *LabelDeclarationContext {
	var p = new(LabelDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_labelDeclaration
	return p
}

func InitEmptyLabelDeclarationContext(p *LabelDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_labelDeclaration
}

func (*LabelDeclarationContext) IsLabelDeclarationContext() {}

func NewLabelDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelDeclarationContext {
	var p = new(LabelDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_labelDeclaration

	return p
}

func (s *LabelDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelDeclarationContext) LABEL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLABEL, 0)
}

func (s *LabelDeclarationContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *LabelDeclarationContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
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

	return t.(INumberContext)
}

func (s *LabelDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *LabelDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOMMA)
}

func (s *LabelDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOMMA, i)
}

func (s *LabelDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterLabelDeclaration(s)
	}
}

func (s *LabelDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitLabelDeclaration(s)
	}
}

func (s *LabelDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitLabelDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) LabelDeclaration() (localctx ILabelDeclarationContext) {
	localctx = NewLabelDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, scopepascalParserRULE_labelDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(scopepascalParserLABEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Number()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserCOMMA {
		{
			p.SetState(215)
			p.Match(scopepascalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Number()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(scopepascalParserSEMI)
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

// IConstSectionContext is an interface to support dynamic dispatch.
type IConstSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	AllIdentifierList() []IIdentifierListContext
	IdentifierList(i int) IIdentifierListContext
	AllEQUAL() []antlr.TerminalNode
	EQUAL(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsConstSectionContext differentiates from other interfaces.
	IsConstSectionContext()
}

type ConstSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstSectionContext() *ConstSectionContext {
	var p = new(ConstSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_constSection
	return p
}

func InitEmptyConstSectionContext(p *ConstSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_constSection
}

func (*ConstSectionContext) IsConstSectionContext() {}

func NewConstSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstSectionContext {
	var p = new(ConstSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_constSection

	return p
}

func (s *ConstSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstSectionContext) CONST() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCONST, 0)
}

func (s *ConstSectionContext) AllIdentifierList() []IIdentifierListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierListContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierListContext); ok {
			tst[i] = t.(IIdentifierListContext)
			i++
		}
	}

	return tst
}

func (s *ConstSectionContext) IdentifierList(i int) IIdentifierListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
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

	return t.(IIdentifierListContext)
}

func (s *ConstSectionContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserEQUAL)
}

func (s *ConstSectionContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, i)
}

func (s *ConstSectionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConstSectionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ConstSectionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *ConstSectionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *ConstSectionContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOLON)
}

func (s *ConstSectionContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, i)
}

func (s *ConstSectionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ConstSectionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *ConstSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterConstSection(s)
	}
}

func (s *ConstSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitConstSection(s)
	}
}

func (s *ConstSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitConstSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ConstSection() (localctx IConstSectionContext) {
	localctx = NewConstSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, scopepascalParserRULE_constSection)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(scopepascalParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(225)
				p.IdentifierList()
			}
			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == scopepascalParserCOLON {
				{
					p.SetState(226)
					p.Match(scopepascalParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(227)
					p.Identifier()
				}

			}
			{
				p.SetState(230)
				p.Match(scopepascalParserEQUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.Expression()
			}
			{
				p.SetState(232)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(236)
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

// IResourceSectionContext is an interface to support dynamic dispatch.
type IResourceSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESOURCESTRING() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllEQUAL() []antlr.TerminalNode
	EQUAL(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsResourceSectionContext differentiates from other interfaces.
	IsResourceSectionContext()
}

type ResourceSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceSectionContext() *ResourceSectionContext {
	var p = new(ResourceSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_resourceSection
	return p
}

func InitEmptyResourceSectionContext(p *ResourceSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_resourceSection
}

func (*ResourceSectionContext) IsResourceSectionContext() {}

func NewResourceSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceSectionContext {
	var p = new(ResourceSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_resourceSection

	return p
}

func (s *ResourceSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceSectionContext) RESOURCESTRING() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRESOURCESTRING, 0)
}

func (s *ResourceSectionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ResourceSectionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *ResourceSectionContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserEQUAL)
}

func (s *ResourceSectionContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, i)
}

func (s *ResourceSectionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ResourceSectionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ResourceSectionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *ResourceSectionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *ResourceSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterResourceSection(s)
	}
}

func (s *ResourceSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitResourceSection(s)
	}
}

func (s *ResourceSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitResourceSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ResourceSection() (localctx IResourceSectionContext) {
	localctx = NewResourceSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, scopepascalParserRULE_resourceSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(scopepascalParserRESOURCESTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == scopepascalParserIDENT {
		{
			p.SetState(239)
			p.Identifier()
		}
		{
			p.SetState(240)
			p.Match(scopepascalParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Expression()
		}
		{
			p.SetState(242)
			p.Match(scopepascalParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(246)
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

// ITypeSectionContext is an interface to support dynamic dispatch.
type ITypeSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	TypeBlock() ITypeBlockContext

	// IsTypeSectionContext differentiates from other interfaces.
	IsTypeSectionContext()
}

type TypeSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSectionContext() *TypeSectionContext {
	var p = new(TypeSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_typeSection
	return p
}

func InitEmptyTypeSectionContext(p *TypeSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_typeSection
}

func (*TypeSectionContext) IsTypeSectionContext() {}

func NewTypeSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSectionContext {
	var p = new(TypeSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_typeSection

	return p
}

func (s *TypeSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSectionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserTYPE, 0)
}

func (s *TypeSectionContext) TypeBlock() ITypeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockContext)
}

func (s *TypeSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterTypeSection(s)
	}
}

func (s *TypeSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitTypeSection(s)
	}
}

func (s *TypeSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitTypeSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) TypeSection() (localctx ITypeSectionContext) {
	localctx = NewTypeSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, scopepascalParserRULE_typeSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(scopepascalParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.TypeBlock()
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

// ITypeBlockContext is an interface to support dynamic dispatch.
type ITypeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllEQUAL() []antlr.TerminalNode
	EQUAL(i int) antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsTypeBlockContext differentiates from other interfaces.
	IsTypeBlockContext()
}

type TypeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBlockContext() *TypeBlockContext {
	var p = new(TypeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_typeBlock
	return p
}

func InitEmptyTypeBlockContext(p *TypeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_typeBlock
}

func (*TypeBlockContext) IsTypeBlockContext() {}

func NewTypeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockContext {
	var p = new(TypeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_typeBlock

	return p
}

func (s *TypeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *TypeBlockContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *TypeBlockContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserEQUAL)
}

func (s *TypeBlockContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, i)
}

func (s *TypeBlockContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeBlockContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *TypeBlockContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *TypeBlockContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *TypeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterTypeBlock(s)
	}
}

func (s *TypeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitTypeBlock(s)
	}
}

func (s *TypeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitTypeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) TypeBlock() (localctx ITypeBlockContext) {
	localctx = NewTypeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, scopepascalParserRULE_typeBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(251)
				p.Identifier()
			}
			{
				p.SetState(252)
				p.Match(scopepascalParserEQUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(253)
				p.Type_()
			}
			{
				p.SetState(254)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ClassType() IClassTypeContext
	RecordType() IRecordTypeContext
	ArrayType() IArrayTypeContext
	PointerType() IPointerTypeContext
	SetType() ISetTypeContext
	FileType() IFileTypeContext
	ScalarType() IScalarTypeContext
	SubrangeType() ISubrangeTypeContext
	FunctionDeclaration() IFunctionDeclarationContext
	ProcedureDeclaration() IProcedureDeclarationContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) ClassType() IClassTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassTypeContext)
}

func (s *TypeContext) RecordType() IRecordTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordTypeContext)
}

func (s *TypeContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *TypeContext) PointerType() IPointerTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *TypeContext) SetType() ISetTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetTypeContext)
}

func (s *TypeContext) FileType() IFileTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileTypeContext)
}

func (s *TypeContext) ScalarType() IScalarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarTypeContext)
}

func (s *TypeContext) SubrangeType() ISubrangeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrangeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrangeTypeContext)
}

func (s *TypeContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *TypeContext) ProcedureDeclaration() IProcedureDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureDeclarationContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, scopepascalParserRULE_type)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.ClassType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.RecordType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(263)
			p.PointerType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(264)
			p.SetType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(265)
			p.FileType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(266)
			p.ScalarType()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(267)
			p.SubrangeType()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(268)
			p.FunctionDeclaration()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(269)
			p.ProcedureDeclaration()
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

// IAccessSpecifierContext is an interface to support dynamic dispatch.
type IAccessSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIVATE() antlr.TerminalNode
	STRICTPRIVATE() antlr.TerminalNode
	PROTECTED() antlr.TerminalNode
	STRICTPROTECTED() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode
	PUBLISHED() antlr.TerminalNode

	// IsAccessSpecifierContext differentiates from other interfaces.
	IsAccessSpecifierContext()
}

type AccessSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessSpecifierContext() *AccessSpecifierContext {
	var p = new(AccessSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_accessSpecifier
	return p
}

func InitEmptyAccessSpecifierContext(p *AccessSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_accessSpecifier
}

func (*AccessSpecifierContext) IsAccessSpecifierContext() {}

func NewAccessSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessSpecifierContext {
	var p = new(AccessSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_accessSpecifier

	return p
}

func (s *AccessSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessSpecifierContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPRIVATE, 0)
}

func (s *AccessSpecifierContext) STRICTPRIVATE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTRICTPRIVATE, 0)
}

func (s *AccessSpecifierContext) PROTECTED() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROTECTED, 0)
}

func (s *AccessSpecifierContext) STRICTPROTECTED() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTRICTPROTECTED, 0)
}

func (s *AccessSpecifierContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPUBLIC, 0)
}

func (s *AccessSpecifierContext) PUBLISHED() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPUBLISHED, 0)
}

func (s *AccessSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterAccessSpecifier(s)
	}
}

func (s *AccessSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitAccessSpecifier(s)
	}
}

func (s *AccessSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitAccessSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) AccessSpecifier() (localctx IAccessSpecifierContext) {
	localctx = NewAccessSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, scopepascalParserRULE_accessSpecifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0) {
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

// IProcedureOrFunctionModifiersContext is an interface to support dynamic dispatch.
type IProcedureOrFunctionModifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllABSTRACT() []antlr.TerminalNode
	ABSTRACT(i int) antlr.TerminalNode
	AllVIRTUAL() []antlr.TerminalNode
	VIRTUAL(i int) antlr.TerminalNode
	AllOVERRIDE() []antlr.TerminalNode
	OVERRIDE(i int) antlr.TerminalNode
	AllREINTRODUCE() []antlr.TerminalNode
	REINTRODUCE(i int) antlr.TerminalNode
	AllOVERLOAD() []antlr.TerminalNode
	OVERLOAD(i int) antlr.TerminalNode
	AllINLINE() []antlr.TerminalNode
	INLINE(i int) antlr.TerminalNode
	AllSTDCALL() []antlr.TerminalNode
	STDCALL(i int) antlr.TerminalNode
	AllCDECL() []antlr.TerminalNode
	CDECL(i int) antlr.TerminalNode
	AllSTATIC() []antlr.TerminalNode
	STATIC(i int) antlr.TerminalNode

	// IsProcedureOrFunctionModifiersContext differentiates from other interfaces.
	IsProcedureOrFunctionModifiersContext()
}

type ProcedureOrFunctionModifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureOrFunctionModifiersContext() *ProcedureOrFunctionModifiersContext {
	var p = new(ProcedureOrFunctionModifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureOrFunctionModifiers
	return p
}

func InitEmptyProcedureOrFunctionModifiersContext(p *ProcedureOrFunctionModifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureOrFunctionModifiers
}

func (*ProcedureOrFunctionModifiersContext) IsProcedureOrFunctionModifiersContext() {}

func NewProcedureOrFunctionModifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureOrFunctionModifiersContext {
	var p = new(ProcedureOrFunctionModifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_procedureOrFunctionModifiers

	return p
}

func (s *ProcedureOrFunctionModifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureOrFunctionModifiersContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *ProcedureOrFunctionModifiersContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllABSTRACT() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserABSTRACT)
}

func (s *ProcedureOrFunctionModifiersContext) ABSTRACT(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserABSTRACT, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllVIRTUAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserVIRTUAL)
}

func (s *ProcedureOrFunctionModifiersContext) VIRTUAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserVIRTUAL, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllOVERRIDE() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserOVERRIDE)
}

func (s *ProcedureOrFunctionModifiersContext) OVERRIDE(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserOVERRIDE, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllREINTRODUCE() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserREINTRODUCE)
}

func (s *ProcedureOrFunctionModifiersContext) REINTRODUCE(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserREINTRODUCE, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllOVERLOAD() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserOVERLOAD)
}

func (s *ProcedureOrFunctionModifiersContext) OVERLOAD(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserOVERLOAD, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllINLINE() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserINLINE)
}

func (s *ProcedureOrFunctionModifiersContext) INLINE(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserINLINE, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllSTDCALL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSTDCALL)
}

func (s *ProcedureOrFunctionModifiersContext) STDCALL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTDCALL, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllCDECL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCDECL)
}

func (s *ProcedureOrFunctionModifiersContext) CDECL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCDECL, i)
}

func (s *ProcedureOrFunctionModifiersContext) AllSTATIC() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSTATIC)
}

func (s *ProcedureOrFunctionModifiersContext) STATIC(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTATIC, i)
}

func (s *ProcedureOrFunctionModifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureOrFunctionModifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureOrFunctionModifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterProcedureOrFunctionModifiers(s)
	}
}

func (s *ProcedureOrFunctionModifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitProcedureOrFunctionModifiers(s)
	}
}

func (s *ProcedureOrFunctionModifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitProcedureOrFunctionModifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ProcedureOrFunctionModifiers() (localctx IProcedureOrFunctionModifiersContext) {
	localctx = NewProcedureOrFunctionModifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, scopepascalParserRULE_procedureOrFunctionModifiers)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(274)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(275)
				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&4351) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

// IClassForwardDeclarationContext is an interface to support dynamic dispatch.
type IClassForwardDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsClassForwardDeclarationContext differentiates from other interfaces.
	IsClassForwardDeclarationContext()
}

type ClassForwardDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassForwardDeclarationContext() *ClassForwardDeclarationContext {
	var p = new(ClassForwardDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_classForwardDeclaration
	return p
}

func InitEmptyClassForwardDeclarationContext(p *ClassForwardDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_classForwardDeclaration
}

func (*ClassForwardDeclarationContext) IsClassForwardDeclarationContext() {}

func NewClassForwardDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassForwardDeclarationContext {
	var p = new(ClassForwardDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_classForwardDeclaration

	return p
}

func (s *ClassForwardDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassForwardDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCLASS, 0)
}

func (s *ClassForwardDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *ClassForwardDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassForwardDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassForwardDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterClassForwardDeclaration(s)
	}
}

func (s *ClassForwardDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitClassForwardDeclaration(s)
	}
}

func (s *ClassForwardDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitClassForwardDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ClassForwardDeclaration() (localctx IClassForwardDeclarationContext) {
	localctx = NewClassForwardDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, scopepascalParserRULE_classForwardDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(scopepascalParserCLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(scopepascalParserSEMI)
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

// IClassTypeContext is an interface to support dynamic dispatch.
type IClassTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	END() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Identifier() IIdentifierContext
	RPAREN() antlr.TerminalNode
	ABSTRACT() antlr.TerminalNode
	AllBlockDeclaration() []IBlockDeclarationContext
	BlockDeclaration(i int) IBlockDeclarationContext
	AllAccessSpecifier() []IAccessSpecifierContext
	AccessSpecifier(i int) IAccessSpecifierContext
	COMMA() antlr.TerminalNode
	IdentifierList() IIdentifierListContext

	// IsClassTypeContext differentiates from other interfaces.
	IsClassTypeContext()
}

type ClassTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassTypeContext() *ClassTypeContext {
	var p = new(ClassTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_classType
	return p
}

func InitEmptyClassTypeContext(p *ClassTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_classType
}

func (*ClassTypeContext) IsClassTypeContext() {}

func NewClassTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassTypeContext {
	var p = new(ClassTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_classType

	return p
}

func (s *ClassTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassTypeContext) CLASS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCLASS, 0)
}

func (s *ClassTypeContext) END() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, 0)
}

func (s *ClassTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *ClassTypeContext) Identifier() IIdentifierContext {
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

func (s *ClassTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *ClassTypeContext) ABSTRACT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserABSTRACT, 0)
}

func (s *ClassTypeContext) AllBlockDeclaration() []IBlockDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IBlockDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockDeclarationContext); ok {
			tst[i] = t.(IBlockDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ClassTypeContext) BlockDeclaration(i int) IBlockDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDeclarationContext); ok {
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

	return t.(IBlockDeclarationContext)
}

func (s *ClassTypeContext) AllAccessSpecifier() []IAccessSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessSpecifierContext); ok {
			len++
		}
	}

	tst := make([]IAccessSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessSpecifierContext); ok {
			tst[i] = t.(IAccessSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *ClassTypeContext) AccessSpecifier(i int) IAccessSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessSpecifierContext); ok {
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

	return t.(IAccessSpecifierContext)
}

func (s *ClassTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOMMA, 0)
}

func (s *ClassTypeContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ClassTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterClassType(s)
	}
}

func (s *ClassTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitClassType(s)
	}
}

func (s *ClassTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitClassType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ClassType() (localctx IClassTypeContext) {
	localctx = NewClassTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, scopepascalParserRULE_classType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(scopepascalParserCLASS)
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

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(285)
			p.Match(scopepascalParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Identifier()
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserCOMMA {
			{
				p.SetState(287)
				p.Match(scopepascalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(288)
				p.IdentifierList()
			}

		}
		{
			p.SetState(291)
			p.Match(scopepascalParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserABSTRACT {
		{
			p.SetState(295)
			p.Match(scopepascalParserABSTRACT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140874927308888) != 0) || _la == scopepascalParserPROPERTY || _la == scopepascalParserIDENT {
		{
			p.SetState(298)
			p.BlockDeclaration()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0 {
		{
			p.SetState(304)
			p.AccessSpecifier()
		}
		{
			p.SetState(305)
			p.BlockDeclaration()
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(312)
		p.Match(scopepascalParserEND)
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

// IRecordTypeContext is an interface to support dynamic dispatch.
type IRecordTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RECORD() antlr.TerminalNode
	END() antlr.TerminalNode
	PACKED() antlr.TerminalNode
	AllBlockDeclaration() []IBlockDeclarationContext
	BlockDeclaration(i int) IBlockDeclarationContext
	AllAccessSpecifier() []IAccessSpecifierContext
	AccessSpecifier(i int) IAccessSpecifierContext
	AllRecordVariantDeclaration() []IRecordVariantDeclarationContext
	RecordVariantDeclaration(i int) IRecordVariantDeclarationContext

	// IsRecordTypeContext differentiates from other interfaces.
	IsRecordTypeContext()
}

type RecordTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordTypeContext() *RecordTypeContext {
	var p = new(RecordTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordType
	return p
}

func InitEmptyRecordTypeContext(p *RecordTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordType
}

func (*RecordTypeContext) IsRecordTypeContext() {}

func NewRecordTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeContext {
	var p = new(RecordTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_recordType

	return p
}

func (s *RecordTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeContext) RECORD() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRECORD, 0)
}

func (s *RecordTypeContext) END() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, 0)
}

func (s *RecordTypeContext) PACKED() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPACKED, 0)
}

func (s *RecordTypeContext) AllBlockDeclaration() []IBlockDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IBlockDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockDeclarationContext); ok {
			tst[i] = t.(IBlockDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) BlockDeclaration(i int) IBlockDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDeclarationContext); ok {
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

	return t.(IBlockDeclarationContext)
}

func (s *RecordTypeContext) AllAccessSpecifier() []IAccessSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessSpecifierContext); ok {
			len++
		}
	}

	tst := make([]IAccessSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessSpecifierContext); ok {
			tst[i] = t.(IAccessSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) AccessSpecifier(i int) IAccessSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessSpecifierContext); ok {
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

	return t.(IAccessSpecifierContext)
}

func (s *RecordTypeContext) AllRecordVariantDeclaration() []IRecordVariantDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecordVariantDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IRecordVariantDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecordVariantDeclarationContext); ok {
			tst[i] = t.(IRecordVariantDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) RecordVariantDeclaration(i int) IRecordVariantDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordVariantDeclarationContext); ok {
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

	return t.(IRecordVariantDeclarationContext)
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, scopepascalParserRULE_recordType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserPACKED {
		{
			p.SetState(314)
			p.Match(scopepascalParserPACKED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(317)
		p.Match(scopepascalParserRECORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140874927308888) != 0) || _la == scopepascalParserPROPERTY || _la == scopepascalParserIDENT {
		{
			p.SetState(318)
			p.BlockDeclaration()
		}

	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0 {
		{
			p.SetState(321)
			p.AccessSpecifier()
		}
		{
			p.SetState(322)
			p.BlockDeclaration()
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserCASE {
		{
			p.SetState(329)
			p.RecordVariantDeclaration()
		}

		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(335)
		p.Match(scopepascalParserEND)
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

// IPropertyDeclarationContext is an interface to support dynamic dispatch.
type IPropertyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROPERTY() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	COLON() antlr.TerminalNode
	PropertyIndexParameters() IPropertyIndexParametersContext
	PropertyReadDeclaration() IPropertyReadDeclarationContext
	PropertyWriteDeclaration() IPropertyWriteDeclarationContext
	PropertyDefaultValueDeclaration() IPropertyDefaultValueDeclarationContext
	PropertyIndexDeclaration() IPropertyIndexDeclarationContext

	// IsPropertyDeclarationContext differentiates from other interfaces.
	IsPropertyDeclarationContext()
}

type PropertyDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyDeclarationContext() *PropertyDeclarationContext {
	var p = new(PropertyDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyDeclaration
	return p
}

func InitEmptyPropertyDeclarationContext(p *PropertyDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyDeclaration
}

func (*PropertyDeclarationContext) IsPropertyDeclarationContext() {}

func NewPropertyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyDeclarationContext {
	var p = new(PropertyDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyDeclaration

	return p
}

func (s *PropertyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyDeclarationContext) PROPERTY() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROPERTY, 0)
}

func (s *PropertyDeclarationContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *PropertyDeclarationContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *PropertyDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *PropertyDeclarationContext) PropertyIndexParameters() IPropertyIndexParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyIndexParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyIndexParametersContext)
}

func (s *PropertyDeclarationContext) PropertyReadDeclaration() IPropertyReadDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyReadDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyReadDeclarationContext)
}

func (s *PropertyDeclarationContext) PropertyWriteDeclaration() IPropertyWriteDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyWriteDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyWriteDeclarationContext)
}

func (s *PropertyDeclarationContext) PropertyDefaultValueDeclaration() IPropertyDefaultValueDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyDefaultValueDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyDefaultValueDeclarationContext)
}

func (s *PropertyDeclarationContext) PropertyIndexDeclaration() IPropertyIndexDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyIndexDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyIndexDeclarationContext)
}

func (s *PropertyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyDeclaration(s)
	}
}

func (s *PropertyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyDeclaration(s)
	}
}

func (s *PropertyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyDeclaration() (localctx IPropertyDeclarationContext) {
	localctx = NewPropertyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, scopepascalParserRULE_propertyDeclaration)
	var _la int

	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(337)
			p.Match(scopepascalParserPROPERTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Identifier()
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserLBRACK {
			{
				p.SetState(339)
				p.PropertyIndexParameters()
			}

		}
		{
			p.SetState(342)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Identifier()
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserREAD {
			{
				p.SetState(344)
				p.PropertyReadDeclaration()
			}

		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserWRITE {
			{
				p.SetState(347)
				p.PropertyWriteDeclaration()
			}

		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserDEFAULT {
			{
				p.SetState(350)
				p.PropertyDefaultValueDeclaration()
			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserINDEX {
			{
				p.SetState(353)
				p.PropertyIndexDeclaration()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(scopepascalParserPROPERTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Identifier()
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserDEFAULT {
			{
				p.SetState(358)
				p.PropertyDefaultValueDeclaration()
			}

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

// IPropertyReadDeclarationContext is an interface to support dynamic dispatch.
type IPropertyReadDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsPropertyReadDeclarationContext differentiates from other interfaces.
	IsPropertyReadDeclarationContext()
}

type PropertyReadDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyReadDeclarationContext() *PropertyReadDeclarationContext {
	var p = new(PropertyReadDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyReadDeclaration
	return p
}

func InitEmptyPropertyReadDeclarationContext(p *PropertyReadDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyReadDeclaration
}

func (*PropertyReadDeclarationContext) IsPropertyReadDeclarationContext() {}

func NewPropertyReadDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyReadDeclarationContext {
	var p = new(PropertyReadDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyReadDeclaration

	return p
}

func (s *PropertyReadDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyReadDeclarationContext) READ() antlr.TerminalNode {
	return s.GetToken(scopepascalParserREAD, 0)
}

func (s *PropertyReadDeclarationContext) Identifier() IIdentifierContext {
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

func (s *PropertyReadDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyReadDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyReadDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyReadDeclaration(s)
	}
}

func (s *PropertyReadDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyReadDeclaration(s)
	}
}

func (s *PropertyReadDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyReadDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyReadDeclaration() (localctx IPropertyReadDeclarationContext) {
	localctx = NewPropertyReadDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, scopepascalParserRULE_propertyReadDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(scopepascalParserREAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Identifier()
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

// IPropertyWriteDeclarationContext is an interface to support dynamic dispatch.
type IPropertyWriteDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WRITE() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsPropertyWriteDeclarationContext differentiates from other interfaces.
	IsPropertyWriteDeclarationContext()
}

type PropertyWriteDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyWriteDeclarationContext() *PropertyWriteDeclarationContext {
	var p = new(PropertyWriteDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyWriteDeclaration
	return p
}

func InitEmptyPropertyWriteDeclarationContext(p *PropertyWriteDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyWriteDeclaration
}

func (*PropertyWriteDeclarationContext) IsPropertyWriteDeclarationContext() {}

func NewPropertyWriteDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyWriteDeclarationContext {
	var p = new(PropertyWriteDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyWriteDeclaration

	return p
}

func (s *PropertyWriteDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyWriteDeclarationContext) WRITE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserWRITE, 0)
}

func (s *PropertyWriteDeclarationContext) Identifier() IIdentifierContext {
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

func (s *PropertyWriteDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyWriteDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyWriteDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyWriteDeclaration(s)
	}
}

func (s *PropertyWriteDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyWriteDeclaration(s)
	}
}

func (s *PropertyWriteDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyWriteDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyWriteDeclaration() (localctx IPropertyWriteDeclarationContext) {
	localctx = NewPropertyWriteDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, scopepascalParserRULE_propertyWriteDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(scopepascalParserWRITE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Identifier()
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

// IPropertyDefaultValueDeclarationContext is an interface to support dynamic dispatch.
type IPropertyDefaultValueDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPropertyDefaultValueDeclarationContext differentiates from other interfaces.
	IsPropertyDefaultValueDeclarationContext()
}

type PropertyDefaultValueDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyDefaultValueDeclarationContext() *PropertyDefaultValueDeclarationContext {
	var p = new(PropertyDefaultValueDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyDefaultValueDeclaration
	return p
}

func InitEmptyPropertyDefaultValueDeclarationContext(p *PropertyDefaultValueDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyDefaultValueDeclaration
}

func (*PropertyDefaultValueDeclarationContext) IsPropertyDefaultValueDeclarationContext() {}

func NewPropertyDefaultValueDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyDefaultValueDeclarationContext {
	var p = new(PropertyDefaultValueDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyDefaultValueDeclaration

	return p
}

func (s *PropertyDefaultValueDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyDefaultValueDeclarationContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDEFAULT, 0)
}

func (s *PropertyDefaultValueDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropertyDefaultValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyDefaultValueDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyDefaultValueDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyDefaultValueDeclaration(s)
	}
}

func (s *PropertyDefaultValueDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyDefaultValueDeclaration(s)
	}
}

func (s *PropertyDefaultValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyDefaultValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyDefaultValueDeclaration() (localctx IPropertyDefaultValueDeclarationContext) {
	localctx = NewPropertyDefaultValueDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, scopepascalParserRULE_propertyDefaultValueDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(scopepascalParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Expression()
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

// IPropertyIndexDeclarationContext is an interface to support dynamic dispatch.
type IPropertyIndexDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode
	Number() INumberContext

	// IsPropertyIndexDeclarationContext differentiates from other interfaces.
	IsPropertyIndexDeclarationContext()
}

type PropertyIndexDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyIndexDeclarationContext() *PropertyIndexDeclarationContext {
	var p = new(PropertyIndexDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexDeclaration
	return p
}

func InitEmptyPropertyIndexDeclarationContext(p *PropertyIndexDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexDeclaration
}

func (*PropertyIndexDeclarationContext) IsPropertyIndexDeclarationContext() {}

func NewPropertyIndexDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyIndexDeclarationContext {
	var p = new(PropertyIndexDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyIndexDeclaration

	return p
}

func (s *PropertyIndexDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyIndexDeclarationContext) INDEX() antlr.TerminalNode {
	return s.GetToken(scopepascalParserINDEX, 0)
}

func (s *PropertyIndexDeclarationContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PropertyIndexDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyIndexDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyIndexDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyIndexDeclaration(s)
	}
}

func (s *PropertyIndexDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyIndexDeclaration(s)
	}
}

func (s *PropertyIndexDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyIndexDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyIndexDeclaration() (localctx IPropertyIndexDeclarationContext) {
	localctx = NewPropertyIndexDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, scopepascalParserRULE_propertyIndexDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(scopepascalParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
		p.Number()
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

// IPropertyIndexParametersContext is an interface to support dynamic dispatch.
type IPropertyIndexParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	PropertyIndexParametersList() IPropertyIndexParametersListContext
	RBRACK() antlr.TerminalNode

	// IsPropertyIndexParametersContext differentiates from other interfaces.
	IsPropertyIndexParametersContext()
}

type PropertyIndexParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyIndexParametersContext() *PropertyIndexParametersContext {
	var p = new(PropertyIndexParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexParameters
	return p
}

func InitEmptyPropertyIndexParametersContext(p *PropertyIndexParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexParameters
}

func (*PropertyIndexParametersContext) IsPropertyIndexParametersContext() {}

func NewPropertyIndexParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyIndexParametersContext {
	var p = new(PropertyIndexParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyIndexParameters

	return p
}

func (s *PropertyIndexParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyIndexParametersContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLBRACK, 0)
}

func (s *PropertyIndexParametersContext) PropertyIndexParametersList() IPropertyIndexParametersListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyIndexParametersListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyIndexParametersListContext)
}

func (s *PropertyIndexParametersContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRBRACK, 0)
}

func (s *PropertyIndexParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyIndexParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyIndexParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyIndexParameters(s)
	}
}

func (s *PropertyIndexParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyIndexParameters(s)
	}
}

func (s *PropertyIndexParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyIndexParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyIndexParameters() (localctx IPropertyIndexParametersContext) {
	localctx = NewPropertyIndexParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, scopepascalParserRULE_propertyIndexParameters)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(scopepascalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.PropertyIndexParametersList()
	}
	{
		p.SetState(377)
		p.Match(scopepascalParserRBRACK)
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

// IPropertyIndexParametersListContext is an interface to support dynamic dispatch.
type IPropertyIndexParametersListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifierList() []IIdentifierListContext
	IdentifierList(i int) IIdentifierListContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsPropertyIndexParametersListContext differentiates from other interfaces.
	IsPropertyIndexParametersListContext()
}

type PropertyIndexParametersListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyIndexParametersListContext() *PropertyIndexParametersListContext {
	var p = new(PropertyIndexParametersListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexParametersList
	return p
}

func InitEmptyPropertyIndexParametersListContext(p *PropertyIndexParametersListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_propertyIndexParametersList
}

func (*PropertyIndexParametersListContext) IsPropertyIndexParametersListContext() {}

func NewPropertyIndexParametersListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyIndexParametersListContext {
	var p = new(PropertyIndexParametersListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_propertyIndexParametersList

	return p
}

func (s *PropertyIndexParametersListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyIndexParametersListContext) AllIdentifierList() []IIdentifierListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierListContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierListContext); ok {
			tst[i] = t.(IIdentifierListContext)
			i++
		}
	}

	return tst
}

func (s *PropertyIndexParametersListContext) IdentifierList(i int) IIdentifierListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
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

	return t.(IIdentifierListContext)
}

func (s *PropertyIndexParametersListContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOLON)
}

func (s *PropertyIndexParametersListContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, i)
}

func (s *PropertyIndexParametersListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PropertyIndexParametersListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PropertyIndexParametersListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *PropertyIndexParametersListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *PropertyIndexParametersListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyIndexParametersListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyIndexParametersListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPropertyIndexParametersList(s)
	}
}

func (s *PropertyIndexParametersListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPropertyIndexParametersList(s)
	}
}

func (s *PropertyIndexParametersListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPropertyIndexParametersList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PropertyIndexParametersList() (localctx IPropertyIndexParametersListContext) {
	localctx = NewPropertyIndexParametersListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, scopepascalParserRULE_propertyIndexParametersList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.IdentifierList()
	}
	{
		p.SetState(380)
		p.Match(scopepascalParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Expression()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserSEMI {
		{
			p.SetState(382)
			p.Match(scopepascalParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.IdentifierList()
		}
		{
			p.SetState(384)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Expression()
		}

		p.SetState(391)
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

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	IdentifierList() IIdentifierListContext
	RBRACK() antlr.TerminalNode
	OF() antlr.TerminalNode
	Type_() ITypeContext
	Identifier() IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	DOUBLEDOT() antlr.TerminalNode
	CONST() antlr.TerminalNode

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(scopepascalParserARRAY, 0)
}

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLBRACK, 0)
}

func (s *ArrayTypeContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRBRACK, 0)
}

func (s *ArrayTypeContext) OF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOF, 0)
}

func (s *ArrayTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayTypeContext) Identifier() IIdentifierContext {
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

func (s *ArrayTypeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayTypeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArrayTypeContext) DOUBLEDOT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDOUBLEDOT, 0)
}

func (s *ArrayTypeContext) CONST() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCONST, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, scopepascalParserRULE_arrayType)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Match(scopepascalParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(scopepascalParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.IdentifierList()
		}
		{
			p.SetState(395)
			p.Match(scopepascalParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(scopepascalParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(397)
				p.Type_()
			}

		case 2:
			{
				p.SetState(398)
				p.Identifier()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.Match(scopepascalParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(scopepascalParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Expression()
		}
		{
			p.SetState(404)
			p.Match(scopepascalParserDOUBLEDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Expression()
		}
		{
			p.SetState(406)
			p.Match(scopepascalParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(407)
				p.Type_()
			}

		case 2:
			{
				p.SetState(408)
				p.Identifier()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.Match(scopepascalParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(scopepascalParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(scopepascalParserCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(414)
			p.Match(scopepascalParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Match(scopepascalParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(416)
				p.Type_()
			}

		case 2:
			{
				p.SetState(417)
				p.Identifier()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
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

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEREFERENCE() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_pointerType
	return p
}

func InitEmptyPointerTypeContext(p *PointerTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_pointerType
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) DEREFERENCE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDEREFERENCE, 0)
}

func (s *PointerTypeContext) Identifier() IIdentifierContext {
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

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterPointerType(s)
	}
}

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitPointerType(s)
	}
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitPointerType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, scopepascalParserRULE_pointerType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(scopepascalParserDEREFERENCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.Identifier()
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

// ISetTypeContext is an interface to support dynamic dispatch.
type ISetTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	OF() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsSetTypeContext differentiates from other interfaces.
	IsSetTypeContext()
}

type SetTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetTypeContext() *SetTypeContext {
	var p = new(SetTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_setType
	return p
}

func InitEmptySetTypeContext(p *SetTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_setType
}

func (*SetTypeContext) IsSetTypeContext() {}

func NewSetTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetTypeContext {
	var p = new(SetTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_setType

	return p
}

func (s *SetTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SetTypeContext) SET() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSET, 0)
}

func (s *SetTypeContext) OF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOF, 0)
}

func (s *SetTypeContext) Identifier() IIdentifierContext {
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

func (s *SetTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterSetType(s)
	}
}

func (s *SetTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitSetType(s)
	}
}

func (s *SetTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitSetType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) SetType() (localctx ISetTypeContext) {
	localctx = NewSetTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, scopepascalParserRULE_setType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(scopepascalParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(scopepascalParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Identifier()
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

// IFileTypeContext is an interface to support dynamic dispatch.
type IFileTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILE() antlr.TerminalNode
	OF() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsFileTypeContext differentiates from other interfaces.
	IsFileTypeContext()
}

type FileTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileTypeContext() *FileTypeContext {
	var p = new(FileTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_fileType
	return p
}

func InitEmptyFileTypeContext(p *FileTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_fileType
}

func (*FileTypeContext) IsFileTypeContext() {}

func NewFileTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileTypeContext {
	var p = new(FileTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_fileType

	return p
}

func (s *FileTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FileTypeContext) FILE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFILE, 0)
}

func (s *FileTypeContext) OF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOF, 0)
}

func (s *FileTypeContext) Identifier() IIdentifierContext {
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

func (s *FileTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFileType(s)
	}
}

func (s *FileTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFileType(s)
	}
}

func (s *FileTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFileType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FileType() (localctx IFileTypeContext) {
	localctx = NewFileTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, scopepascalParserRULE_fileType)
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.Match(scopepascalParserFILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(scopepascalParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.Match(scopepascalParserFILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IScalarTypeContext is an interface to support dynamic dispatch.
type IScalarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	ExpressionList() IExpressionListContext
	RPAREN() antlr.TerminalNode

	// IsScalarTypeContext differentiates from other interfaces.
	IsScalarTypeContext()
}

type ScalarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeContext() *ScalarTypeContext {
	var p = new(ScalarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_scalarType
	return p
}

func InitEmptyScalarTypeContext(p *ScalarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_scalarType
}

func (*ScalarTypeContext) IsScalarTypeContext() {}

func NewScalarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeContext {
	var p = new(ScalarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_scalarType

	return p
}

func (s *ScalarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *ScalarTypeContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ScalarTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *ScalarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterScalarType(s)
	}
}

func (s *ScalarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitScalarType(s)
	}
}

func (s *ScalarTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitScalarType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ScalarType() (localctx IScalarTypeContext) {
	localctx = NewScalarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, scopepascalParserRULE_scalarType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Match(scopepascalParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)
		p.ExpressionList()
	}
	{
		p.SetState(437)
		p.Match(scopepascalParserRPAREN)
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

// ISubrangeTypeContext is an interface to support dynamic dispatch.
type ISubrangeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	DOUBLEDOT() antlr.TerminalNode

	// IsSubrangeTypeContext differentiates from other interfaces.
	IsSubrangeTypeContext()
}

type SubrangeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubrangeTypeContext() *SubrangeTypeContext {
	var p = new(SubrangeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_subrangeType
	return p
}

func InitEmptySubrangeTypeContext(p *SubrangeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_subrangeType
}

func (*SubrangeTypeContext) IsSubrangeTypeContext() {}

func NewSubrangeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubrangeTypeContext {
	var p = new(SubrangeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_subrangeType

	return p
}

func (s *SubrangeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubrangeTypeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubrangeTypeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *SubrangeTypeContext) DOUBLEDOT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDOUBLEDOT, 0)
}

func (s *SubrangeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubrangeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubrangeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterSubrangeType(s)
	}
}

func (s *SubrangeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitSubrangeType(s)
	}
}

func (s *SubrangeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitSubrangeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) SubrangeType() (localctx ISubrangeTypeContext) {
	localctx = NewSubrangeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, scopepascalParserRULE_subrangeType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Expression()
	}
	{
		p.SetState(440)
		p.Match(scopepascalParserDOUBLEDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Expression()
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

// IBlockDeclarationContext is an interface to support dynamic dispatch.
type IBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSection() ITypeSectionContext
	ConstSection() IConstSectionContext
	VarDeclaration() IVarDeclarationContext
	SEMI() antlr.TerminalNode
	FunctionOrProcedureDeclaration() IFunctionOrProcedureDeclarationContext
	PropertyDeclaration() IPropertyDeclarationContext

	// IsBlockDeclarationContext differentiates from other interfaces.
	IsBlockDeclarationContext()
}

type BlockDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockDeclarationContext() *BlockDeclarationContext {
	var p = new(BlockDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_blockDeclaration
	return p
}

func InitEmptyBlockDeclarationContext(p *BlockDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_blockDeclaration
}

func (*BlockDeclarationContext) IsBlockDeclarationContext() {}

func NewBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockDeclarationContext {
	var p = new(BlockDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_blockDeclaration

	return p
}

func (s *BlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockDeclarationContext) TypeSection() ITypeSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSectionContext)
}

func (s *BlockDeclarationContext) ConstSection() IConstSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstSectionContext)
}

func (s *BlockDeclarationContext) VarDeclaration() IVarDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *BlockDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *BlockDeclarationContext) FunctionOrProcedureDeclaration() IFunctionOrProcedureDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionOrProcedureDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionOrProcedureDeclarationContext)
}

func (s *BlockDeclarationContext) PropertyDeclaration() IPropertyDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyDeclarationContext)
}

func (s *BlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterBlockDeclaration(s)
	}
}

func (s *BlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitBlockDeclaration(s)
	}
}

func (s *BlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) BlockDeclaration() (localctx IBlockDeclarationContext) {
	localctx = NewBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, scopepascalParserRULE_blockDeclaration)
	var _la int

	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case scopepascalParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.TypeSection()
		}

	case scopepascalParserCONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.ConstSection()
		}

	case scopepascalParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(445)
			p.VarDeclaration()
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserSEMI {
			{
				p.SetState(446)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case scopepascalParserFUNCTION, scopepascalParserPROCEDURE, scopepascalParserCLASS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(449)
			p.FunctionOrProcedureDeclaration()
		}

	case scopepascalParserPROPERTY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(450)
			p.PropertyDeclaration()
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

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	ProcedureOrFunctionModifiers() IProcedureOrFunctionModifiersContext
	SEMI() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	ParamsDeclaration() IParamsDeclarationContext
	OFOBJECT() antlr.TerminalNode

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserIDENT)
}

func (s *FunctionDeclarationContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserIDENT, i)
}

func (s *FunctionDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *FunctionDeclarationContext) ProcedureOrFunctionModifiers() IProcedureOrFunctionModifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureOrFunctionModifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureOrFunctionModifiersContext)
}

func (s *FunctionDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *FunctionDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCLASS, 0)
}

func (s *FunctionDeclarationContext) ParamsDeclaration() IParamsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclarationContext)
}

func (s *FunctionDeclarationContext) OFOBJECT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOFOBJECT, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, scopepascalParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserCLASS {
		{
			p.SetState(453)
			p.Match(scopepascalParserCLASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(456)
		p.Match(scopepascalParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(scopepascalParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(458)
			p.ParamsDeclaration()
		}

	}
	{
		p.SetState(461)
		p.Match(scopepascalParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.Match(scopepascalParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserOFOBJECT {
		{
			p.SetState(463)
			p.Match(scopepascalParserOFOBJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(466)
		p.ProcedureOrFunctionModifiers()
	}
	{
		p.SetState(467)
		p.Match(scopepascalParserSEMI)
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

// IProcedureDeclarationContext is an interface to support dynamic dispatch.
type IProcedureDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDURE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ProcedureOrFunctionModifiers() IProcedureOrFunctionModifiersContext
	SEMI() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	ParamsDeclaration() IParamsDeclarationContext
	OFOBJECT() antlr.TerminalNode

	// IsProcedureDeclarationContext differentiates from other interfaces.
	IsProcedureDeclarationContext()
}

type ProcedureDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureDeclarationContext() *ProcedureDeclarationContext {
	var p = new(ProcedureDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureDeclaration
	return p
}

func InitEmptyProcedureDeclarationContext(p *ProcedureDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureDeclaration
}

func (*ProcedureDeclarationContext) IsProcedureDeclarationContext() {}

func NewProcedureDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureDeclarationContext {
	var p = new(ProcedureDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_procedureDeclaration

	return p
}

func (s *ProcedureDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureDeclarationContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROCEDURE, 0)
}

func (s *ProcedureDeclarationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserIDENT, 0)
}

func (s *ProcedureDeclarationContext) ProcedureOrFunctionModifiers() IProcedureOrFunctionModifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureOrFunctionModifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureOrFunctionModifiersContext)
}

func (s *ProcedureDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *ProcedureDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCLASS, 0)
}

func (s *ProcedureDeclarationContext) ParamsDeclaration() IParamsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclarationContext)
}

func (s *ProcedureDeclarationContext) OFOBJECT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOFOBJECT, 0)
}

func (s *ProcedureDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterProcedureDeclaration(s)
	}
}

func (s *ProcedureDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitProcedureDeclaration(s)
	}
}

func (s *ProcedureDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitProcedureDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ProcedureDeclaration() (localctx IProcedureDeclarationContext) {
	localctx = NewProcedureDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, scopepascalParserRULE_procedureDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserCLASS {
		{
			p.SetState(469)
			p.Match(scopepascalParserCLASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(472)
		p.Match(scopepascalParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
		p.Match(scopepascalParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(474)
			p.ParamsDeclaration()
		}

	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserOFOBJECT {
		{
			p.SetState(477)
			p.Match(scopepascalParserOFOBJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(480)
		p.ProcedureOrFunctionModifiers()
	}
	{
		p.SetState(481)
		p.Match(scopepascalParserSEMI)
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

// IFunctionOrProcedureDeclarationContext is an interface to support dynamic dispatch.
type IFunctionOrProcedureDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDeclaration() IFunctionDeclarationContext
	ProcedureDeclaration() IProcedureDeclarationContext

	// IsFunctionOrProcedureDeclarationContext differentiates from other interfaces.
	IsFunctionOrProcedureDeclarationContext()
}

type FunctionOrProcedureDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionOrProcedureDeclarationContext() *FunctionOrProcedureDeclarationContext {
	var p = new(FunctionOrProcedureDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionOrProcedureDeclaration
	return p
}

func InitEmptyFunctionOrProcedureDeclarationContext(p *FunctionOrProcedureDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionOrProcedureDeclaration
}

func (*FunctionOrProcedureDeclarationContext) IsFunctionOrProcedureDeclarationContext() {}

func NewFunctionOrProcedureDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionOrProcedureDeclarationContext {
	var p = new(FunctionOrProcedureDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_functionOrProcedureDeclaration

	return p
}

func (s *FunctionOrProcedureDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionOrProcedureDeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *FunctionOrProcedureDeclarationContext) ProcedureDeclaration() IProcedureDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureDeclarationContext)
}

func (s *FunctionOrProcedureDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOrProcedureDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionOrProcedureDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFunctionOrProcedureDeclaration(s)
	}
}

func (s *FunctionOrProcedureDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFunctionOrProcedureDeclaration(s)
	}
}

func (s *FunctionOrProcedureDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFunctionOrProcedureDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FunctionOrProcedureDeclaration() (localctx IFunctionOrProcedureDeclarationContext) {
	localctx = NewFunctionOrProcedureDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, scopepascalParserRULE_functionOrProcedureDeclaration)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(484)
			p.ProcedureDeclaration()
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

// IFunctionOrProcedureContext is an interface to support dynamic dispatch.
type IFunctionOrProcedureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionOrProcedureDeclaration() IFunctionOrProcedureDeclarationContext
	BlockStatement() IBlockStatementContext
	SEMI() antlr.TerminalNode
	AllFunctionOrProcedure() []IFunctionOrProcedureContext
	FunctionOrProcedure(i int) IFunctionOrProcedureContext
	AllVarSection() []IVarSectionContext
	VarSection(i int) IVarSectionContext
	FORWARD() antlr.TerminalNode

	// IsFunctionOrProcedureContext differentiates from other interfaces.
	IsFunctionOrProcedureContext()
}

type FunctionOrProcedureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionOrProcedureContext() *FunctionOrProcedureContext {
	var p = new(FunctionOrProcedureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionOrProcedure
	return p
}

func InitEmptyFunctionOrProcedureContext(p *FunctionOrProcedureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionOrProcedure
}

func (*FunctionOrProcedureContext) IsFunctionOrProcedureContext() {}

func NewFunctionOrProcedureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionOrProcedureContext {
	var p = new(FunctionOrProcedureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_functionOrProcedure

	return p
}

func (s *FunctionOrProcedureContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionOrProcedureContext) FunctionOrProcedureDeclaration() IFunctionOrProcedureDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionOrProcedureDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionOrProcedureDeclarationContext)
}

func (s *FunctionOrProcedureContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunctionOrProcedureContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *FunctionOrProcedureContext) AllFunctionOrProcedure() []IFunctionOrProcedureContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionOrProcedureContext); ok {
			len++
		}
	}

	tst := make([]IFunctionOrProcedureContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionOrProcedureContext); ok {
			tst[i] = t.(IFunctionOrProcedureContext)
			i++
		}
	}

	return tst
}

func (s *FunctionOrProcedureContext) FunctionOrProcedure(i int) IFunctionOrProcedureContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionOrProcedureContext); ok {
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

	return t.(IFunctionOrProcedureContext)
}

func (s *FunctionOrProcedureContext) AllVarSection() []IVarSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarSectionContext); ok {
			len++
		}
	}

	tst := make([]IVarSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarSectionContext); ok {
			tst[i] = t.(IVarSectionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionOrProcedureContext) VarSection(i int) IVarSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
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

	return t.(IVarSectionContext)
}

func (s *FunctionOrProcedureContext) FORWARD() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFORWARD, 0)
}

func (s *FunctionOrProcedureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOrProcedureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionOrProcedureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFunctionOrProcedure(s)
	}
}

func (s *FunctionOrProcedureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFunctionOrProcedure(s)
	}
}

func (s *FunctionOrProcedureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFunctionOrProcedure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FunctionOrProcedure() (localctx IFunctionOrProcedureContext) {
	localctx = NewFunctionOrProcedureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, scopepascalParserRULE_functionOrProcedure)
	var _la int

	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(487)
			p.FunctionOrProcedureDeclaration()
		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737488355384) != 0 {
			p.SetState(490)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case scopepascalParserFUNCTION, scopepascalParserPROCEDURE, scopepascalParserCLASS:
				{
					p.SetState(488)
					p.FunctionOrProcedure()
				}

			case scopepascalParserVAR:
				{
					p.SetState(489)
					p.VarSection()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(494)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(495)
			p.BlockStatement()
		}
		{
			p.SetState(496)
			p.Match(scopepascalParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
			p.FunctionOrProcedureDeclaration()
		}
		{
			p.SetState(499)
			p.Match(scopepascalParserFORWARD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.Match(scopepascalParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN() antlr.TerminalNode
	END() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllStatementError() []IStatementErrorContext
	StatementError(i int) IStatementErrorContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserBEGIN, 0)
}

func (s *BlockStatementContext) END() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, 0)
}

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockStatementContext) AllStatementError() []IStatementErrorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementErrorContext); ok {
			len++
		}
	}

	tst := make([]IStatementErrorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementErrorContext); ok {
			tst[i] = t.(IStatementErrorContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) StatementError(i int) IStatementErrorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementErrorContext); ok {
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

	return t.(IStatementErrorContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, scopepascalParserRULE_blockStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.Match(scopepascalParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-1032) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&33554431) != 0) {
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(505)
				p.Statement()
			}

		case 2:
			{
				p.SetState(506)
				p.StatementError()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(512)
		p.Match(scopepascalParserEND)
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

// IRecordVariantDeclarationContext is an interface to support dynamic dispatch.
type IRecordVariantDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	OF() antlr.TerminalNode
	AllRecordVariant() []IRecordVariantContext
	RecordVariant(i int) IRecordVariantContext
	COLON() antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsRecordVariantDeclarationContext differentiates from other interfaces.
	IsRecordVariantDeclarationContext()
}

type RecordVariantDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordVariantDeclarationContext() *RecordVariantDeclarationContext {
	var p = new(RecordVariantDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordVariantDeclaration
	return p
}

func InitEmptyRecordVariantDeclarationContext(p *RecordVariantDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordVariantDeclaration
}

func (*RecordVariantDeclarationContext) IsRecordVariantDeclarationContext() {}

func NewRecordVariantDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordVariantDeclarationContext {
	var p = new(RecordVariantDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_recordVariantDeclaration

	return p
}

func (s *RecordVariantDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordVariantDeclarationContext) CASE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCASE, 0)
}

func (s *RecordVariantDeclarationContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RecordVariantDeclarationContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *RecordVariantDeclarationContext) OF() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOF, 0)
}

func (s *RecordVariantDeclarationContext) AllRecordVariant() []IRecordVariantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecordVariantContext); ok {
			len++
		}
	}

	tst := make([]IRecordVariantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecordVariantContext); ok {
			tst[i] = t.(IRecordVariantContext)
			i++
		}
	}

	return tst
}

func (s *RecordVariantDeclarationContext) RecordVariant(i int) IRecordVariantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordVariantContext); ok {
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

	return t.(IRecordVariantContext)
}

func (s *RecordVariantDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *RecordVariantDeclarationContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *RecordVariantDeclarationContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *RecordVariantDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordVariantDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordVariantDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterRecordVariantDeclaration(s)
	}
}

func (s *RecordVariantDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitRecordVariantDeclaration(s)
	}
}

func (s *RecordVariantDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitRecordVariantDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) RecordVariantDeclaration() (localctx IRecordVariantDeclarationContext) {
	localctx = NewRecordVariantDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, scopepascalParserRULE_recordVariantDeclaration)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(514)
		p.Match(scopepascalParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)
		p.Identifier()
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserCOLON {
		{
			p.SetState(516)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Identifier()
		}

	}
	{
		p.SetState(520)
		p.Match(scopepascalParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(521)
		p.RecordVariant()
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(522)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(523)
				p.RecordVariant()
			}

		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserSEMI {
		{
			p.SetState(529)
			p.Match(scopepascalParserSEMI)
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

// IRecordVariantContext is an interface to support dynamic dispatch.
type IRecordVariantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	COLON() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext

	// IsRecordVariantContext differentiates from other interfaces.
	IsRecordVariantContext()
}

type RecordVariantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordVariantContext() *RecordVariantContext {
	var p = new(RecordVariantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordVariant
	return p
}

func InitEmptyRecordVariantContext(p *RecordVariantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_recordVariant
}

func (*RecordVariantContext) IsRecordVariantContext() {}

func NewRecordVariantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordVariantContext {
	var p = new(RecordVariantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_recordVariant

	return p
}

func (s *RecordVariantContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordVariantContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *RecordVariantContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *RecordVariantContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *RecordVariantContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *RecordVariantContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *RecordVariantContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
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

	return t.(IVarDeclarationContext)
}

func (s *RecordVariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordVariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordVariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterRecordVariant(s)
	}
}

func (s *RecordVariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitRecordVariant(s)
	}
}

func (s *RecordVariantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitRecordVariant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) RecordVariant() (localctx IRecordVariantContext) {
	localctx = NewRecordVariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, scopepascalParserRULE_recordVariant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.IdentifierList()
	}
	{
		p.SetState(533)
		p.Match(scopepascalParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(534)
		p.Match(scopepascalParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == scopepascalParserIDENT {
		{
			p.SetState(535)
			p.VarDeclaration()
		}

		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(540)
		p.Match(scopepascalParserRPAREN)
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

// IStatementErrorContext is an interface to support dynamic dispatch.
type IStatementErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEND() []antlr.TerminalNode
	END(i int) antlr.TerminalNode
	AllBEGIN() []antlr.TerminalNode
	BEGIN(i int) antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsStatementErrorContext differentiates from other interfaces.
	IsStatementErrorContext()
}

type StatementErrorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementErrorContext() *StatementErrorContext {
	var p = new(StatementErrorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_statementError
	return p
}

func InitEmptyStatementErrorContext(p *StatementErrorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_statementError
}

func (*StatementErrorContext) IsStatementErrorContext() {}

func NewStatementErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementErrorContext {
	var p = new(StatementErrorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_statementError

	return p
}

func (s *StatementErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementErrorContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserEND)
}

func (s *StatementErrorContext) END(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, i)
}

func (s *StatementErrorContext) AllBEGIN() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserBEGIN)
}

func (s *StatementErrorContext) BEGIN(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserBEGIN, i)
}

func (s *StatementErrorContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *StatementErrorContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *StatementErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterStatementError(s)
	}
}

func (s *StatementErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitStatementError(s)
	}
}

func (s *StatementErrorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitStatementError(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) StatementError() (localctx IStatementErrorContext) {
	localctx = NewStatementErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, scopepascalParserRULE_statementError)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(542)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarSection() IVarSectionContext
	InlinedVarDeclaration() IInlinedVarDeclarationContext
	StatementError() IStatementErrorContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarSection() IVarSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarSectionContext)
}

func (s *StatementContext) InlinedVarDeclaration() IInlinedVarDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlinedVarDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlinedVarDeclarationContext)
}

func (s *StatementContext) StatementError() IStatementErrorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementErrorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementErrorContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, scopepascalParserRULE_statement)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(547)
			p.VarSection()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)
			p.InlinedVarDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(549)
			p.StatementError()
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

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
	p.RuleIndex = scopepascalParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserIDENT)
}

func (s *IdentifierContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserIDENT, i)
}

func (s *IdentifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserDOT)
}

func (s *IdentifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserDOT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, scopepascalParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(552)
		p.Match(scopepascalParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserDOT {
		{
			p.SetState(553)
			p.Match(scopepascalParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.Match(scopepascalParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(559)
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

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierListContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, scopepascalParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)
		p.Identifier()
	}
	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserCOMMA {
		{
			p.SetState(561)
			p.Match(scopepascalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Identifier()
		}

		p.SetState(567)
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, scopepascalParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(568)
		p.Expression()
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserCOMMA {
		{
			p.SetState(569)
			p.Match(scopepascalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.Expression()
		}

		p.SetState(575)
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

// IParamsDeclarationContext is an interface to support dynamic dispatch.
type IParamsDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllParamsDeclarationSection() []IParamsDeclarationSectionContext
	ParamsDeclarationSection(i int) IParamsDeclarationSectionContext
	RPAREN() antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsParamsDeclarationContext differentiates from other interfaces.
	IsParamsDeclarationContext()
}

type ParamsDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsDeclarationContext() *ParamsDeclarationContext {
	var p = new(ParamsDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramsDeclaration
	return p
}

func InitEmptyParamsDeclarationContext(p *ParamsDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramsDeclaration
}

func (*ParamsDeclarationContext) IsParamsDeclarationContext() {}

func NewParamsDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsDeclarationContext {
	var p = new(ParamsDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_paramsDeclaration

	return p
}

func (s *ParamsDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *ParamsDeclarationContext) AllParamsDeclarationSection() []IParamsDeclarationSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamsDeclarationSectionContext); ok {
			len++
		}
	}

	tst := make([]IParamsDeclarationSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamsDeclarationSectionContext); ok {
			tst[i] = t.(IParamsDeclarationSectionContext)
			i++
		}
	}

	return tst
}

func (s *ParamsDeclarationContext) ParamsDeclarationSection(i int) IParamsDeclarationSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationSectionContext); ok {
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

	return t.(IParamsDeclarationSectionContext)
}

func (s *ParamsDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *ParamsDeclarationContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *ParamsDeclarationContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *ParamsDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterParamsDeclaration(s)
	}
}

func (s *ParamsDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitParamsDeclaration(s)
	}
}

func (s *ParamsDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitParamsDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ParamsDeclaration() (localctx IParamsDeclarationContext) {
	localctx = NewParamsDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, scopepascalParserRULE_paramsDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.Match(scopepascalParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(577)
		p.ParamsDeclarationSection()
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == scopepascalParserSEMI {
		{
			p.SetState(578)
			p.Match(scopepascalParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.ParamsDeclarationSection()
		}

		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(585)
		p.Match(scopepascalParserRPAREN)
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

// IParamsDeclarationSectionContext is an interface to support dynamic dispatch.
type IParamsDeclarationSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	ParamSpecifier() IParamSpecifierContext
	COLON() antlr.TerminalNode
	Identifier() IIdentifierContext
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext
	FUNCTION() antlr.TerminalNode
	ParamsDeclaration() IParamsDeclarationContext
	PROCEDURE() antlr.TerminalNode

	// IsParamsDeclarationSectionContext differentiates from other interfaces.
	IsParamsDeclarationSectionContext()
}

type ParamsDeclarationSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsDeclarationSectionContext() *ParamsDeclarationSectionContext {
	var p = new(ParamsDeclarationSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramsDeclarationSection
	return p
}

func InitEmptyParamsDeclarationSectionContext(p *ParamsDeclarationSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramsDeclarationSection
}

func (*ParamsDeclarationSectionContext) IsParamsDeclarationSectionContext() {}

func NewParamsDeclarationSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsDeclarationSectionContext {
	var p = new(ParamsDeclarationSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_paramsDeclarationSection

	return p
}

func (s *ParamsDeclarationSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsDeclarationSectionContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ParamsDeclarationSectionContext) ParamSpecifier() IParamSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamSpecifierContext)
}

func (s *ParamsDeclarationSectionContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *ParamsDeclarationSectionContext) Identifier() IIdentifierContext {
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

func (s *ParamsDeclarationSectionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, 0)
}

func (s *ParamsDeclarationSectionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParamsDeclarationSectionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFUNCTION, 0)
}

func (s *ParamsDeclarationSectionContext) ParamsDeclaration() IParamsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclarationContext)
}

func (s *ParamsDeclarationSectionContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROCEDURE, 0)
}

func (s *ParamsDeclarationSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsDeclarationSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsDeclarationSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterParamsDeclarationSection(s)
	}
}

func (s *ParamsDeclarationSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitParamsDeclarationSection(s)
	}
}

func (s *ParamsDeclarationSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitParamsDeclarationSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ParamsDeclarationSection() (localctx IParamsDeclarationSectionContext) {
	localctx = NewParamsDeclarationSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, scopepascalParserRULE_paramsDeclarationSection)
	var _la int

	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case scopepascalParserVAR, scopepascalParserCONST, scopepascalParserOUT, scopepascalParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(588)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0 {
			{
				p.SetState(587)
				p.ParamSpecifier()
			}

		}
		{
			p.SetState(590)
			p.IdentifierList()
		}
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserCOLON {
			{
				p.SetState(591)
				p.Match(scopepascalParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(592)
				p.Identifier()
			}

		}
		p.SetState(597)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserEQUAL {
			{
				p.SetState(595)
				p.Match(scopepascalParserEQUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(596)
				p.Expression()
			}

		}

	case scopepascalParserFUNCTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(599)
			p.Match(scopepascalParserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)
			p.ParamsDeclaration()
		}
		{
			p.SetState(601)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.Identifier()
		}

	case scopepascalParserPROCEDURE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(604)
			p.Match(scopepascalParserPROCEDURE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.ParamsDeclaration()
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

// IParamSpecifierContext is an interface to support dynamic dispatch.
type IParamSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	CONST() antlr.TerminalNode
	OUT() antlr.TerminalNode

	// IsParamSpecifierContext differentiates from other interfaces.
	IsParamSpecifierContext()
}

type ParamSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamSpecifierContext() *ParamSpecifierContext {
	var p = new(ParamSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramSpecifier
	return p
}

func InitEmptyParamSpecifierContext(p *ParamSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_paramSpecifier
}

func (*ParamSpecifierContext) IsParamSpecifierContext() {}

func NewParamSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamSpecifierContext {
	var p = new(ParamSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_paramSpecifier

	return p
}

func (s *ParamSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamSpecifierContext) VAR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserVAR, 0)
}

func (s *ParamSpecifierContext) CONST() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCONST, 0)
}

func (s *ParamSpecifierContext) OUT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOUT, 0)
}

func (s *ParamSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterParamSpecifier(s)
	}
}

func (s *ParamSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitParamSpecifier(s)
	}
}

func (s *ParamSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitParamSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ParamSpecifier() (localctx IParamSpecifierContext) {
	localctx = NewParamSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, scopepascalParserRULE_paramSpecifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
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

// IVarSectionContext is an interface to support dynamic dispatch.
type IVarSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsVarSectionContext differentiates from other interfaces.
	IsVarSectionContext()
}

type VarSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSectionContext() *VarSectionContext {
	var p = new(VarSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_varSection
	return p
}

func InitEmptyVarSectionContext(p *VarSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_varSection
}

func (*VarSectionContext) IsVarSectionContext() {}

func NewVarSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSectionContext {
	var p = new(VarSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_varSection

	return p
}

func (s *VarSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSectionContext) VAR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserVAR, 0)
}

func (s *VarSectionContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *VarSectionContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
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

	return t.(IVarDeclarationContext)
}

func (s *VarSectionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *VarSectionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *VarSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterVarSection(s)
	}
}

func (s *VarSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitVarSection(s)
	}
}

func (s *VarSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitVarSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) VarSection() (localctx IVarSectionContext) {
	localctx = NewVarSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, scopepascalParserRULE_varSection)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(610)
		p.Match(scopepascalParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(611)
				p.VarDeclaration()
			}
			{
				p.SetState(612)
				p.Match(scopepascalParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(616)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
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

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	COLON() antlr.TerminalNode
	Identifier() IIdentifierContext
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *VarDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *VarDeclarationContext) Identifier() IIdentifierContext {
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

func (s *VarDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, 0)
}

func (s *VarDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, scopepascalParserRULE_varDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		p.IdentifierList()
	}
	p.SetState(621)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserCOLON {
		{
			p.SetState(619)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Identifier()
		}

	}
	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserEQUAL {
		{
			p.SetState(623)
			p.Match(scopepascalParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Expression()
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

// IInlinedVarDeclarationContext is an interface to support dynamic dispatch.
type IInlinedVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	IdentifierList() IIdentifierListContext
	SEMI() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Identifier() IIdentifierContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsInlinedVarDeclarationContext differentiates from other interfaces.
	IsInlinedVarDeclarationContext()
}

type InlinedVarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlinedVarDeclarationContext() *InlinedVarDeclarationContext {
	var p = new(InlinedVarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_inlinedVarDeclaration
	return p
}

func InitEmptyInlinedVarDeclarationContext(p *InlinedVarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_inlinedVarDeclaration
}

func (*InlinedVarDeclarationContext) IsInlinedVarDeclarationContext() {}

func NewInlinedVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlinedVarDeclarationContext {
	var p = new(InlinedVarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_inlinedVarDeclaration

	return p
}

func (s *InlinedVarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InlinedVarDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserVAR, 0)
}

func (s *InlinedVarDeclarationContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *InlinedVarDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, 0)
}

func (s *InlinedVarDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *InlinedVarDeclarationContext) Identifier() IIdentifierContext {
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

func (s *InlinedVarDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserASSIGN, 0)
}

func (s *InlinedVarDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InlinedVarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlinedVarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlinedVarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterInlinedVarDeclaration(s)
	}
}

func (s *InlinedVarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitInlinedVarDeclaration(s)
	}
}

func (s *InlinedVarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitInlinedVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) InlinedVarDeclaration() (localctx IInlinedVarDeclarationContext) {
	localctx = NewInlinedVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, scopepascalParserRULE_inlinedVarDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(627)
		p.Match(scopepascalParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(628)
		p.IdentifierList()
	}
	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserCOLON {
		{
			p.SetState(629)
			p.Match(scopepascalParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Identifier()
		}

	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserASSIGN {
		{
			p.SetState(633)
			p.Match(scopepascalParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)
			p.Expression()
		}

	}
	{
		p.SetState(637)
		p.Match(scopepascalParserSEMI)
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

// IRelationaloperatorContext is an interface to support dynamic dispatch.
type IRelationaloperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	GT() antlr.TerminalNode
	IN() antlr.TerminalNode

	// IsRelationaloperatorContext differentiates from other interfaces.
	IsRelationaloperatorContext()
}

type RelationaloperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationaloperatorContext() *RelationaloperatorContext {
	var p = new(RelationaloperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_relationaloperator
	return p
}

func InitEmptyRelationaloperatorContext(p *RelationaloperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_relationaloperator
}

func (*RelationaloperatorContext) IsRelationaloperatorContext() {}

func NewRelationaloperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationaloperatorContext {
	var p = new(RelationaloperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_relationaloperator

	return p
}

func (s *RelationaloperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationaloperatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserEQUAL, 0)
}

func (s *RelationaloperatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserNOT_EQUAL, 0)
}

func (s *RelationaloperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLT, 0)
}

func (s *RelationaloperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLE, 0)
}

func (s *RelationaloperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserGE, 0)
}

func (s *RelationaloperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserGT, 0)
}

func (s *RelationaloperatorContext) IN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserIN, 0)
}

func (s *RelationaloperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationaloperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationaloperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterRelationaloperator(s)
	}
}

func (s *RelationaloperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitRelationaloperator(s)
	}
}

func (s *RelationaloperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitRelationaloperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Relationaloperator() (localctx IRelationaloperatorContext) {
	localctx = NewRelationaloperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, scopepascalParserRULE_relationaloperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(639)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1040384) != 0) {
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

// IAdditiveoperatorContext is an interface to support dynamic dispatch.
type IAdditiveoperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsAdditiveoperatorContext differentiates from other interfaces.
	IsAdditiveoperatorContext()
}

type AdditiveoperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveoperatorContext() *AdditiveoperatorContext {
	var p = new(AdditiveoperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_additiveoperator
	return p
}

func InitEmptyAdditiveoperatorContext(p *AdditiveoperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_additiveoperator
}

func (*AdditiveoperatorContext) IsAdditiveoperatorContext() {}

func NewAdditiveoperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveoperatorContext {
	var p = new(AdditiveoperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_additiveoperator

	return p
}

func (s *AdditiveoperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveoperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPLUS, 0)
}

func (s *AdditiveoperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(scopepascalParserMINUS, 0)
}

func (s *AdditiveoperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserOR, 0)
}

func (s *AdditiveoperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveoperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveoperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterAdditiveoperator(s)
	}
}

func (s *AdditiveoperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitAdditiveoperator(s)
	}
}

func (s *AdditiveoperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitAdditiveoperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Additiveoperator() (localctx IAdditiveoperatorContext) {
	localctx = NewAdditiveoperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, scopepascalParserRULE_additiveoperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&436207616) != 0) {
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

// IMultiplicativeoperatorContext is an interface to support dynamic dispatch.
type IMultiplicativeoperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	AND() antlr.TerminalNode
	SHR() antlr.TerminalNode
	SHL() antlr.TerminalNode

	// IsMultiplicativeoperatorContext differentiates from other interfaces.
	IsMultiplicativeoperatorContext()
}

type MultiplicativeoperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeoperatorContext() *MultiplicativeoperatorContext {
	var p = new(MultiplicativeoperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_multiplicativeoperator
	return p
}

func InitEmptyMultiplicativeoperatorContext(p *MultiplicativeoperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_multiplicativeoperator
}

func (*MultiplicativeoperatorContext) IsMultiplicativeoperatorContext() {}

func NewMultiplicativeoperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeoperatorContext {
	var p = new(MultiplicativeoperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_multiplicativeoperator

	return p
}

func (s *MultiplicativeoperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeoperatorContext) STAR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTAR, 0)
}

func (s *MultiplicativeoperatorContext) SLASH() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSLASH, 0)
}

func (s *MultiplicativeoperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(scopepascalParserDIV, 0)
}

func (s *MultiplicativeoperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(scopepascalParserMOD, 0)
}

func (s *MultiplicativeoperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(scopepascalParserAND, 0)
}

func (s *MultiplicativeoperatorContext) SHR() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSHR, 0)
}

func (s *MultiplicativeoperatorContext) SHL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserSHL, 0)
}

func (s *MultiplicativeoperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeoperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeoperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterMultiplicativeoperator(s)
	}
}

func (s *MultiplicativeoperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitMultiplicativeoperator(s)
	}
}

func (s *MultiplicativeoperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitMultiplicativeoperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Multiplicativeoperator() (localctx IMultiplicativeoperatorContext) {
	localctx = NewMultiplicativeoperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, scopepascalParserRULE_multiplicativeoperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(643)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1643118592) != 0) {
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

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Relationaloperator() IRelationaloperatorContext
	Additiveoperator() IAdditiveoperatorContext
	Multiplicativeoperator() IMultiplicativeoperatorContext

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) Relationaloperator() IRelationaloperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationaloperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationaloperatorContext)
}

func (s *OperatorContext) Additiveoperator() IAdditiveoperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveoperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveoperatorContext)
}

func (s *OperatorContext) Multiplicativeoperator() IMultiplicativeoperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeoperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeoperatorContext)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, scopepascalParserRULE_operator)
	p.SetState(648)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case scopepascalParserEQUAL, scopepascalParserNOT_EQUAL, scopepascalParserLT, scopepascalParserLE, scopepascalParserGT, scopepascalParserGE, scopepascalParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(645)
			p.Relationaloperator()
		}

	case scopepascalParserOR, scopepascalParserPLUS, scopepascalParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(646)
			p.Additiveoperator()
		}

	case scopepascalParserSTAR, scopepascalParserSLASH, scopepascalParserDIV, scopepascalParserMOD, scopepascalParserAND, scopepascalParserSHR, scopepascalParserSHL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(647)
			p.Multiplicativeoperator()
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

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING_LITERAL() []antlr.TerminalNode
	STRING_LITERAL(i int) antlr.TerminalNode
	AllSTRING_CROSSHATCH_LITERAL() []antlr.TerminalNode
	STRING_CROSSHATCH_LITERAL(i int) antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSTRING_LITERAL)
}

func (s *StringContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTRING_LITERAL, i)
}

func (s *StringContext) AllSTRING_CROSSHATCH_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSTRING_CROSSHATCH_LITERAL)
}

func (s *StringContext) STRING_CROSSHATCH_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSTRING_CROSSHATCH_LITERAL, i)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, scopepascalParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(651)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == scopepascalParserSTRING_LITERAL || _la == scopepascalParserSTRING_CROSSHATCH_LITERAL {
		{
			p.SetState(650)
			_la = p.GetTokenStream().LA(1)

			if !(_la == scopepascalParserSTRING_LITERAL || _la == scopepascalParserSTRING_CROSSHATCH_LITERAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(653)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUM_INT() antlr.TerminalNode
	NUM_REAL() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserNUM_INT, 0)
}

func (s *NumberContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(scopepascalParserNUM_REAL, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, scopepascalParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(655)
		_la = p.GetTokenStream().LA(1)

		if !(_la == scopepascalParserNUM_INT || _la == scopepascalParserNUM_REAL) {
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllOperator() []IOperatorContext
	Operator(i int) IOperatorContext
	FunctionExpression() IFunctionExpressionContext
	ErrorExpression() IErrorExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *ExpressionContext) AllOperator() []IOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperatorContext); ok {
			len++
		}
	}

	tst := make([]IOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperatorContext); ok {
			tst[i] = t.(IOperatorContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Operator(i int) IOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
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

	return t.(IOperatorContext)
}

func (s *ExpressionContext) FunctionExpression() IFunctionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *ExpressionContext) ErrorExpression() IErrorExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, scopepascalParserRULE_expression)
	var _la int

	p.SetState(668)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 77, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(657)
			p.Term()
		}
		p.SetState(663)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2080366592) != 0 {
			{
				p.SetState(658)
				p.Operator()
			}
			{
				p.SetState(659)
				p.Term()
			}

			p.SetState(665)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(666)
			p.FunctionExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(667)
			p.ErrorExpression()
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

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	AT() antlr.TerminalNode
	Number() INumberContext
	String_() IStringContext
	FunctionDesignator() IFunctionDesignatorContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	ErrorExpression() IErrorExpressionContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) IDENT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserIDENT, 0)
}

func (s *TermContext) AT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserAT, 0)
}

func (s *TermContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *TermContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *TermContext) FunctionDesignator() IFunctionDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDesignatorContext)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *TermContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *TermContext) ErrorExpression() IErrorExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorExpressionContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, scopepascalParserRULE_term)
	var _la int

	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 79, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(671)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == scopepascalParserAT {
			{
				p.SetState(670)
				p.Match(scopepascalParserAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(673)
			p.Match(scopepascalParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(674)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(675)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(676)
			p.FunctionDesignator()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(677)
			p.Match(scopepascalParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(678)
			p.Expression()
		}
		{
			p.SetState(679)
			p.Match(scopepascalParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(681)
			p.ErrorExpression()
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

// IFunctionExpressionContext is an interface to support dynamic dispatch.
type IFunctionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	COLON() antlr.TerminalNode
	BlockStatement() IBlockStatementContext
	ParamsDeclaration() IParamsDeclarationContext
	VarSection() IVarSectionContext

	// IsFunctionExpressionContext differentiates from other interfaces.
	IsFunctionExpressionContext()
}

type FunctionExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionExpressionContext() *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionExpression
	return p
}

func InitEmptyFunctionExpressionContext(p *FunctionExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionExpression
}

func (*FunctionExpressionContext) IsFunctionExpressionContext() {}

func NewFunctionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_functionExpression

	return p
}

func (s *FunctionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionExpressionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(scopepascalParserFUNCTION, 0)
}

func (s *FunctionExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOLON, 0)
}

func (s *FunctionExpressionContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunctionExpressionContext) ParamsDeclaration() IParamsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclarationContext)
}

func (s *FunctionExpressionContext) VarSection() IVarSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarSectionContext)
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FunctionExpression() (localctx IFunctionExpressionContext) {
	localctx = NewFunctionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, scopepascalParserRULE_functionExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(684)
		p.Match(scopepascalParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(686)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(685)
			p.ParamsDeclaration()
		}

	}
	{
		p.SetState(688)
		p.Match(scopepascalParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserVAR {
		{
			p.SetState(689)
			p.VarSection()
		}

	}
	{
		p.SetState(692)
		p.BlockStatement()
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

// IProcedureExpressionContext is an interface to support dynamic dispatch.
type IProcedureExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDURE() antlr.TerminalNode
	BlockStatement() IBlockStatementContext
	ParamsDeclaration() IParamsDeclarationContext
	VarSection() IVarSectionContext

	// IsProcedureExpressionContext differentiates from other interfaces.
	IsProcedureExpressionContext()
}

type ProcedureExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureExpressionContext() *ProcedureExpressionContext {
	var p = new(ProcedureExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureExpression
	return p
}

func InitEmptyProcedureExpressionContext(p *ProcedureExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_procedureExpression
}

func (*ProcedureExpressionContext) IsProcedureExpressionContext() {}

func NewProcedureExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureExpressionContext {
	var p = new(ProcedureExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_procedureExpression

	return p
}

func (s *ProcedureExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureExpressionContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(scopepascalParserPROCEDURE, 0)
}

func (s *ProcedureExpressionContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ProcedureExpressionContext) ParamsDeclaration() IParamsDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclarationContext)
}

func (s *ProcedureExpressionContext) VarSection() IVarSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarSectionContext)
}

func (s *ProcedureExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterProcedureExpression(s)
	}
}

func (s *ProcedureExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitProcedureExpression(s)
	}
}

func (s *ProcedureExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitProcedureExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ProcedureExpression() (localctx IProcedureExpressionContext) {
	localctx = NewProcedureExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, scopepascalParserRULE_procedureExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(694)
		p.Match(scopepascalParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(696)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(695)
			p.ParamsDeclaration()
		}

	}
	p.SetState(699)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserVAR {
		{
			p.SetState(698)
			p.VarSection()
		}

	}
	{
		p.SetState(701)
		p.BlockStatement()
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

// IFunctionDesignatorContext is an interface to support dynamic dispatch.
type IFunctionDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExpressionList() IExpressionListContext

	// IsFunctionDesignatorContext differentiates from other interfaces.
	IsFunctionDesignatorContext()
}

type FunctionDesignatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDesignatorContext() *FunctionDesignatorContext {
	var p = new(FunctionDesignatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionDesignator
	return p
}

func InitEmptyFunctionDesignatorContext(p *FunctionDesignatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_functionDesignator
}

func (*FunctionDesignatorContext) IsFunctionDesignatorContext() {}

func NewFunctionDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDesignatorContext {
	var p = new(FunctionDesignatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_functionDesignator

	return p
}

func (s *FunctionDesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDesignatorContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDesignatorContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *FunctionDesignatorContext) LT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLT, 0)
}

func (s *FunctionDesignatorContext) GT() antlr.TerminalNode {
	return s.GetToken(scopepascalParserGT, 0)
}

func (s *FunctionDesignatorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserLPAREN, 0)
}

func (s *FunctionDesignatorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, 0)
}

func (s *FunctionDesignatorContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionDesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterFunctionDesignator(s)
	}
}

func (s *FunctionDesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitFunctionDesignator(s)
	}
}

func (s *FunctionDesignatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitFunctionDesignator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) FunctionDesignator() (localctx IFunctionDesignatorContext) {
	localctx = NewFunctionDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, scopepascalParserRULE_functionDesignator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(703)
		p.Identifier()
	}
	p.SetState(708)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(704)
			p.Match(scopepascalParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(705)
			p.Identifier()
		}
		{
			p.SetState(706)
			p.Match(scopepascalParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(715)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == scopepascalParserLPAREN {
		{
			p.SetState(710)
			p.Match(scopepascalParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(712)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-17179872264) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&33554431) != 0) {
			{
				p.SetState(711)
				p.ExpressionList()
			}

		}
		{
			p.SetState(714)
			p.Match(scopepascalParserRPAREN)
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

// IErrorExpressionContext is an interface to support dynamic dispatch.
type IErrorExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllEND() []antlr.TerminalNode
	END(i int) antlr.TerminalNode
	AllBEGIN() []antlr.TerminalNode
	BEGIN(i int) antlr.TerminalNode

	// IsErrorExpressionContext differentiates from other interfaces.
	IsErrorExpressionContext()
}

type ErrorExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorExpressionContext() *ErrorExpressionContext {
	var p = new(ErrorExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_errorExpression
	return p
}

func InitEmptyErrorExpressionContext(p *ErrorExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = scopepascalParserRULE_errorExpression
}

func (*ErrorExpressionContext) IsErrorExpressionContext() {}

func NewErrorExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorExpressionContext {
	var p = new(ErrorExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = scopepascalParserRULE_errorExpression

	return p
}

func (s *ErrorExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorExpressionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserSEMI)
}

func (s *ErrorExpressionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserSEMI, i)
}

func (s *ErrorExpressionContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserRPAREN)
}

func (s *ErrorExpressionContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserRPAREN, i)
}

func (s *ErrorExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserCOMMA)
}

func (s *ErrorExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserCOMMA, i)
}

func (s *ErrorExpressionContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserEND)
}

func (s *ErrorExpressionContext) END(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserEND, i)
}

func (s *ErrorExpressionContext) AllBEGIN() []antlr.TerminalNode {
	return s.GetTokens(scopepascalParserBEGIN)
}

func (s *ErrorExpressionContext) BEGIN(i int) antlr.TerminalNode {
	return s.GetToken(scopepascalParserBEGIN, i)
}

func (s *ErrorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.EnterErrorExpression(s)
	}
}

func (s *ErrorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scopepascalListener); ok {
		listenerT.ExitErrorExpression(s)
	}
}

func (s *ErrorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case scopepascalVisitor:
		return t.VisitErrorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *scopepascalParser) ErrorExpression() (localctx IErrorExpressionContext) {
	localctx = NewErrorExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, scopepascalParserRULE_errorExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(718)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(717)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17179872262) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(720)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 87, p.GetParserRuleContext())
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
