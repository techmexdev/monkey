package ast

// Node is an AST node.
type Node interface {
	TokenLiteral() string
	String() string
}
