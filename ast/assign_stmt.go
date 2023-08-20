package ast

import (
	"bytes"

	"github.com/alejandro31118/slb-lang/token"
)

type AssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) TokenLiteral() string {
	return as.Token.Literal
}

func (as *AssignStatement) String() string {
	var output bytes.Buffer

	output.WriteString(as.TokenLiteral() + " ")
	output.WriteString(as.Name.String())
	output.WriteString(" = ")

	if as.Value != nil {
		output.WriteString(as.Value.String())
	}

	output.WriteString(";")

	return output.String()
}
