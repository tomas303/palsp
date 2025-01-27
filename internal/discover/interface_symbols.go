package discover

import (
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

func HandleDidOpen(path string, content string) {
	// fmt.Println("File opened:", path)
	// fmt.Println("File content:", content)

	// Create an input stream from the file content
	input := antlr.NewInputStream(content)

	// Create a lexer and parser
	lexer := parser.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpascalParser(stream)
	p.SetTrace(new(antlr.TraceListener))
	p.AddErrorListener(new(antlr.DiagnosticErrorListener))

	// Create a custom listener
	listener := &CustomPascalListener{}

	// Attach the listener to the parser
	p.AddParseListener(listener)

	// Walk the AST with the custom listener
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
}

type CustomPascalListener struct {
	parser.BasepascalListener
}
