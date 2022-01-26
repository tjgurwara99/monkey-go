package token

import (
	"strconv"
	"unicode"
)

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Pos     int
	Col     int
}

const (
	// special tokens
	ILLEGAL TokenType = iota
	EOF

	// identifiers and literals
	literal_beggining
	IDENT // main, foobar, x, y, ...
	INT   // 12345
	literal_end

	// Operators
	operator_beggining
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT // <
	GT // >

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;

	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }
	operator_end

	// Keywords
	keyword_beggining
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT: "IDENT",
	INT:   "INT",

	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",

	LT: "<",
	GT: ">",

	COMMA:     ",",
	SEMICOLON: ";",

	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	FUNCTION: "fn",
	LET:      "let",
	TRUE:     "true",
	FALSE:    "false",
	IF:       "if",
	ELSE:     "else",
	RETURN:   "return",
}

func (t TokenType) String() string {
	if t == literal_beggining || t == literal_end || t == operator_beggining || t == operator_end || t == keyword_beggining || t == keyword_end {
		return "token(" + strconv.Itoa(int(t)) + ")"
	}
	if 0 <= t && t < keyword_end {
		return tokens[t]
	}
	return "token(" + strconv.Itoa(int(t)) + ")"
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)
	for i := keyword_beggining + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func (t TokenType) IsLiteral() bool {
	return literal_beggining < t && t < literal_end
}

func (t TokenType) IsOperator() bool {
	return operator_beggining < t && t < operator_end
}

func (t TokenType) IsKeyword() bool {
	return keyword_beggining < t && t < keyword_end
}

func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

func IsIdentifier(name string) bool {
	for i, c := range name {
		if !unicode.IsLetter(c) && (i == 0 || !unicode.IsDigit(c)) && c != '_' {
			return false
		}
	}
	return name != "" && !IsKeyword(name)
}

func LookupIdent(name string) TokenType {
	if tok, ok := keywords[name]; ok {
		return tok
	}
	return IDENT
}
