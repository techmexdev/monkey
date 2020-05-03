package ast

import "monkey/token"

// Node is an AST node.
type Node interface {
	TokenLiteral() string
}

// Program is the root AST node
type Program struct {
	Node
	Statements []Statement
}

// Statement is can be a let, or return
type Statement interface {
	Node
}

// Expression can be: (4 * 5)
type Expression struct {
	Value string
}

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

// Identifier contains name of variable, and token.IDENT
type Identifier struct {
	Token token.Token
	Value string
}

// ReturnStatement contains return value
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

// TokenLiteral allows rs to be an AST node
func (rs ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
