package ast

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	var output bytes.Buffer

	for _, item := range p.Statements {
		output.WriteString(item.String())
	}

	return output.String()
}
