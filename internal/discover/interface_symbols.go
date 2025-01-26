package discover

import (
	"fmt"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

func HandleDidOpen(path string, content string) {
	fmt.Println("File opened:", path)
	// fmt.Println("File content:", content)

	// Create an input stream from the file content
	input := antlr.NewInputStream(content)

	// Create a lexer and parser
	lexer := parser.NewpascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpascalParser(stream)

	// Create a custom listener
	listener := &CustomPascalListener{}

	// Walk the AST with the custom listener
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
}

type CustomPascalListener struct {
	parser.BasepascalListener
}

func (l *CustomPascalListener) EnterProgram(ctx *parser.ProgramContext) {
	fmt.Println("Entering Program:", ctx.GetText())
}
