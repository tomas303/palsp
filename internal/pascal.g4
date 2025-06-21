/*
 BSD License
 
 Copyright (c) 2013, Tom Everett All rights reserved.
 
 Redistribution and use in source and binary forms, with or without modification, are permitted
 provided that the following conditions are met:
 
 1. Redistributions of source code must retain the above copyright notice, this list of conditions
 and the following disclaimer. 2. Redistributions in binary form must reproduce the above copyright
 notice, this list of conditions and the following disclaimer in the documentation and/or other
 materials provided with the distribution. 3. Neither the name of Tom Everett nor the names of its
 contributors may be used to endorse or promote products derived from this software without specific
 prior written permission.
 
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
 FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
 CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
 IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
/*
 Adapted from pascal.g by Hakki Dogusan, Piet Schoutteten and Marton Papp
 */

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

grammar pascal;

options {
    caseInsensitive = true;
}

source
    : program
    | unit
    ;

program
    : PROGRAM identifier (LPAREN identifierList RPAREN)? SEMI usesUnits? implementationBlock BEGIN statements END DOT EOF
    ;

unit
    : UNIT identifier SEMI interfaceSection implementationSection? initializationSection? finalizationSection? END DOT EOF
    ;

interfaceSection
    : INTERFACE usesUnits? interfaceBlock?
    ;

implementationSection
    : IMPLEMENTATION usesUnits? implementationBlock?
    ;

initializationSection
    : INITIALIZATION statements
    ;

finalizationSection
    : FINALIZATION statements
    ;

identifier
    : identifierPart (DOT identifierPart)*
    ;

identifierPart
    : (
        IDENT
        | INDEX
        | READ
        | WRITE
        | CHAR
        | BOOLEAN
        | INTEGER
        | REAL
        | STRING
        | CARDINAL
        | LONGBOOL
        | LONGINT
        | LONGWORD
        | WORD
        | BYTE
        | SHORTINT
        | SMALLINT
        | INT64
        | UINT64
        | SINGLE
        | DOUBLE
        | EXTENDED
        | COMP
        | CURRENCY
        | ANSICHAR
        | WIDECHAR
        | ANSISTRING
        | WIDESTRING
        | UNICODESTRING
        | RAWBYTESTRING
        | UTF8STRING
        | VARIANT
        | OLEVARIANT
        | POINTER
        | PCHAR
        | PANSICHAR
        | PWIDECHAR
        | PUNICODECHAR
        | THANDLE
        | HWND
        | HDC
        | HICON
        | HBITMAP
        | HMENU
        | HINSTANCE
        | HMODULE
        | HKEY
        | DWORD
        | QWORD
        | NATIVEINT
        | NATIVEUINT
        | CODEPAGE
        | TGUID
        | PGUID
        | TEXTFILE
        | TEXT
        | SHORTSTRING
        | OPENSTRING
    ) genericTemplate?
    ;

interfaceBlockMember
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | resourceDefinitionPart
        | typeDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionHeader
    )
    ;

interfaceBlock
    : interfaceBlockMember (SEMI interfaceBlockMember)* SEMI?
    ;

implementationBlockMember
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | resourceDefinitionPart
        | typeDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionDeclaration
        | procedureOrFunctionHeader SEMI FORWARD
        | classOperatorDeclaration
    )
    ;

implementationBlock
    : implementationBlockMember (SEMI implementationBlockMember)* SEMI?
    ;

funcBlockMemeber
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionDeclaration
        | classOperatorDeclaration
    )
    ;

funcBlock
    : funcBlockMemeber (SEMI funcBlockMemeber)* SEMI?
    ;

usesUnits
    : USES identifierList SEMI
    ;

labelDeclarationPart
    : LABEL label (COMMA label)*
    ;

label
    : unsignedInteger
    ;

constantDefinitionPart
    : CONST constantDefinition (SEMI constantDefinition)*
    ;

constantDefinition
    : identifier (COLON typeIdentifier)? EQUAL constant platformHint? deprecatedHint?
    | identifier COLON arrayType EQUAL constant platformHint? deprecatedHint?
    | identifier COLON setType EQUAL constant platformHint? deprecatedHint?
    ;

constantChr
    : CHR LPAREN unsignedInteger RPAREN
    ;

hexConstant
    : HEX_LITERAL
    ;

constant
    : unsignedConstant
    | identifier
    | scalarType
    | arrayConstant (PLUS (arrayConstant | identifier))*
    | recordConstant
    | simpleExpression
    ;

arrayConstant
    : LBRACK constant (COMMA constant)* RBRACK
    | LPAREN constant (COMMA constant)* RPAREN
    ;

recordConstant
    : LPAREN recordField (SEMI recordField)* RPAREN
    ;

recordField
    : identifier COLON constant
    | identifier COLON constant DOTDOT constant
    | identifier COLON arrayConstant
    ;

unsignedNumber
    : unsignedInteger
    | unsignedReal
    ;

unsignedInteger
    : NUM_INT
    ;

unsignedReal
    : NUM_REAL
    ;

sign
    : PLUS
    | MINUS
    ;

bool_
    : TRUE
    | FALSE
    ;

string
    : (STRING_LITERAL | STRING_CROSSHATCH_LITERAL)+
    ;

stringExpression
    : string (PLUS (string | identifier))*
    ;

resourceDefinitionPart
    : RESOURCESTRING resourceDefinition (SEMI resourceDefinition)*
    ;

resourceDefinition
    : identifier EQUAL stringExpression
    ;

deprecatedHint
    : DEPRECATED stringExpression?
    ;

platformHint
    : PLATFORM (EQUAL (unsignedConstant | FALSE | TRUE))?
    ;

alignHint
    : ALIGN unsignedInteger?
    ;

typeDefinitionPart
    : TYPE typeDefinition (SEMI typeDefinition)*
    ;

typeDefinition
    : attributeSection? identifier EQUAL (
        | aliasDistinctType
        | aliasType
        | type_
    )
    ;

classType
    : CLASS OF typeIdentifier
    | CLASS (LPAREN identifier RPAREN)
    | CLASS (
        (LPAREN identifier classImplementsInterfaces RPAREN)? ABSTRACT? classTypeBlock? (
            SEMI? accessSpecifier classTypeBlock?
        )* SEMI? END
    )?
    ;

classTypeBlock
    : classDeclarationPart (SEMI classDeclarationPart)*
    ;

interfaceType
    : INTERFACE (
        (LPAREN identifier RPAREN)? GUID_LITERAL? interfaceDeclarationPart? (
            SEMI interfaceDeclarationPart
        )* SEMI? END
    )?
    ;

functionType
    : FUNCTION (formalParameterList)? COLON resultType (OF OBJECT)? procedureOrFunctionHeaderModifiers
    | REFERENCE TO FUNCTION (formalParameterList)? COLON resultType procedureOrFunctionHeaderModifiers
    ;

procedureType
    : PROCEDURE (formalParameterList)? (OF OBJECT)? procedureOrFunctionHeaderModifiers
    | REFERENCE TO PROCEDURE (formalParameterList)? procedureOrFunctionHeaderModifiers
    ;

aliasDistinctType
    : TYPE typeIdentifier
    ;

aliasType
    : typeIdentifier
    ;

classImplementsInterfaces
    : (COMMA typeIdentifier)*
    ;

accessSpecifier
    : STRICT? PRIVATE
    | STRICT? PROTECTED
    | PUBLIC
    | PUBLISHED
    ;

classDeclarationPart
    : CLASS? VAR? attributeSection? typedIdentifierList
    | typeDefinitionPart
    | constantDefinitionPart
    | functionHeader
    | procedureHeader
    | CLASS? propertyDeclaration (SEMI DEFAULT)?
    ;

GUID_LITERAL
    : '[' '\'' '{' HEX_CHAR_SEQ8 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ12 '}' '\'' ']'
    ;

interfaceGuidConst
    : GUID_LITERAL
    ;

interfaceDeclaration
    : interfaceDeclarationPart*
    ;

interfaceDeclarationPart
    : functionHeader
    | procedureHeader
    | propertyDeclaration
    | errorInterfaceDeclarationPart
    ;

errorInterfaceDeclarationPart
    : ~(END)+
    ;

errorClassDeclarationPart
    : ~(STRICT | PRIVATE | PROTECTED | PUBLIC | PUBLISHED | END)+
    ;

propertyDeclaration
    : PROPERTY identifier propertyIndexParameters? COLON typeIdentifier propertyReadDeclaration? propertyWriteDeclaration?
        propertyDefaultValueDeclaration? propertyIndexDeclaration? (SEMI DEFAULT)?
    | PROPERTY identifier propertyDefaultValueDeclaration?
    ;

propertyReadDeclaration
    : READ identifier
    ;

propertyWriteDeclaration
    : WRITE identifier
    ;

propertyDefaultValueDeclaration
    : DEFAULT expression
    ;

propertyIndexDeclaration
    : INDEX unsignedNumber
    ;

propertyIndexParameters
    : LBRACK formalParameterSection (COMMA formalParameterSection)* RBRACK
    ;

propertyIndexParametersList
    : identifierList COLON indexType (SEMI identifierList COLON indexType)*
    ;

genericTemplate
    : '<' genericTemplateList '>'
    ;

genericTemplateList
    : genericTypeParameter (COMMA genericTypeParameter)*
    ;

genericTypeParameter
    : typeIdentifier (COLON genericConstraints)?
    ;

genericConstraints
    : genericConstraint (COMMA genericConstraint)*
    ;

genericConstraint
    : CONSTRUCTOR
    | CLASS
    | RECORD
    | CLASS OF typeIdentifier
    | typeIdentifier
    ;

type_
    : simpleType
    | structuredType
    | pointerType
    | functionType
    | procedureType
    ;

simpleType
    : scalarType
    | subrangeType
    | typeIdentifier
    | stringtype
    ;

scalarType
    : LPAREN scalerList RPAREN
    ;

scalerList
    : scalerMember (COMMA scalerMember)*
    ;

scalerMember
    : identifier (EQUAL expression)?
    ;

subrangeType
    : simpleExpression DOTDOT simpleExpression
    ;

stringTypeIdentifier
    : (STRING | ANSISTRING | WIDESTRING | UNICODESTRING | RAWBYTESTRING | UTF8STRING | SHORTSTRING)
    ;

typeIdentifier
    : identifier
    | (
        CHAR
        | BOOLEAN
        | INTEGER
        | REAL
        | CARDINAL
        | LONGBOOL
        | LONGINT
        | LONGWORD
        | WORD
        | BYTE
        | SHORTINT
        | SMALLINT
        | INT64
        | UINT64
        | SINGLE
        | DOUBLE
        | EXTENDED
        | COMP
        | CURRENCY
        | ANSICHAR
        | WIDECHAR
        | VARIANT
        | OLEVARIANT
        | POINTER
        | PCHAR
        | PANSICHAR
        | PWIDECHAR
        | PUNICODECHAR
        | THANDLE
        | HWND
        | HDC
        | HICON
        | HBITMAP
        | HMENU
        | HINSTANCE
        | HMODULE
        | HKEY
        | DWORD
        | QWORD
        | NATIVEINT
        | NATIVEUINT
        | CODEPAGE
        | TGUID
        | PGUID
        | TEXTFILE
        | TEXT
        | OPENSTRING
    )
    | stringtype
    | arrayType
    ;

structuredType
    : PACKED unpackedStructuredType
    | unpackedStructuredType
    | helperType
    | classType
    | PACKED? recordType
    | interfaceType
    ;

unpackedStructuredType
    : arrayType
    | setType
    | fileType
    ;

stringtype
    : stringTypeIdentifier (
        LBRACK (identifier | unsignedNumber | hexConstant) RBRACK     // Square brackets
        | LPAREN (identifier | unsignedNumber | hexConstant) RPAREN   // Parentheses
    )?
    ;

arrayType
    : ARRAY LBRACK typeList RBRACK OF type_ (EQUAL scalarType)?
    | ARRAY LBRACK2 typeList RBRACK2 OF type_ (EQUAL scalarType)?
    | ARRAY OF CONST
    | ARRAY OF type_
    ;

typeList
    : indexType (COMMA indexType)*
    ;

indexType
    : simpleType
    ;

recordType
    : RECORD recordParts? SEMI? END alignHint? platformHint? deprecatedHint?
    | RECORD recordTypeBlock? (SEMI? accessSpecifier recordTypeBlock?)* SEMI? END alignHint? platformHint? deprecatedHint?
    ;

recordTypeBlock
    : recordDeclarationPart (SEMI recordDeclarationPart)*
    ;

recordDeclarationPart
    : CLASS? VAR? attributeSection? typedIdentifierList
    | typeDefinitionPart
    | constantDefinitionPart
    | functionHeader
    | procedureHeader
    | classOperatorHeader
    | CLASS? propertyDeclaration (SEMI DEFAULT)?
    ;

recordParts
    : recordFixedPart (SEMI recordVariantPart)?
    | recordVariantPart
    ;

recordFixedPart
    : typedIdentifierList (SEMI typedIdentifierList)* SEMI?
    ;

recordVariantPart
    : CASE tag OF recordVariant (SEMI recordVariant)*
    ;

tag
    : identifier COLON typeIdentifier
    | typeIdentifier
    ;

recordVariant
    : constList COLON LPAREN recordParts SEMI? RPAREN
    ;

helperType
    : CLASS HELPER FOR typeIdentifier accessSpecifier? helperDeclarationPart? (
        SEMI helperDeclarationPart
    )* (accessSpecifier helperDeclarationPart? (SEMI helperDeclarationPart)*) END
    ;

helperDeclarationPart
    : typeDefinitionPart
    | constantDefinitionPart
    | functionHeader
    | procedureHeader
    | propertyDeclaration (SEMI DEFAULT)?
    | errorHelperDeclarationPart
    ;

errorHelperDeclarationPart
    : ~(STRICT | PRIVATE | PROTECTED | PUBLIC | PUBLISHED | END)+
    ;

setType
    : SET OF simpleType
    ;

fileType
    : FILE OF type_
    | FILE
    ;

pointerType
    : DEREFERENCE typeIdentifier
    ;

variableDeclarationPart
    : VAR variableDeclaration (SEMI variableDeclaration)*
    ;

variableDeclaration
    : attributeSection? typedIdentifierList (EQUAL simpleExpression)?
    ;

procedureHeader
    : attributeSection? CLASS? (PROCEDURE | CONSTRUCTOR | DESTRUCTOR) identifier (
        formalParameterList
    )? (SEMI deprecatedHint)? procedureOrFunctionHeaderModifiers (SEMI deprecatedHint)?
    ;

functionHeader
    : attributeSection? CLASS? FUNCTION identifier (formalParameterList)? COLON resultType (
        SEMI deprecatedHint
    )? procedureOrFunctionHeaderModifiers (SEMI deprecatedHint)?
    ;

procedureOrFunctionHeader
    : procedureHeader
    | functionHeader
    ;

procedureOrFunctionHeaderModifiers
    : (
        SEMI? (
            CDECL
            | STDCALL
            | ABSTRACT
            | VIRTUAL
            | OVERRIDE
            | REINTRODUCE
            | OVERLOAD
            | INLINE
            | STATIC
            | PLATFORM
        )
    )*
    ;

procedureOrFunctionDeclaration
    : procedureDeclaration
    | functionDeclaration
    ;

procedureDeclaration
    : procedureHeader procedureOrFunctionBody SEMI
    ;

functionDeclaration
    : functionHeader procedureOrFunctionBody SEMI
    ;

procedureLambdaDeclaration
    : PROCEDURE (formalParameterList)? procedureOrFunctionBody
    ;

functionLambdaDeclaration
    : FUNCTION (formalParameterList)? COLON resultType procedureOrFunctionBody
    ;

resultType
    : typeIdentifier
    ;

procedureOrFunctionBody
    : funcBlock? compoundStatement
    ;

classOperatorHeader
    : CLASS OPERATOR identifier (formalParameterList)? COLON resultType procedureOrFunctionHeaderModifiers
    ;

classOperatorDeclaration
    : classOperatorHeader procedureOrFunctionBody
    ;

formalParameterList
    : LPAREN formalParameterSection (SEMI formalParameterSection)* RPAREN
    ;

formalParameterSection
    : attributeSection? parameterGroup
    | attributeSection? VAR parameterGroup
    | attributeSection? CONST parameterGroup
    | attributeSection? OUT parameterGroup
    | attributeSection? FUNCTION parameterGroup
    | attributeSection? PROCEDURE parameterGroup
    ;

parameterGroup
    : identifierList (COLON typeIdentifier)? defaultValue?
    ;

identifierList
    : identifier (COMMA identifier)*
    ;

constList
    : constant (COMMA constant)*
    ;

defaultValue
    : EQUAL expression
    ;

typedIdentifierList
    : identifierList COLON type_ platformHint? deprecatedHint?
    ;

statement
    : label COLON unlabelledStatement
    | unlabelledStatement
    | errorStatement
    ;

errorStatement
    : ~('END' | ';')+
    ;

unlabelledStatement
    : simpleStatement
    | structuredStatement
    ;

simpleStatement
    : assignmentStatement
    | methodCallStatement
    | procedureStatement
    | gotoStatement
    | inheritedStatement
    | typeCast
    | emptyStatement_
    | raiseExceptionStatement
    | variableDeclarationStatement
    ;

assignmentStatement
    : variableDesignator ASSIGN expression
    | propertyDesignator ASSIGN expression
    ;

raiseExceptionStatement
    : RAISE expression?
    ;

variableDeclarationStatement
    : VAR identifierList (COLON type_)? (ASSIGN expression)?
    ;

variableDesignator
    : (typeCast | AT identifier | functionDesignator) (
        LBRACK expression (COMMA expression)* RBRACK
        | LBRACK2 expression (COMMA expression)* RBRACK2
        | DOT functionDesignator
        | DEREFERENCE+
    )*
    ;

typeCast
    : typeIdentifier LPAREN expression RPAREN
    | LPAREN expression AS typeIdentifier RPAREN
    ;

propertyDesignator
    : functionDesignator (DOT functionDesignator)* DOT identifier (
        LBRACK expression (COMMA expression)* RBRACK
    )?
    ;

expression
    : simpleExpression (relationaloperator expression)?
    ;

relationaloperator
    : EQUAL
    | NOT_EQUAL
    | LT
    | LE
    | GE
    | GT
    | IN
    ;

simpleExpression
    : term (additiveoperator simpleExpression)?
    ;

additiveoperator
    : PLUS
    | MINUS
    | OR
    ;

term
    : signedFactor (multiplicativeoperator term)?
    ;

multiplicativeoperator
    : STAR
    | SLASH
    | DIV
    | MOD
    | AND
    | SHR
    | SHL
    | XOR
    ;

signedFactor
    : (PLUS | MINUS)? factor
    ;

factor
    : INHERITED? functionDesignator
    | defaultDesignator
    | variableDesignator (AS identifier)?
    | LPAREN expression RPAREN
    | unsignedConstant
    | set_
    | NOT factor
    | bool_
    | factor LBRACK expression (COMMA expression)* RBRACK
    | AT? typeIdentifier (LPAREN expression RPAREN)? (DEREFERENCE)*
    | factor (DOT expression)+
    | identifier
    | functionLambdaDeclaration
    | procedureLambdaDeclaration
    ;

unsignedConstant
    : sign? unsignedNumber
    | constantChr
    | hexConstant
    | string
    | NIL
    ;

functionDesignator
    : (identifier) (LPAREN parameterList? RPAREN)?
    ;

defaultDesignator
    : DEFAULT (LPAREN parameterList RPAREN)?
    ;

parameterList
    : actualParameter (COMMA actualParameter)*
    ;

set_
    : LBRACK elementList RBRACK
    | LBRACK2 elementList RBRACK2
    ;

elementList
    : element (COMMA element)*
    |
    ;

element
    : expression (DOTDOT expression)?
    ;

procedureStatement
    : identifier (LPAREN parameterList? RPAREN)?
    ;

methodCallStatement
    : variableDesignator (DOT variableDesignator)* DOT identifier (LPAREN parameterList RPAREN)?
    ;

actualParameter
    : expression parameterwidth*
    | procedureLambdaDeclaration
    | functionLambdaDeclaration
    ;

parameterwidth
    : COLON expression
    ;

gotoStatement
    : GOTO label
    ;

inheritedStatement
    : INHERITED (identifier (LPAREN parameterList RPAREN)?)?
    ;

emptyStatement_
    :
    ;

empty_
    :
    /* empty */
    ;

structuredStatement
    : compoundStatement
    | conditionalStatement
    | repetetiveStatement
    | withStatement
    | tryExceptStatement
    | tryFinallyStatement
    ;

compoundStatement
    : BEGIN statements END
    ;

statements
    : statement (SEMI statement)*
    ;

conditionalStatement
    : ifStatement
    | caseStatement
    ;

ifStatement
    : IF expression THEN statement (ELSE statement)? SEMI?
    ;

caseStatement
    : CASE expression OF caseListElement (SEMI caseListElement)* (SEMI ELSE statements)? SEMI? END
    ;

caseListElement
    : constList COLON statement
    ;

repetetiveStatement
    : whileStatement
    | repeatStatement
    | forStatement
    ;

whileStatement
    : WHILE expression DO statement
    ;

repeatStatement
    : REPEAT statements UNTIL expression
    ;

forStatement
    : FOR VAR? identifier ASSIGN forList DO statement
    | FOR VAR? identifier IN expression DO statement
    ;

forList
    : initialValue (TO | DOWNTO) finalValue
    ;

initialValue
    : expression
    ;

finalValue
    : expression
    ;

withStatement
    : WITH withStatementVariableList DO statement
    | WITH expression DO statement
    ;

tryExceptStatement
    : TRY statements EXCEPT statements END
    | TRY statements EXCEPT (exceptionCase SEMI)+ exceptionElse? END
    ;

exceptionCase
    : ON identifier COLON typeIdentifier DO statements
    ;

exceptionElse
    : ELSE statements
    ;

tryFinallyStatement
    : TRY statements FINALLY statements END
    ;

withStatementVariableList
    : variableDesignator (COMMA variableDesignator)*
    ;

attributeSection
    : LBRACK attributeList RBRACK
    ;

attributeList
    : attributeItem (COMMA attributeItem)*
    ;

attributeItem
    : identifier (LPAREN parameterList RPAREN)?
    ;



AND
    : 'AND'
    ;

ARRAY
    : 'ARRAY'
    ;

BEGIN
    : 'BEGIN'
    ;

BOOLEAN
    : 'BOOLEAN'
    ;

CASE
    : 'CASE'
    ;

CHAR
    : 'CHAR'
    ;

CHR
    : 'CHR'
    ;

CONST
    : 'CONST'
    ;

DEPRECATED
    : 'DEPRECATED'
    ;

DIV
    : 'DIV'
    ;

DO
    : 'DO'
    ;

DOWNTO
    : 'DOWNTO'
    ;

ELSE
    : 'ELSE'
    ;

END
    : 'END'
    ;

FILE
    : 'FILE'
    ;

FOR
    : 'FOR'
    ;

FUNCTION
    : 'FUNCTION'
    ;

GOTO
    : 'GOTO'
    ;

HELPER
    : 'HELPER'
    ;

IF
    : 'IF'
    ;

IN
    : 'IN'
    ;

INTEGER
    : 'INTEGER'
    ;

LABEL
    : 'LABEL'
    ;

MOD
    : 'MOD'
    ;

NIL
    : 'NIL'
    ;

NOT
    : 'NOT'
    ;

OF
    : 'OF'
    ;

OR
    : 'OR'
    ;

PACKED
    : 'PACKED'
    ;

PROCEDURE
    : 'PROCEDURE'
    ;

PROGRAM
    : 'PROGRAM'
    ;

REAL
    : 'REAL'
    ;

RECORD
    : 'RECORD'
    ;

REFERENCE
    : 'REFERENCE'
    ;

REPEAT
    : 'REPEAT'
    ;

SET
    : 'SET'
    ;

THEN
    : 'THEN'
    ;

TO
    : 'TO'
    ;

TYPE
    : 'TYPE'
    ;

UNTIL
    : 'UNTIL'
    ;

VAR
    : 'VAR'
    ;

WHILE
    : 'WHILE'
    ;

WITH
    : 'WITH'
    ;

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

STAR
    : '*'
    ;

SLASH
    : '/'
    ;

ASSIGN
    : ':='
    ;

COMMA
    : ','
    ;

SEMI
    : ';'
    ;

COLON
    : ':'
    ;

EQUAL
    : '='
    ;

NOT_EQUAL
    : '<>'
    ;

LT
    : '<'
    ;

LE
    : '<='
    ;

GE
    : '>='
    ;

GT
    : '>'
    ;

LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

LBRACK
    : '['
    ;

LBRACK2
    : '(.'
    ;

RBRACK
    : ']'
    ;

RBRACK2
    : '.)'
    ;

DEREFERENCE
    : '^'
    ;

AT
    : '@'
    ;

DOT
    : '.'
    ;

DOTDOT
    : '..'
    ;

LCURLY
    : '{'
    ;

RCURLY
    : '}'
    ;

UNIT
    : 'UNIT'
    ;

INTERFACE
    : 'INTERFACE'
    ;

USES
    : 'USES'
    ;

STRING
    : 'STRING'
    ;

IMPLEMENTATION
    : 'IMPLEMENTATION'
    ;

TRUE
    : 'TRUE'
    ;

FALSE
    : 'FALSE'
    ;

CLASS
    : 'CLASS'
    ;

PRIVATE
    : 'PRIVATE'
    ;

PROTECTED
    : 'PROTECTED'
    ;

PUBLIC
    : 'PUBLIC'
    ;

PUBLISHED
    : 'PUBLISHED'
    ;

STRICT
    : 'STRICT'
    ;

OUT
    : 'OUT'
    ;

PROPERTY
    : 'PROPERTY'
    ;

READ
    : 'READ'
    ;

WRITE
    : 'WRITE'
    ;

DEFAULT
    : 'DEFAULT'
    ;

INDEX
    : 'INDEX'
    ;

AS
    : 'AS'
    ;

TRY
    : 'TRY'
    ;

FINALLY
    : 'FINALLY'
    ;

EXCEPT
    : 'EXCEPT'
    ;

INITIALIZATION
    : 'INITIALIZATION'
    ;

FINALIZATION
    : 'FINALIZATION'
    ;

OBJECT
    : 'OBJECT'
    ;

INHERITED
    : 'INHERITED'
    ;

ABSTRACT
    : 'ABSTRACT'
    ;

REINTRODUCE
    : 'REINTRODUCE'
    ;

VIRTUAL
    : 'VIRTUAL'
    ;

OVERRIDE
    : 'OVERRIDE'
    ;

OVERLOAD
    : 'OVERLOAD'
    ;

INLINE
    : 'INLINE'
    ;

CDECL
    : 'CDECL'
    ;

STDCALL
    : 'stdcall'
    ;

STATIC
    : 'STATIC'
    ;

CONSTRUCTOR
    : 'CONSTRUCTOR'
    ;

DESTRUCTOR
    : 'DESTRUCTOR'
    ;

RESOURCESTRING
    : 'resourcestring'
    ;

FORWARD
    : 'FORWARD'
    ;

RAISE
    : 'RAISE'
    ;

SHR
    : 'SHR'
    ;

SHL
    : 'SHL'
    ;

XOR
    : 'XOR'
    ;

CARDINAL
    : 'Cardinal'
    ;

LONGBOOL
    : 'LONGBOOL'
    ;

LONGINT
    : 'LONGINT'
    ;

LONGWORD
    : 'LONGWORD'
    ;

WORD
    : 'WORD'
    ;

OPERATOR
    : 'operator'
    ;

ON
    : 'ON'
    ;

BYTE
    : 'BYTE'
    ;

SHORTINT
    : 'SHORTINT'
    ;

SMALLINT
    : 'SMALLINT'
    ;

INT64
    : 'INT64'
    ;

UINT64
    : 'UINT64'
    ;

SINGLE
    : 'SINGLE'
    ;

DOUBLE
    : 'DOUBLE'
    ;

EXTENDED
    : 'EXTENDED'
    ;

COMP
    : 'COMP'
    ;

CURRENCY
    : 'CURRENCY'
    ;

ANSICHAR
    : 'ANSICHAR'
    ;

WIDECHAR
    : 'WIDECHAR'
    ;

ANSISTRING
    : 'ANSISTRING'
    ;

WIDESTRING
    : 'WIDESTRING'
    ;

UNICODESTRING
    : 'UNICODESTRING'
    ;

RAWBYTESTRING
    : 'RAWBYTESTRING'
    ;

UTF8STRING
    : 'UTF8STRING'
    ;

VARIANT
    : 'VARIANT'
    ;

OLEVARIANT
    : 'OLEVARIANT'
    ;

POINTER
    : 'POINTER'
    ;

PCHAR
    : 'PCHAR'
    ;

PANSICHAR
    : 'PANSICHAR'
    ;

PWIDECHAR
    : 'PWIDECHAR'
    ;

PUNICODECHAR
    : 'PUNICODECHAR'
    ;

THANDLE
    : 'THANDLE'
    ;

HWND
    : 'HWND'
    ;

HDC
    : 'HDC'
    ;

HICON
    : 'HICON'
    ;

HBITMAP
    : 'HBITMAP'
    ;

HMENU
    : 'HMENU'
    ;

HINSTANCE
    : 'HINSTANCE'
    ;

HMODULE
    : 'HMODULE'
    ;

HKEY
    : 'HKEY'
    ;

DWORD
    : 'DWORD'
    ;

QWORD
    : 'QWORD'
    ;

NATIVEINT
    : 'NATIVEINT'
    ;

NATIVEUINT
    : 'NATIVEUINT'
    ;

CODEPAGE
    : 'CODEPAGE'
    ;

TGUID
    : 'TGUID'
    ;

PGUID
    : 'PGUID'
    ;

TEXTFILE
    : 'TEXTFILE'
    ;

TEXT
    : 'TEXT'
    ;

SHORTSTRING
    : 'SHORTSTRING'
    ;

OPENSTRING
    : 'OPENSTRING'
    ;

PLATFORM
    : 'PLATFORM'
    ;

ALIGN
    : 'ALIGN'
    ;

fragment WHITESPACE
    : [ \t\r\n]+
    ;

WS
    : [ \t\r\n]+ -> skip
    ;

COMMENT_1
    : '(*' .*? '*)' -> skip
    ;

COMMENT_2
    : '{' .*? '}' -> skip
    ;

COMMENT_3
    : '//' ~[\r\n]* -> skip
    ;

IDENT
    : ('A' .. 'Z' | '_') ('A' .. 'Z' | '0' .. '9' | '_')*
    ;

HEX_LITERAL
    : '$' ('A' .. 'F' | '0' .. '9')+
    ;

STRING_LITERAL
    : '\'' ('\'\'' | ~ ('\''))* '\''
    ;

STRING_CROSSHATCH_LITERAL
    : '#' ([0-9]+ | HEX_LITERAL)
    ;

NUM_INT
    : ('0' .. '9')+
    ;

NUM_REAL
    : ('0' .. '9')+ (('.' ('0' .. '9')+ (EXPONENT)?)? | EXPONENT)
    ;

fragment EXPONENT
    : ('E') ('+' | '-')? ('0' .. '9')+
    ;

UTF8BOM
    : '\uFEFF' -> skip
    ;

fragment HEX_CHAR
    : ('A' .. 'F' | '0' .. '9')
    ;

fragment HEX_CHAR_SEQ12
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;

fragment HEX_CHAR_SEQ8
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;

fragment HEX_CHAR_SEQ4
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;