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
	usb              *UnitScopeBuilder
	sbs              *scopeBuilderStack
	inImplementation bool
	currentScope     stack[string]
}

// NewMemorySymbolCollector creates a new collector for in-memory model
func NewMemorySymbolCollector() *MemorySymbolCollector {
	csb := commonScopeBuilder{cmsc: commonScope{}}
	sbs := newScopeBuilderStack()
	sbs.push(&csb)
	usb := &UnitScopeBuilder{
		commonScopeBuilder: &csb,
		interfaceUses:      *newStack[string](),
		implementationUses: *newStack[string](),
	}
	currentScope := stack[string]{}
	return &MemorySymbolCollector{
		usb:          usb,
		sbs:          sbs,
		currentScope: currentScope,
	}
}

func (mc *MemorySymbolCollector) BeginScope(name string, position Position) {
	newsb := commonScopeBuilder{cmsc: commonScope{name: name, position: position}}
	mc.sbs.push(&newsb)
	if mc.currentScope.length() == 0 {
		mc.currentScope.push(strings.ToLower(name))
	} else {
		mc.currentScope.push(mc.currentScope.peek() + "." + strings.ToLower(name))
	}
}

func (mc *MemorySymbolCollector) EndScope(name string) {
	sb := mc.sbs.pop()
	sb.setName(name)
	sb.parentSWM(mc.sbs.peek().symbolStackLast())
	mc.sbs.peek().addScope(sb.finish())
	mc.currentScope.pop()
}

func (mc *MemorySymbolCollector) EnterImplementation(position Position) {
	mc.inImplementation = true
	mc.usb.setImplementationPos(position)
}

func (mc *MemorySymbolCollector) AddUseUnit(unit string) {
	if mc.inImplementation {
		mc.usb.addImplementationUses(unit)
	} else {
		mc.usb.addInterfaceUses(unit)
	}
}

func (mc *MemorySymbolCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	currentBuilder := mc.sbs.current()
	currentBuilder.addSymbol(name, definition, int(kind), position, mc.currentScope.peek())
}

func (mc *MemorySymbolCollector) GetScope() TopScope {
	return mc.usb.finish()
}

func (dc *MemorySymbolCollector) AccessSpecifier(as AccessSpec) {

}

// scopeBuilderStack manages a stack of scope builders
type scopeBuilderStack struct {
	items []*commonScopeBuilder
}

// newScopeBuilderStack creates a new scope builder stack
func newScopeBuilderStack() *scopeBuilderStack {
	return &scopeBuilderStack{
		items: make([]*commonScopeBuilder, 0),
	}
}

// push adds a scope builder to the top of the stack
func (s *scopeBuilderStack) push(item *commonScopeBuilder) {
	s.items = append(s.items, item)
}

// pop removes and returns the top scope builder from the stack
func (s *scopeBuilderStack) pop() *commonScopeBuilder {
	if len(s.items) == 0 {
		return nil
	}

	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

// peek returns the top scope builder without removing it
func (s *scopeBuilderStack) peek() *commonScopeBuilder {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// current returns the current scope builder (alias for peek)
func (s *scopeBuilderStack) current() *commonScopeBuilder {
	return s.peek()
}

// length returns the number of items in the stack
func (s *scopeBuilderStack) length() int {
	return len(s.items)
}

// UnifiedListener is a parse tree listener that can collect symbols for different purposes
type UnifiedListener struct {
	parser.BasepascalListener
	collector SymbolCollector
	// unitName         string
	// inUsesList       bool
	// inImplementation bool
	// currentSymbol    string
	// currentType      string
	// scope            string
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
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) EnterProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.collector.BeginScope(buildIdentifier(ctx.ProcedureHeader().Identifier()), newPosition(ctx.ProcedureHeader().Identifier()))
}

func (s *UnifiedListener) ExitProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.collector.EndScope(buildIdentifier(ctx.ProcedureHeader().Identifier()))
}

func (s *UnifiedListener) EnterFunctionHeader(ctx *parser.FunctionHeaderContext) {
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if ctx.ResultType() != nil {
		s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), newPosition(ctx.Identifier()))
	}
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.collector.BeginScope(buildIdentifier(ctx.FunctionHeader().Identifier()), newPosition(ctx.FunctionHeader().Identifier()))
}

func (s *UnifiedListener) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	if ctx.FunctionHeader().ResultType() != nil {
		s.collector.AddSymbol(buildIdentifier(ctx.FunctionHeader().Identifier()), FunctionResult, buildTypeIdentifier(ctx.FunctionHeader().ResultType().TypeIdentifier()), newPosition(ctx.FunctionHeader().Identifier()))
	}
	s.collector.EndScope(buildIdentifier(ctx.FunctionHeader().Identifier()))
}

func (s *UnifiedListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), newPosition(ctx.Identifier()))
}

func (s *UnifiedListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), TypeSymbol, buildTypeDef(ctx), newPosition(ctx.Identifier()))
}
