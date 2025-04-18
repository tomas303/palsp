package discover

import (
	"fmt"
	"palsp/internal/parser"
	"strings"
)

func buildUnderscoreTypeDef(ctx parser.IType_Context) string {
	if ctx == nil {
		return ""
	}
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
	if ctx == nil {
		return result
	}
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
	if ctx != nil {
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

func buildGenericTemplate(ctx parser.IGenericTemplateContext) string {
	if ctx == nil {
		return ""
	}
	result := "<"
	templateItems := ctx.GenericTemplateList().AllGenericTemplateItem()
	for i, templateItem := range templateItems {
		if templateItem.Identifier() != nil {
			result += buildIdentifier(templateItem.Identifier())
		} else if templateItem.GenericTemplate() != nil {
			result += buildGenericTemplate(templateItem.GenericTemplate())
		}
		if i < len(templateItems)-1 {
			result += ","
		}
	}
	result += ">"
	return result
}

func buildTypeIdentifier(ctx parser.ITypeIdentifierContext) string {
	// if ctx.LT() != nil && ctx.GT() != nil {
	// 	return buildIdentifier(ctx.Identifier()) + "<" + buildTypeIdentifier(ctx.TypeIdentifier()) + ">"
	// }

	if ctx.Identifier() != nil {
		result := buildIdentifier(ctx.Identifier())
		if ctx.GenericTemplate() != nil {
			result += buildGenericTemplate(ctx.GenericTemplate())
		}
		return result
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

func buildProcedureHeader(ctx *parser.ProcedureHeaderContext) string {
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
	return definition
}

func buildFunctionHeader(ctx *parser.FunctionHeaderContext) string {
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
	if ctx.ResultType() == nil || ctx.ResultType().TypeIdentifier() == nil {
		definition += ": unknown"
	} else {
		definition += ": " + buildTypeIdentifier(ctx.ResultType().TypeIdentifier())
	}
	definition += buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers())
	return definition
}
