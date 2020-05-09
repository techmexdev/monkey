package ast

import (
	"monkey/token"
	"strings"
)

// Node is an AST node.
type Node interface {
	TokenLiteral() string
	String() string
}

// Program is the root AST node
type Program struct {
	Node
	Statements []Statement
}

func (p Program) String() string {
	var ss []string
	for _, s := range p.Statements {
		ss = append(ss, s.String())
	}
	return strings.Join(ss, "\n")
}

// Statement is can be a let, or return
type Statement interface {
	Node
	String() string
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
	return i.Token.Literal + " " + i.Value
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
