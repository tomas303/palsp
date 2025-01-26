package discover

import (
	"fmt"
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

func (l *CustomPascalListener) EnterProgram(ctx *parser.ProgramContext) {
	fmt.Println("Entering Program:", ctx.GetText())
}

func (l *CustomPascalListener) ExitProgram(ctx *parser.ProgramContext) {
	fmt.Println("Exiting Program:", ctx.GetText())
}

// Add methods for all rules of the class
func (l *CustomPascalListener) EnterBlock(ctx *parser.BlockContext) {
	fmt.Println("Entering Block:", ctx.GetText())
}

func (l *CustomPascalListener) ExitBlock(ctx *parser.BlockContext) {
	fmt.Println("Exiting Block:", ctx.GetText())
}

func (l *CustomPascalListener) EnterStatement(ctx *parser.StatementContext) {
	fmt.Println("Entering Statement:", ctx.GetText())
}

func (l *CustomPascalListener) ExitStatement(ctx *parser.StatementContext) {
	fmt.Println("Exiting Statement:", ctx.GetText())
}

func (l *CustomPascalListener) EnterExpression(ctx *parser.ExpressionContext) {
	fmt.Println("Entering Expression:", ctx.GetText())
}

func (l *CustomPascalListener) ExitExpression(ctx *parser.ExpressionContext) {
	fmt.Println("Exiting Expression:", ctx.GetText())
}

// Add methods related to class
func (l *CustomPascalListener) EnterClassType(ctx *parser.ClassTypeContext) {
	fmt.Println("Entering ClassType:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassType(ctx *parser.ClassTypeContext) {
	fmt.Println("Exiting ClassType:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassImplementsInterfaces(ctx *parser.ClassImplementsInterfacesContext) {
	fmt.Println("Entering ClassImplementsInterfaces:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassImplementsInterfaces(ctx *parser.ClassImplementsInterfacesContext) {
	fmt.Println("Exiting ClassImplementsInterfaces:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	fmt.Println("Entering ClassDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	fmt.Println("Exiting ClassDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassDeclarationRow(ctx *parser.ClassDeclarationRowContext) {
	fmt.Println("Entering ClassDeclarationRow:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassDeclarationRow(ctx *parser.ClassDeclarationRowContext) {
	fmt.Println("Exiting ClassDeclarationRow:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassPrivateDeclaration(ctx *parser.ClassPrivateDeclarationContext) {
	fmt.Println("Entering ClassPrivateDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassPrivateDeclaration(ctx *parser.ClassPrivateDeclarationContext) {
	fmt.Println("Exiting ClassPrivateDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassStrictPrivateDeclaration(ctx *parser.ClassStrictPrivateDeclarationContext) {
	fmt.Println("Entering ClassStrictPrivateDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassStrictPrivateDeclaration(ctx *parser.ClassStrictPrivateDeclarationContext) {
	fmt.Println("Exiting ClassStrictPrivateDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassProtectedDeclaration(ctx *parser.ClassProtectedDeclarationContext) {
	fmt.Println("Entering ClassProtectedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassProtectedDeclaration(ctx *parser.ClassProtectedDeclarationContext) {
	fmt.Println("Exiting ClassProtectedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassStrictProtectedDeclaration(ctx *parser.ClassStrictProtectedDeclarationContext) {
	fmt.Println("Entering ClassStrictProtectedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassStrictProtectedDeclaration(ctx *parser.ClassStrictProtectedDeclarationContext) {
	fmt.Println("Exiting ClassStrictProtectedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassPublicDeclaration(ctx *parser.ClassPublicDeclarationContext) {
	fmt.Println("Entering ClassPublicDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassPublicDeclaration(ctx *parser.ClassPublicDeclarationContext) {
	fmt.Println("Exiting ClassPublicDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassPublishedDeclaration(ctx *parser.ClassPublishedDeclarationContext) {
	fmt.Println("Entering ClassPublishedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassPublishedDeclaration(ctx *parser.ClassPublishedDeclarationContext) {
	fmt.Println("Exiting ClassPublishedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) EnterClassImplicitPublishedDeclaration(ctx *parser.ClassImplicitPublishedDeclarationContext) {
	fmt.Println("Entering ClassImplicitPublishedDeclaration:", ctx.GetText())
}

func (l *CustomPascalListener) ExitClassImplicitPublishedDeclaration(ctx *parser.ClassImplicitPublishedDeclarationContext) {
	fmt.Println("Exiting ClassImplicitPublishedDeclaration:", ctx.GetText())
}

// ...add more methods for other rules as needed...
