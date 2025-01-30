// Code generated from pascal.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

// EnterTopLevelDeclarations is called when production topLevelDeclarations is entered.
func (s *BasepascalListener) EnterTopLevelDeclarations(ctx *TopLevelDeclarationsContext) {}

// ExitTopLevelDeclarations is called when production topLevelDeclarations is exited.
func (s *BasepascalListener) ExitTopLevelDeclarations(ctx *TopLevelDeclarationsContext) {}

// EnterBlock is called when production block is entered.
func (s *BasepascalListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasepascalListener) ExitBlock(ctx *BlockContext) {}

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

// EnterConstant is called when production constant is entered.
func (s *BasepascalListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasepascalListener) ExitConstant(ctx *ConstantContext) {}

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

// EnterTypeDefinitionPart is called when production typeDefinitionPart is entered.
func (s *BasepascalListener) EnterTypeDefinitionPart(ctx *TypeDefinitionPartContext) {}

// ExitTypeDefinitionPart is called when production typeDefinitionPart is exited.
func (s *BasepascalListener) ExitTypeDefinitionPart(ctx *TypeDefinitionPartContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BasepascalListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BasepascalListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BasepascalListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BasepascalListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterProcedureType is called when production procedureType is entered.
func (s *BasepascalListener) EnterProcedureType(ctx *ProcedureTypeContext) {}

// ExitProcedureType is called when production procedureType is exited.
func (s *BasepascalListener) ExitProcedureType(ctx *ProcedureTypeContext) {}

// EnterClassType is called when production classType is entered.
func (s *BasepascalListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BasepascalListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterClassImplementsInterfaces is called when production classImplementsInterfaces is entered.
func (s *BasepascalListener) EnterClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) {}

// ExitClassImplementsInterfaces is called when production classImplementsInterfaces is exited.
func (s *BasepascalListener) ExitClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasepascalListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasepascalListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassPrivateDeclaration is called when production classPrivateDeclaration is entered.
func (s *BasepascalListener) EnterClassPrivateDeclaration(ctx *ClassPrivateDeclarationContext) {}

// ExitClassPrivateDeclaration is called when production classPrivateDeclaration is exited.
func (s *BasepascalListener) ExitClassPrivateDeclaration(ctx *ClassPrivateDeclarationContext) {}

// EnterClassStrictPrivateDeclaration is called when production classStrictPrivateDeclaration is entered.
func (s *BasepascalListener) EnterClassStrictPrivateDeclaration(ctx *ClassStrictPrivateDeclarationContext) {
}

// ExitClassStrictPrivateDeclaration is called when production classStrictPrivateDeclaration is exited.
func (s *BasepascalListener) ExitClassStrictPrivateDeclaration(ctx *ClassStrictPrivateDeclarationContext) {
}

// EnterClassProtectedDeclaration is called when production classProtectedDeclaration is entered.
func (s *BasepascalListener) EnterClassProtectedDeclaration(ctx *ClassProtectedDeclarationContext) {}

// ExitClassProtectedDeclaration is called when production classProtectedDeclaration is exited.
func (s *BasepascalListener) ExitClassProtectedDeclaration(ctx *ClassProtectedDeclarationContext) {}

// EnterClassStrictProtectedDeclaration is called when production classStrictProtectedDeclaration is entered.
func (s *BasepascalListener) EnterClassStrictProtectedDeclaration(ctx *ClassStrictProtectedDeclarationContext) {
}

// ExitClassStrictProtectedDeclaration is called when production classStrictProtectedDeclaration is exited.
func (s *BasepascalListener) ExitClassStrictProtectedDeclaration(ctx *ClassStrictProtectedDeclarationContext) {
}

// EnterClassPublicDeclaration is called when production classPublicDeclaration is entered.
func (s *BasepascalListener) EnterClassPublicDeclaration(ctx *ClassPublicDeclarationContext) {}

// ExitClassPublicDeclaration is called when production classPublicDeclaration is exited.
func (s *BasepascalListener) ExitClassPublicDeclaration(ctx *ClassPublicDeclarationContext) {}

// EnterClassPublishedDeclaration is called when production classPublishedDeclaration is entered.
func (s *BasepascalListener) EnterClassPublishedDeclaration(ctx *ClassPublishedDeclarationContext) {}

// ExitClassPublishedDeclaration is called when production classPublishedDeclaration is exited.
func (s *BasepascalListener) ExitClassPublishedDeclaration(ctx *ClassPublishedDeclarationContext) {}

// EnterClassImplicitPublishedDeclaration is called when production classImplicitPublishedDeclaration is entered.
func (s *BasepascalListener) EnterClassImplicitPublishedDeclaration(ctx *ClassImplicitPublishedDeclarationContext) {
}

// ExitClassImplicitPublishedDeclaration is called when production classImplicitPublishedDeclaration is exited.
func (s *BasepascalListener) ExitClassImplicitPublishedDeclaration(ctx *ClassImplicitPublishedDeclarationContext) {
}

// EnterClassDeclarationPart is called when production classDeclarationPart is entered.
func (s *BasepascalListener) EnterClassDeclarationPart(ctx *ClassDeclarationPartContext) {}

// ExitClassDeclarationPart is called when production classDeclarationPart is exited.
func (s *BasepascalListener) ExitClassDeclarationPart(ctx *ClassDeclarationPartContext) {}

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

// EnterMethodIdentifier is called when production methodIdentifier is entered.
func (s *BasepascalListener) EnterMethodIdentifier(ctx *MethodIdentifierContext) {}

// ExitMethodIdentifier is called when production methodIdentifier is exited.
func (s *BasepascalListener) ExitMethodIdentifier(ctx *MethodIdentifierContext) {}

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

// EnterSubrangeType is called when production subrangeType is entered.
func (s *BasepascalListener) EnterSubrangeType(ctx *SubrangeTypeContext) {}

// ExitSubrangeType is called when production subrangeType is exited.
func (s *BasepascalListener) ExitSubrangeType(ctx *SubrangeTypeContext) {}

// EnterTypeIdentifier is called when production typeIdentifier is entered.
func (s *BasepascalListener) EnterTypeIdentifier(ctx *TypeIdentifierContext) {}

// ExitTypeIdentifier is called when production typeIdentifier is exited.
func (s *BasepascalListener) ExitTypeIdentifier(ctx *TypeIdentifierContext) {}

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

// EnterComponentType is called when production componentType is entered.
func (s *BasepascalListener) EnterComponentType(ctx *ComponentTypeContext) {}

// ExitComponentType is called when production componentType is exited.
func (s *BasepascalListener) ExitComponentType(ctx *ComponentTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BasepascalListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BasepascalListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BasepascalListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BasepascalListener) ExitFieldList(ctx *FieldListContext) {}

// EnterFixedPart is called when production fixedPart is entered.
func (s *BasepascalListener) EnterFixedPart(ctx *FixedPartContext) {}

// ExitFixedPart is called when production fixedPart is exited.
func (s *BasepascalListener) ExitFixedPart(ctx *FixedPartContext) {}

// EnterRecordSection is called when production recordSection is entered.
func (s *BasepascalListener) EnterRecordSection(ctx *RecordSectionContext) {}

// ExitRecordSection is called when production recordSection is exited.
func (s *BasepascalListener) ExitRecordSection(ctx *RecordSectionContext) {}

// EnterVariantPart is called when production variantPart is entered.
func (s *BasepascalListener) EnterVariantPart(ctx *VariantPartContext) {}

// ExitVariantPart is called when production variantPart is exited.
func (s *BasepascalListener) ExitVariantPart(ctx *VariantPartContext) {}

// EnterTag is called when production tag is entered.
func (s *BasepascalListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasepascalListener) ExitTag(ctx *TagContext) {}

// EnterVariant is called when production variant is entered.
func (s *BasepascalListener) EnterVariant(ctx *VariantContext) {}

// ExitVariant is called when production variant is exited.
func (s *BasepascalListener) ExitVariant(ctx *VariantContext) {}

// EnterSetType is called when production setType is entered.
func (s *BasepascalListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BasepascalListener) ExitSetType(ctx *SetTypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BasepascalListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BasepascalListener) ExitBaseType(ctx *BaseTypeContext) {}

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

// EnterProcedureAndFunctionDeclarationPart is called when production procedureAndFunctionDeclarationPart is entered.
func (s *BasepascalListener) EnterProcedureAndFunctionDeclarationPart(ctx *ProcedureAndFunctionDeclarationPartContext) {
}

// ExitProcedureAndFunctionDeclarationPart is called when production procedureAndFunctionDeclarationPart is exited.
func (s *BasepascalListener) ExitProcedureAndFunctionDeclarationPart(ctx *ProcedureAndFunctionDeclarationPartContext) {
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

// EnterProcedureMethodDeclaration is called when production procedureMethodDeclaration is entered.
func (s *BasepascalListener) EnterProcedureMethodDeclaration(ctx *ProcedureMethodDeclarationContext) {
}

// ExitProcedureMethodDeclaration is called when production procedureMethodDeclaration is exited.
func (s *BasepascalListener) ExitProcedureMethodDeclaration(ctx *ProcedureMethodDeclarationContext) {}

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

// EnterConstList is called when production constList is entered.
func (s *BasepascalListener) EnterConstList(ctx *ConstListContext) {}

// ExitConstList is called when production constList is exited.
func (s *BasepascalListener) ExitConstList(ctx *ConstListContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasepascalListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasepascalListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterResultType is called when production resultType is entered.
func (s *BasepascalListener) EnterResultType(ctx *ResultTypeContext) {}

// ExitResultType is called when production resultType is exited.
func (s *BasepascalListener) ExitResultType(ctx *ResultTypeContext) {}

// EnterFunctionMethodDeclaration is called when production functionMethodDeclaration is entered.
func (s *BasepascalListener) EnterFunctionMethodDeclaration(ctx *FunctionMethodDeclarationContext) {}

// ExitFunctionMethodDeclaration is called when production functionMethodDeclaration is exited.
func (s *BasepascalListener) ExitFunctionMethodDeclaration(ctx *FunctionMethodDeclarationContext) {}

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

// EnterVariable is called when production variable is entered.
func (s *BasepascalListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasepascalListener) ExitVariable(ctx *VariableContext) {}

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

// EnterProcedureStatement is called when production procedureStatement is entered.
func (s *BasepascalListener) EnterProcedureStatement(ctx *ProcedureStatementContext) {}

// ExitProcedureStatement is called when production procedureStatement is exited.
func (s *BasepascalListener) ExitProcedureStatement(ctx *ProcedureStatementContext) {}

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

// EnterTryFinallyStatement is called when production tryFinallyStatement is entered.
func (s *BasepascalListener) EnterTryFinallyStatement(ctx *TryFinallyStatementContext) {}

// ExitTryFinallyStatement is called when production tryFinallyStatement is exited.
func (s *BasepascalListener) ExitTryFinallyStatement(ctx *TryFinallyStatementContext) {}

// EnterRecordVariableList is called when production recordVariableList is entered.
func (s *BasepascalListener) EnterRecordVariableList(ctx *RecordVariableListContext) {}

// ExitRecordVariableList is called when production recordVariableList is exited.
func (s *BasepascalListener) ExitRecordVariableList(ctx *RecordVariableListContext) {}
