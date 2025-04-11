package discover

import (
	"fmt"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
	// "github.com/rs/zerolog/log" // Import your configured logger
	"palsp/internal/log"
)

// Custom error listener that sends errors to zerolog
type ZerologErrorListener struct {
	antlr.DefaultErrorListener // Embed default implementation
}

// SyntaxError is called by ANTLR when a syntax error occurs
func (l *ZerologErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	// log.Logger.Error().
	// 	Int("line", line).
	// 	Int("column", column).
	// 	Str("err", msg).
	// 	Msg("ANTLR syntax error")

	errorMsg := fmt.Sprintf("ANTLR syntax error at line %d, column %d: %s", line, column, msg)
	log.Logger.Error().Msg(errorMsg)
}

// trace listener that logs enter/exit events(based on original ANTLR TraceListener)
type ZerologTraceListener struct {
	parser antlr.Parser
}

func NewZerologTraceListener(parser antlr.Parser) *ZerologTraceListener {
	tl := new(ZerologTraceListener)
	tl.parser = parser
	return tl
}

func (t *ZerologTraceListener) VisitErrorNode(_ antlr.ErrorNode) {
}

func (t *ZerologTraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Logger.Debug().Str("enter   ", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).Str("LT(1)", t.parser.GetTokenStream().LT(1).GetText()).Msg("ANTLR enter rule")
}

func (t *ZerologTraceListener) VisitTerminal(node antlr.TerminalNode) {
	log.Logger.Debug().Str("consume ", fmt.Sprint(node.GetSymbol())).Str("rule", t.parser.GetRuleNames()[t.parser.GetParserRuleContext().GetRuleIndex()]).Msg("ANTLR consume token")
}

func (t *ZerologTraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	log.Logger.Debug().Str("exit    ", t.parser.GetRuleNames()[ctx.GetRuleIndex()]).Str("LT(1)=", t.parser.GetTokenStream().LT(1).GetText()).Msg("ANTLR exit rule")
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

type parseOptions struct {
	Trace       bool
	HandleError bool
}

func defaultOptions() parseOptions {
	return parseOptions{
		Trace:       false,
		HandleError: false,
	}
}

func parseFromContent(content string, listener antlr.ParseTreeListener, options parseOptions) {

	defer func() {
		if r := recover(); r != nil {
			if r == ErrListenerBreak {
				return
			}
			panic(r) // Re-panic for all other errors
		}
	}()

	input := antlr.NewInputStream(content)
	lexer := parser.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpascalParser(stream)
	if options.Trace {
		p.SetTrace(new(antlr.TraceListener))
	}
	if options.HandleError {
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	}
	p.AddParseListener(listener)
	p.Source()
}

// Modify your ParseCST function to use these listeners:
func ParseCST(content string) antlr.Tree {
	input := antlr.NewInputStream(content)
	lexer := parser.NewpascalLexer(input)

	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&ZerologErrorListener{})

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpascalParser(stream)

	// Remove default error listeners and add custom one
	p.RemoveErrorListeners()
	p.AddErrorListener(&ZerologErrorListener{})

	p.SetErrorHandler(NewResilientErrorStrategy())

	if log.IsDebugEnabled() {
		p.AddParseListener(NewZerologTraceListener(p))
	}

	// Return the AST by invoking the Source rule
	return p.Source()
}
