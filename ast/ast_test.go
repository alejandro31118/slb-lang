package ast

import (
	"testing"

	"github.com/alejandro31118/slb-lang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AssignStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}
}
