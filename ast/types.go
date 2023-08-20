package ast

import "github.com/alejandro31118/slb-lang/token"

type Node interface {
	TokenLiteral() string
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

type Ident struct {
	Token token.Token
	Value string
}

type AssignStatement struct {
	Token token.Token
	Name  *Ident
	Value Expression
}
