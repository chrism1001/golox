package lexer

import (
	"golox/token"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("simple first test", func(t *testing.T) {
		input := `//
		(( )){}
		!*+-/=<> <= ==`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.LEFT_PAREN, "("},
			{token.LEFT_PAREN, "("},
			{token.RIGHT_PAREN, ")"},
			{token.RIGHT_PAREN, ")"},
			{token.LEFT_BRACE, "{"},
			{token.RIGHT_BRACE, "}"},
			{token.BANG, "!"},
			{token.STAR, "*"},
			{token.PLUS, "+"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.EQUAL, "="},
			{token.LESS, "<"},
			{token.GREATER, ">"},
			{token.LESS_EQUAL, "<="},
			{token.EQUAL_EQUAL, "=="},
		}

		l := New(input)

		tokens := l.scanTokens()

		for i, tt := range tests {
			if tokens[i].Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected %q got %q", i, tt.expectedType, tokens[i].Type)
			}
			if string(tokens[i].Type) != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tokens[i].Literal)
			}
		}
	})
}
