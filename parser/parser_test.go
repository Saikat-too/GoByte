package parser

import (
	"gobyte/ast"
	"gobyte/lexer"
	"testing"
)

func TestAssignStatement(t *testing.T) {
	input := `
  x = 5
  y = 10
  add = 16

  `

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 Statements . got= %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"add"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAssignStatement(t *testing.T, as ast.Statement, name string) bool {
	astmt, ok := as.(*ast.AssignStatement)

	if !ok {
		t.Errorf(" s not *ast.LetStatement. got=%T", as)
	}

	if astmt.Name.Value != name {
		t.Errorf("astmt.Name.Value not '%s'. got = %s", name, astmt.Name.Value)
		return false
	}

	if astmt.Name.TokenLiteral() != name {
		t.Errorf("as.Name not '%s'. got=%s", name, astmt.Name)
		return false
	}

	return true
}
