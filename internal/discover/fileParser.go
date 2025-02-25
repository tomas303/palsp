package discover

import (
	"io"
	parser1 "palsp/internal/parser"

	parser2 "palsp/internal/scopeparser/parser"

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
	lexer := parser1.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser1.NewpascalParser(stream)
	if options.Trace {
		p.SetTrace(new(antlr.TraceListener))
	}
	if options.HandleError {
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	}
	p.AddParseListener(listener)
	p.Source()
}

func scopeparseFromContent(content string, listener antlr.ParseTreeListener, options parseOptions) {
	input := antlr.NewInputStream(content)
	lexer := parser2.NewscopepascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser2.NewscopepascalParser(stream)
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
	lexer := parser1.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser1.NewpascalParser(stream)
	if options.Trace {
		p.SetTrace(new(antlr.TraceListener))
	}
	if options.HandleError {
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	}
	p.AddParseListener(listener)
	p.Source()
}

func parseAST(content string) antlr.Tree {
	input := antlr.NewInputStream(content)
	lexer := parser1.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser1.NewpascalParser(stream)
	// Return the AST by invoking the Source rule
	return p.Source()
}
