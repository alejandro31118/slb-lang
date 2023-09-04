package parser

import (
	"fmt"
	"strconv"

	"github.com/alejandro31118/slb-lang/ast"
	"github.com/alejandro31118/slb-lang/lexer"
	"github.com/alejandro31118/slb-lang/token"
)

const (
	_ = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	PrefixParseFunc func() ast.Expression
	InfixParseFunc  func(ast.Expression) ast.Expression
)

type Parser struct {
	Lexer            *lexer.Lexer
	CurrentToken     token.Token
	FollowingToken   token.Token
	Errors           []string
	PrefixParseFuncs map[token.Type]PrefixParseFunc
	InfixParseFuncs  map[token.Type]InfixParseFunc
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		Lexer:            lexer,
		Errors:           []string{},
		PrefixParseFuncs: make(map[token.Type]PrefixParseFunc),
	}

	parser.RegisterPrefix(token.IDENT, parser.ParseIndentifierExpression)
	parser.RegisterPrefix(token.INT, parser.ParseIntegerExpression)

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
	case token.RETURN:
		return parser.ParseReturnStatement()
	default:
		return parser.ParseExpressionStatement()
	}
}

func (parser *Parser) ParseExpression(precedence int) ast.Expression {
	prefix := parser.PrefixParseFuncs[parser.CurrentToken.Type]
	if prefix == nil {
		return nil
	}

	return prefix()
}

func (parser *Parser) ParseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: parser.CurrentToken}
	stmt.Expression = parser.ParseExpression(LOWEST)

	if parser.NextTokenTypeIs(token.SEMICOLON) {
		parser.NextToken()
	}

	return stmt
}

func (parser *Parser) ParseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{
		Token: parser.CurrentToken,
	}

	if !parser.ExpectNextTokenTypeBe(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
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

func (parser *Parser) ParseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: parser.CurrentToken,
	}

	parser.NextToken()

	if !parser.CurrentTokenTypeIs(token.SEMICOLON) {
		parser.NextToken()
	}

	return stmt
}

func (parser *Parser) ParseIndentifierExpression() ast.Expression {
	return &ast.Identifier{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

func (parser *Parser) ParseIntegerExpression() ast.Expression {
	integer := &ast.Integer{
		Token: parser.CurrentToken,
	}

	value, err := strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", parser.CurrentToken.Literal)
		parser.Errors = append(parser.Errors, msg)

		return nil
	}

	integer.Value = value
	return integer
}

func (parser *Parser) RegisterPrefix(tokenType token.Type, fn PrefixParseFunc) {
	parser.PrefixParseFuncs[tokenType] = fn
}
