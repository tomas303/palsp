package discover

import (
	"fmt"
	"palsp/internal/log"
	"palsp/internal/parser"
	"strings"

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
	debugInfo string
}

func NewZerologTraceListener(parser antlr.Parser, debugInfo string) *ZerologTraceListener {
	tl := new(ZerologTraceListener)
	tl.parser = parser
	tl.debugInfo = debugInfo
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
		Str("di", t.debugInfo).
		Str("errorNode", node.GetText()).
		Str("interval", fmt.Sprintf("%v", node.GetSourceInterval())).
		Str("rule", t.parser.GetRuleNames()[t.parser.GetParserRuleContext().GetRuleIndex()]).
		Send()
}

func (t *ZerologTraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	currentToken := t.parser.GetTokenStream().LT(1)
	t.getEvent(ctx).
		Str("di", t.debugInfo).
		Str("enter", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).
		Str("token", currentToken.GetText()).
		Int("tokenType", currentToken.GetTokenType()).
		Send()
}

func (t *ZerologTraceListener) VisitTerminal(node antlr.TerminalNode) {
	t.getEvent(t.parser.GetParserRuleContext()).
		Str("di", t.debugInfo).
		Str("consume", fmt.Sprint(node.GetSymbol())).
		Str("rule", t.parser.GetRuleNames()[t.parser.GetParserRuleContext().GetRuleIndex()]).
		Send()
}

func (t *ZerologTraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	currentToken := t.parser.GetTokenStream().LT(1)
	t.getEvent(ctx).
		Str("di", t.debugInfo).
		Str("exit", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).
		Str("token", currentToken.GetText()).
		Int("tokenType", currentToken.GetTokenType()).
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

// FindOriginalLine finds the original line number corresponding to the parsed line number
// and its filecontext.
// It just keep tracks of regions beeing pushed and popped and calculates cumulative delta

func (pd *ParsedData) FindOriginalLine(line int) (int, bool, *FileContext) {

	type sourceStackItem struct {
		start           *Region
		cumulativeDelta int
	}

	if line < 0 || len(pd.Regions) == 0 {
		return -1, false, nil
	}

	sourceStack := []sourceStackItem{
		{
			start:           &pd.Regions[0],
			cumulativeDelta: 0,
		},
	}
	actualSource := &sourceStack[0]
	for i := 1; i < len(pd.Regions); i++ {
		region := &pd.Regions[i]

		if line < region.mainLine {
			return line - actualSource.start.mainLine - actualSource.cumulativeDelta, true, actualSource.start.fileCtx
		}

		if region.rsmove == rsPush {
			sourceStack = append(sourceStack, sourceStackItem{
				start:           region,
				cumulativeDelta: 0,
			})
			actualSource = &sourceStack[len(sourceStack)-1]
		}

		if region.rsmove == rsPop {
			cumulated := actualSource.cumulativeDelta
			span := region.mainLine - actualSource.start.mainLine
			sourceStack = sourceStack[:len(sourceStack)-1]
			actualSource = &sourceStack[len(sourceStack)-1]
			actualSource.cumulativeDelta += cumulated + span
		}
	}

	return line - actualSource.start.mainLine - actualSource.cumulativeDelta, true, actualSource.start.fileCtx
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

	parsedLine := originalLine + cumulativeDelta
	return parsedLine, true
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
	pascalParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)

	// Remove default error listeners and add custom one
	pascalParser.RemoveErrorListeners()
	pascalParser.AddErrorListener(NewZerologErrorListener(debugInfo))
	pascalParser.AddErrorListener(NewAmbiguityErrorListener(debugInfo))

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

type AmbiguityErrorListener struct {
	antlr.DefaultErrorListener
	debugInfo string
}

func NewAmbiguityErrorListener(debugInfo string) *AmbiguityErrorListener {
	return &AmbiguityErrorListener{debugInfo: debugInfo}
}

func (l *AmbiguityErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {

	tokenStream := recognizer.GetTokenStream()

	// Extract ambiguous content
	ambiguousText := l.getTokenText(tokenStream, startIndex, stopIndex)
	contextBefore := l.getContextBefore(tokenStream, startIndex, 3)
	contextAfter := l.getContextAfter(tokenStream, stopIndex, 3)

	// Get line/column info if available
	startToken := tokenStream.Get(startIndex)
	line := startToken.GetLine()
	column := startToken.GetColumn()

	// Get current rule name
	ruleName := "unknown"
	if ctx := recognizer.GetParserRuleContext(); ctx != nil && ctx.GetRuleIndex() >= 0 {
		ruleName = recognizer.GetRuleNames()[ctx.GetRuleIndex()]
	}

	log.AntlrTrace.Warn().
		Str("di", l.debugInfo).
		Str("pos", fmt.Sprintf("[%d:%d]", line, column)).
		Str("rule", ruleName).
		Str("ambigAlts", ambigAlts.String()).
		Str("exact", fmt.Sprintf("%t", exact)).
		Str("context", fmt.Sprintf("...%s [%s] %s...", contextBefore, ambiguousText, contextAfter)).
		Msg("AMBIGUITY DETECTED")
}

// Helper methods
func (l *AmbiguityErrorListener) getTokenText(stream antlr.TokenStream, start, stop int) string {
	if start < 0 || stop >= stream.Size() || start > stop {
		return ""
	}

	var tokens []string
	for i := start; i <= stop; i++ {
		tokens = append(tokens, stream.Get(i).GetText())
	}
	return strings.Join(tokens, " ")
}

func (l *AmbiguityErrorListener) getContextBefore(stream antlr.TokenStream, start, size int) string {
	if start <= size {
		return ""
	}
	return l.getTokenText(stream, start-size, start-1)
}

func (l *AmbiguityErrorListener) getContextAfter(stream antlr.TokenStream, stop, size int) string {
	if stop+size >= stream.Size() {
		return ""
	}
	return l.getTokenText(stream, stop+1, stop+size)
}

func (l *AmbiguityErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {

	log.AntlrTrace.Info().
		Str("di", l.debugInfo).
		Int("startIndex", startIndex).
		Int("stopIndex", stopIndex).
		Str("conflictingAlts", conflictingAlts.String()).
		Msg("ATTEMPTING FULL CONTEXT")
}

func (l *AmbiguityErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {

	log.AntlrTrace.Info().
		Str("di", l.debugInfo).
		Int("startIndex", startIndex).
		Int("stopIndex", stopIndex).
		Int("prediction", prediction).
		Msg("CONTEXT SENSITIVITY")
}
