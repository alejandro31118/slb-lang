package ast

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) == 0 {
		return ""
	}

	return prog.Statements[0].TokenLiteral()
}

func (assign *AssignStatement) TokenLiteral() string {
	return assign.Token.Literal
}

func (ident *Ident) TokenLiteral() string {
	return ident.Token.Literal
}
