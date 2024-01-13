package token

// Token is a lexical token of the Go programming language.
type TokenType string

type Token struct {
	Type     TokenType
	Value    string
	FileName string
	Line     int
	Column   int
}

// const tokens = []string{
// 	// Special tokens
// 	ILLEGAL: "ILLEGAL",
// 	EOF     : "EOF",
// 	// Identifiers + literals
// 	IDENT : "IDENT", // add, foobar, x, y, ...
// 	INT   : "INT",   // 1234567890
// 	// Operators
// 	ASSIGN   : "=",
// 	PLUS     : "+",
// 	MINUS    : "-",
// 	BANG     : "!",
// 	ASTERISK : "*",
// 	COMMA     : ",",
// 	SEMICOLON : ";",
// 	LPAREN   : "(",
// 	RPAREN   : ")",
// 	LBRACE   : "{",
// 	RBRACE   : "}",
// 	LBRACKET : "[",
// 	RBRACKET : "]",
// 	// Keywords
// 	FUNCTION : "FUNCTION",
// 	LET      : "LET",
// }

const (
	// Special tokens
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	EQ        = "=="
	NOT_EQ    = "!="
	ARROW     = "=>"
	PLUS      = "+"
	PLUS_EQ   = "+="
	MINUS     = "-"
	MINUS_EQ  = "-="
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	DOT       = "."
	SEMICOLON = ";"

	UNION       = "|"
	ARROW_SMALL = "->" //TODO: find better name

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	LT       = "<"
	GT       = ">"
	PIPE     = ">>"

	MODULO = "%"
	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	AND  = "AND"
	OR   = "OR"
	NOT  = "NOT"
	IF   = "IF"
	ELSE = "ELSE"

	RETURN = "RETURN"
	MATCH  = "MATCH"
	SELF   = "SELF"
	ENUM   = "ENUM"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	"func":  FUNCTION,
	"let":   LET,
	"const": CONST,
	"true":  TRUE,
	"false": FALSE,

	"and":  AND,
	"or":   OR,
	"not":  NOT,
	"if":   IF,
	"else": ELSE,

	"return": RETURN,
	"match":  MATCH,
	"self":   SELF,
	"enum":   ENUM,
}
