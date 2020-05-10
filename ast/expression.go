package ast

import "monkey/token"

// Expression can be: (4 * 5)
type Expression interface {
	Node
}

// ExpressionStatement is an ast node with an Expression.
type ExpressionStatement struct {
	Token token.Token
	Expression
}

// TokenLiteral allows es to be an AST node
func (es ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns value of es's Expression
func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier contains name of variable, and token.IDENT
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral allows i to be an AST node
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns token, and literal value
func (i Identifier) String() string {
	return i.Token.Literal
}

// Integer contains a number
type Integer struct {
	Token token.Token
	Value int64
}

// TokenLiteral allows i to be an AST node
func (i Integer) TokenLiteral() string {
	return i.Token.Literal
}

// String returns token's literal value
func (i Integer) String() string {
	return i.Token.Literal
}

// PrefixExpression contains a number
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Expression
}

// TokenLiteral allows pe to be an AST node
func (pe PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns token's literal value
func (pe PrefixExpression) String() string {
	str := "(" + pe.Operator + ")"
	if pe.Expression != nil {
		str += " " + pe.Expression.String()
	}

	return str
}
