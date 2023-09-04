package ast

import "github.com/alejandro31118/slb-lang/token"

type Integer struct {
	Token token.Token
	Value int64
}

func (i *Integer) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Integer) String() string {
	return i.TokenLiteral()
}
