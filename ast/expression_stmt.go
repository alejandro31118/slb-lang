package ast

import "github.com/alejandro31118/slb-lang/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression == nil {
		return ""
	}

	return es.Expression.String()
}
