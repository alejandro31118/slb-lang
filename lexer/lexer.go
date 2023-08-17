package lexer

import (
	"github.com/alejandro31118/slb-lang/token"
)

type Lexer struct {
	Source       string
	Position     int
	ReadPosition int
	Char         byte
}

type lexer interface {
	NextToken() token.Token
}

func New(source string) Lexer {
	lex := Lexer{Source: source}
	lex.ChopChar()
	return lex
}

func (lex *Lexer) NextToken() token.Token {
	lex.TrimLeft()

	if lex.Char == '/' && lex.NextChar() == '/' {
		lex.DropLine()
	}

	var tok token.Token
	switch lex.Char {
	case '+':
		tok = token.New(token.PLUS, lex.Char)
	case '-':
		tok = token.New(token.MINUS, lex.Char)
	case '*':
		tok = token.New(token.ASTARISK, lex.Char)
	case '/':
		tok = token.New(token.SLASH, lex.Char)
	case '<':
		tok = token.New(token.LT, lex.Char)
	case '>':
		tok = token.New(token.GT, lex.Char)
	case ',':
		tok = token.New(token.COMMA, lex.Char)
	case ';':
		tok = token.New(token.SEMICOLON, lex.Char)
	case ':':
		tok = token.New(token.COLON, lex.Char)
	case '{':
		tok = token.New(token.LBRACE, lex.Char)
	case '}':
		tok = token.New(token.RBRACE, lex.Char)
	case '[':
		tok = token.New(token.LBRACKET, lex.Char)
	case ']':
		tok = token.New(token.RBRACKET, lex.Char)
	case '(':
		tok = token.New(token.LPAREN, lex.Char)
	case ')':
		tok = token.New(token.RPAREN, lex.Char)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lex.GetString()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case '=':
		if lex.NextChar() == '=' {
			prevChar := lex.Char
			lex.ChopChar()

			tok.Type = token.EQ
			tok.Literal = string(prevChar) + string(lex.Char)
		} else {
			tok = token.New(token.ASSIGN, lex.Char)
		}
	case '!':
		if lex.NextChar() == '=' {
			prevChar := lex.Char
			lex.ChopChar()

			tok.Type = token.NEQ
			tok.Literal = string(prevChar) + string(lex.Char)
		} else {
			tok = token.New(token.BANG, lex.Char)
		}
	default:
		if isDigit(lex.Char) {
			return lex.GetNumberToken()
		}

		if isLetter(lex.Char) {
			tok.Literal = lex.GetIdent()
			tok.Type = token.IdentifyToken(tok.Literal)

			return tok
		}

		tok = token.New(token.ILLEGAL, lex.Char)
	}

	lex.ChopChar()
	return tok
}

func (lex *Lexer) TrimLeft() {
	for lex.Char == ' ' || lex.Char == '\t' || lex.Char == '\n' || lex.Char == '\r' {
		lex.ChopChar()
	}
}

func (lex *Lexer) ChopChar() {
	if lex.IsNotEmpty() {
		lex.Char = lex.Source[lex.ReadPosition]
	} else {
		lex.Char = 0
	}

	lex.Position = lex.ReadPosition
	lex.ReadPosition += 1
}

func (lex *Lexer) DropLine() {
	for lex.Char != '\n' && lex.Char != '\r' {
		lex.ChopChar()
	}

	lex.TrimLeft()
}

func (lex *Lexer) GetString() string {
	pos := lex.Position + 1

	for {
		lex.ChopChar()
		if lex.Char == '"' || lex.Char == 0 {
			break
		}
	}

	return lex.Source[pos:lex.Position]
}

func (lex *Lexer) NextChar() byte {
	if lex.IsEmpty() {
		return 0
	}

	return lex.Source[lex.ReadPosition]
}

func (lex *Lexer) GetAbstract(conditionCallback func(byte) bool) string {
	oldPosition := lex.Position

	for conditionCallback(lex.Char) {
		lex.ChopChar()
	}

	return lex.Source[oldPosition:lex.Position]
}

func (lex *Lexer) GetNumber() string {
	return lex.GetAbstract(isDigit)
}

func (lex *Lexer) GetIdent() string {
	return lex.GetAbstract(isLetter)
}

func (lex *Lexer) GetNumberToken() token.Token {
	intNumberPart := lex.GetNumber()

	if lex.Char != '.' {
		return token.Token{Type: token.INT, Literal: intNumberPart}
	}

	lex.ChopChar()
	decimalNumberPart := lex.GetNumber()
	return token.Token{Type: token.FLOAT, Literal: intNumberPart + "." + decimalNumberPart}
}

func (lex *Lexer) IsNotEmpty() bool {
	return lex.ReadPosition < len(lex.Source)
}

func (lex *Lexer) IsEmpty() bool {
	return lex.ReadPosition >= len(lex.Source)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}
