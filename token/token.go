package token

// Type is a token's type.
type Type string

// Token is a lexical token
type Token struct {
	Type
	Literal string
}

const (
	// ILLEGAL is an unsupported token
	ILLEGAL Type = "ILLEGAL"
	// EOF is end of file
	EOF = "EOF"
	// IDENT is a variable name
	IDENT     = "IDENT"
	INT       = "INT"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Equality
	EQ     = "=="
	NOT_EQ = "!="

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywordType = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// IdentType returns a keyword token type, or IDENT
func IdentType(ident string) Type {
	if typ, ok := keywordType[ident]; ok {
		return typ
	}
	return IDENT
}
