package ast

import "token"

// Node base node
type Node interface {
	TokenLiteral() string
}

// Expression expression node
type Expression interface {
	Node
	expressionNode()
}

// Statement statement node
type Statement interface {
	Node
	statementNode()
}

// Program the top node
type Program struct {
	Statements []Statement
}

// TokenLiteral implement Node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement let statement
type LetStatement struct {
	// LET token
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral implement Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier represent identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral implement Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
