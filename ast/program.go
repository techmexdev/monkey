package ast

import "strings"

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
