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
	err := SymDB().insertSymbol(s.unit_id, symbol, s.scope(), kind, definition)
	if err != nil {
		log.Printf("Non-fatal error encountered: %v", err)
	}
}

func (s *publicSymbolsListener) canPublish() bool {
	return s.accessSpecifiersStack.isEmpty() || s.accessSpecifiersStack.peek() == AccPublic || s.accessSpecifiersStack.peek() == AccPublished
}

func buildUnderscoreTypeDef(ctx parser.IType_Context) string {
	if ctx.SimpleType() != nil {
		return buildSimpleTypeDef(ctx.SimpleType())
	}
	if ctx.StructuredType() != nil {
		return buildStructuredTypeDef(ctx.StructuredType())
	}
	if ctx.PointerType() != nil {
		return "^" + buildTypeIdentifier(ctx.PointerType().TypeIdentifier())
	}
	return ""
}

func buildSimpleTypeDef(ctx parser.ISimpleTypeContext) string {
	result := ""
	if ctx.ScalarType() != nil {
		result = "(" + buildIdentifiers(ctx.ScalarType().IdentifierList()) + ")"
	}
	if ctx.SubrangeType() != nil {
		if len(ctx.SubrangeType().AllSimpleExpression()) > 0 {
			result = result + buildSimpleExpression(ctx.SubrangeType().SimpleExpression(0))
		}
		result = result + ".."
		if len(ctx.SubrangeType().AllSimpleExpression()) > 1 {
			result = result + buildSimpleExpression(ctx.SubrangeType().SimpleExpression(1))
		}
	}
	if ctx.TypeIdentifier() != nil {
		result = buildTypeIdentifier(ctx.TypeIdentifier())
	}
	if ctx.Stringtype() != nil {
		result = "string"
		if ctx.Stringtype().Identifier() != nil {
			result = result + "[" + buildIdentifier(ctx.Stringtype().Identifier()) + "]"
		}
		if ctx.Stringtype().UnsignedNumber() != nil {
			result = result + "[" + ctx.Stringtype().UnsignedNumber().GetText() + "]"
		}
	}
	if result == "" {
		result = "undetected: " + ctx.GetText()
	}
	return result
}

func buildStructuredTypeDef(ctx parser.IStructuredTypeContext) string {
	result := ""
	if ctx.PACKED() != nil {
		result += "packed "
	}
	if ctx.UnpackedStructuredType() != nil {
		if ctx.UnpackedStructuredType().RecordType() != nil {
			result += buildRecordTypeDef(ctx.UnpackedStructuredType().RecordType())
		}
		if ctx.UnpackedStructuredType().SetType() != nil {
			result += buildSetTypeDef(ctx.UnpackedStructuredType().SetType())
		}
		if ctx.UnpackedStructuredType().FileType() != nil {
			result += buildFileType(ctx.UnpackedStructuredType().FileType())
		}
	}
	if ctx.ClassType() != nil {
		result += buildClassTypeDef(ctx.ClassType())
	}
	return result
}

func buildRecordTypeDef(ctx parser.IRecordTypeContext) string {
	result := "record\n"
	result += buildRecordParts(ctx.RecordParts())
	result += "end"
	return result
}

func buildRecordParts(ctx parser.IRecordPartsContext) string {
	result := ""
	if ctx != nil {
		if ctx.RecordFixedPart() != nil {
			result += buildRecordFixedPart(ctx.RecordFixedPart())
			if ctx.RecordVariantPart() != nil {
				result += buildRecordVariantPart(ctx.RecordVariantPart())
			}
		} else if ctx.RecordVariantPart() != nil {
			result += buildRecordVariantPart(ctx.RecordVariantPart())
		}
	}
	return result
}

func buildRecordFixedPart(ctx parser.IRecordFixedPartContext) string {
	result := ""
	for _, typeIdlist := range ctx.AllTypedIdentifierList() {
		names, typedef := buildTypedIdentifierList(typeIdlist)
		result += strings.Join(names, ",") + ": " + typedef + ";\n"
	}
	return result
}

func buildRecordVariantPart(ctx parser.IRecordVariantPartContext) string {
	result := "case " + ctx.Tag().GetText() + " of\n"
	for _, variant := range ctx.AllRecordVariant() {
		result += buildRecordVariant(variant)
	}
	return result
}

func buildRecordVariant(ctx parser.IRecordVariantContext) string {
	result := ""
	result += buildConstList(ctx.ConstList())
	if result != "" {
		result += ": "
	}
	parts := buildRecordParts(ctx.RecordParts())
	if parts != "" {
		result += "(\n" + parts + ");\n"
	}
	return result
}

func buildSetTypeDef(ctx parser.ISetTypeContext) string {
	result := "set of "
	if ctx.SimpleType() != nil {
		result += buildSimpleTypeDef(ctx.SimpleType())
	}
	return result
}

func buildFileType(ctx parser.IFileTypeContext) string {
	if ctx.Type_() != nil {
		return "file of " + buildUnderscoreTypeDef(ctx.Type_())
	}
	return "file"
}

func buildClassTypeDef(ctx parser.IClassTypeContext) string {
	result := "class"
	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		implements := []string{}
		if ctx.ClassImplementsInterfaces() != nil {
			for _, identifier := range ctx.ClassImplementsInterfaces().AllTypeIdentifier() {
				implements = append(implements, identifier.GetText())
			}
		}
		if ctx.Identifier() != nil {
			result += "(" + buildIdentifier(ctx.Identifier())
			if len(implements) > 0 {
				result += ", " + strings.Join(implements, ", ")
			}
			result += ")"
		}

	}
	if ctx.ABSTRACT() != nil {
		result += " abstract"
	}
	result += "\n"
	if ctx.ClassImplicitPublishedDeclaration() != nil {
		result += buildClassImplicitPublishedDeclaration(ctx.ClassImplicitPublishedDeclaration())
	}
	result += "end"
	return result
}

func buildClassImplicitPublishedDeclaration(ctx parser.IClassImplicitPublishedDeclarationContext) string {
	result := "published\n"
	for _, part := range ctx.AllClassDeclarationPart() {
		result += buildClassDeclarationPart(part)
	}
	return result
}

func buildClassDeclarationPart(ctx parser.IClassDeclarationPartContext) string {
	result := ""
	if ctx.TypedIdentifierList() != nil {
		list, typedef := buildTypedIdentifierList(ctx.TypedIdentifierList())
		result += strings.Join(list, ", ") + ": " + typedef + ";\n"
	}
	return result
}

func buildFunctionTypeDef(ctx parser.IFunctionTypeContext) string {
	result := ""
	params := buildParameterList(ctx.FormalParameterList())
	if ctx.ResultType() != nil {
		result = ctx.ResultType().TypeIdentifier().GetText()
	}
	if result != "" {
		result += fmt.Sprintf("function(%s): %s", params, result)
	} else {
		result += fmt.Sprintf("function(%s)", params)
	}
	result += buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers())
	return result
}
func buildProcedureTypeDef(ctx parser.IProcedureTypeContext) string {
	params := buildParameterList(ctx.FormalParameterList())
	return fmt.Sprintf("procedure(%s)", params) + buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers())
}

func buildTypeDef(ctx *parser.TypeDefinitionContext) string {
	// Attempt to use a more descriptive definition based on the context type.
	// Note: Replace the context type names with those in your actual generated parser.

	if ctx.Type_() != nil {
		return buildUnderscoreTypeDef(ctx.Type_())
	}

	if ctx.FunctionType() != nil {
		return buildFunctionTypeDef(ctx.FunctionType())
	}

	if ctx.ProcedureType() != nil {
		return buildProcedureTypeDef(ctx.ProcedureType())
	}

	return ""
}

func buildParameterGroup(ctx parser.IParameterGroupContext) string {
	result := ""
	if ctx.IdentifierList() != nil {
		identifiers := ctx.IdentifierList().AllIdentifier()
		for i, identifier := range identifiers {
			result += identifier.GetText()
			if i < len(identifiers)-1 {
				result += ", "
			}
		}
		if ctx.TypeIdentifier() != nil {
			result += ": " + ctx.TypeIdentifier().GetText()
		}
		if ctx.DefaultValue() != nil {
			result += " = " + ctx.DefaultValue().GetText()
		}
	}
	return result
}

func buildParameterSection(ctx parser.IFormalParameterSectionContext) string {
	result := ""
	if ctx.VAR() != nil {
		result += "var "
	}
	if ctx.CONST() != nil {
		result += "const "
	}
	if ctx.OUT() != nil {
		result += "out "
	}
	if ctx.FUNCTION() != nil {
		result += "function "
	}
	if ctx.PROCEDURE() != nil {
		result += "procedure "
	}
	if ctx.ParameterGroup() != nil {
		result += buildParameterGroup(ctx.ParameterGroup())
	}
	return result
}

func buildParameterList(ctx parser.IFormalParameterListContext) string {
	var params []string
	if ctx != nil {
		for _, paramsection := range ctx.AllFormalParameterSection() {
			params = append(params, buildParameterSection(paramsection))
		}
	}
	return strings.Join(params, "; ")
}

func buildProcedureOrFunctionHeaderModifiers(ctx parser.IProcedureOrFunctionHeaderModifiersContext) string {
	result := ""
	if len(ctx.AllABSTRACT()) > 0 {
		result += "abstract;"
	}
	if len(ctx.AllVIRTUAL()) > 0 {
		result += "virtual;"
	}
	if len(ctx.AllOVERRIDE()) > 0 {
		result += "override;"
	}
	if len(ctx.AllREINTRODUCE()) > 0 {
		result += "reintroduce;"
	}
	if len(ctx.AllOVERLOAD()) > 0 {
		result += "overload;"
	}
	if len(ctx.AllINLINE()) > 0 {
		result += "inline;"
	}
	if len(ctx.AllSTDCALL()) > 0 {
		result += "stdcall;"
	}
	if len(ctx.AllCDECL()) > 0 {
		result += "cdecl;"
	}
	if len(result) > 0 {
		result = "; " + result
	}
	return result
}

func buildIdentifier(ctx parser.IIdentifierContext) string {
	nodes := ctx.AllIDENT()
	if len(nodes) > 0 {
		var texts []string
		for _, node := range nodes {
			texts = append(texts, node.GetText())
		}
		return strings.Join(texts, ".")
	}
	if ctx.INDEX() != nil {
		return ctx.INDEX().GetText()
	}
	if ctx.READ() != nil {
		return ctx.READ().GetText()
	}
	if ctx.WRITE() != nil {
		return ctx.WRITE().GetText()
	}
	return ""
}

func buildIdentifiers(ctx parser.IIdentifierListContext) string {
	var ids []string
	for _, identifier := range ctx.AllIdentifier() {
		ids = append(ids, buildIdentifier(identifier))
	}
	return strings.Join(ids, ", ")
}

func buildSimpleExpression(ctx parser.ISimpleExpressionContext) string {
	return ctx.GetText()
}

func buildTypeIdentifier(ctx parser.ITypeIdentifierContext) string {
	if ctx.LT() != nil && ctx.GT() != nil {
		return buildIdentifier(ctx.Identifier()) + "<" + buildTypeIdentifier(ctx.TypeIdentifier()) + ">"
	}
	if ctx.Identifier() != nil {
		return buildIdentifier(ctx.Identifier())
	}
	if ctx.CHAR() != nil {
		return "char"
	}
	if ctx.BOOLEAN() != nil {
		return "boolean"
	}
	if ctx.INTEGER() != nil {
		return "integer"
	}
	if ctx.REAL() != nil {
		return "real"
	}
	if ctx.STRING() != nil {
		return "string"
	}
	if ctx.CARDINAL() != nil {
		return "cardinal"
	}
	if ctx.LONGBOOL() != nil {
		return "longbool"
	}
	if ctx.LONGINT() != nil {
		return "longint"
	}
	if ctx.ArrayType() != nil {
		return buildArrayType(ctx.ArrayType())
	}
	return ""
}

func buildTypedIdentifierList(ctx parser.ITypedIdentifierListContext) ([]string, string) {
	if ctx == nil {
		return []string{}, ""
	}
	typedef := buildUnderscoreTypeDef(ctx.Type_())
	list := []string{}
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		list = append(list, buildIdentifier(identifier))
	}
	return list, typedef
}

func buildArrayType(ctx parser.IArrayTypeContext) string {
	if ctx.TypeList() != nil {
		return "array[" + buildTypeList(ctx.TypeList()) + "] of " + buildUnderscoreTypeDef(ctx.Type_())
	}
	if ctx.Type_() != nil {
		return "array of " + buildUnderscoreTypeDef(ctx.Type_())
	}
	if ctx.CONST() != nil {
		return "array of const"
	}
	return ""
}

func buildTypeList(ctx parser.ITypeListContext) string {
	var types []string
	for _, indexType := range ctx.AllIndexType() {
		if indexType.SimpleType() != nil {
			types = append(types, buildSimpleTypeDef(indexType.SimpleType()))
		}
	}
	return strings.Join(types, ",")
}

func buildConstList(ctx parser.IConstListContext) string {
	consts := []string{}
	for _, constant := range ctx.AllConstant() {
		consts = append(consts, constant.GetText())
	}
	return strings.Join(consts, ",")
}
