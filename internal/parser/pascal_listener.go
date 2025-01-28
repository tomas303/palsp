// Code generated from pascal.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

	// EnterFooter is called when entering the footer production.
	EnterFooter(c *FooterContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterTopLevelDeclarations is called when entering the topLevelDeclarations production.
	EnterTopLevelDeclarations(c *TopLevelDeclarationsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterUsesUnitsPart is called when entering the usesUnitsPart production.
	EnterUsesUnitsPart(c *UsesUnitsPartContext)

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

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

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

	// EnterTypeDefinitionPart is called when entering the typeDefinitionPart production.
	EnterTypeDefinitionPart(c *TypeDefinitionPartContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterProcedureType is called when entering the procedureType production.
	EnterProcedureType(c *ProcedureTypeContext)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterClassImplementsInterfaces is called when entering the classImplementsInterfaces production.
	EnterClassImplementsInterfaces(c *ClassImplementsInterfacesContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassPrivateDeclaration is called when entering the classPrivateDeclaration production.
	EnterClassPrivateDeclaration(c *ClassPrivateDeclarationContext)

	// EnterClassStrictPrivateDeclaration is called when entering the classStrictPrivateDeclaration production.
	EnterClassStrictPrivateDeclaration(c *ClassStrictPrivateDeclarationContext)

	// EnterClassProtectedDeclaration is called when entering the classProtectedDeclaration production.
	EnterClassProtectedDeclaration(c *ClassProtectedDeclarationContext)

	// EnterClassStrictProtectedDeclaration is called when entering the classStrictProtectedDeclaration production.
	EnterClassStrictProtectedDeclaration(c *ClassStrictProtectedDeclarationContext)

	// EnterClassPublicDeclaration is called when entering the classPublicDeclaration production.
	EnterClassPublicDeclaration(c *ClassPublicDeclarationContext)

	// EnterClassPublishedDeclaration is called when entering the classPublishedDeclaration production.
	EnterClassPublishedDeclaration(c *ClassPublishedDeclarationContext)

	// EnterClassImplicitPublishedDeclaration is called when entering the classImplicitPublishedDeclaration production.
	EnterClassImplicitPublishedDeclaration(c *ClassImplicitPublishedDeclarationContext)

	// EnterClassDeclarationPart is called when entering the classDeclarationPart production.
	EnterClassDeclarationPart(c *ClassDeclarationPartContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyReadDeclaration is called when entering the propertyReadDeclaration production.
	EnterPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// EnterPropertyWriteDeclaration is called when entering the propertyWriteDeclaration production.
	EnterPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// EnterPropertyIndexDeclaration is called when entering the propertyIndexDeclaration production.
	EnterPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// EnterPropertyIndexParameters is called when entering the propertyIndexParameters production.
	EnterPropertyIndexParameters(c *PropertyIndexParametersContext)

	// EnterPropertyIndexParametersList is called when entering the propertyIndexParametersList production.
	EnterPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// EnterMethodIdentifier is called when entering the methodIdentifier production.
	EnterMethodIdentifier(c *MethodIdentifierContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterSimpleType is called when entering the simpleType production.
	EnterSimpleType(c *SimpleTypeContext)

	// EnterScalarType is called when entering the scalarType production.
	EnterScalarType(c *ScalarTypeContext)

	// EnterSubrangeType is called when entering the subrangeType production.
	EnterSubrangeType(c *SubrangeTypeContext)

	// EnterTypeIdentifier is called when entering the typeIdentifier production.
	EnterTypeIdentifier(c *TypeIdentifierContext)

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

	// EnterComponentType is called when entering the componentType production.
	EnterComponentType(c *ComponentTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterFixedPart is called when entering the fixedPart production.
	EnterFixedPart(c *FixedPartContext)

	// EnterRecordSection is called when entering the recordSection production.
	EnterRecordSection(c *RecordSectionContext)

	// EnterVariantPart is called when entering the variantPart production.
	EnterVariantPart(c *VariantPartContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterVariant is called when entering the variant production.
	EnterVariant(c *VariantContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterFileType is called when entering the fileType production.
	EnterFileType(c *FileTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterVariableDeclarationPart is called when entering the variableDeclarationPart production.
	EnterVariableDeclarationPart(c *VariableDeclarationPartContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterProcedureAndFunctionDeclarationPart is called when entering the procedureAndFunctionDeclarationPart production.
	EnterProcedureAndFunctionDeclarationPart(c *ProcedureAndFunctionDeclarationPartContext)

	// EnterProcedureOrFunctionDeclaration is called when entering the procedureOrFunctionDeclaration production.
	EnterProcedureOrFunctionDeclaration(c *ProcedureOrFunctionDeclarationContext)

	// EnterProcedureDeclaration is called when entering the procedureDeclaration production.
	EnterProcedureDeclaration(c *ProcedureDeclarationContext)

	// EnterProcedureMethodDeclaration is called when entering the procedureMethodDeclaration production.
	EnterProcedureMethodDeclaration(c *ProcedureMethodDeclarationContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameterSection is called when entering the formalParameterSection production.
	EnterFormalParameterSection(c *FormalParameterSectionContext)

	// EnterParameterGroup is called when entering the parameterGroup production.
	EnterParameterGroup(c *ParameterGroupContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterConstList is called when entering the constList production.
	EnterConstList(c *ConstListContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterResultType is called when entering the resultType production.
	EnterResultType(c *ResultTypeContext)

	// EnterFunctionMethodDeclaration is called when entering the functionMethodDeclaration production.
	EnterFunctionMethodDeclaration(c *FunctionMethodDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterUnlabelledStatement is called when entering the unlabelledStatement production.
	EnterUnlabelledStatement(c *UnlabelledStatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

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

	// EnterProcedureStatement is called when entering the procedureStatement production.
	EnterProcedureStatement(c *ProcedureStatementContext)

	// EnterActualParameter is called when entering the actualParameter production.
	EnterActualParameter(c *ActualParameterContext)

	// EnterParameterwidth is called when entering the parameterwidth production.
	EnterParameterwidth(c *ParameterwidthContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

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

	// EnterTryFinallyStatement is called when entering the tryFinallyStatement production.
	EnterTryFinallyStatement(c *TryFinallyStatementContext)

	// EnterRecordVariableList is called when entering the recordVariableList production.
	EnterRecordVariableList(c *RecordVariableListContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitFooter is called when exiting the footer production.
	ExitFooter(c *FooterContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitTopLevelDeclarations is called when exiting the topLevelDeclarations production.
	ExitTopLevelDeclarations(c *TopLevelDeclarationsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitUsesUnitsPart is called when exiting the usesUnitsPart production.
	ExitUsesUnitsPart(c *UsesUnitsPartContext)

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

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

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

	// ExitTypeDefinitionPart is called when exiting the typeDefinitionPart production.
	ExitTypeDefinitionPart(c *TypeDefinitionPartContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitProcedureType is called when exiting the procedureType production.
	ExitProcedureType(c *ProcedureTypeContext)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitClassImplementsInterfaces is called when exiting the classImplementsInterfaces production.
	ExitClassImplementsInterfaces(c *ClassImplementsInterfacesContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassPrivateDeclaration is called when exiting the classPrivateDeclaration production.
	ExitClassPrivateDeclaration(c *ClassPrivateDeclarationContext)

	// ExitClassStrictPrivateDeclaration is called when exiting the classStrictPrivateDeclaration production.
	ExitClassStrictPrivateDeclaration(c *ClassStrictPrivateDeclarationContext)

	// ExitClassProtectedDeclaration is called when exiting the classProtectedDeclaration production.
	ExitClassProtectedDeclaration(c *ClassProtectedDeclarationContext)

	// ExitClassStrictProtectedDeclaration is called when exiting the classStrictProtectedDeclaration production.
	ExitClassStrictProtectedDeclaration(c *ClassStrictProtectedDeclarationContext)

	// ExitClassPublicDeclaration is called when exiting the classPublicDeclaration production.
	ExitClassPublicDeclaration(c *ClassPublicDeclarationContext)

	// ExitClassPublishedDeclaration is called when exiting the classPublishedDeclaration production.
	ExitClassPublishedDeclaration(c *ClassPublishedDeclarationContext)

	// ExitClassImplicitPublishedDeclaration is called when exiting the classImplicitPublishedDeclaration production.
	ExitClassImplicitPublishedDeclaration(c *ClassImplicitPublishedDeclarationContext)

	// ExitClassDeclarationPart is called when exiting the classDeclarationPart production.
	ExitClassDeclarationPart(c *ClassDeclarationPartContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyReadDeclaration is called when exiting the propertyReadDeclaration production.
	ExitPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// ExitPropertyWriteDeclaration is called when exiting the propertyWriteDeclaration production.
	ExitPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// ExitPropertyIndexDeclaration is called when exiting the propertyIndexDeclaration production.
	ExitPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// ExitPropertyIndexParameters is called when exiting the propertyIndexParameters production.
	ExitPropertyIndexParameters(c *PropertyIndexParametersContext)

	// ExitPropertyIndexParametersList is called when exiting the propertyIndexParametersList production.
	ExitPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// ExitMethodIdentifier is called when exiting the methodIdentifier production.
	ExitMethodIdentifier(c *MethodIdentifierContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitSimpleType is called when exiting the simpleType production.
	ExitSimpleType(c *SimpleTypeContext)

	// ExitScalarType is called when exiting the scalarType production.
	ExitScalarType(c *ScalarTypeContext)

	// ExitSubrangeType is called when exiting the subrangeType production.
	ExitSubrangeType(c *SubrangeTypeContext)

	// ExitTypeIdentifier is called when exiting the typeIdentifier production.
	ExitTypeIdentifier(c *TypeIdentifierContext)

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

	// ExitComponentType is called when exiting the componentType production.
	ExitComponentType(c *ComponentTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitFixedPart is called when exiting the fixedPart production.
	ExitFixedPart(c *FixedPartContext)

	// ExitRecordSection is called when exiting the recordSection production.
	ExitRecordSection(c *RecordSectionContext)

	// ExitVariantPart is called when exiting the variantPart production.
	ExitVariantPart(c *VariantPartContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitVariant is called when exiting the variant production.
	ExitVariant(c *VariantContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitFileType is called when exiting the fileType production.
	ExitFileType(c *FileTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitVariableDeclarationPart is called when exiting the variableDeclarationPart production.
	ExitVariableDeclarationPart(c *VariableDeclarationPartContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitProcedureAndFunctionDeclarationPart is called when exiting the procedureAndFunctionDeclarationPart production.
	ExitProcedureAndFunctionDeclarationPart(c *ProcedureAndFunctionDeclarationPartContext)

	// ExitProcedureOrFunctionDeclaration is called when exiting the procedureOrFunctionDeclaration production.
	ExitProcedureOrFunctionDeclaration(c *ProcedureOrFunctionDeclarationContext)

	// ExitProcedureDeclaration is called when exiting the procedureDeclaration production.
	ExitProcedureDeclaration(c *ProcedureDeclarationContext)

	// ExitProcedureMethodDeclaration is called when exiting the procedureMethodDeclaration production.
	ExitProcedureMethodDeclaration(c *ProcedureMethodDeclarationContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameterSection is called when exiting the formalParameterSection production.
	ExitFormalParameterSection(c *FormalParameterSectionContext)

	// ExitParameterGroup is called when exiting the parameterGroup production.
	ExitParameterGroup(c *ParameterGroupContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitConstList is called when exiting the constList production.
	ExitConstList(c *ConstListContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitResultType is called when exiting the resultType production.
	ExitResultType(c *ResultTypeContext)

	// ExitFunctionMethodDeclaration is called when exiting the functionMethodDeclaration production.
	ExitFunctionMethodDeclaration(c *FunctionMethodDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitUnlabelledStatement is called when exiting the unlabelledStatement production.
	ExitUnlabelledStatement(c *UnlabelledStatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

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

	// ExitProcedureStatement is called when exiting the procedureStatement production.
	ExitProcedureStatement(c *ProcedureStatementContext)

	// ExitActualParameter is called when exiting the actualParameter production.
	ExitActualParameter(c *ActualParameterContext)

	// ExitParameterwidth is called when exiting the parameterwidth production.
	ExitParameterwidth(c *ParameterwidthContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

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

	// ExitTryFinallyStatement is called when exiting the tryFinallyStatement production.
	ExitTryFinallyStatement(c *TryFinallyStatementContext)

	// ExitRecordVariableList is called when exiting the recordVariableList production.
	ExitRecordVariableList(c *RecordVariableListContext)
}
