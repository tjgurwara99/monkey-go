package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/tjgurwara99/monkey-go/pkg/token"
)

type Lexer struct {
	start   int
	pos     int
	width   int
	line    int
	col     int
	input   string
	Tokens  chan token.Token
	prevCol []int
}

const eof = -1

func Lex(input string) *Lexer {
	l := &Lexer{
		input:  input,
		Tokens: make(chan token.Token),
		line:   1,
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.Tokens)
}

func (l *Lexer) emit(t token.TokenType) {
	if t == token.EOF {
		l.Tokens <- token.Token{
			Type:    t,
			Literal: "",
			Line:    l.line,
			Pos:     l.pos,
			Col:     l.col + (l.pos - l.start),
		}
		return
	}
	l.Tokens <- token.Token{
		Type:    t,
		Literal: l.input[l.start:l.pos],
		Line:    l.line,
		Pos:     l.start,
		Col:     l.col - (l.pos - l.start),
	}
	l.start = l.pos
}

func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		if l.col == 0 {
			l.col = l.prevCol[len(l.prevCol)-1]
			l.line -= 1
		}
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.prevCol = append(l.prevCol, l.col)
	l.col += l.width
	l.pos += l.width
	if r == '\n' {
		l.line++
		l.col = 0
	}
	return r
}

// peek is usually supposed to be used for double character operators.
// Current syntax of Monkey doesn't have these double character operators yet.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *Lexer) backup() {
	l.pos -= l.width
	l.col = l.prevCol[len(l.prevCol)-1]
	l.prevCol = l.prevCol[:len(l.prevCol)-1]
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

type stateFn func(*Lexer) stateFn

func (l *Lexer) doubleCharOperator(secondChar rune, failToken token.TokenType, passToken token.TokenType) token.TokenType {
	ch := l.peek()
	if ch == secondChar {
		l.next()
		return passToken
	}
	return failToken
}

func lexText(l *Lexer) stateFn {
	switch r := l.next(); {
	case r == '\n' || r == ' ' || r == '\t' || r == '\r':
		l.ignore()
	case r == '=':
		l.emit(l.doubleCharOperator('=', token.ASSIGN, token.EQ))
	case r == '+':
		l.emit(token.PLUS)
	case r == '-':
		l.emit(token.MINUS)
	case r == '!':
		l.emit(l.doubleCharOperator('=', token.BANG, token.NOT_LEQ))
	case r == '*':
		l.emit(token.ASTERISK)
	case r == '/':
		l.emit(token.SLASH)
	case r == '<':
		l.emit(l.doubleCharOperator('=', token.LT, token.LTEQ))
	case r == '>':
		l.emit(l.doubleCharOperator('=', token.GT, token.GTEQ))
	case r == ',':
		l.emit(token.COMMA)
	case r == ';':
		l.emit(token.SEMICOLON)
	case r == '(':
		l.emit(token.LPAREN)
	case r == ')':
		l.emit(token.RPAREN)
	case r == '{':
		l.emit(token.LBRACE)
	case r == '}':
		l.emit(token.RBRACE)
	case r == '|':
		l.emit(l.doubleCharOperator('|', token.OR, token.ILLEGAL))
	case r == '&':
		l.emit(l.doubleCharOperator('&', token.AND, token.ILLEGAL))
	case isIdent(r):
		l.backup()
		return lexIdent
	case r == eof:
		l.pos += 1
		l.col += 1
		l.emit(token.EOF)
		return nil
	}
	return lexText
}

func isIdent(r rune) bool {
	return r == '_' || unicode.IsDigit(r) || unicode.IsLetter(r)
}

func lexIdent(l *Lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isIdent(r):
			// absorb
		default:
			l.backup()
			l.emit(token.LookupIdent(l.input[l.start:l.pos]))
			break Loop
		}
	}
	return lexText
}
