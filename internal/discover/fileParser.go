package discover

import (
	"palsp/internal/log"
	"palsp/internal/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// Custom error listener that sends errors to zerolog
type CustomErrorListener struct {
	debugInfo string
}

func (cel *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.AntlrError.Error().
		Str("di", cel.debugInfo).
		Int("line", line).
		Int("column", column).
		Str("msg", msg).
		Msg("Syntax error")
}

func (cel *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Optional: handle ambiguity reports
}

func (cel *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Optional: handle full context attempts
}

func (cel *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	// Optional: handle context sensitivity
}

// Custom trace listener that logs enter/exit events (based on original ANTLR TraceListener)
type CustomTraceListener struct {
	debugInfo string
}

func (ctl *CustomTraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.AntlrTrace.Debug().
		Str("di", ctl.debugInfo).
		Str("enter", "rule").
		Str("rule", strconv.Itoa(ctx.GetRuleContext().GetRuleIndex())).
		Msg("Enter rule")
}

func (ctl *CustomTraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	log.AntlrTrace.Debug().
		Str("di", ctl.debugInfo).
		Str("exit", "rule").
		Str("rule", strconv.Itoa(ctx.GetRuleContext().GetRuleIndex())).
		Msg("Exit rule")
}

func (ctl *CustomTraceListener) VisitTerminal(node antlr.TerminalNode) {
	log.AntlrTrace.Debug().
		Str("di", ctl.debugInfo).
		Str("token", node.GetText()).
		Msg("Visit terminal")
}

func (ctl *CustomTraceListener) VisitErrorNode(node antlr.ErrorNode) {
	log.AntlrTrace.Debug().
		Str("di", ctl.debugInfo).
		Str("token", node.GetText()).
		Msg("Visit error node")
}

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
	Tree   antlr.Tree
	Stream antlr.TokenStream
	// todo: make it non sparse so just map line to line (or otherwise regions should map line count too)
	Regions []Region
}

func (pd *ParsedData) FindOriginalLine(line int) (string, int) {
	if line < 0 || len(pd.Regions) == 0 {
		return "unknown", line
	}

	cumulativeDelta := 0
	var targetRegion *Region

	// Find the region where mainLine is last lesser or equal to line parameter
	// Sum deltas from regions we skip
	for i := 0; i < len(pd.Regions); i++ {
		region := &pd.Regions[i]

		if region.mainLine <= line {
			// This region could be our target, but keep looking for a better match
			targetRegion = region
		} else {
			// This region's mainLine is greater than our line, so we stop here
			break
		}

		// Add delta from this region (movement relative to previous region)
		cumulativeDelta += region.delta
	}

	// If no suitable region found, use the first region
	if targetRegion == nil {
		targetRegion = &pd.Regions[0]
		cumulativeDelta = pd.Regions[0].delta
	}

	// Calculate the original line by subtracting cumulative delta
	originalLine := line - cumulativeDelta

	// Ensure we don't return negative line numbers
	if originalLine < 1 {
		originalLine = 1
	}

	filename := "unknown"
	if targetRegion.fileCtx != nil {
		filename = targetRegion.fileCtx.Filename
	}

	return filename, originalLine
}

func (pd *ParsedData) FindParsedLine(originalLine int) (int, bool) {
	if originalLine < 1 || len(pd.Regions) == 0 {
		return -1, false
	}

	cumulativeDelta := 0

	// Iterate through regions to find where this original line maps to
	for i := 0; i < len(pd.Regions); i++ {
		region := &pd.Regions[i]

		// Add delta from this region (movement relative to previous region)
		cumulativeDelta += region.delta

		// Check if this region contains our original line
		// The original line + cumulative delta should equal or be less than the mainLine
		parsedLine := originalLine + cumulativeDelta

		// If this parsed line is within or at the region boundary, we found our match
		if parsedLine <= region.mainLine {
			return parsedLine, true
		}

	}

	// If not found in any region, check if it falls after the last region
	if len(pd.Regions) > 0 {
		parsedLine := originalLine + cumulativeDelta
		return parsedLine, true
	}

	return -1, false
}

func ParseCST(content string, debugInfo string) *ParsedData {
	// Get file path from debugInfo for preprocessing
	filePath := debugInfo
	if filePath == "" {
		filePath = "unknown"
	}

	// Create input stream from preprocessed content
	//input := antlr.NewInputStream(preprocessed.Content)
	input := newPascalCharStream(content, filePath, SymDB().GetSearchPaths(), SymDB().GetDefines())
	lexer := parser.NewpascalLexer(input)

	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&CustomErrorListener{debugInfo: debugInfo})

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser - use lowercase constructor
	pascalParser := parser.NewpascalParser(stream)

	// Remove default error listeners and add custom one
	pascalParser.RemoveErrorListeners()
	pascalParser.AddErrorListener(&CustomErrorListener{debugInfo: debugInfo})

	// Optionally add trace listener for debugging
	if log.AntlrTrace.Debug().Enabled() {
		pascalParser.AddParseListener(&CustomTraceListener{debugInfo: debugInfo})
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
