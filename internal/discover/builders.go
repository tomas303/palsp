package discover

import (
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
		var builder strings.Builder
		builder.WriteString("^")
		builder.WriteString(buildTypeIdentifier(ctx.PointerType().TypeIdentifier()))
		return builder.String()
	}
	return ""
}

func buildSimpleTypeDef(ctx parser.ISimpleTypeContext) string {
	var builder strings.Builder // Use var instead of &strings.Builder{}
	if ctx.ScalarType() != nil {
		builder.WriteString("(")
		scalarMembers := ctx.ScalarType().ScalerList().AllScalerMember()
		for i, scalarMember := range scalarMembers {
			builder.WriteString(scalarMember.Identifier().GetText())
			if scalarMember.Expression() != nil {
				builder.WriteString(" = ")
				builder.WriteString(scalarMember.Expression().GetText())
			}
			if i < len(scalarMembers)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(")")
	}
	if ctx.SubrangeType() != nil {
		if len(ctx.SubrangeType().AllSimpleExpression()) > 0 {
			builder.WriteString(buildSimpleExpression(ctx.SubrangeType().SimpleExpression(0)))
		}
		builder.WriteString("..")
		if len(ctx.SubrangeType().AllSimpleExpression()) > 1 {
			builder.WriteString(buildSimpleExpression(ctx.SubrangeType().SimpleExpression(1)))
		}
	}
	if ctx.TypeIdentifier() != nil {
		builder.WriteString(buildTypeIdentifier(ctx.TypeIdentifier()))
	}
	if ctx.Stringtype() != nil {
		builder.WriteString("string")
		if ctx.Stringtype().Identifier() != nil {
			builder.WriteString("[")
			builder.WriteString(buildIdentifier(ctx.Stringtype().Identifier()))
			builder.WriteString("]")
		}
		if ctx.Stringtype().UnsignedNumber() != nil {
			builder.WriteString("[")
			builder.WriteString(ctx.Stringtype().UnsignedNumber().GetText())
			builder.WriteString("]")
		}
	}
	result := builder.String()
	if result == "" {
		var fallbackBuilder strings.Builder
		fallbackBuilder.WriteString("undetected: ")
		fallbackBuilder.WriteString(ctx.GetText())
		return fallbackBuilder.String()
	}
	return result
}

func buildStructuredTypeDef(ctx parser.IStructuredTypeContext) string {
	var builder strings.Builder
	if ctx.PACKED() != nil {
		builder.WriteString("packed ")
	}
	if ctx.UnpackedStructuredType() != nil {
		if ctx.UnpackedStructuredType().SetType() != nil {
			builder.WriteString(buildSetTypeDef(ctx.UnpackedStructuredType().SetType()))
		}
		if ctx.UnpackedStructuredType().FileType() != nil {
			builder.WriteString(buildFileType(ctx.UnpackedStructuredType().FileType()))
		}
	}
	if ctx.ClassType() != nil {
		builder.WriteString(buildClassTypeDef(ctx.ClassType()))
	}
	if ctx.RecordType() != nil {
		builder.WriteString(buildRecordTypeDef(ctx.RecordType()))
	}
	if ctx.InterfaceType() != nil {
		builder.WriteString(buildInterfaceTypeDef(ctx.InterfaceType()))
	}
	if ctx.HelperType() != nil {
		builder.WriteString(buildHelperTypeDef(ctx.HelperType()))
	}
	return builder.String()
}

func buildRecordTypeDef(ctx parser.IRecordTypeContext) string {
	var builder strings.Builder
	builder.WriteString("record\n")
	builder.WriteString(buildRecordParts(ctx.RecordParts()))
	builder.WriteString("end")
	return builder.String()
}

func buildRecordParts(ctx parser.IRecordPartsContext) string {
	var builder strings.Builder
	if ctx != nil {
		if ctx.RecordFixedPart() != nil {
			builder.WriteString(buildRecordFixedPart(ctx.RecordFixedPart()))
			if ctx.RecordVariantPart() != nil {
				builder.WriteString(buildRecordVariantPart(ctx.RecordVariantPart()))
			}
		} else if ctx.RecordVariantPart() != nil {
			builder.WriteString(buildRecordVariantPart(ctx.RecordVariantPart()))
		}
	}
	return builder.String()
}

func buildRecordFixedPart(ctx parser.IRecordFixedPartContext) string {
	var builder strings.Builder
	for _, typeIdlist := range ctx.AllTypedIdentifierList() {
		names, typedef := buildTypedIdentifierList(typeIdlist)
		builder.WriteString(strings.Join(names, ","))
		builder.WriteString(": ")
		builder.WriteString(typedef)
		builder.WriteString(";\n")
	}
	return builder.String()
}

func buildRecordVariantPart(ctx parser.IRecordVariantPartContext) string {
	var builder strings.Builder
	builder.WriteString("case ")
	builder.WriteString(ctx.Tag().GetText())
	builder.WriteString(" of\n")
	for _, variant := range ctx.AllRecordVariant() {
		builder.WriteString(buildRecordVariant(variant))
	}
	return builder.String()
}

func buildRecordVariant(ctx parser.IRecordVariantContext) string {
	var builder strings.Builder
	constList := buildConstList(ctx.ConstList())
	if constList != "" {
		builder.WriteString(constList)
		builder.WriteString(": ")
	}
	parts := buildRecordParts(ctx.RecordParts())
	if parts != "" {
		builder.WriteString("(\n")
		builder.WriteString(parts)
		builder.WriteString(");\n")
	}
	return builder.String()
}

func buildHelperTypeDef(ctx parser.IHelperTypeContext) string {
	var builder strings.Builder
	builder.WriteString("class helper for ")
	builder.WriteString(buildTypeIdentifier(ctx.TypeIdentifier()))
	builder.WriteString(" end")
	return builder.String()
}

func buildSetTypeDef(ctx parser.ISetTypeContext) string {
	var builder strings.Builder
	builder.WriteString("set of ")
	if ctx.SimpleType() != nil {
		builder.WriteString(buildSimpleTypeDef(ctx.SimpleType()))
	}
	return builder.String()
}

func buildFileType(ctx parser.IFileTypeContext) string {
	if ctx.Type_() != nil {
		var builder strings.Builder
		builder.WriteString("file of ")
		builder.WriteString(buildUnderscoreTypeDef(ctx.Type_()))
		return builder.String()
	}
	return "file"
}

func buildInterfaceTypeDef(ctx parser.IInterfaceTypeContext) string {
	var builder strings.Builder
	builder.WriteString("interface")
	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		if ctx.Identifier() != nil {
			builder.WriteString("(")
			builder.WriteString(buildIdentifier(ctx.Identifier()))
			builder.WriteString(")")
		}
	}
	builder.WriteString("\n")
	if ctx.GUID_LITERAL() != nil {
		builder.WriteString(ctx.GUID_LITERAL().GetText())
	}
	builder.WriteString("end")
	return builder.String()
}

func buildClassTypeDef(ctx parser.IClassTypeContext) string {
	var builder strings.Builder
	builder.WriteString("class")
	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		implements := []string{}
		if ctx.ClassImplementsInterfaces() != nil {
			for _, identifier := range ctx.ClassImplementsInterfaces().AllTypeIdentifier() {
				implements = append(implements, identifier.GetText())
			}
		}
		if ctx.Identifier() != nil {
			builder.WriteString("(")
			builder.WriteString(buildIdentifier(ctx.Identifier()))
			if len(implements) > 0 {
				builder.WriteString(", ")
				builder.WriteString(strings.Join(implements, ", "))
			}
			builder.WriteString(")")
		}
	}
	if ctx.ABSTRACT() != nil {
		builder.WriteString(" abstract")
	}
	return builder.String()
}

func buildFunctionTypeDef(ctx parser.IFunctionTypeContext) string {
	var builder strings.Builder
	if ctx.REFERENCE() != nil {
		builder.WriteString("reference to ")
	}
	params := buildParameterList(ctx.FormalParameterList())
	resultType := ""
	if ctx.ResultType() != nil {
		resultType = ctx.ResultType().TypeIdentifier().GetText()
	}

	if resultType != "" {
		builder.WriteString("function(")
		builder.WriteString(params)
		builder.WriteString("): ")
		builder.WriteString(resultType)
	} else {
		builder.WriteString("function(")
		builder.WriteString(params)
		builder.WriteString(")")
	}
	if ctx.OF() != nil {
		builder.WriteString(" of object ")
	}
	builder.WriteString(buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers()))
	return builder.String()
}

func buildProcedureTypeDef(ctx parser.IProcedureTypeContext) string {
	var builder strings.Builder
	if ctx.REFERENCE() != nil {
		builder.WriteString("reference to ")
	}
	params := buildParameterList(ctx.FormalParameterList())
	builder.WriteString("procedure(")
	builder.WriteString(params)
	builder.WriteString(")")
	builder.WriteString(buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers()))
	if ctx.OF() != nil {
		builder.WriteString(" of object ")
	}
	return builder.String()
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
	var builder strings.Builder
	if ctx.IdentifierList() != nil {
		identifiers := ctx.IdentifierList().AllIdentifier()
		for i, identifier := range identifiers {
			builder.WriteString(identifier.GetText())
			if i < len(identifiers)-1 {
				builder.WriteString(", ")
			}
		}
		if ctx.TypeIdentifier() != nil {
			builder.WriteString(": ")
			builder.WriteString(ctx.TypeIdentifier().GetText())
		}
		if ctx.DefaultValue() != nil {
			builder.WriteString(" = ")
			builder.WriteString(ctx.DefaultValue().GetText())
		}
	}
	return builder.String()
}

func getParameterSpecifier(ctx parser.IFormalParameterSectionContext) string {
	var builder strings.Builder
	if ctx.VAR() != nil {
		builder.WriteString("var")
	}
	if ctx.CONST() != nil {
		builder.WriteString("const")
	}
	if ctx.OUT() != nil {
		builder.WriteString("out")
	}
	if ctx.FUNCTION() != nil {
		builder.WriteString("function")
	}
	if ctx.PROCEDURE() != nil {
		builder.WriteString("procedure")
	}
	return builder.String()
}

func buildParameterSection(ctx parser.IFormalParameterSectionContext) string {
	var builder strings.Builder
	specifier := getParameterSpecifier(ctx)
	if len(specifier) > 0 {
		builder.WriteString(specifier)
		builder.WriteString(" ")
	}
	if ctx.ParameterGroup() != nil {
		builder.WriteString(buildParameterGroup(ctx.ParameterGroup()))
	}
	return builder.String()
}

func buildOneParameter(paramSectionCtx parser.IFormalParameterSectionContext, paramCtx parser.IIdentifierContext) string {
	var builder strings.Builder
	specifier := getParameterSpecifier(paramSectionCtx)
	if len(specifier) > 0 {
		builder.WriteString(specifier)
		builder.WriteString(" ")
	}
	builder.WriteString(paramCtx.GetText())
	if paramSectionCtx.ParameterGroup() != nil {
		if paramSectionCtx.ParameterGroup().TypeIdentifier() != nil {
			builder.WriteString(": ")
			builder.WriteString(paramSectionCtx.ParameterGroup().TypeIdentifier().GetText())
		}
		if paramSectionCtx.ParameterGroup().DefaultValue() != nil {
			builder.WriteString(" = ")
			builder.WriteString(paramSectionCtx.ParameterGroup().DefaultValue().GetText())
		}
	}
	return builder.String()
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
	var builder strings.Builder
	if ctx == nil {
		return ""
	}
	if len(ctx.AllABSTRACT()) > 0 {
		builder.WriteString("abstract;")
	}
	if len(ctx.AllVIRTUAL()) > 0 {
		builder.WriteString("virtual;")
	}
	if len(ctx.AllOVERRIDE()) > 0 {
		builder.WriteString("override;")
	}
	if len(ctx.AllREINTRODUCE()) > 0 {
		builder.WriteString("reintroduce;")
	}
	if len(ctx.AllOVERLOAD()) > 0 {
		builder.WriteString("overload;")
	}
	if len(ctx.AllINLINE()) > 0 {
		builder.WriteString("inline;")
	}
	if len(ctx.AllSTDCALL()) > 0 {
		builder.WriteString("stdcall;")
	}
	if len(ctx.AllCDECL()) > 0 {
		builder.WriteString("cdecl;")
	}
	result := builder.String()
	if len(result) > 0 {
		var finalBuilder strings.Builder
		finalBuilder.WriteString("; ")
		finalBuilder.WriteString(result)
		return finalBuilder.String()
	}
	return result
}

func getLastIdent(ctx parser.IIdentifierContext) string {
	if ctx != nil {
		identTokens := ctx.AllIdentifierPart()
		if len(identTokens) == 0 {
			return ""
		}
		lastIdent := identTokens[len(identTokens)-1]
		return lastIdent.GetText()
	}
	return ""
}

func buildIdentifierPart(ctx parser.IIdentifierPartContext) string {
	var builder strings.Builder
	if ctx.IDENT() != nil {
		builder.WriteString(ctx.IDENT().GetText())
	}
	if ctx.INDEX() != nil {
		builder.WriteString(ctx.INDEX().GetText())
	}
	if ctx.READ() != nil {
		builder.WriteString(ctx.READ().GetText())
	}
	if ctx.WRITE() != nil {
		builder.WriteString(ctx.WRITE().GetText())
	}
	if ctx.GenericTemplate() != nil {
		builder.WriteString(buildGenericTemplate(ctx.GenericTemplate()))
	}
	return builder.String()
}

func buildIdentifier(ctx parser.IIdentifierContext) string {
	var builder strings.Builder
	if ctx != nil {
		for i, part := range ctx.AllIdentifierPart() {
			if i > 0 {
				builder.WriteString(".")
			}
			builder.WriteString(buildIdentifierPart(part))
		}
	}
	return builder.String()
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
	var builder strings.Builder
	builder.WriteString("<")
	typeidentifiers := ctx.GenericTemplateList().AllGenericTypeParameter()
	for i, typeParam := range typeidentifiers {
		builder.WriteString(buildTypeIdentifier(typeParam.TypeIdentifier()))
		if i < len(typeidentifiers)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString(">")
	return builder.String()
}

func buildTypeIdentifier(ctx parser.ITypeIdentifierContext) string {
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
	var builder strings.Builder
	if ctx.TypeList() != nil {
		builder.WriteString("array[")
		builder.WriteString(buildTypeList(ctx.TypeList()))
		builder.WriteString("] of ")
		builder.WriteString(buildUnderscoreTypeDef(ctx.Type_()))
		return builder.String()
	}
	if ctx.Type_() != nil {
		builder.WriteString("array of ")
		builder.WriteString(buildUnderscoreTypeDef(ctx.Type_()))
		return builder.String()
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

func buildProcedureHeader(ctx parser.IProcedureHeaderContext) string {
	var builder strings.Builder
	if ctx.CLASS() != nil {
		builder.WriteString("class ")
	}
	if ctx.PROCEDURE() != nil {
		builder.WriteString("procedure ")
	} else if ctx.CONSTRUCTOR() != nil {
		builder.WriteString("constructor ")
	} else if ctx.DESTRUCTOR() != nil {
		builder.WriteString("destructor ")
	}
	if ctx.Identifier() != nil {
		builder.WriteString(buildIdentifier(ctx.Identifier()))
	}
	builder.WriteString("(")
	builder.WriteString(buildParameterList(ctx.FormalParameterList()))
	builder.WriteString(")")
	builder.WriteString(buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers()))
	return builder.String()
}

func buildFunctionHeader(ctx parser.IFunctionHeaderContext) string {
	var builder strings.Builder
	if ctx.CLASS() != nil {
		builder.WriteString("class ")
	}
	builder.WriteString("function ")
	if ctx.Identifier() != nil {
		builder.WriteString(buildIdentifier(ctx.Identifier()))
	}
	builder.WriteString("(")
	builder.WriteString(buildParameterList(ctx.FormalParameterList()))
	builder.WriteString(")")
	if ctx.ResultType() == nil || ctx.ResultType().TypeIdentifier() == nil {
		builder.WriteString(": unknown")
	} else {
		builder.WriteString(": ")
		builder.WriteString(buildTypeIdentifier(ctx.ResultType().TypeIdentifier()))
	}
	builder.WriteString(buildProcedureOrFunctionHeaderModifiers(ctx.ProcedureOrFunctionHeaderModifiers()))
	return builder.String()
}

func buildProperty(ctx *parser.PropertyDeclarationContext) string {
	var builder strings.Builder
	builder.WriteString("property ")
	if ctx.Identifier() != nil {
		builder.WriteString(buildIdentifier(ctx.Identifier()))
	}
	if ctx.TypeIdentifier() != nil {
		builder.WriteString(": ")
		builder.WriteString(buildTypeIdentifier(ctx.TypeIdentifier()))
	}
	return builder.String()
}
