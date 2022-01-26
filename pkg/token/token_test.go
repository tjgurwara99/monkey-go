package token_test

import (
	"testing"

	"github.com/tjgurwara99/monkey-go/pkg/token"
)

func TestIsLiteral(t *testing.T) {
	testCases := []struct {
		tokenType token.TokenType
		expected  bool
	}{
		{token.ILLEGAL, false},
		{token.EOF, false},
		{token.IDENT, true},
		{token.INT, true},
		{token.ASSIGN, false},
		{token.PLUS, false},
		{token.MINUS, false},
		{token.COMMA, false},
		{token.SEMICOLON, false},
		{token.LPAREN, false},
		{token.RPAREN, false},
		{token.LBRACE, false},
		{token.RBRACE, false},
		{token.FUNCTION, false},
		{token.LET, false},
		{token.TRUE, false},
		{token.FALSE, false},
		{token.IF, false},
		{token.ELSE, false},
		{token.RETURN, false},
	}

	for _, tc := range testCases {
		if tc.expected != tc.tokenType.IsLiteral() {
			t.Errorf("%s is literal: %t", tc.tokenType, tc.expected)
		}
	}
}

func TestIsOperator(t *testing.T) {
	testCases := []struct {
		tokenType token.TokenType
		expected  bool
	}{
		{token.ILLEGAL, false},
		{token.EOF, false},
		{token.IDENT, false},
		{token.INT, false},
		{token.ASSIGN, true},
		{token.PLUS, true},
		{token.MINUS, true},
		{token.COMMA, true},
		{token.SEMICOLON, true},
		{token.LPAREN, true},
		{token.RPAREN, true},
		{token.LBRACE, true},
		{token.RBRACE, true},
		{token.FUNCTION, false},
		{token.LET, false},
		{token.TRUE, false},
		{token.FALSE, false},
		{token.IF, false},
		{token.ELSE, false},
		{token.RETURN, false},
	}

	for _, tc := range testCases {
		if tc.expected != tc.tokenType.IsOperator() {
			t.Errorf("%s is literal: %t", tc.tokenType, tc.expected)
		}
	}
}

func TestTokenType_IsKeyword(t *testing.T) {
	testCases := []struct {
		tokenType token.TokenType
		expected  bool
	}{
		{token.ILLEGAL, false},
		{token.EOF, false},
		{token.IDENT, false},
		{token.INT, false},
		{token.ASSIGN, false},
		{token.PLUS, false},
		{token.MINUS, false},
		{token.COMMA, false},
		{token.SEMICOLON, false},
		{token.LPAREN, false},
		{token.RPAREN, false},
		{token.LBRACE, false},
		{token.RBRACE, false},
		{token.FUNCTION, true},
		{token.LET, true},
		{token.TRUE, true},
		{token.FALSE, true},
		{token.IF, true},
		{token.ELSE, true},
		{token.RETURN, true},
	}

	for _, tc := range testCases {
		if tc.expected != tc.tokenType.IsKeyword() {
			t.Errorf("%s is literal: %t", tc.tokenType, tc.expected)
		}
	}
}

func TestIsKeyword(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected bool
	}{
		{"", false},
		{"a", false},
		{"fn", true},
		{"let", true},
		{"true", true},
		{"false", true},
		{"if", true},
		{"else", true},
		{"return", true},
	}

	for _, tc := range testCases {
		if tc.expected != token.IsKeyword(tc.keyword) {
			t.Errorf("%s is keyword: %t", tc.keyword, tc.expected)
		}
	}
}

func TestIsIdentifier(t *testing.T) {
	testCases := []struct {
		identifier string
		expected   bool
	}{
		{"", false},
		{"a", true},
		{"fn", false},
		{"let", false},
		{"true", false},
		{"false", false},
		{"if", false},
		{"else", false},
		{"return", false},
		{"_abc", true},
		{"a-bc", false},
	}

	for _, tc := range testCases {
		if tc.expected != token.IsIdentifier(tc.identifier) {
			t.Errorf("%s is identifier: %t", tc.identifier, tc.expected)
		}
	}
}
