// Code generated from /home/tomas/development/tomas303/projects/pascallsp/palsp/internal/pascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // pascal

import "github.com/antlr4-go/antlr/v4"

// pascalListener is a complete listener for a parse tree produced by pascalParser.
type pascalListener interface {
	antlr.ParseTreeListener

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterInterfaceSection is called when entering the interfaceSection production.
	EnterInterfaceSection(c *InterfaceSectionContext)

	// EnterImplementationSection is called when entering the implementationSection production.
	EnterImplementationSection(c *ImplementationSectionContext)

	// EnterInitializationSection is called when entering the initializationSection production.
	EnterInitializationSection(c *InitializationSectionContext)

	// EnterFinalizationSection is called when entering the finalizationSection production.
	EnterFinalizationSection(c *FinalizationSectionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// EnterIdentifierPart is called when entering the identifierPart production.
	EnterIdentifierPart(c *IdentifierPartContext)

	// EnterInterfaceBlockMember is called when entering the interfaceBlockMember production.
	EnterInterfaceBlockMember(c *InterfaceBlockMemberContext)

	// EnterInterfaceBlock is called when entering the interfaceBlock production.
	EnterInterfaceBlock(c *InterfaceBlockContext)

	// EnterImplementationBlockMember is called when entering the implementationBlockMember production.
	EnterImplementationBlockMember(c *ImplementationBlockMemberContext)

	// EnterImplementationBlock is called when entering the implementationBlock production.
	EnterImplementationBlock(c *ImplementationBlockContext)

	// EnterFuncBlockMemeber is called when entering the funcBlockMemeber production.
	EnterFuncBlockMemeber(c *FuncBlockMemeberContext)

	// EnterFuncBlock is called when entering the funcBlock production.
	EnterFuncBlock(c *FuncBlockContext)

	// EnterUsesUnits is called when entering the usesUnits production.
	EnterUsesUnits(c *UsesUnitsContext)

	// EnterLabelDeclarationPart is called when entering the labelDeclarationPart production.
	EnterLabelDeclarationPart(c *LabelDeclarationPartContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterConstantDefinitionPart is called when entering the constantDefinitionPart production.
	EnterConstantDefinitionPart(c *ConstantDefinitionPartContext)

	// EnterConstantDefinition is called when entering the constantDefinition production.
	EnterConstantDefinition(c *ConstantDefinitionContext)

	// EnterConstantChr is called when entering the constantChr production.
	EnterConstantChr(c *ConstantChrContext)

	// EnterHexConstant is called when entering the hexConstant production.
	EnterHexConstant(c *HexConstantContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterArrayConstant is called when entering the arrayConstant production.
	EnterArrayConstant(c *ArrayConstantContext)

	// EnterRecordConstant is called when entering the recordConstant production.
	EnterRecordConstant(c *RecordConstantContext)

	// EnterRecordField is called when entering the recordField production.
	EnterRecordField(c *RecordFieldContext)

	// EnterUnsignedNumber is called when entering the unsignedNumber production.
	EnterUnsignedNumber(c *UnsignedNumberContext)

	// EnterUnsignedInteger is called when entering the unsignedInteger production.
	EnterUnsignedInteger(c *UnsignedIntegerContext)

	// EnterUnsignedReal is called when entering the unsignedReal production.
	EnterUnsignedReal(c *UnsignedRealContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterBool_ is called when entering the bool_ production.
	EnterBool_(c *Bool_Context)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterStringExpression is called when entering the stringExpression production.
	EnterStringExpression(c *StringExpressionContext)

	// EnterResourceDefinitionPart is called when entering the resourceDefinitionPart production.
	EnterResourceDefinitionPart(c *ResourceDefinitionPartContext)

	// EnterResourceDefinition is called when entering the resourceDefinition production.
	EnterResourceDefinition(c *ResourceDefinitionContext)

	// EnterDeprecatedHint is called when entering the deprecatedHint production.
	EnterDeprecatedHint(c *DeprecatedHintContext)

	// EnterPlatformHint is called when entering the platformHint production.
	EnterPlatformHint(c *PlatformHintContext)

	// EnterAlignHint is called when entering the alignHint production.
	EnterAlignHint(c *AlignHintContext)

	// EnterTypeDefinitionPart is called when entering the typeDefinitionPart production.
	EnterTypeDefinitionPart(c *TypeDefinitionPartContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterClassSection is called when entering the classSection production.
	EnterClassSection(c *ClassSectionContext)

	// EnterClassTypeBlock is called when entering the classTypeBlock production.
	EnterClassTypeBlock(c *ClassTypeBlockContext)

	// EnterInterfaceType is called when entering the interfaceType production.
	EnterInterfaceType(c *InterfaceTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterProcedureType is called when entering the procedureType production.
	EnterProcedureType(c *ProcedureTypeContext)

	// EnterAliasDistinctType is called when entering the aliasDistinctType production.
	EnterAliasDistinctType(c *AliasDistinctTypeContext)

	// EnterClassImplementsInterfaces is called when entering the classImplementsInterfaces production.
	EnterClassImplementsInterfaces(c *ClassImplementsInterfacesContext)

	// EnterAccessSpecifier is called when entering the accessSpecifier production.
	EnterAccessSpecifier(c *AccessSpecifierContext)

	// EnterClassDeclarationPart is called when entering the classDeclarationPart production.
	EnterClassDeclarationPart(c *ClassDeclarationPartContext)

	// EnterInterfaceGuidConst is called when entering the interfaceGuidConst production.
	EnterInterfaceGuidConst(c *InterfaceGuidConstContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterInterfaceDeclarationPart is called when entering the interfaceDeclarationPart production.
	EnterInterfaceDeclarationPart(c *InterfaceDeclarationPartContext)

	// EnterErrorInterfaceDeclarationPart is called when entering the errorInterfaceDeclarationPart production.
	EnterErrorInterfaceDeclarationPart(c *ErrorInterfaceDeclarationPartContext)

	// EnterErrorClassDeclarationPart is called when entering the errorClassDeclarationPart production.
	EnterErrorClassDeclarationPart(c *ErrorClassDeclarationPartContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyReadDeclaration is called when entering the propertyReadDeclaration production.
	EnterPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// EnterPropertyWriteDeclaration is called when entering the propertyWriteDeclaration production.
	EnterPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// EnterPropertyDefaultValueDeclaration is called when entering the propertyDefaultValueDeclaration production.
	EnterPropertyDefaultValueDeclaration(c *PropertyDefaultValueDeclarationContext)

	// EnterPropertyIndexDeclaration is called when entering the propertyIndexDeclaration production.
	EnterPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// EnterPropertyIndexParameters is called when entering the propertyIndexParameters production.
	EnterPropertyIndexParameters(c *PropertyIndexParametersContext)

	// EnterPropertyIndexParametersList is called when entering the propertyIndexParametersList production.
	EnterPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// EnterGenericTemplate is called when entering the genericTemplate production.
	EnterGenericTemplate(c *GenericTemplateContext)

	// EnterGenericTemplateList is called when entering the genericTemplateList production.
	EnterGenericTemplateList(c *GenericTemplateListContext)

	// EnterGenericTypeParameter is called when entering the genericTypeParameter production.
	EnterGenericTypeParameter(c *GenericTypeParameterContext)

	// EnterGenericConstraints is called when entering the genericConstraints production.
	EnterGenericConstraints(c *GenericConstraintsContext)

	// EnterGenericConstraint is called when entering the genericConstraint production.
	EnterGenericConstraint(c *GenericConstraintContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterSimpleType is called when entering the simpleType production.
	EnterSimpleType(c *SimpleTypeContext)

	// EnterScalarType is called when entering the scalarType production.
	EnterScalarType(c *ScalarTypeContext)

	// EnterScalerList is called when entering the scalerList production.
	EnterScalerList(c *ScalerListContext)

	// EnterScalerMember is called when entering the scalerMember production.
	EnterScalerMember(c *ScalerMemberContext)

	// EnterSubrangeType is called when entering the subrangeType production.
	EnterSubrangeType(c *SubrangeTypeContext)

	// EnterStringTypeIdentifier is called when entering the stringTypeIdentifier production.
	EnterStringTypeIdentifier(c *StringTypeIdentifierContext)

	// EnterStructuredType is called when entering the structuredType production.
	EnterStructuredType(c *StructuredTypeContext)

	// EnterUnpackedStructuredType is called when entering the unpackedStructuredType production.
	EnterUnpackedStructuredType(c *UnpackedStructuredTypeContext)

	// EnterStringtype is called when entering the stringtype production.
	EnterStringtype(c *StringtypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterRecordContent is called when entering the recordContent production.
	EnterRecordContent(c *RecordContentContext)

	// EnterRecordSection is called when entering the recordSection production.
	EnterRecordSection(c *RecordSectionContext)

	// EnterRecordFieldsSection is called when entering the recordFieldsSection production.
	EnterRecordFieldsSection(c *RecordFieldsSectionContext)

	// EnterRecordVariantSection is called when entering the recordVariantSection production.
	EnterRecordVariantSection(c *RecordVariantSectionContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterRecordVariant is called when entering the recordVariant production.
	EnterRecordVariant(c *RecordVariantContext)

	// EnterHelperType is called when entering the helperType production.
	EnterHelperType(c *HelperTypeContext)

	// EnterHelperDeclarationPart is called when entering the helperDeclarationPart production.
	EnterHelperDeclarationPart(c *HelperDeclarationPartContext)

	// EnterErrorHelperDeclarationPart is called when entering the errorHelperDeclarationPart production.
	EnterErrorHelperDeclarationPart(c *ErrorHelperDeclarationPartContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterFileType is called when entering the fileType production.
	EnterFileType(c *FileTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterVariableDeclarationPart is called when entering the variableDeclarationPart production.
	EnterVariableDeclarationPart(c *VariableDeclarationPartContext)

	// EnterThreadvarDeclarationPart is called when entering the threadvarDeclarationPart production.
	EnterThreadvarDeclarationPart(c *ThreadvarDeclarationPartContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterProcedureHeader is called when entering the procedureHeader production.
	EnterProcedureHeader(c *ProcedureHeaderContext)

	// EnterFunctionHeader is called when entering the functionHeader production.
	EnterFunctionHeader(c *FunctionHeaderContext)

	// EnterProcedureOrFunctionHeader is called when entering the procedureOrFunctionHeader production.
	EnterProcedureOrFunctionHeader(c *ProcedureOrFunctionHeaderContext)

	// EnterProcedureOrFunctionHeaderModifiers is called when entering the procedureOrFunctionHeaderModifiers production.
	EnterProcedureOrFunctionHeaderModifiers(c *ProcedureOrFunctionHeaderModifiersContext)

	// EnterProcedureOrFunctionDeclaration is called when entering the procedureOrFunctionDeclaration production.
	EnterProcedureOrFunctionDeclaration(c *ProcedureOrFunctionDeclarationContext)

	// EnterProcedureDeclaration is called when entering the procedureDeclaration production.
	EnterProcedureDeclaration(c *ProcedureDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterProcedureLambdaDeclaration is called when entering the procedureLambdaDeclaration production.
	EnterProcedureLambdaDeclaration(c *ProcedureLambdaDeclarationContext)

	// EnterFunctionLambdaDeclaration is called when entering the functionLambdaDeclaration production.
	EnterFunctionLambdaDeclaration(c *FunctionLambdaDeclarationContext)

	// EnterResultType is called when entering the resultType production.
	EnterResultType(c *ResultTypeContext)

	// EnterProcedureOrFunctionBody is called when entering the procedureOrFunctionBody production.
	EnterProcedureOrFunctionBody(c *ProcedureOrFunctionBodyContext)

	// EnterClassOperatorHeader is called when entering the classOperatorHeader production.
	EnterClassOperatorHeader(c *ClassOperatorHeaderContext)

	// EnterClassOperatorDeclaration is called when entering the classOperatorDeclaration production.
	EnterClassOperatorDeclaration(c *ClassOperatorDeclarationContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameterSection is called when entering the formalParameterSection production.
	EnterFormalParameterSection(c *FormalParameterSectionContext)

	// EnterParameterGroup is called when entering the parameterGroup production.
	EnterParameterGroup(c *ParameterGroupContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterQualifiedIdentifierList is called when entering the qualifiedIdentifierList production.
	EnterQualifiedIdentifierList(c *QualifiedIdentifierListContext)

	// EnterConstList is called when entering the constList production.
	EnterConstList(c *ConstListContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterTypedIdentifierList is called when entering the typedIdentifierList production.
	EnterTypedIdentifierList(c *TypedIdentifierListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterUnlabelledStatement is called when entering the unlabelledStatement production.
	EnterUnlabelledStatement(c *UnlabelledStatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterRaiseExceptionStatement is called when entering the raiseExceptionStatement production.
	EnterRaiseExceptionStatement(c *RaiseExceptionStatementContext)

	// EnterVariableDeclarationStatement is called when entering the variableDeclarationStatement production.
	EnterVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// EnterTypeCast is called when entering the typeCast production.
	EnterTypeCast(c *TypeCastContext)

	// EnterPropertyDesignator is called when entering the propertyDesignator production.
	EnterPropertyDesignator(c *PropertyDesignatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelationaloperator is called when entering the relationaloperator production.
	EnterRelationaloperator(c *RelationaloperatorContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterAdditiveoperator is called when entering the additiveoperator production.
	EnterAdditiveoperator(c *AdditiveoperatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMultiplicativeoperator is called when entering the multiplicativeoperator production.
	EnterMultiplicativeoperator(c *MultiplicativeoperatorContext)

	// EnterSignedFactor is called when entering the signedFactor production.
	EnterSignedFactor(c *SignedFactorContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPostfixOp is called when entering the postfixOp production.
	EnterPostfixOp(c *PostfixOpContext)

	// EnterUnsignedConstant is called when entering the unsignedConstant production.
	EnterUnsignedConstant(c *UnsignedConstantContext)

	// EnterFunctionDesignator is called when entering the functionDesignator production.
	EnterFunctionDesignator(c *FunctionDesignatorContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterActualParameter is called when entering the actualParameter production.
	EnterActualParameter(c *ActualParameterContext)

	// EnterParameterwidth is called when entering the parameterwidth production.
	EnterParameterwidth(c *ParameterwidthContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

	// EnterInheritedStatement is called when entering the inheritedStatement production.
	EnterInheritedStatement(c *InheritedStatementContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterEmpty_ is called when entering the empty_ production.
	EnterEmpty_(c *Empty_Context)

	// EnterStructuredStatement is called when entering the structuredStatement production.
	EnterStructuredStatement(c *StructuredStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterConditionalStatement is called when entering the conditionalStatement production.
	EnterConditionalStatement(c *ConditionalStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterCaseConstRange is called when entering the caseConstRange production.
	EnterCaseConstRange(c *CaseConstRangeContext)

	// EnterCaseConstList is called when entering the caseConstList production.
	EnterCaseConstList(c *CaseConstListContext)

	// EnterCaseListElement is called when entering the caseListElement production.
	EnterCaseListElement(c *CaseListElementContext)

	// EnterRepetetiveStatement is called when entering the repetetiveStatement production.
	EnterRepetetiveStatement(c *RepetetiveStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterRepeatStatement is called when entering the repeatStatement production.
	EnterRepeatStatement(c *RepeatStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForList is called when entering the forList production.
	EnterForList(c *ForListContext)

	// EnterInitialValue is called when entering the initialValue production.
	EnterInitialValue(c *InitialValueContext)

	// EnterFinalValue is called when entering the finalValue production.
	EnterFinalValue(c *FinalValueContext)

	// EnterWithStatement is called when entering the withStatement production.
	EnterWithStatement(c *WithStatementContext)

	// EnterTryExceptStatement is called when entering the tryExceptStatement production.
	EnterTryExceptStatement(c *TryExceptStatementContext)

	// EnterExceptionCase is called when entering the exceptionCase production.
	EnterExceptionCase(c *ExceptionCaseContext)

	// EnterExceptionElse is called when entering the exceptionElse production.
	EnterExceptionElse(c *ExceptionElseContext)

	// EnterTryFinallyStatement is called when entering the tryFinallyStatement production.
	EnterTryFinallyStatement(c *TryFinallyStatementContext)

	// EnterWithStatementVariableList is called when entering the withStatementVariableList production.
	EnterWithStatementVariableList(c *WithStatementVariableListContext)

	// EnterAttributeSection is called when entering the attributeSection production.
	EnterAttributeSection(c *AttributeSectionContext)

	// EnterAttributeList is called when entering the attributeList production.
	EnterAttributeList(c *AttributeListContext)

	// EnterAttributeItem is called when entering the attributeItem production.
	EnterAttributeItem(c *AttributeItemContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitInterfaceSection is called when exiting the interfaceSection production.
	ExitInterfaceSection(c *InterfaceSectionContext)

	// ExitImplementationSection is called when exiting the implementationSection production.
	ExitImplementationSection(c *ImplementationSectionContext)

	// ExitInitializationSection is called when exiting the initializationSection production.
	ExitInitializationSection(c *InitializationSectionContext)

	// ExitFinalizationSection is called when exiting the finalizationSection production.
	ExitFinalizationSection(c *FinalizationSectionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitIdentifierPart is called when exiting the identifierPart production.
	ExitIdentifierPart(c *IdentifierPartContext)

	// ExitInterfaceBlockMember is called when exiting the interfaceBlockMember production.
	ExitInterfaceBlockMember(c *InterfaceBlockMemberContext)

	// ExitInterfaceBlock is called when exiting the interfaceBlock production.
	ExitInterfaceBlock(c *InterfaceBlockContext)

	// ExitImplementationBlockMember is called when exiting the implementationBlockMember production.
	ExitImplementationBlockMember(c *ImplementationBlockMemberContext)

	// ExitImplementationBlock is called when exiting the implementationBlock production.
	ExitImplementationBlock(c *ImplementationBlockContext)

	// ExitFuncBlockMemeber is called when exiting the funcBlockMemeber production.
	ExitFuncBlockMemeber(c *FuncBlockMemeberContext)

	// ExitFuncBlock is called when exiting the funcBlock production.
	ExitFuncBlock(c *FuncBlockContext)

	// ExitUsesUnits is called when exiting the usesUnits production.
	ExitUsesUnits(c *UsesUnitsContext)

	// ExitLabelDeclarationPart is called when exiting the labelDeclarationPart production.
	ExitLabelDeclarationPart(c *LabelDeclarationPartContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitConstantDefinitionPart is called when exiting the constantDefinitionPart production.
	ExitConstantDefinitionPart(c *ConstantDefinitionPartContext)

	// ExitConstantDefinition is called when exiting the constantDefinition production.
	ExitConstantDefinition(c *ConstantDefinitionContext)

	// ExitConstantChr is called when exiting the constantChr production.
	ExitConstantChr(c *ConstantChrContext)

	// ExitHexConstant is called when exiting the hexConstant production.
	ExitHexConstant(c *HexConstantContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitArrayConstant is called when exiting the arrayConstant production.
	ExitArrayConstant(c *ArrayConstantContext)

	// ExitRecordConstant is called when exiting the recordConstant production.
	ExitRecordConstant(c *RecordConstantContext)

	// ExitRecordField is called when exiting the recordField production.
	ExitRecordField(c *RecordFieldContext)

	// ExitUnsignedNumber is called when exiting the unsignedNumber production.
	ExitUnsignedNumber(c *UnsignedNumberContext)

	// ExitUnsignedInteger is called when exiting the unsignedInteger production.
	ExitUnsignedInteger(c *UnsignedIntegerContext)

	// ExitUnsignedReal is called when exiting the unsignedReal production.
	ExitUnsignedReal(c *UnsignedRealContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitBool_ is called when exiting the bool_ production.
	ExitBool_(c *Bool_Context)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitStringExpression is called when exiting the stringExpression production.
	ExitStringExpression(c *StringExpressionContext)

	// ExitResourceDefinitionPart is called when exiting the resourceDefinitionPart production.
	ExitResourceDefinitionPart(c *ResourceDefinitionPartContext)

	// ExitResourceDefinition is called when exiting the resourceDefinition production.
	ExitResourceDefinition(c *ResourceDefinitionContext)

	// ExitDeprecatedHint is called when exiting the deprecatedHint production.
	ExitDeprecatedHint(c *DeprecatedHintContext)

	// ExitPlatformHint is called when exiting the platformHint production.
	ExitPlatformHint(c *PlatformHintContext)

	// ExitAlignHint is called when exiting the alignHint production.
	ExitAlignHint(c *AlignHintContext)

	// ExitTypeDefinitionPart is called when exiting the typeDefinitionPart production.
	ExitTypeDefinitionPart(c *TypeDefinitionPartContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitClassSection is called when exiting the classSection production.
	ExitClassSection(c *ClassSectionContext)

	// ExitClassTypeBlock is called when exiting the classTypeBlock production.
	ExitClassTypeBlock(c *ClassTypeBlockContext)

	// ExitInterfaceType is called when exiting the interfaceType production.
	ExitInterfaceType(c *InterfaceTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitProcedureType is called when exiting the procedureType production.
	ExitProcedureType(c *ProcedureTypeContext)

	// ExitAliasDistinctType is called when exiting the aliasDistinctType production.
	ExitAliasDistinctType(c *AliasDistinctTypeContext)

	// ExitClassImplementsInterfaces is called when exiting the classImplementsInterfaces production.
	ExitClassImplementsInterfaces(c *ClassImplementsInterfacesContext)

	// ExitAccessSpecifier is called when exiting the accessSpecifier production.
	ExitAccessSpecifier(c *AccessSpecifierContext)

	// ExitClassDeclarationPart is called when exiting the classDeclarationPart production.
	ExitClassDeclarationPart(c *ClassDeclarationPartContext)

	// ExitInterfaceGuidConst is called when exiting the interfaceGuidConst production.
	ExitInterfaceGuidConst(c *InterfaceGuidConstContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitInterfaceDeclarationPart is called when exiting the interfaceDeclarationPart production.
	ExitInterfaceDeclarationPart(c *InterfaceDeclarationPartContext)

	// ExitErrorInterfaceDeclarationPart is called when exiting the errorInterfaceDeclarationPart production.
	ExitErrorInterfaceDeclarationPart(c *ErrorInterfaceDeclarationPartContext)

	// ExitErrorClassDeclarationPart is called when exiting the errorClassDeclarationPart production.
	ExitErrorClassDeclarationPart(c *ErrorClassDeclarationPartContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyReadDeclaration is called when exiting the propertyReadDeclaration production.
	ExitPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// ExitPropertyWriteDeclaration is called when exiting the propertyWriteDeclaration production.
	ExitPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// ExitPropertyDefaultValueDeclaration is called when exiting the propertyDefaultValueDeclaration production.
	ExitPropertyDefaultValueDeclaration(c *PropertyDefaultValueDeclarationContext)

	// ExitPropertyIndexDeclaration is called when exiting the propertyIndexDeclaration production.
	ExitPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// ExitPropertyIndexParameters is called when exiting the propertyIndexParameters production.
	ExitPropertyIndexParameters(c *PropertyIndexParametersContext)

	// ExitPropertyIndexParametersList is called when exiting the propertyIndexParametersList production.
	ExitPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// ExitGenericTemplate is called when exiting the genericTemplate production.
	ExitGenericTemplate(c *GenericTemplateContext)

	// ExitGenericTemplateList is called when exiting the genericTemplateList production.
	ExitGenericTemplateList(c *GenericTemplateListContext)

	// ExitGenericTypeParameter is called when exiting the genericTypeParameter production.
	ExitGenericTypeParameter(c *GenericTypeParameterContext)

	// ExitGenericConstraints is called when exiting the genericConstraints production.
	ExitGenericConstraints(c *GenericConstraintsContext)

	// ExitGenericConstraint is called when exiting the genericConstraint production.
	ExitGenericConstraint(c *GenericConstraintContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitSimpleType is called when exiting the simpleType production.
	ExitSimpleType(c *SimpleTypeContext)

	// ExitScalarType is called when exiting the scalarType production.
	ExitScalarType(c *ScalarTypeContext)

	// ExitScalerList is called when exiting the scalerList production.
	ExitScalerList(c *ScalerListContext)

	// ExitScalerMember is called when exiting the scalerMember production.
	ExitScalerMember(c *ScalerMemberContext)

	// ExitSubrangeType is called when exiting the subrangeType production.
	ExitSubrangeType(c *SubrangeTypeContext)

	// ExitStringTypeIdentifier is called when exiting the stringTypeIdentifier production.
	ExitStringTypeIdentifier(c *StringTypeIdentifierContext)

	// ExitStructuredType is called when exiting the structuredType production.
	ExitStructuredType(c *StructuredTypeContext)

	// ExitUnpackedStructuredType is called when exiting the unpackedStructuredType production.
	ExitUnpackedStructuredType(c *UnpackedStructuredTypeContext)

	// ExitStringtype is called when exiting the stringtype production.
	ExitStringtype(c *StringtypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitRecordContent is called when exiting the recordContent production.
	ExitRecordContent(c *RecordContentContext)

	// ExitRecordSection is called when exiting the recordSection production.
	ExitRecordSection(c *RecordSectionContext)

	// ExitRecordFieldsSection is called when exiting the recordFieldsSection production.
	ExitRecordFieldsSection(c *RecordFieldsSectionContext)

	// ExitRecordVariantSection is called when exiting the recordVariantSection production.
	ExitRecordVariantSection(c *RecordVariantSectionContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitRecordVariant is called when exiting the recordVariant production.
	ExitRecordVariant(c *RecordVariantContext)

	// ExitHelperType is called when exiting the helperType production.
	ExitHelperType(c *HelperTypeContext)

	// ExitHelperDeclarationPart is called when exiting the helperDeclarationPart production.
	ExitHelperDeclarationPart(c *HelperDeclarationPartContext)

	// ExitErrorHelperDeclarationPart is called when exiting the errorHelperDeclarationPart production.
	ExitErrorHelperDeclarationPart(c *ErrorHelperDeclarationPartContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitFileType is called when exiting the fileType production.
	ExitFileType(c *FileTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitVariableDeclarationPart is called when exiting the variableDeclarationPart production.
	ExitVariableDeclarationPart(c *VariableDeclarationPartContext)

	// ExitThreadvarDeclarationPart is called when exiting the threadvarDeclarationPart production.
	ExitThreadvarDeclarationPart(c *ThreadvarDeclarationPartContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitProcedureHeader is called when exiting the procedureHeader production.
	ExitProcedureHeader(c *ProcedureHeaderContext)

	// ExitFunctionHeader is called when exiting the functionHeader production.
	ExitFunctionHeader(c *FunctionHeaderContext)

	// ExitProcedureOrFunctionHeader is called when exiting the procedureOrFunctionHeader production.
	ExitProcedureOrFunctionHeader(c *ProcedureOrFunctionHeaderContext)

	// ExitProcedureOrFunctionHeaderModifiers is called when exiting the procedureOrFunctionHeaderModifiers production.
	ExitProcedureOrFunctionHeaderModifiers(c *ProcedureOrFunctionHeaderModifiersContext)

	// ExitProcedureOrFunctionDeclaration is called when exiting the procedureOrFunctionDeclaration production.
	ExitProcedureOrFunctionDeclaration(c *ProcedureOrFunctionDeclarationContext)

	// ExitProcedureDeclaration is called when exiting the procedureDeclaration production.
	ExitProcedureDeclaration(c *ProcedureDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitProcedureLambdaDeclaration is called when exiting the procedureLambdaDeclaration production.
	ExitProcedureLambdaDeclaration(c *ProcedureLambdaDeclarationContext)

	// ExitFunctionLambdaDeclaration is called when exiting the functionLambdaDeclaration production.
	ExitFunctionLambdaDeclaration(c *FunctionLambdaDeclarationContext)

	// ExitResultType is called when exiting the resultType production.
	ExitResultType(c *ResultTypeContext)

	// ExitProcedureOrFunctionBody is called when exiting the procedureOrFunctionBody production.
	ExitProcedureOrFunctionBody(c *ProcedureOrFunctionBodyContext)

	// ExitClassOperatorHeader is called when exiting the classOperatorHeader production.
	ExitClassOperatorHeader(c *ClassOperatorHeaderContext)

	// ExitClassOperatorDeclaration is called when exiting the classOperatorDeclaration production.
	ExitClassOperatorDeclaration(c *ClassOperatorDeclarationContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameterSection is called when exiting the formalParameterSection production.
	ExitFormalParameterSection(c *FormalParameterSectionContext)

	// ExitParameterGroup is called when exiting the parameterGroup production.
	ExitParameterGroup(c *ParameterGroupContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitQualifiedIdentifierList is called when exiting the qualifiedIdentifierList production.
	ExitQualifiedIdentifierList(c *QualifiedIdentifierListContext)

	// ExitConstList is called when exiting the constList production.
	ExitConstList(c *ConstListContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitTypedIdentifierList is called when exiting the typedIdentifierList production.
	ExitTypedIdentifierList(c *TypedIdentifierListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitUnlabelledStatement is called when exiting the unlabelledStatement production.
	ExitUnlabelledStatement(c *UnlabelledStatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitRaiseExceptionStatement is called when exiting the raiseExceptionStatement production.
	ExitRaiseExceptionStatement(c *RaiseExceptionStatementContext)

	// ExitVariableDeclarationStatement is called when exiting the variableDeclarationStatement production.
	ExitVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// ExitTypeCast is called when exiting the typeCast production.
	ExitTypeCast(c *TypeCastContext)

	// ExitPropertyDesignator is called when exiting the propertyDesignator production.
	ExitPropertyDesignator(c *PropertyDesignatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelationaloperator is called when exiting the relationaloperator production.
	ExitRelationaloperator(c *RelationaloperatorContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitAdditiveoperator is called when exiting the additiveoperator production.
	ExitAdditiveoperator(c *AdditiveoperatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMultiplicativeoperator is called when exiting the multiplicativeoperator production.
	ExitMultiplicativeoperator(c *MultiplicativeoperatorContext)

	// ExitSignedFactor is called when exiting the signedFactor production.
	ExitSignedFactor(c *SignedFactorContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPostfixOp is called when exiting the postfixOp production.
	ExitPostfixOp(c *PostfixOpContext)

	// ExitUnsignedConstant is called when exiting the unsignedConstant production.
	ExitUnsignedConstant(c *UnsignedConstantContext)

	// ExitFunctionDesignator is called when exiting the functionDesignator production.
	ExitFunctionDesignator(c *FunctionDesignatorContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitActualParameter is called when exiting the actualParameter production.
	ExitActualParameter(c *ActualParameterContext)

	// ExitParameterwidth is called when exiting the parameterwidth production.
	ExitParameterwidth(c *ParameterwidthContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

	// ExitInheritedStatement is called when exiting the inheritedStatement production.
	ExitInheritedStatement(c *InheritedStatementContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitEmpty_ is called when exiting the empty_ production.
	ExitEmpty_(c *Empty_Context)

	// ExitStructuredStatement is called when exiting the structuredStatement production.
	ExitStructuredStatement(c *StructuredStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitConditionalStatement is called when exiting the conditionalStatement production.
	ExitConditionalStatement(c *ConditionalStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitCaseConstRange is called when exiting the caseConstRange production.
	ExitCaseConstRange(c *CaseConstRangeContext)

	// ExitCaseConstList is called when exiting the caseConstList production.
	ExitCaseConstList(c *CaseConstListContext)

	// ExitCaseListElement is called when exiting the caseListElement production.
	ExitCaseListElement(c *CaseListElementContext)

	// ExitRepetetiveStatement is called when exiting the repetetiveStatement production.
	ExitRepetetiveStatement(c *RepetetiveStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitRepeatStatement is called when exiting the repeatStatement production.
	ExitRepeatStatement(c *RepeatStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForList is called when exiting the forList production.
	ExitForList(c *ForListContext)

	// ExitInitialValue is called when exiting the initialValue production.
	ExitInitialValue(c *InitialValueContext)

	// ExitFinalValue is called when exiting the finalValue production.
	ExitFinalValue(c *FinalValueContext)

	// ExitWithStatement is called when exiting the withStatement production.
	ExitWithStatement(c *WithStatementContext)

	// ExitTryExceptStatement is called when exiting the tryExceptStatement production.
	ExitTryExceptStatement(c *TryExceptStatementContext)

	// ExitExceptionCase is called when exiting the exceptionCase production.
	ExitExceptionCase(c *ExceptionCaseContext)

	// ExitExceptionElse is called when exiting the exceptionElse production.
	ExitExceptionElse(c *ExceptionElseContext)

	// ExitTryFinallyStatement is called when exiting the tryFinallyStatement production.
	ExitTryFinallyStatement(c *TryFinallyStatementContext)

	// ExitWithStatementVariableList is called when exiting the withStatementVariableList production.
	ExitWithStatementVariableList(c *WithStatementVariableListContext)

	// ExitAttributeSection is called when exiting the attributeSection production.
	ExitAttributeSection(c *AttributeSectionContext)

	// ExitAttributeList is called when exiting the attributeList production.
	ExitAttributeList(c *AttributeListContext)

	// ExitAttributeItem is called when exiting the attributeItem production.
	ExitAttributeItem(c *AttributeItemContext)
}
