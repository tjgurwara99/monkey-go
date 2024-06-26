package lexer_test

import (
	"testing"

	"github.com/tjgurwara99/monkey-go/pkg/lexer"
	"github.com/tjgurwara99/monkey-go/pkg/token"
)

func TestNextOperatorsDelimiters(t *testing.T) {
	input := `=+-(){},;!/*<> let true false if else
return
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedPos     int
		expectedCol     int
	}{
		{token.ASSIGN, "=", 1, 0, 0},
		{token.PLUS, "+", 1, 1, 1},
		{token.MINUS, "-", 1, 2, 2},
		{token.LPAREN, "(", 1, 3, 3},
		{token.RPAREN, ")", 1, 4, 4},
		{token.LBRACE, "{", 1, 5, 5},
		{token.RBRACE, "}", 1, 6, 6},
		{token.COMMA, ",", 1, 7, 7},
		{token.SEMICOLON, ";", 1, 8, 8},
		{token.BANG, "!", 1, 9, 9},
		{token.SLASH, "/", 1, 10, 10},
		{token.ASTERISK, "*", 1, 11, 11},
		{token.LT, "<", 1, 12, 12},
		{token.GT, ">", 1, 13, 13},
		{token.LET, "let", 1, 15, 15},
		{token.TRUE, "true", 1, 19, 19},
		{token.FALSE, "false", 1, 24, 24},
		{token.IF, "if", 1, 30, 30},
		{token.ELSE, "else", 1, 33, 33},
		{token.RETURN, "return", 2, 38, 0},
		{token.EOF, "", 2, 46, 8},
	}

	lexed := lexer.New(input)
	index := 0
	for token := range lexed.Tokens {
		if token.Type != tests[index].expectedType {
			t.Fatalf("expected token type %q, got %q", tests[index].expectedType, token.Type)
		}
		if token.Literal != tests[index].expectedLiteral {
			t.Fatalf("expected token literal %q, got %q", tests[index].expectedLiteral, token.Literal)
		}
		if token.Line != tests[index].expectedLine {
			t.Fatalf("expected token line %d, got %d", tests[index].expectedLine, token.Line)
		}
		if token.Pos != tests[index].expectedPos {
			t.Fatalf("expected token pos %d, got %d", tests[index].expectedPos, token.Pos)
		}
		if token.Col != tests[index].expectedCol {
			t.Fatalf("expected token col %d, got %d", tests[index].expectedCol, token.Col)
		}
		index++
	}

}
