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
    : INTERFACE usesUnits? interfaceBlock
    ;

implementationSection
    : IMPLEMENTATION usesUnits? implementationBlock
    ;

initializationSection
    : INITIALIZATION blockStatement
    ;

finalizationSection
    : FINALIZATION blockStatement
    ;


interfaceBlock
    : (
        labelDeclaration
        | constDeclaration
        | resourceDeclaration
        | typeDeclaration
        | varDeclaration
        | functionOrProcedureDeclaration
    )*
    ;

implementationBlock
    : (
        labelDeclaration
        | constDeclaration
        | resourceDeclaration
        | typeDeclaration
        | varDeclaration
        | functionOrProcedure
    )*
    ;


usesUnits
    : USES identifierList SEMI
    ;

labelDeclaration
    : LABEL number (COMMA number)* SEMI
    ;

constDeclaration
    : CONST identifierList (COLON identifier)? EQUAL expression SEMI
    ;

resourceDeclaration
    : RESOURCESTRING (identifier EQUAL expression SEMI)+
    ;

typeDeclaration
    : TYPE (identifier EQUAL (
        classDeclaration
        | recordDeclaration
        | functionDeclaration
        | procedureDeclaration
        ) SEMI)+
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

classDeclaration
    : CLASS (LPAREN identifier (COMMA identifierList)? RPAREN)? ABSTRACT? structuredDeclarationSection? (accessSpecifier structuredDeclarationSection)* END SEMI
    ;

recordDeclaration
    : PACKED? RECORD identifier structuredDeclarationSection? (accessSpecifier structuredDeclarationSection)* END SEMI
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

structuredDeclarationSection
    : typeDeclaration
    | varDeclaration
    | constDeclaration
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
    : functionOrProcedureDeclaration (functionOrProcedure | varDeclaration)* blockStatement SEMI
    | functionOrProcedureDeclaration FORWARD SEMI
    ;

blockStatement
    : BEGIN (statement | statementError)* END
    ;

// Error-handling nodes (to tolerate broken syntax)
statementError
    : ~('end' | 'begin' | ';')+
    ; // Captures unknown text to skip over it



statement
    : varDeclaration
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

varDeclaration
    : VAR varDeclarationSection (SEMI varDeclarationSection)* SEMI
    ;

varDeclarationSection
    : identifierList ((COLON identifier) | (EQUAL expression))
    ;

inlinedVarDeclaration
    : VAR identifierList COLON identifier (EQUAL expression)? SEMI
    | VAR identifier ASSIGN expression SEMI
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

term: IDENT  
    | number  
    | string  
    | functionDesignator
    | '(' expression ')' // Allow grouping
    | errorExpression // Capture malformed terms
    ;


// Lambda expression (Pascal anonymous function)
functionExpression
    : FUNCTION paramsDeclaration? COLON varDeclaration? blockStatement
    ;

// Lambda expression (Pascal anonymous procedure)
procedureExpression
    : PROCEDURE paramsDeclaration? varDeclaration? blockStatement
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
