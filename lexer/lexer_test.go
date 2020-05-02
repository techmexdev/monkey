package lexer_test

import (
	"testing"

	"monkey/lexer"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;

	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	
	10 == 10
	10 != 9
	`
	wantToks := []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.ASSIGN, Literal: token.ASSIGN},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.ASSIGN, Literal: token.ASSIGN},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.ASSIGN, Literal: token.ASSIGN},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: token.LPAREN},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.COMMA, Literal: token.COMMA},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.RPAREN, Literal: token.RPAREN},
		{Type: token.LBRACE, Literal: token.LBRACE},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: token.PLUS},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},
		{Type: token.RBRACE, Literal: token.RBRACE},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "result"},
		{Type: token.ASSIGN, Literal: token.ASSIGN},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.LPAREN, Literal: token.LPAREN},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.COMMA, Literal: token.COMMA},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.RPAREN, Literal: token.RPAREN},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.BANG, Literal: token.BANG},
		{Type: token.MINUS, Literal: token.MINUS},
		{Type: token.SLASH, Literal: token.SLASH},
		{Type: token.ASTERISK, Literal: token.ASTERISK},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: token.LT},
		{Type: token.INT, Literal: "10"},
		{Type: token.GT, Literal: token.GT},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},

		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: token.LPAREN},
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: token.LT},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: token.RPAREN},
		{Type: token.LBRACE, Literal: token.LBRACE},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},
		{Type: token.RBRACE, Literal: token.RBRACE},

		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: token.LBRACE},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: token.SEMICOLON},
		{Type: token.RBRACE, Literal: token.RBRACE},

		{Type: token.INT, Literal: "10"},
		{Type: token.EQ, Literal: token.EQ},
		{Type: token.INT, Literal: "10"},
		{Type: token.INT, Literal: "10"},
		{Type: token.NOT_EQ, Literal: token.NOT_EQ},
		{Type: token.INT, Literal: "9"},

		{Type: token.EOF, Literal: ""},
	}

	lex := lexer.New(input)
	for i, want := range wantToks {
		tok := lex.NextToken()

		if tok.Type != want.Type {
			t.Fatalf("wrong token %v: have type %s want %s", i, tok.Type, want.Type)
		}
		if tok.Literal != want.Literal {
			t.Fatalf("wrong token %v: have literal %v want %v", i, tok.Literal, want.Literal)
		}
	}

}

func BenchmarkNextToken(b *testing.B) {
	input := `
	let five = 5;

	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	
	10 == 10
	10 != 9
	`
	b.StartTimer()
	lex := lexer.New(input)
	for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
	}

	b.StopTimer()

}
