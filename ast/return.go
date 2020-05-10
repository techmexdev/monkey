package ast

import "monkey/token"

// ReturnStatement contains return value
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

// TokenLiteral allows rs to be an AST node
func (rs ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns token, and literal value
func (rs ReturnStatement) String() string {
	str := rs.Token.Literal
	if rs.Value != nil {
		str += " " + rs.Value.String()
	}

	return str
}
