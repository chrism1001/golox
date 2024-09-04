package lexer

import (
	"golox/token"
)

type Lexer struct {
	input   string
	start   int
	current int
	line    int
	tokens  []token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.scanTokens()
	return l
}

func (l *Lexer) scanTokens() []token.Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, token.Token{Type: token.EOF, Literal: "", Line: l.line})
	return l.tokens
}

func (l *Lexer) scanToken() {
	ch := l.nextChar()
	switch ch {
	case '(':
		l.addToken(token.LEFT_PAREN, ch)
	case ')':
		l.addToken(token.RIGHT_PAREN, ch)
	case '{':
		l.addToken(token.LEFT_BRACE, ch)
	case '}':
		l.addToken(token.RIGHT_BRACE, ch)
	case ',':
		l.addToken(token.COMMA, ch)
	case '.':
		l.addToken(token.DOT, ch)
	case '-':
		l.addToken(token.MINUS, ch)
	case '+':
		l.addToken(token.PLUS, ch)
	case ';':
		l.addToken(token.SEMICOLON, ch)
	case '*':
		l.addToken(token.STAR, ch)
	case '!':
		if l.match('=') {
			l.addToken(token.BANG_EQUAL, ch)
		} else {
			l.addToken(token.BANG, ch)
		}
	case '=':
		if l.match('=') {
			l.addToken(token.EQUAL_EQUAL, ch)
		} else {
			l.addToken(token.EQUAL, ch)
		}
	case '<':
		if l.match('=') {
			l.addToken(token.LESS_EQUAL, ch)
		} else {
			l.addToken(token.LESS, ch)
		}
	case '>':
		if l.match('=') {
			l.addToken(token.GREATER_EQUAL, ch)
		} else {
			l.addToken(token.GREATER, ch)
		}
	default:
		l.addToken(token.ILLEGAL, ch)
	}
}

func (l *Lexer) addToken(tokenType token.TokenType, ch byte) {
	l.tokens = append(l.tokens, token.Token{Type: tokenType, Literal: string(ch), Line: l.line})
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.input)
}

func (l *Lexer) nextChar() byte {
	ch := l.input[l.current]
	l.current += 1
	return ch
}

func (l *Lexer) match(expected byte) bool {
	if l.isAtEnd() {
		return false
	}
	if l.input[l.current] != expected {
		return false
	}

	l.current += 1
	return true
}
