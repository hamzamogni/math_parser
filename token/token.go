package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// user defined identifiers + literals
	IDENT  = "IDENT"
	NUMBER = "NUMBER"

	// Delimiters
	COMMA = ","

	LPAREN = "("
	RPAREN = ")"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	// Exponentiation
	EXPONENT = "**"
)
