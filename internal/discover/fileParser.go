package discover

import (
	"io"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

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

func fullDebugOptions() parseOptions {
	return parseOptions{
		Trace:       true,
		HandleError: true,
	}
}

func parseFromContent(content string, listener antlr.ParseTreeListener, options parseOptions) {
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

func parseFromReader(reader io.Reader, listener antlr.ParseTreeListener, options parseOptions) {
	input := antlr.NewIoStream(reader)
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
