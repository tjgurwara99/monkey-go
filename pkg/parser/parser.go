package parser

import (
	"github.com/tjgurwara99/monkey-go/pkg/ast"
	"github.com/tjgurwara99/monkey-go/pkg/lexer"
	"github.com/tjgurwara99/monkey-go/pkg/token"
)

type Parser struct {
	l *lexer.Lexer

	current token.Token
	peek    token.Token
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = <-p.l.Tokens
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for !p.currentIs(token.EOF) {
		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.current.Type {
	case token.LET:
		return p.ParseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.current}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.current,
		Value: p.current.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currentIs(token.SEMICOLON) {
		p.nextToken()
	}

	// stmt.Value = p.ParseExpression();
	return stmt
}

// func (p *Parser) ParseExpression() ast.Expression {}

func (p *Parser) currentIs(t token.TokenType) bool {
	return p.current.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peek.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}
