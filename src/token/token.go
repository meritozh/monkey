// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package token

const (
	// ILLEGAL illegal token
	ILLEGAL = "ILLEGAL"
	// EOF end of file
	EOF = "EOF"

	// IDENT identifier
	IDENT = "IDENT"
	// INT literal
	INT = "INT"

	// ASSIGN assign operation
	ASSIGN = "="
	// PLUS plus operation
	PLUS = "+"
	// MINUS minus operation
	MINUS = "-"
	// BANG boolean operation
	BANG = "!"
	// ASTERISK multiple operation
	ASTERISK = "*"
	// SLASH divide operation
	SLASH = "/"
	// EQ equal operation
	EQ = "=="
	// NOTEQ not equal operation
	NOTEQ = "!="
	// LT less than operation
	LT = "<"
	// GT great than operation
	GT = ">"

	// COMMA comma
	COMMA = ","
	// SEMICOLON semicolon
	SEMICOLON = ";"

	// LPAREN left parenthesis
	LPAREN = "("
	// RPAREN right parenthesis
	RPAREN = ")"

	// LBRACE left brace
	LBRACE = "{"
	// RBRACE right brace
	RBRACE = "}"

	// FUNCTION function keyword
	FUNCTION = "FUNCTION"
	// LET let keyword
	LET = "LET"
	// TRUE true keyword
	TRUE = "TRUE"
	// FALSE false keyword
	FALSE = "FALSE"
	// IF if keyword
	IF = "IF"
	// ELSE else keyword
	ELSE = "ELSE"
	// RETURN return keyword
	RETURN = "RETURN"

	// STRING string literal
	STRING = "STRING"
	// LBRACKET left bracket
	LBRACKET = "["
	// RBRACKET right bracket
	RBRACKET = "]"
)

// Type string alias
type Type string

// Token core token structure
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent look up current string token type
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
