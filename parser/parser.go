package parser

import (
	"fmt"

	"github.com/alejandro31118/slb-lang/ast"
	"github.com/alejandro31118/slb-lang/lexer"
	"github.com/alejandro31118/slb-lang/token"
)

type Parser struct {
	Lexer          *lexer.Lexer
	CurrentToken   token.Token
	FollowingToken token.Token
	Errors         []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		Lexer:  lexer,
		Errors: []string{},
	}

	// Get 2 tokens to set "CurrentToken" and "FollowingToken"
	parser.NextToken()
	parser.NextToken()

	return parser
}

func (parser *Parser) NextToken() {
	parser.CurrentToken = parser.FollowingToken
	parser.FollowingToken = parser.Lexer.NextToken()
}

func (parser *Parser) ThrowError(typ token.Type) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", typ, parser.FollowingToken.Type)
	parser.Errors = append(parser.Errors, msg)
}

func (parser *Parser) CurrentTokenTypeIs(typ token.Type) bool {
	return parser.CurrentToken.Type == typ
}

func (parser *Parser) NextTokenTypeIs(typ token.Type) bool {
	return parser.FollowingToken.Type == typ
}

func (parser *Parser) ExpectNextTokenTypeBe(typ token.Type) bool {
	if parser.NextTokenTypeIs(typ) {
		parser.NextToken()
		return true
	}

	parser.ThrowError(typ)
	return false
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for !parser.CurrentTokenTypeIs(token.EOF) {
		stmt := parser.ParseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		parser.NextToken()
	}

	return program
}

func (parser *Parser) ParseStatement() ast.Statement {
	switch parser.CurrentToken.Type {
	case token.LET:
		return parser.ParseAssignStatement()
	default:
		return nil
	}
}

func (parser *Parser) ParseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{
		Token: parser.CurrentToken,
	}

	if !parser.ExpectNextTokenTypeBe(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}

	if !parser.ExpectNextTokenTypeBe(token.ASSIGN) {
		return nil
	}

	if !parser.CurrentTokenTypeIs(token.SEMICOLON) {
		parser.NextToken()
	}

	return stmt
}
