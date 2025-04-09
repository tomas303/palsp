package discover

import (
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rs/zerolog/log" // Import your configured logger
)

// Custom error listener that sends errors to zerolog
type ZerologErrorListener struct {
	antlr.DefaultErrorListener // Embed default implementation
}

// SyntaxError is called by ANTLR when a syntax error occurs
func (l *ZerologErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	log.Error().
		Int("line", line).
		Int("column", column).
		Str("message", msg).
		Msg("ANTLR syntax error")
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

	// Conditionally add trace listener based on configuration
	// This could be controlled by a debug flag
	// if isTraceEnabled() { // Implement this function based on your config
	// p.SetTrace(new(antlr.TraceListener))
	// }

	// Return the AST by invoking the Source rule
	return p.Source()
}
