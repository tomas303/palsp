package discover

import (
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

// func fullDebugOptions() parseOptions {
// 	return parseOptions{
// 		Trace:       true,
// 		HandleError: true,
// 	}
// }

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

func ParseCST(content string) antlr.Tree {
	input := antlr.NewInputStream(content)
	lexer := parser.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpascalParser(stream)
	p.SetTrace(new(antlr.TraceListener))
	// Return the AST by invoking the Source rule
	return p.Source()
}
