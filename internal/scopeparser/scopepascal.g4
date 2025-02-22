grammar scopepascal;

options {
    caseInsensitive = true;
}


// Parser rules

source
    : program
    | unit
    ;

program
    : PROGRAM identifier (LPAREN identifierList RPAREN)? SEMI 
        (USES identifierList)? implementationBlock 
        blockStatement DOT EOF
    ;

unit
    : UNIT identifier SEMI 
      interfaceSection
      implementationSection?
      initializationSection?
      finalizationSection?
      END DOT EOF
    ;

interfaceSection
    : INTERFACE unitList? interfaceBlock
    ;

implementationSection
    : IMPLEMENTATION unitList? implementationBlock
    ;

initializationSection
    : INITIALIZATION blockStatement
    ;

finalizationSection
    : FINALIZATION blockStatement
    ;


interfaceBlock
    : (
        typeSection
        | labelDeclaration
        | constSection
        | resourceSection
        | varSection
        | functionOrProcedureDeclaration
    )*
    ;

implementationBlock
    : (
        labelDeclaration
        | constSection
        | resourceSection
        | typeSection
        | varSection
        | functionOrProcedure
    )*
    ;


unitList
    : USES identifierList SEMI
    ;

labelDeclaration
    : LABEL number (COMMA number)* SEMI
    ;

constSection
    : CONST (identifierList (COLON identifier)? EQUAL expression SEMI)+
    ;

resourceSection
    : RESOURCESTRING (identifier EQUAL expression SEMI)+
    ;

typeSection
    : TYPE typeBlock
    ;

typeBlock
    : (identifier EQUAL type SEMI)+
    ;

type
    :classType
    | recordType
    | arrayType
    | pointerType
    | setType
    | fileType
    | scalarType
    | subrangeType
    | functionDeclaration
    | procedureDeclaration
    ;

accessSpecifier
    : PRIVATE
    | STRICTPRIVATE
    | PROTECTED
    | STRICTPROTECTED
    | PUBLIC
    | PUBLISHED
    ;

procedureOrFunctionModifiers: (
		SEMI (ABSTRACT | VIRTUAL | OVERRIDE | REINTRODUCE | OVERLOAD | INLINE | STDCALL | CDECL | STATIC)
	)*;


classForwardDeclaration
    : CLASS SEMI
    ;

classType
    : CLASS (LPAREN identifier (COMMA identifierList)? RPAREN)? ABSTRACT? blockDeclaration* (accessSpecifier blockDeclaration)* END
    ;

recordType
    : PACKED? RECORD blockDeclaration? (accessSpecifier blockDeclaration)* recordVariantDeclaration* END
    ;

propertyDeclaration
    : PROPERTY identifier propertyIndexParameters? COLON identifier propertyReadDeclaration? propertyWriteDeclaration? propertyDefaultValueDeclaration? propertyIndexDeclaration?
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
    : INDEX number
    ;

propertyIndexParameters
    : LBRACK propertyIndexParametersList RBRACK
    ;

propertyIndexParametersList
    : identifierList COLON expression (SEMI identifierList COLON expression)*
    ;

arrayType
    : ARRAY LBRACK identifierList RBRACK OF (type|identifier)
    | ARRAY LBRACK expression DOUBLEDOT expression OF (type|identifier)
    | ARRAY OF CONST
    | ARRAY OF (type|identifier)
    ;

pointerType
    : DEREFERENCE identifier
    ;

setType
    : SET OF identifier
    ;

fileType
    : FILE OF identifier
    | FILE
    ;

scalarType
    : LPAREN expressionList RPAREN
    ;

subrangeType
    : expression DOUBLEDOT expression
    ;

blockDeclaration
    : typeSection
    | constSection
    | varDeclaration SEMI?
    | functionOrProcedureDeclaration
    | propertyDeclaration
    ;

functionDeclaration
    : CLASS? FUNCTION IDENT paramsDeclaration? COLON IDENT OFOBJECT? procedureOrFunctionModifiers SEMI
    ;

procedureDeclaration
    : CLASS? PROCEDURE IDENT paramsDeclaration? OFOBJECT? procedureOrFunctionModifiers SEMI
    ;

functionOrProcedureDeclaration
    : functionDeclaration
    | procedureDeclaration
    ;    

functionOrProcedure
    : functionOrProcedureDeclaration (functionOrProcedure | varSection)* blockStatement SEMI
    | functionOrProcedureDeclaration FORWARD SEMI
    ;

blockStatement
    : BEGIN (statement | statementError)* END
    ;


recordVariantDeclaration
    : CASE identifier(COLON identifier)? OF recordVariant (SEMI recordVariant)* SEMI?
    ;

recordVariant
    : identifierList COLON LPAREN varDeclaration+ RPAREN
    ;


// Error-handling nodes (to tolerate broken syntax)
statementError
    : ~('end' | 'begin' | ';')+
    ; // Captures unknown text to skip over it



statement
    : varSection
    | inlinedVarDeclaration
    | statementError
    ;


identifier
    : IDENT (DOT IDENT)*
    ;

identifierList
    : identifier (COMMA identifier)*
    ;

expressionList
    : expression (COMMA expression)*
    ;


paramsDeclaration
    : LPAREN paramsDeclarationSection (SEMI paramsDeclarationSection)* RPAREN
    ;

paramsDeclarationSection
    : paramSpecifier? identifierList (COLON identifier)? (EQUAL expression)?
    | FUNCTION paramsDeclaration COLON identifier
    | PROCEDURE paramsDeclaration
    ;

paramSpecifier
    : VAR
    | CONST
    | OUT
    ;

varSection
    : VAR (varDeclaration SEMI)+
    ;

varDeclaration
    : identifierList (COLON identifier)? (EQUAL expression)?
    ;

inlinedVarDeclaration
    : VAR identifierList (COLON identifier)? (ASSIGN expression)? SEMI
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

additiveoperator
    : PLUS
    | MINUS
    | OR
    ;

multiplicativeoperator
    : STAR
    | SLASH
    | DIV
    | MOD
    | AND
    | SHR
    | SHL
    ;

operator
    : relationaloperator
    | additiveoperator
    | multiplicativeoperator
    ;

string
    : (STRING_LITERAL | STRING_CROSSHATCH_LITERAL)+
    ;

number
    : NUM_INT
    | NUM_REAL
    ;


expression
    : term (operator term)* | functionExpression | errorExpression ;

term
    : AT? IDENT  
    | number  
    | string  
    | functionDesignator
    | '(' expression ')' // Allow grouping
    | errorExpression // Capture malformed terms
    ;


// Lambda expression (Pascal anonymous function)
functionExpression
    : FUNCTION paramsDeclaration? COLON varSection? blockStatement
    ;

// Lambda expression (Pascal anonymous procedure)
procedureExpression
    : PROCEDURE paramsDeclaration? varSection? blockStatement
    ;

// Function call inside expression
functionDesignator
    : identifier (LT identifier GT)? (LPAREN expressionList? RPAREN)?
    ;

// Error-handling nodes -  Captures anything unexpected
errorExpression
    : ~(SEMI | RPAREN | COMMA | END | BEGIN)+
    ;


// Lexer rules
BEGIN
    : 'begin'
    ;

END
    : 'end'
    ;

FUNCTION
    : 'function'
    ;

PROCEDURE
    : 'procedure'
    ;

VAR
    : 'var'
    ;

CONST
    : 'const'
    ;

OUT
    : 'out'
    ;

DOT
    : '.'
    ;

DOUBLEDOT
    : '..'
    ;

SEMI
    : ';'
    ;

COMMA
    : ','
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

GT
    : '>'
    ;

GE
    : '>='
    ;

IN
    : 'in'
    ;
    
STAR
    : '*'
    ;

SLASH
    : '/'
    ;

DIV
    : 'div'
    ;

MOD
    : 'mod'
    ;

AND
    : 'and'
    ;

OR
    : 'or'
    ;

NOT
    : 'not'
    ;

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

SHR
    : 'shr'
    ;

SHL
    : 'shl'
    ;

LBRACK
    : '['
    ;

RBRACK
    : ']'
    ;

LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

ASSIGN
    : ':='
    ;

LABEL
    : 'label'
    ;

TYPE
    : 'type'
    ;

RESOURCESTRING
    : 'resourcestring'
    ;

PROGRAM
    : 'program'
    ;

UNIT
    : 'unit'
    ;

INTERFACE
    : 'interface'
    ;   

IMPLEMENTATION
    : 'implementation'
    ;

INITIALIZATION  
    : 'initialization'
    ;   

FINALIZATION
    : 'finalization'    
    ;

USES
    : 'uses'
    ;

FORWARD
    : 'forward'
    ;

CLASS
    : 'class'
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

STRICTPRIVATE
    : 'STRICT' WHITESPACE 'PRIVATE'
    ;

STRICTPROTECTED
    : 'STRICT' WHITESPACE 'PROTECTED'
    ;

OFOBJECT
    : 'OF' WHITESPACE 'OBJECT'
    ;

VIRTUAL
    : 'VIRTUAL'
    ;

OVERRIDE
    : 'OVERRIDE'
    ;   

REINTRODUCE
    : 'REINTRODUCE'
    ;   

OVERLOAD
    : 'OVERLOAD'
    ;   

INLINE  
    : 'INLINE'
    ;   

STDCALL
    : 'STDCALL'
    ;

CDECL   
    : 'CDECL'
    ;

STATIC
    : 'STATIC'
    ;

PACKED
    : 'PACKED'
    ;

RECORD
    : 'RECORD'
    ;    

PROPERTY
    : 'PROPERTY'
    ;

DEFAULT
    : 'DEFAULT'
    ;

fragment WHITESPACE : [ \t\r\n]+ ;

ABSTRACT
    : 'abstract'
    ;

READ
    : 'read'
    ;

WRITE
    : 'write'
    ;

INDEX
    : 'index'
    ;

ARRAY
    : 'array'
    ;

OF
    : 'of'
    ;

DEREFERENCE
    : '^'
    ;

AT
    : '@'
    ;

SET
    : 'set'
    ;

FILE
    : 'file'
    ;

CASE
    : 'case'
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
    : '#' [0-9]+
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

UTF8BOM
    : '\uFEFF' -> skip
    ;
