package token

type Token struct {
	Type    string
	Literal string
}

const (
	INT    = "INT"    // Token type for integers.
	FLOAT  = "FLOAT"  // Token type for floating point numbers.
	STRING = "STRING" // Token type for strings.

	ASSIGN   = "=" // Token type for assignment operators.
	PLUS     = "+" // Token type for addition.
	MINUS    = "-" // Token type for subtraction.
	ASTARISK = "*" // Token type for multiplication.
	SLASH    = "/" // Token type for division.

	ILLEGAL = "ILLEGAL" // Token type for illegal tokens.
	EOF     = "EOF"     // Token type that represents end of file.
	IDENT   = "IDENT"   // Token type for identifiers.

	BANG = "!"  // Token type for NOT operator.
	LT   = "<"  // Token type for 'less than' operator.
	GT   = ">"  // Token type for 'greater than' operator.
	EQ   = "==" // Token type for equality operator.
	NEQ  = "!=" // Token type for not equality operator.

	COMMA     = "," // Token type for commas.
	SEMICOLON = ";" // Token type for semicolons.
	COLON     = ":" // Token type for colons.

	LPAREN   = "(" // Token type for left parentheses.
	RPAREN   = ")" // Token type for right parentheses.
	LBRACE   = "{" // Token type for left braces.
	RBRACE   = "}" // Token type for right braces.
	LBRACKET = "[" // Token type for left brackets.
	RBRACKET = "]" // Token type for right brackets.

	FUNCTION = "FUNCTION" // Token type for functions.
	LET      = "LET"      // Token type for lets.
	TRUE     = "TRUE"     // Token type for true.
	FALSE    = "FALSE"    // Token type for false.
	IF       = "IF"       // Token type for if.
	ELSE     = "ELSE"     // Token type for else.
	RETURN   = "RETURN"   // Token type for return.
)

var keywords = map[string]string{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func New(typ string, lit byte) Token {
	return Token{Type: typ, Literal: string(lit)}
}

func IdentifyToken(identifier string) string {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
