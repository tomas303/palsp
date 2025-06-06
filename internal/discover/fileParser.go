package discover

import (
	"fmt"
	"palsp/internal/log"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rs/zerolog"
)

// Custom error listener that sends errors to zerolog
type ZerologErrorListener struct {
	antlr.DefaultErrorListener // Embed default implementation
	debugInfo                  string
}

func NewZerologErrorListener(debugInfo string) *ZerologErrorListener {
	l := new(ZerologErrorListener)
	l.debugInfo = debugInfo
	return l
}

// SyntaxError is called by ANTLR when a syntax error occurs
func (l *ZerologErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	log.AntlrError.Error().
		Int("line", line).
		Int("column", column).
		Str("msg", msg).
		Str("di", l.debugInfo).
		Send()
}

// trace listener that logs enter/exit events(based on original ANTLR TraceListener)
type ZerologTraceListener struct {
	parser    antlr.Parser
	degubInfo string
}

func NewZerologTraceListener(parser antlr.Parser, debugInfo string) *ZerologTraceListener {
	tl := new(ZerologTraceListener)
	tl.parser = parser
	tl.degubInfo = debugInfo
	return tl
}
func (t *ZerologTraceListener) getEvent(ctx antlr.ParserRuleContext) *zerolog.Event {
	ruleName := t.parser.GetRuleNames()[ctx.GetRuleIndex()]
	if len(ruleName) >= 5 && ruleName[0:5] == "error" {
		return log.AntlrTrace.Error()
	}
	return log.AntlrTrace.Debug()
}

func (t *ZerologTraceListener) VisitErrorNode(node antlr.ErrorNode) {
	log.AntlrTrace.Error().
		Str("di", t.degubInfo).
		Str("errorNode", node.GetText()).
		Str("interval", fmt.Sprintf("%v", node.GetSourceInterval())).
		Str("rule", t.parser.GetRuleNames()[t.parser.GetParserRuleContext().GetRuleIndex()]).
		Send()
}

func (t *ZerologTraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	t.getEvent(ctx).
		Str("di", t.degubInfo).
		Str("enter", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).
		Str("token", t.parser.GetTokenStream().LT(1).GetText()).
		Send()
}

func (t *ZerologTraceListener) VisitTerminal(node antlr.TerminalNode) {
	t.getEvent(t.parser.GetParserRuleContext()).
		Str("di", t.degubInfo).
		Str("consume", fmt.Sprint(node.GetSymbol())).
		Str("rule", t.parser.GetRuleNames()[t.parser.GetParserRuleContext().GetRuleIndex()]).
		Send()
}

func (t *ZerologTraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	t.getEvent(ctx).
		Str("di", t.degubInfo).
		Str("exit", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).
		Str("token", t.parser.GetTokenStream().LT(1).GetText()).
		Send()
}

type ResilientErrorStrategy struct {
	*antlr.DefaultErrorStrategy
	endTokenType int
}

// Override Recover to skip tokens until 'end' or EOF
func (es *ResilientErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	es.ReportError(recognizer, e)

	for {
		t := recognizer.GetTokenStream().LA(1)
		if t == antlr.TokenEOF || t == es.endTokenType {
			break
		}
		recognizer.Consume()
	}
}

// Optionally override RecoverInline for single-token errors
func (es *ResilientErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	es.ReportMatch(recognizer)
	return recognizer.GetCurrentToken()
}

func NewResilientErrorStrategy() *ResilientErrorStrategy {
	var endTokenType int
	for i, n := range parser.PascalLexerLexerStaticData.SymbolicNames {
		if n == "END" {
			endTokenType = i
			break
		}
	}

	return &ResilientErrorStrategy{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
		endTokenType:         endTokenType,
	}
}

// ParseFile parses the given content as a Pascal source file and returns the parse tree.
func ParseFile(content string) (antlr.Tree, error) {
	// Create input stream from content
	input := antlr.NewInputStream(content)

	// Create lexer
	lexer := parser.NewpascalLexer(input)

	// Create token stream
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser
	p := parser.NewpascalParser(tokens)

	// Parse starting from source rule
	tree := p.Source()

	return tree, nil
}

type ParsedData struct {
	Tree    antlr.Tree
	Stream  antlr.TokenStream
	Regions []Region
}

// FindOriginalLine finds the original line number corresponding to the parsed line number.
// Thats because parsed data contains includes too which move original lines.
// Each region contains source - so each source from external file just cumulate movement
// and when region of line is found, we can calculate original line by subtracting it.
func (pd *ParsedData) FindOriginalLine(line int) (int, bool) {
	if line < 0 || len(pd.Regions) == 0 {
		return -1, false
	}

	prevRegion := &pd.Regions[0]
	mainFile := prevRegion.fileCtx.Filename
	cumulativeDelta := 0

	for i := 1; i < len(pd.Regions); i++ {
		region := &pd.Regions[i]
		if prevRegion.fileCtx.Filename != mainFile {
			cumulativeDelta += region.mainLine - prevRegion.mainLine
		}
		if line <= region.mainLine {
			return line - cumulativeDelta, true
		}
		prevRegion = region
	}

	if len(pd.Regions) > 0 {
		return line - cumulativeDelta, true
	}

	return line, false
}

// FindParsedLine finds the parsed line number corresponding to the original line number.
// Thats because parsed data contains includes too which move original lines.
// Each region contains source - so each source from external file just move following lines
func (pd *ParsedData) FindParsedLine(originalLine int) (int, bool) {
	if originalLine < 1 || len(pd.Regions) == 0 {
		return -1, false
	}

	prevRegion := &pd.Regions[0]
	mainFile := prevRegion.fileCtx.Filename
	cumulativeDelta := 0

	for i := 1; i < len(pd.Regions); i++ {
		region := &pd.Regions[i]

		if prevRegion.fileCtx.Filename != mainFile {
			cumulativeDelta += region.mainLine - prevRegion.mainLine
		}
		parsedLine := originalLine + cumulativeDelta
		if parsedLine <= region.mainLine {
			return parsedLine, true
		}
		prevRegion = region
	}

	if len(pd.Regions) > 0 {
		parsedLine := originalLine + cumulativeDelta
		return parsedLine, true
	}

	return originalLine, false
}

func ParseCST(content string, debugInfo string, skipImplementation bool) *ParsedData {
	// Get file path from debugInfo for preprocessing
	filePath := debugInfo
	if filePath == "" {
		filePath = "unknown"
	}

	// Create input stream from preprocessed content
	//input := antlr.NewInputStream(preprocessed.Content)
	input := newPascalCharStream(content, filePath, SymDB().GetSearchPaths(), SymDB().GetDefines(), skipImplementation)
	lexer := parser.NewpascalLexer(input)

	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(NewZerologErrorListener(debugInfo))

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser - use lowercase constructor
	pascalParser := parser.NewpascalParser(stream)

	// Remove default error listeners and add custom one
	pascalParser.RemoveErrorListeners()
	pascalParser.AddErrorListener(NewZerologErrorListener(debugInfo))

	// Optionally add trace listener for debugging
	if log.AntlrTrace.Debug().Enabled() {
		pascalParser.AddParseListener(NewZerologTraceListener(pascalParser, debugInfo))
	}

	// Return the AST by invoking the Source rule
	tree := pascalParser.Source()

	ParsedData := &ParsedData{
		Tree:    tree,
		Stream:  stream,
		Regions: input.regions,
	}

	return ParsedData
}
