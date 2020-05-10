package ast

import "monkey/token"

// LetStatement can be the following block: let foo = 4
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral allows ls to be an AST node
func (ls LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns token, and literal value
func (ls LetStatement) String() string {
	str := ls.Token.Literal
	if ls.Name != nil {
		str += " " + ls.Name.Value
	}

	if ls.Value != nil {
		str += " = " + ls.Value.String() + ";"
	}

	return str
}
