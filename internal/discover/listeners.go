package discover

import (
	"errors"
	"fmt"
	"log"
	"palsp/internal/parser" // Ensure this import is correct

	// added
	"github.com/antlr4-go/antlr/v4"
)

// Interface for any object that has a GetStart() method returning something with GetLine() and GetColumn()
type positionable interface {
	GetStart() antlr.Token
}

// Helper function to create a Position from any context that has GetStart()
func newPosition(ctx positionable) Position {
	if ctx == nil || ctx.GetStart() == nil {
		return Position{Line: 0, Character: 0}
	}
	return Position{
		Line:      ctx.GetStart().GetLine(),
		Character: ctx.GetStart().GetColumn(),
	}
}

type AccessSpec int

const (
	AccUnknown AccessSpec = iota
	AccPrivate
	AccStrictPrivate
	AccProtected
	AccStrictProtected
	AccPublic
	AccPublished
)

// public and all symbols ... only difference that public stops on implementation, probably in different table
// and probably less rules. So private shoudl wrap it and build on top of it
// this will be used to scan units form interface section of opended files. That can be recursive to fill up class inheritance(later if deemed necessary)
// this will be problaby heavy load but with more units already parsed it will be faster

type publicSymbolsListener struct {
	parser.BasepascalListener

	unitName              string
	unit_id               int
	scopeStack            stack[string]
	identExitActionStack  stack[func(identifier string)]
	accessSpecifiersStack stack[AccessSpec]
}

type scopeListener struct {
	parser.BasepascalListener
	usb              UnitScopeBuilder
	sbs              stack[*commonScopeBuilder]
	inImplementation bool
}

var ErrListenerBreak = errors.New("listener breaks gracefully")

func NewScopeListener(unit string) *scopeListener {
	csb := commonScopeBuilder{cmsc: commonScope{name: unit}}
	sbs := *newStack[*commonScopeBuilder]()
	sbs.push(&csb)
	usb := UnitScopeBuilder{
		commonScopeBuilder: &csb,
		interfaceUses:      *newStack[string](),
		implementationUses: *newStack[string](),
	}
	return &scopeListener{
		usb: usb,
		sbs: sbs,
	}
}

func (s *publicSymbolsListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	panic(ErrListenerBreak)
}

func (s *publicSymbolsListener) ExitIdentifier(ctx *parser.IdentifierContext) {
	action := s.identExitActionStack.pop()
	if action != nil {
		identifier := buildIdentifier(ctx)
		action(identifier)
	}
}

func (s *publicSymbolsListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.identExitActionStack.push(func(identifier string) {
		fmt.Printf("Entering type: %s\n", identifier)
		s.scopeStack.push(identifier)
	})
}

func (s *publicSymbolsListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	if ctx.ForwardClassType() != nil {
		s.scopeStack.pop()
	} else {
		identifier := s.scopeStack.pop()
		if s.canPublish() {
			fmt.Printf("Exiting type: %s\n", identifier)
			s.insertSymbol(identifier, int(TypeSymbol), buildTypeDef(ctx))
		}
	}
}

func (s *publicSymbolsListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if !s.canPublish() {
		return
	}
	definition := ""
	name := ""
	if ctx.CLASS() != nil {
		definition = "class "
	}
	if ctx.PROCEDURE() != nil {
		definition += "procedure "
	} else if ctx.CONSTRUCTOR() != nil {
		definition += "constructor "
	} else if ctx.DESTRUCTOR() != nil {
		definition += "destructor "
	}
	if ctx.MethodIdentifier() != nil {
		name = buildIdentifier(ctx.MethodIdentifier().Identifier())
	} else if ctx.Identifier() != nil {
		name = buildIdentifier(ctx.Identifier())
	}
	definition += name
	definition += "(" + buildParameterList(ctx.FormalParameterList()) + ")"
	definition += buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers())
	if definition != "" {
		s.insertSymbol(name, int(ProcedureSymbol), definition)
	}
}

func (s *publicSymbolsListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if !s.canPublish() {
		return
	}
	definition := ""
	name := ""
	if ctx.CLASS() != nil {
		definition = "class "
	}
	definition += "function "
	if ctx.MethodIdentifier() != nil {
		name = buildIdentifier(ctx.MethodIdentifier().Identifier())
	} else if ctx.Identifier() != nil {
		name = buildIdentifier(ctx.Identifier())
	}
	definition += name
	definition += "(" + buildParameterList(ctx.FormalParameterList()) + ")"
	definition += ": " + buildTypeIdentifier(ctx.ResultType().TypeIdentifier())
	definition += buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers())
	if definition != "" {
		s.insertSymbol(name, int(FunctionSymbol), definition)
	}
}

func (s *publicSymbolsListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
	if !s.canPublish() {
		return
	}
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
	if !s.canPublish() {
		return
	}
	defaultvalue := ""
	if ctx.SimpleExpression() != nil {
		defaultvalue += "=" + buildSimpleExpression(ctx.SimpleExpression())
	}
	fields, fieldtype := buildTypedIdentifierList(ctx.TypedIdentifierList())
	fieldtype += defaultvalue
	for _, field := range fields {
		var definition string
		if fieldtype == "" {
			definition = field
		} else {
			definition = fmt.Sprintf("%s: %s", field, fieldtype)
		}
		s.insertSymbol(field, int(VariableSymbol), definition)
	}
}

// func (s *publicSymbolsListener) EnterRecordType(ctx *parser.RecordTypeContext) {
// 	s.scopeStack.push("record to be done")
// }

// func (s *publicSymbolsListener) ExitRecordType(ctx *parser.RecordTypeContext) {
// 	s.scopeStack.pop()
// }clo

func (s *publicSymbolsListener) ExitAccessSpecifier(ctx *parser.AccessSpecifierContext) {
	if ctx.PRIVATE() != nil {
		s.accessSpecifiersStack.push(AccPrivate)
	} else if ctx.STRICTPRIVATE() != nil {
		s.accessSpecifiersStack.push(AccStrictPrivate)
	} else if ctx.PROTECTED() != nil {
		s.accessSpecifiersStack.push(AccProtected)
	} else if ctx.STRICTPROTECTED() != nil {
		s.accessSpecifiersStack.push(AccStrictProtected)
	} else if ctx.PUBLIC() != nil {
		s.accessSpecifiersStack.push(AccPublic)
	} else if ctx.PUBLISHED() != nil {
		s.accessSpecifiersStack.push(AccPublished)
	} else {
		s.accessSpecifiersStack.push(AccPublished)
	}
}

func (s *publicSymbolsListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.accessSpecifiersStack.pop()
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
	err := SymDB().InsertSymbol(s.unit_id, symbol, s.scope(), kind, definition)
	if err != nil {
		log.Printf("Non-fatal error encountered: %v", err)
	}
}

func (s *publicSymbolsListener) canPublish() bool {
	return s.accessSpecifiersStack.isEmpty() || s.accessSpecifiersStack.peek() == AccPublic || s.accessSpecifiersStack.peek() == AccPublished
}

func (s *scopeListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	s.inImplementation = true
	s.usb.setImplementationPos(newPosition(ctx))
}

// func (s *scopeListener) ExitImplementationSection(ctx *parser.ImplementationSectionContext) {
// }

func (s *scopeListener) GetScope() TopScope {
	return s.usb.finish()
}

func (s *scopeListener) ab() *commonScopeBuilder {
	return s.sbs.peek()
}

func (s *scopeListener) ExitUsesUnits(ctx *parser.UsesUnitsContext) {
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		if s.inImplementation {
			s.usb.addImplementationUses(buildIdentifier(identifier))
		} else {
			s.usb.addInterfaceUses(buildIdentifier(identifier))
		}
	}
}

func (s *scopeListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
	for _, identifier := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
		s.addSymbol(buildIdentifier(identifier), typedef, int(VariableSymbol), newPosition(identifier))
	}
}

func (s *scopeListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
	fieldtype := ""
	if ctx.TypeIdentifier() != nil {
		fieldtype = buildTypeIdentifier(ctx.TypeIdentifier())
	}
	if fieldtype == "" {
		if ctx.Constant().String_() != nil {
			fieldtype = "string"
		} else if ctx.Constant().UnsignedNumber() != nil {
			fieldtype = "integer"
		} else if ctx.Constant().Sign() != nil {
			fieldtype = "integer"
		}
	}
	s.addSymbol(buildIdentifier(ctx.Identifier()), fieldtype, int(ConstantSymbol), newPosition(ctx.Identifier()))
}

func (s *scopeListener) ExitFormalParameterList(ctx *parser.FormalParameterListContext) {
	for _, parSecCtx := range ctx.AllFormalParameterSection() {
		parType := ""
		if parSecCtx.ParameterGroup().TypeIdentifier() != nil {
			parType = parSecCtx.ParameterGroup().TypeIdentifier().GetText()
		}
		// if ctx.DefaultValue() != nil {
		// 	result += " = " + ctx.DefaultValue().GetText()
		// }
		for _, id := range parSecCtx.ParameterGroup().IdentifierList().AllIdentifier() {
			s.addSymbol(buildIdentifier(id), parType, int(ParameterSymbol), newPosition(id))
		}
		//
	}
}

func (s *scopeListener) ExitClassDeclarationPart(ctx *parser.ClassDeclarationPartContext) {
	if ctx.TypedIdentifierList() != nil {
		typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
		for _, id := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
			s.addSymbol(buildIdentifier(id), typedef, int(ClassVariable), newPosition(id))
		}
	}
}

func (s *scopeListener) EnterProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.ab().getName() != "procedure" {
		newsb := commonScopeBuilder{cmsc: commonScope{name: buildIdentifier(ctx.Identifier()), position: newPosition(ctx.Identifier())}}
		s.sbs.push(&newsb)
	}
}

func (s *scopeListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.ab().getName() != "procedure" {
		sb := s.sbs.pop()
		sb.setName(buildIdentifier(ctx.Identifier()))
		sb.parentSWM(s.ab().symbolStackLast())
		s.ab().addScope(sb.finish())
	}
}

func (s *scopeListener) EnterProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	newsb := commonScopeBuilder{cmsc: commonScope{name: buildIdentifier(ctx.ProcedureHeader().Identifier()), position: newPosition(ctx.ProcedureHeader().Identifier())}}
	s.sbs.push(&newsb)
}

func (s *scopeListener) ExitProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	sb := s.sbs.pop()
	sb.setName(buildIdentifier(ctx.ProcedureHeader().Identifier()))
	sb.parentSWM(s.ab().symbolStackLast())
	s.ab().addScope(sb.finish())
}

func (s *scopeListener) EnterFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.ab().getName() != "function" {
		newsb := commonScopeBuilder{cmsc: commonScope{name: buildIdentifier(ctx.Identifier()), position: newPosition(ctx.Identifier())}}
		s.sbs.push(&newsb)
	}
}

func (s *scopeListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.ab().getName() != "function" {
		funcName := buildIdentifier(ctx.Identifier())
		if ctx.ResultType() != nil {
			s.addSymbol(funcName, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), int(FunctionResult), newPosition(ctx.Identifier()))
		}
		sb := s.sbs.pop()
		sb.setName(funcName)
		sb.parentSWM(s.ab().symbolStackLast())
		s.ab().addScope(sb.finish())
	}
}

func (s *scopeListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	newsb := commonScopeBuilder{cmsc: commonScope{name: buildIdentifier(ctx.FunctionHeader().Identifier()), position: newPosition(ctx.FunctionHeader().Identifier())}}
	s.sbs.push(&newsb)
}

func (s *scopeListener) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	funcName := buildIdentifier(ctx.FunctionHeader().Identifier())
	if ctx.FunctionHeader().ResultType() != nil {
		s.addSymbol(funcName, buildTypeIdentifier(ctx.FunctionHeader().ResultType().TypeIdentifier()), int(FunctionResult), newPosition(ctx.FunctionHeader().Identifier()))
	}
	sb := s.sbs.pop()
	sb.setName(funcName)
	sb.parentSWM(s.ab().symbolStackLast())
	s.ab().addScope(sb.finish())
}

func (s *scopeListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	newsb := commonScopeBuilder{cmsc: commonScope{name: buildIdentifier(ctx.Identifier()), position: newPosition(ctx.Identifier())}}
	s.sbs.push(&newsb)
}

func (s *scopeListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	sb := s.sbs.pop()
	sb.setName(buildIdentifier(ctx.Identifier()))
	s.addSymbol(buildIdentifier(ctx.Identifier()), buildTypeDef(ctx), int(TypeSymbol), newPosition(ctx.Identifier()))
	sb.parentSWM(s.ab().symbolStackLast())
	s.ab().addScope(sb.finish())
}

func (s *scopeListener) addSymbol(name string, definition string, kind int, position Position) {
	sscope := ""
	for _, x := range s.sbs.all() {
		sscope += x.getName() + "."
	}
	s.ab().addSymbol(name, definition, kind, position, sscope)
}
