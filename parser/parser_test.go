package parser

import (
	"testing"

	"github.com/alejandro31118/slb-lang/ast"
	"github.com/alejandro31118/slb-lang/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	length := len(program.Statements)
	if length != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", length)
	}
	checkParserErrors(t, p)

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdent) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors
	length := len(errors)

	if length == 0 {
		return
	}

	t.Errorf("parser has %d errors", length)

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testAssignStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	assignStmt, ok := stmt.(*ast.AssignStatement)
	if !ok {
		t.Errorf("stmt not *ast.AssignStatement. got=%T", stmt)
		return false
	}

	if assignStmt.Name.Value != name {
		t.Errorf("assignStmt.Name.Value not '%s'. got=%s", name, assignStmt.Name.Value)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name not '%s'. got=%s", name, assignStmt.Name)
		return false
	}

	return true
}
