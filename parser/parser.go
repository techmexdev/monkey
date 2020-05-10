package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
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
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parsePrefixExpression() (ast.Expression, error) {
	preExp := ast.PrefixExpression{}
	switch p.currTok.Type {
	case token.BANG:
		preExp.Operator = p.currTok.Literal
	case token.MINUS:
		preExp.Operator = p.currTok.Literal
	}
	p.readToken()

	expr, err := p.parseExpression()
	if err != nil {
		return ast.PrefixExpression{}, err
	}

	preExp.Expression = expr

	return preExp, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	var expr ast.Expression
	switch p.currTok.Type {
	case token.BANG, token.MINUS:
		preExp := ast.PrefixExpression{Token: p.currTok, Operator: p.currTok.Literal}
		p.readToken()
		exp, err := p.parseExpression()
		if err != nil {
			return nil, fmt.Errorf("failed parsing prefix expression's expression: %s", err)
		}
		preExp.Expression = exp
		return &preExp, nil
	case token.IDENT:
		expr = &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	case token.INT:
		num, err := strconv.ParseInt(p.currTok.Literal, 0, 64)
		if err != nil {
			return &ast.Identifier{}, fmt.Errorf("could not parse token: %v into int64: %s", p.currTok, err)
		}

		expr = &ast.Integer{Token: p.currTok, Value: num}
	}
	return expr, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := ast.ExpressionStatement{Token: p.currTok}
	expr, err := p.parseExpression()
	if err != nil {
		return &ast.ExpressionStatement{}, err
	}
	stmt.Expression = expr

	// expressions can end with a semicolon. Ex 5+5;
	if p.nextTok.Type == token.SEMICOLON {
		p.readToken()
	}

	return &stmt, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := ast.ReturnStatement{Token: p.currTok}
	if p.currTok.Type != token.RETURN {
		return &ast.ReturnStatement{}, fmt.Errorf("have token type %s in beginning of return statement, want %s", p.currTok.Type, token.RETURN)
	}
	p.readToken()

	expr, err := p.parseExpression()
	if err != nil {
		return &ast.ReturnStatement{}, fmt.Errorf("failed parsing expression in return statement: %s", err)
	}
	stmt.Value = expr

	if p.nextTok.Type != token.SEMICOLON {
		return nil, fmt.Errorf("have token %v, want %s", p.nextTok.Type, token.SEMICOLON)
	}
	p.readToken()

	return &stmt, nil
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := ast.LetStatement{Token: p.currTok}

	if p.currTok.Type != token.LET {
		return &ast.LetStatement{}, fmt.Errorf("have token type %s in beginning of let statement, want %s", p.currTok.Type, token.LET)
	}
	p.readToken()

	if p.currTok.Type != token.IDENT {
		return nil, fmt.Errorf("have next token type %s, want %s", p.currTok.Type, token.IDENT)
	}
	stmt.Name = &ast.Identifier{Token: p.currTok, Value: p.currTok.Literal}
	p.readToken()

	if p.currTok.Type != token.ASSIGN {
		return nil, fmt.Errorf("have next token type %s, want %s", p.currTok.Type, token.ASSIGN)
	}
	p.readToken()

	expr, err := p.parseExpression()
	if err != nil {
		return &ast.LetStatement{}, fmt.Errorf("failed parsing expression in let statement: %s", err)
	}
	stmt.Value = expr

	if p.nextTok.Type != token.SEMICOLON {
		return nil, fmt.Errorf("have token %v, want %s", p.nextTok.Type, token.SEMICOLON)
	}
	p.readToken()

	return &stmt, nil
}
