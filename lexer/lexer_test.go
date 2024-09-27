package lexer

import (
	"gobyte/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `five = 5
    
  ten = 10 

  add = five+ten 



  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.IDENT, "five"},
		{token.PLUS, "+"},
		{token.IDENT, "ten"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q , got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong .expected = %q , got = %q ", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
