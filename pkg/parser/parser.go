package parser

import (
	"fmt"

	"github.com/tjgurwara99/monkey-go/pkg/ast"
	"github.com/tjgurwara99/monkey-go/pkg/token"
)

type Parser struct {
	tokens chan token.Token

	current token.Token
	peek    token.Token

	errors []string
}

func New(t chan token.Token) *Parser {
	return &Parser{
		tokens: t,
		errors: []string{},
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = <-p.tokens
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

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}
