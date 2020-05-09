package parser_test

import (
	"log"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
		let one = 1;
		let two = 4930;
		let three = 5;
	`
	wantSS := []ast.LetStatement{
		{
			Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "one"}, Value: "one"},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "1"}, Value: 1},
		},
		{
			Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "two"}, Value: "two"},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "4930"}, Value: 4930},
		},
		{
			Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "three"}, Value: "three"},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		},
	}

	par := parser.New(lexer.New(input))
	prog, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}

	stmts := prog.Statements
	if len(stmts) != len(wantSS) {
		t.Fatalf("have %v statements, want %v", len(stmts), len(wantSS))
	}

	for i, stmt := range stmts {
		stmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("have statement type %T, want %T", stmt, ast.LetStatement{})
		}

		if stmt.Name.Value != wantSS[i].Name.Value {
			t.Fatalf("have identifier name: %s, want %s", stmt.Name.Value, wantSS[i].Name.Value)
		}

		if stmt.Value.TokenLiteral() != wantSS[i].Value.TokenLiteral() {
			t.Fatalf("have identifier value: %s, want %s", stmt.Value.TokenLiteral(), wantSS[i].Value.TokenLiteral())
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
		return 1;
		return 4930;
		return 5;
	`
	wantSS := []ast.ReturnStatement{
		{
			Token: token.Token{Type: token.RETURN},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "1"}, Value: 1},
		},
		{
			Token: token.Token{Type: token.RETURN},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "4930"}, Value: 4930},
		},
		{
			Token: token.Token{Type: token.RETURN},
			Value: ast.Integer{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		},
	}

	par := parser.New(lexer.New(input))
	prog, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}

	stmts := prog.Statements
	if len(stmts) != len(wantSS) {
		t.Fatalf("have %v statements, want %v", len(stmts), len(wantSS))
	}

	for i, stmt := range stmts {
		stmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("have statement type %T, want %T", stmt, &ast.ReturnStatement{})
		}

		if stmt.Value.TokenLiteral() != wantSS[i].Value.TokenLiteral() {
			t.Fatalf("have return value: %s, want %s", stmt.Value.TokenLiteral(), wantSS[i].Value.TokenLiteral())
		}
	}
}

func TestExpression(t *testing.T) {
	input := `5; foo`
	wantSS := []ast.Expression{
		ast.Integer{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "foo"}, Value: "foo"},
	}

	par := parser.New(lexer.New(input))
	prog, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}

	stmts := prog.Statements

	if len(stmts) != len(wantSS) {
		t.Fatalf("have %v statements, want %v", len(stmts), len(wantSS))
	}

	for i, stmt := range stmts {
		stmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("have statement type %T, want %T", stmt, ast.ExpressionStatement{})
		}
		log.Printf("stmt = %+v\n", stmt)

		if stmt.Expression.TokenLiteral() != wantSS[i].TokenLiteral() {
			t.Fatalf("have expression value: %s, want %s", stmt.String(), wantSS[i].String())
		}
	}
}

func TestString(t *testing.T) {
	input := `let foo = 5;`
	par := parser.New(lexer.New(input))
	prog, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}

	str := prog.String()
	if want := "let foo = 5;"; str != want {
		t.Fatalf("have program string %s, want %s", str, want)
	}
}
