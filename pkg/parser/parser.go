package parser

import (
	"fmt"

	"github.com/tjgurwara99/monkey-go/pkg/ast"
	"github.com/tjgurwara99/monkey-go/pkg/token"
)

type parser struct {
	tokens chan token.Token

	current token.Token
	peek    token.Token

	errors []error
}

func New(t chan token.Token) *parser {
	return &parser{
		tokens: t,
		errors: []error{},
	}
}

func (p *parser) Errors() []error {
	return p.errors
}

func (p *parser) nextToken() {
	p.current = p.peek
	p.peek = <-p.tokens
}

func (p *parser) ParseProgram() *ast.Program {
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

func (p *parser) ParseStatement() ast.Statement {
	switch p.current.Type {
	case token.LET:
		return p.ParseLetStatement()
	default:
		return nil
	}
}

func (p *parser) ParseLetStatement() *ast.LetStatement {
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

// func (p *parser) ParseExpression() ast.Expression {}

func (p *parser) currentIs(t token.TokenType) bool {
	return p.current.Type == t
}

func (p *parser) peekTokenIs(t token.TokenType) bool {
	return p.peek.Type == t
}

func (p *parser) peekError(t token.TokenType) {
	msg := fmt.Errorf("expected next token to be %s, got %s instead", t, p.peek.Type)
	p.errors = append(p.errors, msg)
}

func (p *parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}
