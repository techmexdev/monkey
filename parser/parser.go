package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser makes statements and expressions from a lexer's tokens
type Parser struct {
	l       *lexer.Lexer
	currTok token.Token
	nextTok token.Token
	errors  []string
}

// New creates a parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.readToken()
	p.readToken()
	return p
}

func (p *Parser) readToken() {
	p.currTok = p.nextTok
	p.nextTok = p.l.NextToken()
}

// Parse reads lexer's tokens, and creates AST Nodes from them
func (p *Parser) Parse() (*ast.Program, error) {
	pro := &ast.Program{}
	var err error
	for p.currTok.Type != token.EOF {
		stmt, stmtErr := p.parseStatement()
		if stmtErr != nil {
			err = fmt.Errorf("%s\n%s", err, stmtErr)
			continue
		}
		if stmt != nil {
			pro.Statements = append(pro.Statements, stmt)
		}
		p.readToken()
	}

	return pro, err
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.currTok.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return nil, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := ast.ReturnStatement{Token: p.currTok}

	// TODO: Allow expressions, not just ints as stmt value
	if p.nextTok.Type != token.INT {
		return nil, fmt.Errorf("have next token type %s, want %s", p.nextTok.Type, token.INT)
	}
	stmt.Value.Value = p.nextTok.Literal
	p.readToken()

	return &stmt, nil
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := ast.LetStatement{Token: p.currTok}

	if p.nextTok.Type != token.IDENT {
		return nil, fmt.Errorf("have next token type %s, want %s", p.nextTok.Type, token.IDENT)
	}
	p.readToken()
	stmt.Name = &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}

	if p.nextTok.Type != token.ASSIGN {
		return nil, fmt.Errorf("have next token type %s, want %s", p.nextTok.Type, token.ASSIGN)
	}
	p.readToken()

	// TODO: Allow expressions, not just ints as stmt value
	if p.nextTok.Type != token.INT {
		return nil, fmt.Errorf("have next token type %s, want %s", p.nextTok.Type, token.INT)
	}
	stmt.Value = ast.Expression{Value: p.currTok.Literal}
	p.readToken()

	// TODO: Handle expressions
	if p.nextTok.Type != token.SEMICOLON {
		return nil, fmt.Errorf("have token %v, want %s", p.nextTok.Type, token.SEMICOLON)
	}
	p.readToken()

	return &stmt, nil
}
