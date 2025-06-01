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

func ParseCST(content string, debugInfo string) (antlr.Tree, antlr.TokenStream) {
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

	return tree, stream
}
