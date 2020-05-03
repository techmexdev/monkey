package parser_test

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
		let one = 1;
		let two = 4930;
		let three = 5;
	`
	wantSS := []ast.LetStatement{
		{Name: &ast.Identifier{Value: "one"}, Value: ast.Expression{Value: "1"}},
		{Name: &ast.Identifier{Value: "two"}, Value: ast.Expression{Value: "4930"}},
		{Name: &ast.Identifier{Value: "three"}, Value: ast.Expression{Value: "5"}},
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
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
		return 1;
		return 4930;
		return 5;
	`
	wantSS := []ast.ReturnStatement{
		{Value: ast.Expression{Value: "1"}},
		{Value: ast.Expression{Value: "4930"}},
		{Value: ast.Expression{Value: "5"}},
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
			t.Fatalf("have statement type %T, want %T", stmt, ast.ReturnStatement{})
		}

		if stmt.Value != wantSS[i].Value {
			t.Fatalf("have return value: %s, want %s", stmt.Value, wantSS[i].Value)
		}
	}
}
