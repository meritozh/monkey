package parser

import (
	"ast"
	"lexer"
	"token"
)

// Parser parser structure
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New Parser constructor
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two token, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parse AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
