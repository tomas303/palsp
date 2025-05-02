package discover

import (
	"errors"
	"palsp/internal/parser" // Ensure this import is correct
	"strings"

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

var ErrListenerBreak = errors.New("listener breaks gracefully")

// SymbolCollector defines an interface for collecting symbols
type SymbolCollector interface {
	AddSymbol(name string, kind SymbolKind, definition string, position Position)
	BeginScope(name string, position Position)
	EndScope(name string)
	EnterImplementation(position Position)
	AddUseUnit(unit string)
	AccessSpecifier(as AccessSpec)
}

// DBSymbolCollector implements SymbolCollector for database storage
type DBSymbolCollector struct {
	unitID           int
	db               *symDB
	currentScope     stack[string]
	curentAccessSpec AccessSpec
}

// NewDBSymbolCollector creates a new collector for database storage
func NewDBSymbolCollector(unitID int, db *symDB) *DBSymbolCollector {
	return &DBSymbolCollector{
		unitID:       unitID,
		db:           db,
		currentScope: *newStack[string](),
	}
}

func (dc *DBSymbolCollector) BeginScope(name string, position Position) {
	if dc.currentScope.length() == 0 {
		dc.currentScope.push(strings.ToLower(name))
	} else {
		dc.currentScope.push(dc.currentScope.peek() + "." + strings.ToLower(name))
	}
}

func (dc *DBSymbolCollector) EndScope(name string) {
	dc.currentScope.pop()
}

func (dc *DBSymbolCollector) EnterImplementation(position Position) {
	panic(ErrListenerBreak)
}

func (dc *DBSymbolCollector) AddUseUnit(unit string) {
}

func (dc *DBSymbolCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	dc.db.InsertSymbol(dc.unitID, name, dc.currentScope.peek(), int(kind), definition)
}

func (dc *DBSymbolCollector) AccessSpecifier(as AccessSpec) {
	dc.curentAccessSpec = as
}

// MemorySymbolCollector implements SymbolCollector for in-memory model
type MemorySymbolCollector struct {
	unitScope        *UnitScope
	scopeStack       *scopeStack
	inImplementation bool
	currentScope     stack[string]
}

// NewMemorySymbolCollector creates a new collector for in-memory model
func NewMemorySymbolCollector() *MemorySymbolCollector {
	rootScope := commonScope{name: "root", position: Position{}}
	scopeStack := newScopeStack()
	scopeStack.push(&rootScope)
	unitScope := &UnitScope{
		interfaceUses:      *newStack[string](),
		implementationUses: *newStack[string](),
	}
	currentScope := stack[string]{}
	return &MemorySymbolCollector{
		unitScope:    unitScope,
		scopeStack:   scopeStack,
		currentScope: currentScope,
	}
}

func (mc *MemorySymbolCollector) BeginScope(name string, position Position) {
	// log.Logger.Debug().Str("begin scope", name).Int("line", position.Line).Int("chr", position.Character).Send()
	newScope := commonScope{name: strings.ToLower(name), position: position}
	mc.scopeStack.push(&newScope)
	if mc.currentScope.length() == 0 {
		mc.currentScope.push(strings.ToLower(name))
	} else {
		mc.currentScope.push(mc.currentScope.peek() + "." + strings.ToLower(name))
	}
}

func (mc *MemorySymbolCollector) EndScope(name string) {
	// log.Logger.Debug().Str("end scope", name).Send()
	scope := mc.scopeStack.pop()
	parentscope := mc.scopeStack.peek()
	scope.parentSWM = parentscope.symbolStack.length() - 1
	parentscope.scopeStack.push(scope)
	mc.currentScope.pop()
}

func (mc *MemorySymbolCollector) EnterImplementation(position Position) {
	mc.inImplementation = true
	mc.unitScope.implementationPos = position
}

func (mc *MemorySymbolCollector) AddUseUnit(unit string) {
	if mc.inImplementation {
		mc.unitScope.implementationUses.push(unit)
	} else {
		mc.unitScope.interfaceUses.push(unit)
	}
}

func (mc *MemorySymbolCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	name = strings.ToLower(name)
	smb := Symbol{Name: name, Definition: definition, Kind: int(kind), Position: position, Scope: mc.currentScope.peek()}
	scope := mc.scopeStack.current()
	scope.symbolStack.push(smb)
	if scope.scopeStack.length() > 0 && scope.scopeStack.peek().getName() == name {
		scope.scopeStack.peek().setParentSWM(scope.symbolStack.length() - 1)
	}
}

func (mc *MemorySymbolCollector) GetScope() TopScope {
	mc.unitScope.Scope = mc.scopeStack.peek().scopeStack.peek()
	return mc.unitScope
}

func (dc *MemorySymbolCollector) AccessSpecifier(as AccessSpec) {

}

// scopeStack manages a stack of scope builders
type scopeStack struct {
	items []*commonScope
}

// newScopeStack creates a new scope builder stack
func newScopeStack() *scopeStack {
	return &scopeStack{
		items: make([]*commonScope, 0),
	}
}

// push adds a scope builder to the top of the stack
func (s *scopeStack) push(item *commonScope) {
	s.items = append(s.items, item)
}

// pop removes and returns the top scope builder from the stack
func (s *scopeStack) pop() *commonScope {
	if len(s.items) == 0 {
		return nil
	}

	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

// peek returns the top scope builder without removing it
func (s *scopeStack) peek() *commonScope {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// current returns the current scope builder (alias for peek)
func (s *scopeStack) current() *commonScope {
	return s.peek()
}

// length returns the number of items in the stack
func (s *scopeStack) length() int {
	return len(s.items)
}

// UnifiedListener is a parse tree listener that can collect symbols for different purposes
type UnifiedListener struct {
	parser.BasepascalListener
	collector     SymbolCollector
	inDeclaration bool
}

// NewUnifiedListener creates a new unified listener with the specified collector
func NewUnifiedListener(collector SymbolCollector) *UnifiedListener {
	return &UnifiedListener{
		collector: collector,
	}
}

func (s *UnifiedListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	s.collector.EnterImplementation(newPosition(ctx))
}

// func (s *UnifiedListener) GetScope() TopScope {
// 	return s.usb.finish()
// }

// func (s *UnifiedListener) ab() *commonScopeBuilder {
// 	return s.sbs.peek()
// }

func (s *UnifiedListener) ExitUsesUnits(ctx *parser.UsesUnitsContext) {
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		s.collector.AddUseUnit(buildIdentifier(identifier))
		s.collector.AddSymbol(buildIdentifier(identifier), UnitReference, identifier.GetText(), newPosition(identifier))
	}
}

func (s *UnifiedListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
	for _, identifier := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
		s.collector.AddSymbol(buildIdentifier(identifier), VariableSymbol, typedef, newPosition(identifier))
	}
}

func (s *UnifiedListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
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
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), ConstantSymbol, fieldtype, newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitFormalParameterList(ctx *parser.FormalParameterListContext) {
	for _, parSecCtx := range ctx.AllFormalParameterSection() {
		parType := ""
		if parSecCtx.ParameterGroup() == nil {
			continue
		}
		if parSecCtx.ParameterGroup().TypeIdentifier() != nil {
			parType = parSecCtx.ParameterGroup().TypeIdentifier().GetText()
		}
		// if ctx.DefaultValue() != nil {
		// 	result += " = " + ctx.DefaultValue().GetText()
		// }
		for _, id := range parSecCtx.ParameterGroup().IdentifierList().AllIdentifier() {
			s.collector.AddSymbol(buildIdentifier(id), ParameterSymbol, parType, newPosition(id))
		}
		//
	}
}

func (s *UnifiedListener) ExitClassDeclarationPart(ctx *parser.ClassDeclarationPartContext) {
	if ctx.TypedIdentifierList() != nil {
		typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
		for _, id := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
			s.collector.AddSymbol(buildIdentifier(id), ClassVariable, typedef, newPosition(id))
		}
	}
}

func (s *UnifiedListener) EnterUnit(ctx *parser.UnitContext) {
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitUnit(ctx *parser.UnitContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
}

func (s *UnifiedListener) EnterProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), newPosition(ctx.Identifier()))
		return
	}
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) EnterProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.inDeclaration = true
	s.collector.BeginScope(buildIdentifier(ctx.ProcedureHeader().Identifier()), newPosition(ctx.ProcedureHeader().Identifier()))
}

func (s *UnifiedListener) ExitProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.collector.EndScope(buildIdentifier(ctx.ProcedureHeader().Identifier()))
	s.inDeclaration = false
}

func (s *UnifiedListener) EnterFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		if ctx.ResultType() != nil {
			s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), newPosition(ctx.Identifier()))
		}
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), newPosition(ctx.Identifier()))
		return
	}
	if ctx.ResultType() != nil {
		s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), newPosition(ctx.Identifier()))
	}
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.inDeclaration = true
	s.collector.BeginScope(buildIdentifier(ctx.FunctionHeader().Identifier()), newPosition(ctx.FunctionHeader().Identifier()))
}

func (s *UnifiedListener) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.collector.EndScope(buildIdentifier(ctx.FunctionHeader().Identifier()))
	s.inDeclaration = false
}

func (s *UnifiedListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), TypeSymbol, buildTypeDef(ctx), newPosition(ctx.Identifier()))
}
