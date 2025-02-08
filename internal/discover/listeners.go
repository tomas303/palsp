package discover

import (
	"fmt"
	"log"
	"palsp/internal/parser" // Ensure this import is correct
	"strings"               // added

	"github.com/antlr4-go/antlr/v4"
)

// type PascalListener interface {
// 	EnterUnit(ctx *parser.UnitContext)
// 	EnterEveryRule(ctx antlr.ParserRuleContext)
// 	ExitEveryRule(ctx antlr.ParserRuleContext)
// }

type finishError struct {
	Message string
}

type listenerFactory func() antlr.ParseTreeListener

type listenerHandler func(listener antlr.ParseTreeListener, path string)

type unitNameListener struct {
	parser.BasepascalListener

	unitName string
	isUnit   bool
}

// public and all symbols ... only difference that public stops on implementation, probably in different table
// and probably less rules. So private shoudl wrap it and build on top of it
// this will be used to scan units form interface section of opended files. That can be recursive to fill up class inheritance(later if deemed necessary)
// this will be problaby heavy load but with more units already parsed it will be faster
type publicSymbolsListener struct {
	parser.BasepascalListener

	unitName   string
	unit_id    int
	scopeStack stack[string]
}

func (e *finishError) Error() string {
	return e.Message
}

func newFinishError(message string) *finishError {
	return &finishError{Message: message}
}

func (l *unitNameListener) ExitUnit(ctx *parser.UnitContext) {
	identifiers := ctx.AllIdentifier() // get all identifiers
	if len(identifiers) > 0 {          // if there is at least one identifier
		var parts []string
		for i := 0; i < len(identifiers); i++ {
			parts = append(parts, identifiers[i].GetText())
		}
		unitName := strings.Join(parts, ".") // join with dot delimiter
		fmt.Println("Unit identified:", unitName)
		l.unitName = unitName
		l.isUnit = true
		panic(newFinishError("Unit ID rule hit"))
	}
}

// GetUnitName returns the unit name identified by the listener
func (l *unitNameListener) UnitName() string {
	return l.unitName
}

// IsUnit returns whether the listener has identified a unit
func (l *unitNameListener) IsUnit() bool {
	return l.isUnit
}

func (s *publicSymbolsListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	panic(newFinishError("implementation hit, no more public symbols"))
}

func (s *publicSymbolsListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	name := ""
	if ctx.MethodIdentifier() != nil {
		name = ctx.MethodIdentifier().Identifier().GetText()
	} else if ctx.Identifier() != nil {
		name = ctx.Identifier().GetText()
	}
	if name != "" {
		s.insertSymbol(name, int(ProcedureSymbol), name)
	}
}

func (s *publicSymbolsListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	name := ""
	if ctx.MethodIdentifier() != nil {
		name = ctx.MethodIdentifier().Identifier().GetText()
	} else if ctx.Identifier() != nil {
		name = ctx.Identifier().GetText()
	}
	if name != "" {
		s.insertSymbol(name, int(FunctionSymbol), name)
	}
}

func (s *publicSymbolsListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
	name := safeGetText(ctx.Identifier())
	if name != "" {
		typename := safeGetText(ctx.TypeIdentifier())
		value := ctx.Constant().GetText()
		var definition string
		if typename == "" {
			definition = fmt.Sprintf("%s = %s", name, value)
		} else {
			definition = fmt.Sprintf("%s: %s = %s", name, typename, value)
		}
		s.insertSymbol(name, int(ConstantSymbol), definition)
	}
}

func (s *publicSymbolsListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	//identifierList COLON type_
	identifierList := ctx.IdentifierList()
	// probably some custom logic here to get type ... who knows what typename will be
	// typename := safeGetText(ctx.Type_())
	// SimpleType() ISimpleTypeContext
	// StructuredType() IStructuredTypeContext
	// PointerType() IPointerTypeContext
	typename := safeGetText(ctx.Type_())
	for _, identifier := range identifierList.AllIdentifier() {
		name := identifier.GetText()
		var definition string
		if typename == "" {
			definition = name
		} else {
			definition = fmt.Sprintf("%s: %s", name, typename)
		}
		s.insertSymbol(name, int(VariableSymbol), definition)
	}
}

func (s *publicSymbolsListener) EnterClassType(ctx *parser.ClassTypeContext) {
	// necessary solve public, protected, private ... in this case only protexted and public and must be distinguished, strict protected and protected kinda too
	name := safeGetText(ctx.Identifier())
	SymDB().insertSymbol(s.unit_id, name, s.scope(), int(ClassSymbol), "class( ... inherit, impelemntms)")
	s.scopeStack.push(name)
}

func (s *publicSymbolsListener) ExitClassType(ctx *parser.ClassTypeContext) {
	s.scopeStack.pop()
}

func (s *publicSymbolsListener) EnterRecordType(ctx *parser.RecordTypeContext) {
	s.scopeStack.push("record to be done")
}

func (s *publicSymbolsListener) ExitRecordType(ctx *parser.RecordTypeContext) {
	s.scopeStack.pop()
}

func safeGetText(ctx interface{ GetText() string }) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetText()
}

func (s *publicSymbolsListener) scope() string {
	return s.scopeStack.joinByDot()
}

func (s *publicSymbolsListener) insertSymbol(symbol string, kind int, definition string) {
	err := SymDB().insertSymbol(s.unit_id, symbol, s.scope(), kind, definition)
	if err != nil {
		log.Printf("Non-fatal error encountered: %v", err)
	}
}
