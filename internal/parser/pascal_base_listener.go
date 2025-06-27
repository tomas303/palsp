// Code generated from /home/tomas/development/tomas303/projects/pascallsp/palsp/internal/pascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // pascal

import "github.com/antlr4-go/antlr/v4"

// BasepascalListener is a complete listener for a parse tree produced by pascalParser.
type BasepascalListener struct{}

var _ pascalListener = &BasepascalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepascalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepascalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepascalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepascalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSource is called when production source is entered.
func (s *BasepascalListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BasepascalListener) ExitSource(ctx *SourceContext) {}

// EnterProgram is called when production program is entered.
func (s *BasepascalListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasepascalListener) ExitProgram(ctx *ProgramContext) {}

// EnterUnit is called when production unit is entered.
func (s *BasepascalListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BasepascalListener) ExitUnit(ctx *UnitContext) {}

// EnterInterfaceSection is called when production interfaceSection is entered.
func (s *BasepascalListener) EnterInterfaceSection(ctx *InterfaceSectionContext) {}

// ExitInterfaceSection is called when production interfaceSection is exited.
func (s *BasepascalListener) ExitInterfaceSection(ctx *InterfaceSectionContext) {}

// EnterImplementationSection is called when production implementationSection is entered.
func (s *BasepascalListener) EnterImplementationSection(ctx *ImplementationSectionContext) {}

// ExitImplementationSection is called when production implementationSection is exited.
func (s *BasepascalListener) ExitImplementationSection(ctx *ImplementationSectionContext) {}

// EnterInitializationSection is called when production initializationSection is entered.
func (s *BasepascalListener) EnterInitializationSection(ctx *InitializationSectionContext) {}

// ExitInitializationSection is called when production initializationSection is exited.
func (s *BasepascalListener) ExitInitializationSection(ctx *InitializationSectionContext) {}

// EnterFinalizationSection is called when production finalizationSection is entered.
func (s *BasepascalListener) EnterFinalizationSection(ctx *FinalizationSectionContext) {}

// ExitFinalizationSection is called when production finalizationSection is exited.
func (s *BasepascalListener) ExitFinalizationSection(ctx *FinalizationSectionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasepascalListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasepascalListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BasepascalListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BasepascalListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterIdentifierPart is called when production identifierPart is entered.
func (s *BasepascalListener) EnterIdentifierPart(ctx *IdentifierPartContext) {}

// ExitIdentifierPart is called when production identifierPart is exited.
func (s *BasepascalListener) ExitIdentifierPart(ctx *IdentifierPartContext) {}

// EnterInterfaceBlockMember is called when production interfaceBlockMember is entered.
func (s *BasepascalListener) EnterInterfaceBlockMember(ctx *InterfaceBlockMemberContext) {}

// ExitInterfaceBlockMember is called when production interfaceBlockMember is exited.
func (s *BasepascalListener) ExitInterfaceBlockMember(ctx *InterfaceBlockMemberContext) {}

// EnterInterfaceBlock is called when production interfaceBlock is entered.
func (s *BasepascalListener) EnterInterfaceBlock(ctx *InterfaceBlockContext) {}

// ExitInterfaceBlock is called when production interfaceBlock is exited.
func (s *BasepascalListener) ExitInterfaceBlock(ctx *InterfaceBlockContext) {}

// EnterImplementationBlockMember is called when production implementationBlockMember is entered.
func (s *BasepascalListener) EnterImplementationBlockMember(ctx *ImplementationBlockMemberContext) {}

// ExitImplementationBlockMember is called when production implementationBlockMember is exited.
func (s *BasepascalListener) ExitImplementationBlockMember(ctx *ImplementationBlockMemberContext) {}

// EnterImplementationBlock is called when production implementationBlock is entered.
func (s *BasepascalListener) EnterImplementationBlock(ctx *ImplementationBlockContext) {}

// ExitImplementationBlock is called when production implementationBlock is exited.
func (s *BasepascalListener) ExitImplementationBlock(ctx *ImplementationBlockContext) {}

// EnterFuncBlockMemeber is called when production funcBlockMemeber is entered.
func (s *BasepascalListener) EnterFuncBlockMemeber(ctx *FuncBlockMemeberContext) {}

// ExitFuncBlockMemeber is called when production funcBlockMemeber is exited.
func (s *BasepascalListener) ExitFuncBlockMemeber(ctx *FuncBlockMemeberContext) {}

// EnterFuncBlock is called when production funcBlock is entered.
func (s *BasepascalListener) EnterFuncBlock(ctx *FuncBlockContext) {}

// ExitFuncBlock is called when production funcBlock is exited.
func (s *BasepascalListener) ExitFuncBlock(ctx *FuncBlockContext) {}

// EnterUsesUnits is called when production usesUnits is entered.
func (s *BasepascalListener) EnterUsesUnits(ctx *UsesUnitsContext) {}

// ExitUsesUnits is called when production usesUnits is exited.
func (s *BasepascalListener) ExitUsesUnits(ctx *UsesUnitsContext) {}

// EnterLabelDeclarationPart is called when production labelDeclarationPart is entered.
func (s *BasepascalListener) EnterLabelDeclarationPart(ctx *LabelDeclarationPartContext) {}

// ExitLabelDeclarationPart is called when production labelDeclarationPart is exited.
func (s *BasepascalListener) ExitLabelDeclarationPart(ctx *LabelDeclarationPartContext) {}

// EnterLabel is called when production label is entered.
func (s *BasepascalListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BasepascalListener) ExitLabel(ctx *LabelContext) {}

// EnterConstantDefinitionPart is called when production constantDefinitionPart is entered.
func (s *BasepascalListener) EnterConstantDefinitionPart(ctx *ConstantDefinitionPartContext) {}

// ExitConstantDefinitionPart is called when production constantDefinitionPart is exited.
func (s *BasepascalListener) ExitConstantDefinitionPart(ctx *ConstantDefinitionPartContext) {}

// EnterConstantDefinition is called when production constantDefinition is entered.
func (s *BasepascalListener) EnterConstantDefinition(ctx *ConstantDefinitionContext) {}

// ExitConstantDefinition is called when production constantDefinition is exited.
func (s *BasepascalListener) ExitConstantDefinition(ctx *ConstantDefinitionContext) {}

// EnterConstantChr is called when production constantChr is entered.
func (s *BasepascalListener) EnterConstantChr(ctx *ConstantChrContext) {}

// ExitConstantChr is called when production constantChr is exited.
func (s *BasepascalListener) ExitConstantChr(ctx *ConstantChrContext) {}

// EnterHexConstant is called when production hexConstant is entered.
func (s *BasepascalListener) EnterHexConstant(ctx *HexConstantContext) {}

// ExitHexConstant is called when production hexConstant is exited.
func (s *BasepascalListener) ExitHexConstant(ctx *HexConstantContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasepascalListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasepascalListener) ExitConstant(ctx *ConstantContext) {}

// EnterArrayConstant is called when production arrayConstant is entered.
func (s *BasepascalListener) EnterArrayConstant(ctx *ArrayConstantContext) {}

// ExitArrayConstant is called when production arrayConstant is exited.
func (s *BasepascalListener) ExitArrayConstant(ctx *ArrayConstantContext) {}

// EnterRecordConstant is called when production recordConstant is entered.
func (s *BasepascalListener) EnterRecordConstant(ctx *RecordConstantContext) {}

// ExitRecordConstant is called when production recordConstant is exited.
func (s *BasepascalListener) ExitRecordConstant(ctx *RecordConstantContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BasepascalListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BasepascalListener) ExitRecordField(ctx *RecordFieldContext) {}

// EnterUnsignedNumber is called when production unsignedNumber is entered.
func (s *BasepascalListener) EnterUnsignedNumber(ctx *UnsignedNumberContext) {}

// ExitUnsignedNumber is called when production unsignedNumber is exited.
func (s *BasepascalListener) ExitUnsignedNumber(ctx *UnsignedNumberContext) {}

// EnterUnsignedInteger is called when production unsignedInteger is entered.
func (s *BasepascalListener) EnterUnsignedInteger(ctx *UnsignedIntegerContext) {}

// ExitUnsignedInteger is called when production unsignedInteger is exited.
func (s *BasepascalListener) ExitUnsignedInteger(ctx *UnsignedIntegerContext) {}

// EnterUnsignedReal is called when production unsignedReal is entered.
func (s *BasepascalListener) EnterUnsignedReal(ctx *UnsignedRealContext) {}

// ExitUnsignedReal is called when production unsignedReal is exited.
func (s *BasepascalListener) ExitUnsignedReal(ctx *UnsignedRealContext) {}

// EnterSign is called when production sign is entered.
func (s *BasepascalListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BasepascalListener) ExitSign(ctx *SignContext) {}

// EnterBool_ is called when production bool_ is entered.
func (s *BasepascalListener) EnterBool_(ctx *Bool_Context) {}

// ExitBool_ is called when production bool_ is exited.
func (s *BasepascalListener) ExitBool_(ctx *Bool_Context) {}

// EnterString is called when production string is entered.
func (s *BasepascalListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasepascalListener) ExitString(ctx *StringContext) {}

// EnterStringExpression is called when production stringExpression is entered.
func (s *BasepascalListener) EnterStringExpression(ctx *StringExpressionContext) {}

// ExitStringExpression is called when production stringExpression is exited.
func (s *BasepascalListener) ExitStringExpression(ctx *StringExpressionContext) {}

// EnterResourceDefinitionPart is called when production resourceDefinitionPart is entered.
func (s *BasepascalListener) EnterResourceDefinitionPart(ctx *ResourceDefinitionPartContext) {}

// ExitResourceDefinitionPart is called when production resourceDefinitionPart is exited.
func (s *BasepascalListener) ExitResourceDefinitionPart(ctx *ResourceDefinitionPartContext) {}

// EnterResourceDefinition is called when production resourceDefinition is entered.
func (s *BasepascalListener) EnterResourceDefinition(ctx *ResourceDefinitionContext) {}

// ExitResourceDefinition is called when production resourceDefinition is exited.
func (s *BasepascalListener) ExitResourceDefinition(ctx *ResourceDefinitionContext) {}

// EnterDeprecatedHint is called when production deprecatedHint is entered.
func (s *BasepascalListener) EnterDeprecatedHint(ctx *DeprecatedHintContext) {}

// ExitDeprecatedHint is called when production deprecatedHint is exited.
func (s *BasepascalListener) ExitDeprecatedHint(ctx *DeprecatedHintContext) {}

// EnterPlatformHint is called when production platformHint is entered.
func (s *BasepascalListener) EnterPlatformHint(ctx *PlatformHintContext) {}

// ExitPlatformHint is called when production platformHint is exited.
func (s *BasepascalListener) ExitPlatformHint(ctx *PlatformHintContext) {}

// EnterAlignHint is called when production alignHint is entered.
func (s *BasepascalListener) EnterAlignHint(ctx *AlignHintContext) {}

// ExitAlignHint is called when production alignHint is exited.
func (s *BasepascalListener) ExitAlignHint(ctx *AlignHintContext) {}

// EnterTypeDefinitionPart is called when production typeDefinitionPart is entered.
func (s *BasepascalListener) EnterTypeDefinitionPart(ctx *TypeDefinitionPartContext) {}

// ExitTypeDefinitionPart is called when production typeDefinitionPart is exited.
func (s *BasepascalListener) ExitTypeDefinitionPart(ctx *TypeDefinitionPartContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BasepascalListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BasepascalListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterClassType is called when production classType is entered.
func (s *BasepascalListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BasepascalListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterClassSection is called when production classSection is entered.
func (s *BasepascalListener) EnterClassSection(ctx *ClassSectionContext) {}

// ExitClassSection is called when production classSection is exited.
func (s *BasepascalListener) ExitClassSection(ctx *ClassSectionContext) {}

// EnterClassTypeBlock is called when production classTypeBlock is entered.
func (s *BasepascalListener) EnterClassTypeBlock(ctx *ClassTypeBlockContext) {}

// ExitClassTypeBlock is called when production classTypeBlock is exited.
func (s *BasepascalListener) ExitClassTypeBlock(ctx *ClassTypeBlockContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BasepascalListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BasepascalListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BasepascalListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BasepascalListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterProcedureType is called when production procedureType is entered.
func (s *BasepascalListener) EnterProcedureType(ctx *ProcedureTypeContext) {}

// ExitProcedureType is called when production procedureType is exited.
func (s *BasepascalListener) ExitProcedureType(ctx *ProcedureTypeContext) {}

// EnterAliasDistinctType is called when production aliasDistinctType is entered.
func (s *BasepascalListener) EnterAliasDistinctType(ctx *AliasDistinctTypeContext) {}

// ExitAliasDistinctType is called when production aliasDistinctType is exited.
func (s *BasepascalListener) ExitAliasDistinctType(ctx *AliasDistinctTypeContext) {}

// EnterClassImplementsInterfaces is called when production classImplementsInterfaces is entered.
func (s *BasepascalListener) EnterClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) {}

// ExitClassImplementsInterfaces is called when production classImplementsInterfaces is exited.
func (s *BasepascalListener) ExitClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *BasepascalListener) EnterAccessSpecifier(ctx *AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *BasepascalListener) ExitAccessSpecifier(ctx *AccessSpecifierContext) {}

// EnterClassDeclarationPart is called when production classDeclarationPart is entered.
func (s *BasepascalListener) EnterClassDeclarationPart(ctx *ClassDeclarationPartContext) {}

// ExitClassDeclarationPart is called when production classDeclarationPart is exited.
func (s *BasepascalListener) ExitClassDeclarationPart(ctx *ClassDeclarationPartContext) {}

// EnterInterfaceGuidConst is called when production interfaceGuidConst is entered.
func (s *BasepascalListener) EnterInterfaceGuidConst(ctx *InterfaceGuidConstContext) {}

// ExitInterfaceGuidConst is called when production interfaceGuidConst is exited.
func (s *BasepascalListener) ExitInterfaceGuidConst(ctx *InterfaceGuidConstContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BasepascalListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BasepascalListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterInterfaceDeclarationPart is called when production interfaceDeclarationPart is entered.
func (s *BasepascalListener) EnterInterfaceDeclarationPart(ctx *InterfaceDeclarationPartContext) {}

// ExitInterfaceDeclarationPart is called when production interfaceDeclarationPart is exited.
func (s *BasepascalListener) ExitInterfaceDeclarationPart(ctx *InterfaceDeclarationPartContext) {}

// EnterErrorInterfaceDeclarationPart is called when production errorInterfaceDeclarationPart is entered.
func (s *BasepascalListener) EnterErrorInterfaceDeclarationPart(ctx *ErrorInterfaceDeclarationPartContext) {
}

// ExitErrorInterfaceDeclarationPart is called when production errorInterfaceDeclarationPart is exited.
func (s *BasepascalListener) ExitErrorInterfaceDeclarationPart(ctx *ErrorInterfaceDeclarationPartContext) {
}

// EnterErrorClassDeclarationPart is called when production errorClassDeclarationPart is entered.
func (s *BasepascalListener) EnterErrorClassDeclarationPart(ctx *ErrorClassDeclarationPartContext) {}

// ExitErrorClassDeclarationPart is called when production errorClassDeclarationPart is exited.
func (s *BasepascalListener) ExitErrorClassDeclarationPart(ctx *ErrorClassDeclarationPartContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BasepascalListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BasepascalListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterPropertyReadDeclaration is called when production propertyReadDeclaration is entered.
func (s *BasepascalListener) EnterPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) {}

// ExitPropertyReadDeclaration is called when production propertyReadDeclaration is exited.
func (s *BasepascalListener) ExitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) {}

// EnterPropertyWriteDeclaration is called when production propertyWriteDeclaration is entered.
func (s *BasepascalListener) EnterPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) {}

// ExitPropertyWriteDeclaration is called when production propertyWriteDeclaration is exited.
func (s *BasepascalListener) ExitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) {}

// EnterPropertyDefaultValueDeclaration is called when production propertyDefaultValueDeclaration is entered.
func (s *BasepascalListener) EnterPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) {
}

// ExitPropertyDefaultValueDeclaration is called when production propertyDefaultValueDeclaration is exited.
func (s *BasepascalListener) ExitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) {
}

// EnterPropertyIndexDeclaration is called when production propertyIndexDeclaration is entered.
func (s *BasepascalListener) EnterPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) {}

// ExitPropertyIndexDeclaration is called when production propertyIndexDeclaration is exited.
func (s *BasepascalListener) ExitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) {}

// EnterPropertyIndexParameters is called when production propertyIndexParameters is entered.
func (s *BasepascalListener) EnterPropertyIndexParameters(ctx *PropertyIndexParametersContext) {}

// ExitPropertyIndexParameters is called when production propertyIndexParameters is exited.
func (s *BasepascalListener) ExitPropertyIndexParameters(ctx *PropertyIndexParametersContext) {}

// EnterPropertyIndexParametersList is called when production propertyIndexParametersList is entered.
func (s *BasepascalListener) EnterPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) {
}

// ExitPropertyIndexParametersList is called when production propertyIndexParametersList is exited.
func (s *BasepascalListener) ExitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) {
}

// EnterGenericTemplate is called when production genericTemplate is entered.
func (s *BasepascalListener) EnterGenericTemplate(ctx *GenericTemplateContext) {}

// ExitGenericTemplate is called when production genericTemplate is exited.
func (s *BasepascalListener) ExitGenericTemplate(ctx *GenericTemplateContext) {}

// EnterGenericTemplateList is called when production genericTemplateList is entered.
func (s *BasepascalListener) EnterGenericTemplateList(ctx *GenericTemplateListContext) {}

// ExitGenericTemplateList is called when production genericTemplateList is exited.
func (s *BasepascalListener) ExitGenericTemplateList(ctx *GenericTemplateListContext) {}

// EnterGenericTypeParameter is called when production genericTypeParameter is entered.
func (s *BasepascalListener) EnterGenericTypeParameter(ctx *GenericTypeParameterContext) {}

// ExitGenericTypeParameter is called when production genericTypeParameter is exited.
func (s *BasepascalListener) ExitGenericTypeParameter(ctx *GenericTypeParameterContext) {}

// EnterGenericConstraints is called when production genericConstraints is entered.
func (s *BasepascalListener) EnterGenericConstraints(ctx *GenericConstraintsContext) {}

// ExitGenericConstraints is called when production genericConstraints is exited.
func (s *BasepascalListener) ExitGenericConstraints(ctx *GenericConstraintsContext) {}

// EnterGenericConstraint is called when production genericConstraint is entered.
func (s *BasepascalListener) EnterGenericConstraint(ctx *GenericConstraintContext) {}

// ExitGenericConstraint is called when production genericConstraint is exited.
func (s *BasepascalListener) ExitGenericConstraint(ctx *GenericConstraintContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasepascalListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasepascalListener) ExitType_(ctx *Type_Context) {}

// EnterSimpleType is called when production simpleType is entered.
func (s *BasepascalListener) EnterSimpleType(ctx *SimpleTypeContext) {}

// ExitSimpleType is called when production simpleType is exited.
func (s *BasepascalListener) ExitSimpleType(ctx *SimpleTypeContext) {}

// EnterScalarType is called when production scalarType is entered.
func (s *BasepascalListener) EnterScalarType(ctx *ScalarTypeContext) {}

// ExitScalarType is called when production scalarType is exited.
func (s *BasepascalListener) ExitScalarType(ctx *ScalarTypeContext) {}

// EnterScalerList is called when production scalerList is entered.
func (s *BasepascalListener) EnterScalerList(ctx *ScalerListContext) {}

// ExitScalerList is called when production scalerList is exited.
func (s *BasepascalListener) ExitScalerList(ctx *ScalerListContext) {}

// EnterScalerMember is called when production scalerMember is entered.
func (s *BasepascalListener) EnterScalerMember(ctx *ScalerMemberContext) {}

// ExitScalerMember is called when production scalerMember is exited.
func (s *BasepascalListener) ExitScalerMember(ctx *ScalerMemberContext) {}

// EnterSubrangeType is called when production subrangeType is entered.
func (s *BasepascalListener) EnterSubrangeType(ctx *SubrangeTypeContext) {}

// ExitSubrangeType is called when production subrangeType is exited.
func (s *BasepascalListener) ExitSubrangeType(ctx *SubrangeTypeContext) {}

// EnterStringTypeIdentifier is called when production stringTypeIdentifier is entered.
func (s *BasepascalListener) EnterStringTypeIdentifier(ctx *StringTypeIdentifierContext) {}

// ExitStringTypeIdentifier is called when production stringTypeIdentifier is exited.
func (s *BasepascalListener) ExitStringTypeIdentifier(ctx *StringTypeIdentifierContext) {}

// EnterStructuredType is called when production structuredType is entered.
func (s *BasepascalListener) EnterStructuredType(ctx *StructuredTypeContext) {}

// ExitStructuredType is called when production structuredType is exited.
func (s *BasepascalListener) ExitStructuredType(ctx *StructuredTypeContext) {}

// EnterUnpackedStructuredType is called when production unpackedStructuredType is entered.
func (s *BasepascalListener) EnterUnpackedStructuredType(ctx *UnpackedStructuredTypeContext) {}

// ExitUnpackedStructuredType is called when production unpackedStructuredType is exited.
func (s *BasepascalListener) ExitUnpackedStructuredType(ctx *UnpackedStructuredTypeContext) {}

// EnterStringtype is called when production stringtype is entered.
func (s *BasepascalListener) EnterStringtype(ctx *StringtypeContext) {}

// ExitStringtype is called when production stringtype is exited.
func (s *BasepascalListener) ExitStringtype(ctx *StringtypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BasepascalListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BasepascalListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BasepascalListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BasepascalListener) ExitTypeList(ctx *TypeListContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BasepascalListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BasepascalListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BasepascalListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BasepascalListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterRecordContent is called when production recordContent is entered.
func (s *BasepascalListener) EnterRecordContent(ctx *RecordContentContext) {}

// ExitRecordContent is called when production recordContent is exited.
func (s *BasepascalListener) ExitRecordContent(ctx *RecordContentContext) {}

// EnterRecordSection is called when production recordSection is entered.
func (s *BasepascalListener) EnterRecordSection(ctx *RecordSectionContext) {}

// ExitRecordSection is called when production recordSection is exited.
func (s *BasepascalListener) ExitRecordSection(ctx *RecordSectionContext) {}

// EnterRecordFieldsSection is called when production recordFieldsSection is entered.
func (s *BasepascalListener) EnterRecordFieldsSection(ctx *RecordFieldsSectionContext) {}

// ExitRecordFieldsSection is called when production recordFieldsSection is exited.
func (s *BasepascalListener) ExitRecordFieldsSection(ctx *RecordFieldsSectionContext) {}

// EnterRecordVariantSection is called when production recordVariantSection is entered.
func (s *BasepascalListener) EnterRecordVariantSection(ctx *RecordVariantSectionContext) {}

// ExitRecordVariantSection is called when production recordVariantSection is exited.
func (s *BasepascalListener) ExitRecordVariantSection(ctx *RecordVariantSectionContext) {}

// EnterTag is called when production tag is entered.
func (s *BasepascalListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasepascalListener) ExitTag(ctx *TagContext) {}

// EnterRecordVariant is called when production recordVariant is entered.
func (s *BasepascalListener) EnterRecordVariant(ctx *RecordVariantContext) {}

// ExitRecordVariant is called when production recordVariant is exited.
func (s *BasepascalListener) ExitRecordVariant(ctx *RecordVariantContext) {}

// EnterHelperType is called when production helperType is entered.
func (s *BasepascalListener) EnterHelperType(ctx *HelperTypeContext) {}

// ExitHelperType is called when production helperType is exited.
func (s *BasepascalListener) ExitHelperType(ctx *HelperTypeContext) {}

// EnterHelperDeclarationPart is called when production helperDeclarationPart is entered.
func (s *BasepascalListener) EnterHelperDeclarationPart(ctx *HelperDeclarationPartContext) {}

// ExitHelperDeclarationPart is called when production helperDeclarationPart is exited.
func (s *BasepascalListener) ExitHelperDeclarationPart(ctx *HelperDeclarationPartContext) {}

// EnterErrorHelperDeclarationPart is called when production errorHelperDeclarationPart is entered.
func (s *BasepascalListener) EnterErrorHelperDeclarationPart(ctx *ErrorHelperDeclarationPartContext) {
}

// ExitErrorHelperDeclarationPart is called when production errorHelperDeclarationPart is exited.
func (s *BasepascalListener) ExitErrorHelperDeclarationPart(ctx *ErrorHelperDeclarationPartContext) {}

// EnterSetType is called when production setType is entered.
func (s *BasepascalListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BasepascalListener) ExitSetType(ctx *SetTypeContext) {}

// EnterFileType is called when production fileType is entered.
func (s *BasepascalListener) EnterFileType(ctx *FileTypeContext) {}

// ExitFileType is called when production fileType is exited.
func (s *BasepascalListener) ExitFileType(ctx *FileTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BasepascalListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BasepascalListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterVariableDeclarationPart is called when production variableDeclarationPart is entered.
func (s *BasepascalListener) EnterVariableDeclarationPart(ctx *VariableDeclarationPartContext) {}

// ExitVariableDeclarationPart is called when production variableDeclarationPart is exited.
func (s *BasepascalListener) ExitVariableDeclarationPart(ctx *VariableDeclarationPartContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BasepascalListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BasepascalListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterProcedureHeader is called when production procedureHeader is entered.
func (s *BasepascalListener) EnterProcedureHeader(ctx *ProcedureHeaderContext) {}

// ExitProcedureHeader is called when production procedureHeader is exited.
func (s *BasepascalListener) ExitProcedureHeader(ctx *ProcedureHeaderContext) {}

// EnterFunctionHeader is called when production functionHeader is entered.
func (s *BasepascalListener) EnterFunctionHeader(ctx *FunctionHeaderContext) {}

// ExitFunctionHeader is called when production functionHeader is exited.
func (s *BasepascalListener) ExitFunctionHeader(ctx *FunctionHeaderContext) {}

// EnterProcedureOrFunctionHeader is called when production procedureOrFunctionHeader is entered.
func (s *BasepascalListener) EnterProcedureOrFunctionHeader(ctx *ProcedureOrFunctionHeaderContext) {}

// ExitProcedureOrFunctionHeader is called when production procedureOrFunctionHeader is exited.
func (s *BasepascalListener) ExitProcedureOrFunctionHeader(ctx *ProcedureOrFunctionHeaderContext) {}

// EnterProcedureOrFunctionHeaderModifiers is called when production procedureOrFunctionHeaderModifiers is entered.
func (s *BasepascalListener) EnterProcedureOrFunctionHeaderModifiers(ctx *ProcedureOrFunctionHeaderModifiersContext) {
}

// ExitProcedureOrFunctionHeaderModifiers is called when production procedureOrFunctionHeaderModifiers is exited.
func (s *BasepascalListener) ExitProcedureOrFunctionHeaderModifiers(ctx *ProcedureOrFunctionHeaderModifiersContext) {
}

// EnterProcedureOrFunctionDeclaration is called when production procedureOrFunctionDeclaration is entered.
func (s *BasepascalListener) EnterProcedureOrFunctionDeclaration(ctx *ProcedureOrFunctionDeclarationContext) {
}

// ExitProcedureOrFunctionDeclaration is called when production procedureOrFunctionDeclaration is exited.
func (s *BasepascalListener) ExitProcedureOrFunctionDeclaration(ctx *ProcedureOrFunctionDeclarationContext) {
}

// EnterProcedureDeclaration is called when production procedureDeclaration is entered.
func (s *BasepascalListener) EnterProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// ExitProcedureDeclaration is called when production procedureDeclaration is exited.
func (s *BasepascalListener) ExitProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasepascalListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasepascalListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterProcedureLambdaDeclaration is called when production procedureLambdaDeclaration is entered.
func (s *BasepascalListener) EnterProcedureLambdaDeclaration(ctx *ProcedureLambdaDeclarationContext) {
}

// ExitProcedureLambdaDeclaration is called when production procedureLambdaDeclaration is exited.
func (s *BasepascalListener) ExitProcedureLambdaDeclaration(ctx *ProcedureLambdaDeclarationContext) {}

// EnterFunctionLambdaDeclaration is called when production functionLambdaDeclaration is entered.
func (s *BasepascalListener) EnterFunctionLambdaDeclaration(ctx *FunctionLambdaDeclarationContext) {}

// ExitFunctionLambdaDeclaration is called when production functionLambdaDeclaration is exited.
func (s *BasepascalListener) ExitFunctionLambdaDeclaration(ctx *FunctionLambdaDeclarationContext) {}

// EnterResultType is called when production resultType is entered.
func (s *BasepascalListener) EnterResultType(ctx *ResultTypeContext) {}

// ExitResultType is called when production resultType is exited.
func (s *BasepascalListener) ExitResultType(ctx *ResultTypeContext) {}

// EnterProcedureOrFunctionBody is called when production procedureOrFunctionBody is entered.
func (s *BasepascalListener) EnterProcedureOrFunctionBody(ctx *ProcedureOrFunctionBodyContext) {}

// ExitProcedureOrFunctionBody is called when production procedureOrFunctionBody is exited.
func (s *BasepascalListener) ExitProcedureOrFunctionBody(ctx *ProcedureOrFunctionBodyContext) {}

// EnterClassOperatorHeader is called when production classOperatorHeader is entered.
func (s *BasepascalListener) EnterClassOperatorHeader(ctx *ClassOperatorHeaderContext) {}

// ExitClassOperatorHeader is called when production classOperatorHeader is exited.
func (s *BasepascalListener) ExitClassOperatorHeader(ctx *ClassOperatorHeaderContext) {}

// EnterClassOperatorDeclaration is called when production classOperatorDeclaration is entered.
func (s *BasepascalListener) EnterClassOperatorDeclaration(ctx *ClassOperatorDeclarationContext) {}

// ExitClassOperatorDeclaration is called when production classOperatorDeclaration is exited.
func (s *BasepascalListener) ExitClassOperatorDeclaration(ctx *ClassOperatorDeclarationContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BasepascalListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BasepascalListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterSection is called when production formalParameterSection is entered.
func (s *BasepascalListener) EnterFormalParameterSection(ctx *FormalParameterSectionContext) {}

// ExitFormalParameterSection is called when production formalParameterSection is exited.
func (s *BasepascalListener) ExitFormalParameterSection(ctx *FormalParameterSectionContext) {}

// EnterParameterGroup is called when production parameterGroup is entered.
func (s *BasepascalListener) EnterParameterGroup(ctx *ParameterGroupContext) {}

// ExitParameterGroup is called when production parameterGroup is exited.
func (s *BasepascalListener) ExitParameterGroup(ctx *ParameterGroupContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasepascalListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasepascalListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterQualifiedIdentifierList is called when production qualifiedIdentifierList is entered.
func (s *BasepascalListener) EnterQualifiedIdentifierList(ctx *QualifiedIdentifierListContext) {}

// ExitQualifiedIdentifierList is called when production qualifiedIdentifierList is exited.
func (s *BasepascalListener) ExitQualifiedIdentifierList(ctx *QualifiedIdentifierListContext) {}

// EnterConstList is called when production constList is entered.
func (s *BasepascalListener) EnterConstList(ctx *ConstListContext) {}

// ExitConstList is called when production constList is exited.
func (s *BasepascalListener) ExitConstList(ctx *ConstListContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BasepascalListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BasepascalListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterTypedIdentifierList is called when production typedIdentifierList is entered.
func (s *BasepascalListener) EnterTypedIdentifierList(ctx *TypedIdentifierListContext) {}

// ExitTypedIdentifierList is called when production typedIdentifierList is exited.
func (s *BasepascalListener) ExitTypedIdentifierList(ctx *TypedIdentifierListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasepascalListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasepascalListener) ExitStatement(ctx *StatementContext) {}

// EnterUnlabelledStatement is called when production unlabelledStatement is entered.
func (s *BasepascalListener) EnterUnlabelledStatement(ctx *UnlabelledStatementContext) {}

// ExitUnlabelledStatement is called when production unlabelledStatement is exited.
func (s *BasepascalListener) ExitUnlabelledStatement(ctx *UnlabelledStatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BasepascalListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BasepascalListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BasepascalListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BasepascalListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterRaiseExceptionStatement is called when production raiseExceptionStatement is entered.
func (s *BasepascalListener) EnterRaiseExceptionStatement(ctx *RaiseExceptionStatementContext) {}

// ExitRaiseExceptionStatement is called when production raiseExceptionStatement is exited.
func (s *BasepascalListener) ExitRaiseExceptionStatement(ctx *RaiseExceptionStatementContext) {}

// EnterVariableDeclarationStatement is called when production variableDeclarationStatement is entered.
func (s *BasepascalListener) EnterVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// ExitVariableDeclarationStatement is called when production variableDeclarationStatement is exited.
func (s *BasepascalListener) ExitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// EnterTypeCast is called when production typeCast is entered.
func (s *BasepascalListener) EnterTypeCast(ctx *TypeCastContext) {}

// ExitTypeCast is called when production typeCast is exited.
func (s *BasepascalListener) ExitTypeCast(ctx *TypeCastContext) {}

// EnterPropertyDesignator is called when production propertyDesignator is entered.
func (s *BasepascalListener) EnterPropertyDesignator(ctx *PropertyDesignatorContext) {}

// ExitPropertyDesignator is called when production propertyDesignator is exited.
func (s *BasepascalListener) ExitPropertyDesignator(ctx *PropertyDesignatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasepascalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasepascalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelationaloperator is called when production relationaloperator is entered.
func (s *BasepascalListener) EnterRelationaloperator(ctx *RelationaloperatorContext) {}

// ExitRelationaloperator is called when production relationaloperator is exited.
func (s *BasepascalListener) ExitRelationaloperator(ctx *RelationaloperatorContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BasepascalListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BasepascalListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterAdditiveoperator is called when production additiveoperator is entered.
func (s *BasepascalListener) EnterAdditiveoperator(ctx *AdditiveoperatorContext) {}

// ExitAdditiveoperator is called when production additiveoperator is exited.
func (s *BasepascalListener) ExitAdditiveoperator(ctx *AdditiveoperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BasepascalListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasepascalListener) ExitTerm(ctx *TermContext) {}

// EnterMultiplicativeoperator is called when production multiplicativeoperator is entered.
func (s *BasepascalListener) EnterMultiplicativeoperator(ctx *MultiplicativeoperatorContext) {}

// ExitMultiplicativeoperator is called when production multiplicativeoperator is exited.
func (s *BasepascalListener) ExitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) {}

// EnterSignedFactor is called when production signedFactor is entered.
func (s *BasepascalListener) EnterSignedFactor(ctx *SignedFactorContext) {}

// ExitSignedFactor is called when production signedFactor is exited.
func (s *BasepascalListener) ExitSignedFactor(ctx *SignedFactorContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasepascalListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasepascalListener) ExitFactor(ctx *FactorContext) {}

// EnterPostfixOp is called when production postfixOp is entered.
func (s *BasepascalListener) EnterPostfixOp(ctx *PostfixOpContext) {}

// ExitPostfixOp is called when production postfixOp is exited.
func (s *BasepascalListener) ExitPostfixOp(ctx *PostfixOpContext) {}

// EnterUnsignedConstant is called when production unsignedConstant is entered.
func (s *BasepascalListener) EnterUnsignedConstant(ctx *UnsignedConstantContext) {}

// ExitUnsignedConstant is called when production unsignedConstant is exited.
func (s *BasepascalListener) ExitUnsignedConstant(ctx *UnsignedConstantContext) {}

// EnterFunctionDesignator is called when production functionDesignator is entered.
func (s *BasepascalListener) EnterFunctionDesignator(ctx *FunctionDesignatorContext) {}

// ExitFunctionDesignator is called when production functionDesignator is exited.
func (s *BasepascalListener) ExitFunctionDesignator(ctx *FunctionDesignatorContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasepascalListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasepascalListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterSet_ is called when production set_ is entered.
func (s *BasepascalListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BasepascalListener) ExitSet_(ctx *Set_Context) {}

// EnterElementList is called when production elementList is entered.
func (s *BasepascalListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BasepascalListener) ExitElementList(ctx *ElementListContext) {}

// EnterElement is called when production element is entered.
func (s *BasepascalListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasepascalListener) ExitElement(ctx *ElementContext) {}

// EnterActualParameter is called when production actualParameter is entered.
func (s *BasepascalListener) EnterActualParameter(ctx *ActualParameterContext) {}

// ExitActualParameter is called when production actualParameter is exited.
func (s *BasepascalListener) ExitActualParameter(ctx *ActualParameterContext) {}

// EnterParameterwidth is called when production parameterwidth is entered.
func (s *BasepascalListener) EnterParameterwidth(ctx *ParameterwidthContext) {}

// ExitParameterwidth is called when production parameterwidth is exited.
func (s *BasepascalListener) ExitParameterwidth(ctx *ParameterwidthContext) {}

// EnterGotoStatement is called when production gotoStatement is entered.
func (s *BasepascalListener) EnterGotoStatement(ctx *GotoStatementContext) {}

// ExitGotoStatement is called when production gotoStatement is exited.
func (s *BasepascalListener) ExitGotoStatement(ctx *GotoStatementContext) {}

// EnterInheritedStatement is called when production inheritedStatement is entered.
func (s *BasepascalListener) EnterInheritedStatement(ctx *InheritedStatementContext) {}

// ExitInheritedStatement is called when production inheritedStatement is exited.
func (s *BasepascalListener) ExitInheritedStatement(ctx *InheritedStatementContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BasepascalListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BasepascalListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterEmpty_ is called when production empty_ is entered.
func (s *BasepascalListener) EnterEmpty_(ctx *Empty_Context) {}

// ExitEmpty_ is called when production empty_ is exited.
func (s *BasepascalListener) ExitEmpty_(ctx *Empty_Context) {}

// EnterStructuredStatement is called when production structuredStatement is entered.
func (s *BasepascalListener) EnterStructuredStatement(ctx *StructuredStatementContext) {}

// ExitStructuredStatement is called when production structuredStatement is exited.
func (s *BasepascalListener) ExitStructuredStatement(ctx *StructuredStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BasepascalListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BasepascalListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterStatements is called when production statements is entered.
func (s *BasepascalListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BasepascalListener) ExitStatements(ctx *StatementsContext) {}

// EnterConditionalStatement is called when production conditionalStatement is entered.
func (s *BasepascalListener) EnterConditionalStatement(ctx *ConditionalStatementContext) {}

// ExitConditionalStatement is called when production conditionalStatement is exited.
func (s *BasepascalListener) ExitConditionalStatement(ctx *ConditionalStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasepascalListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasepascalListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BasepascalListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BasepascalListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterCaseConstRange is called when production caseConstRange is entered.
func (s *BasepascalListener) EnterCaseConstRange(ctx *CaseConstRangeContext) {}

// ExitCaseConstRange is called when production caseConstRange is exited.
func (s *BasepascalListener) ExitCaseConstRange(ctx *CaseConstRangeContext) {}

// EnterCaseConstList is called when production caseConstList is entered.
func (s *BasepascalListener) EnterCaseConstList(ctx *CaseConstListContext) {}

// ExitCaseConstList is called when production caseConstList is exited.
func (s *BasepascalListener) ExitCaseConstList(ctx *CaseConstListContext) {}

// EnterCaseListElement is called when production caseListElement is entered.
func (s *BasepascalListener) EnterCaseListElement(ctx *CaseListElementContext) {}

// ExitCaseListElement is called when production caseListElement is exited.
func (s *BasepascalListener) ExitCaseListElement(ctx *CaseListElementContext) {}

// EnterRepetetiveStatement is called when production repetetiveStatement is entered.
func (s *BasepascalListener) EnterRepetetiveStatement(ctx *RepetetiveStatementContext) {}

// ExitRepetetiveStatement is called when production repetetiveStatement is exited.
func (s *BasepascalListener) ExitRepetetiveStatement(ctx *RepetetiveStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BasepascalListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BasepascalListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BasepascalListener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BasepascalListener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BasepascalListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BasepascalListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForList is called when production forList is entered.
func (s *BasepascalListener) EnterForList(ctx *ForListContext) {}

// ExitForList is called when production forList is exited.
func (s *BasepascalListener) ExitForList(ctx *ForListContext) {}

// EnterInitialValue is called when production initialValue is entered.
func (s *BasepascalListener) EnterInitialValue(ctx *InitialValueContext) {}

// ExitInitialValue is called when production initialValue is exited.
func (s *BasepascalListener) ExitInitialValue(ctx *InitialValueContext) {}

// EnterFinalValue is called when production finalValue is entered.
func (s *BasepascalListener) EnterFinalValue(ctx *FinalValueContext) {}

// ExitFinalValue is called when production finalValue is exited.
func (s *BasepascalListener) ExitFinalValue(ctx *FinalValueContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BasepascalListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BasepascalListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterTryExceptStatement is called when production tryExceptStatement is entered.
func (s *BasepascalListener) EnterTryExceptStatement(ctx *TryExceptStatementContext) {}

// ExitTryExceptStatement is called when production tryExceptStatement is exited.
func (s *BasepascalListener) ExitTryExceptStatement(ctx *TryExceptStatementContext) {}

// EnterExceptionCase is called when production exceptionCase is entered.
func (s *BasepascalListener) EnterExceptionCase(ctx *ExceptionCaseContext) {}

// ExitExceptionCase is called when production exceptionCase is exited.
func (s *BasepascalListener) ExitExceptionCase(ctx *ExceptionCaseContext) {}

// EnterExceptionElse is called when production exceptionElse is entered.
func (s *BasepascalListener) EnterExceptionElse(ctx *ExceptionElseContext) {}

// ExitExceptionElse is called when production exceptionElse is exited.
func (s *BasepascalListener) ExitExceptionElse(ctx *ExceptionElseContext) {}

// EnterTryFinallyStatement is called when production tryFinallyStatement is entered.
func (s *BasepascalListener) EnterTryFinallyStatement(ctx *TryFinallyStatementContext) {}

// ExitTryFinallyStatement is called when production tryFinallyStatement is exited.
func (s *BasepascalListener) ExitTryFinallyStatement(ctx *TryFinallyStatementContext) {}

// EnterWithStatementVariableList is called when production withStatementVariableList is entered.
func (s *BasepascalListener) EnterWithStatementVariableList(ctx *WithStatementVariableListContext) {}

// ExitWithStatementVariableList is called when production withStatementVariableList is exited.
func (s *BasepascalListener) ExitWithStatementVariableList(ctx *WithStatementVariableListContext) {}

// EnterAttributeSection is called when production attributeSection is entered.
func (s *BasepascalListener) EnterAttributeSection(ctx *AttributeSectionContext) {}

// ExitAttributeSection is called when production attributeSection is exited.
func (s *BasepascalListener) ExitAttributeSection(ctx *AttributeSectionContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *BasepascalListener) EnterAttributeList(ctx *AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *BasepascalListener) ExitAttributeList(ctx *AttributeListContext) {}

// EnterAttributeItem is called when production attributeItem is entered.
func (s *BasepascalListener) EnterAttributeItem(ctx *AttributeItemContext) {}

// ExitAttributeItem is called when production attributeItem is exited.
func (s *BasepascalListener) ExitAttributeItem(ctx *AttributeItemContext) {}
