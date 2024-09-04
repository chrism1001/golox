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
	ch := l.advance()
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
			l.addToken(token.BANG_EQUAL, string(ch)+string(l.input[l.current-1]))
		} else {
			l.addToken(token.BANG, ch)
		}
	case '=':
		if l.match('=') {
			l.addToken(token.EQUAL_EQUAL, string(ch)+string(l.input[l.current-1]))
		} else {
			l.addToken(token.EQUAL, ch)
		}
	case '<':
		if l.match('=') {
			l.addToken(token.LESS_EQUAL, string(ch)+string(l.input[l.current-1]))
		} else {
			l.addToken(token.LESS, ch)
		}
	case '>':
		if l.match('=') {
			l.addToken(token.GREATER_EQUAL, string(ch)+string(l.input[l.current-1]))
		} else {
			l.addToken(token.GREATER, ch)
		}
	case '/':
		if l.match('/') {
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
		} else {
			l.addToken(token.SLASH, ch)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		l.line += 1
	case '"':
		l.readString()
	default:
		if isAlpha(ch) {
			l.readIdentifier()
		} else if isDigit(ch) {
			l.readNumber()
		} else {
			l.addToken(token.ILLEGAL, ch)
		}
	}
}

func (l *Lexer) addToken(tokenType token.TokenType, literal interface{}) {
	val := ""
	switch v := literal.(type) {
	case string:
		val = v
	case byte:
		val = string(v)
	}
	l.tokens = append(l.tokens, token.Token{Type: tokenType, Literal: val, Line: l.line})
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.input)
}

func (l *Lexer) advance() byte {
	ch := l.input[l.current]
	l.current += 1
	return ch
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.input[l.current]
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

func (l *Lexer) readIdentifier() {
	for isAlphaNumeric(l.peek()) {
		l.advance()
	}

	literal := l.input[l.start:l.current]
	tokenType := token.LookupIdent(literal)
	l.addToken(tokenType, literal)
}

func (l *Lexer) readString() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line += 1
		}
		l.advance()
	}

	if l.isAtEnd() {
		l.addToken(token.ILLEGAL, "ILLEGAL")
	}

	l.advance()
	value := l.input[l.start+1 : l.current-1]
	l.addToken(token.STRING, value)
}

func (l *Lexer) readNumber() {
	for isDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && isDigit(l.peekNext()) {
		l.advance()

		for isDigit(l.peek()) {
			l.advance()
		}
	}

	value := l.input[l.start:l.current]
	l.addToken(token.NUMBER, value)
}

func (l *Lexer) peekNext() byte {
	if l.current+1 >= len(l.input) {
		return 0
	}
	return l.input[l.current+1]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func (l *Lexer) GetTokens() []token.Token {
	return l.tokens
}
