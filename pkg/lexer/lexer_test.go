package lexer_test

import (
	"testing"

	"github.com/tjgurwara99/monkey-go/pkg/lexer"
	"github.com/tjgurwara99/monkey-go/pkg/token"
)

func TestNextOperatorsDelimiters(t *testing.T) {
	input := "=+-(){},;"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedPos     int
		expectedCol     int
	}{
		{token.ASSIGN, "=", 1, 1, 1},
		{token.PLUS, "+", 1, 2, 2},
		{token.MINUS, "-", 1, 3, 3},
		{token.LPAREN, "(", 1, 4, 4},
		{token.RPAREN, ")", 1, 5, 5},
		{token.LBRACE, "{", 1, 6, 6},
		{token.RBRACE, "}", 1, 7, 7},
		{token.COMMA, ",", 1, 8, 8},
		{token.SEMICOLON, ";", 1, 9, 9},
		{token.EOF, "", 1, 10, 10},
	}

	lexed := lexer.Lex(input)
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
