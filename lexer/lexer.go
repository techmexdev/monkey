package lexer

import (
	"monkey/token"
)

// Lexer iterates over text, creating tokens
type Lexer struct {
	input        string
	position     int  // current char
	readPosition int  // after current char
	ch           byte // current char
}

const nullChar = 0 // ASCI code for null

// New creates a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken reads token at l.position, and increments pointer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case []byte(token.SEMICOLON)[0]:
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}

	case []byte(token.COMMA)[0]:
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}

	case []byte(token.LPAREN)[0]:
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}

	case []byte(token.RPAREN)[0]:
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}

	case []byte(token.LBRACE)[0]:
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}

	case []byte(token.RBRACE)[0]:
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}

	case []byte(token.ASSIGN)[0]:
		if l.peekChar() == []byte(token.ASSIGN)[0] {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: lit}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		}

	case []byte(token.PLUS)[0]:
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}

	case []byte(token.MINUS)[0]:
		tok = token.Token{Type: token.MINUS, Literal: string(l.ch)}

	case []byte(token.BANG)[0]:
		if l.peekChar() == []byte(token.EQ)[0] {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: lit}
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(l.ch)}
		}

	case []byte(token.ASTERISK)[0]:
		tok = token.Token{Type: token.ASTERISK, Literal: string(l.ch)}

	case []byte(token.SLASH)[0]:
		tok = token.Token{Type: token.SLASH, Literal: string(l.ch)}

	case []byte(token.LT)[0]:
		tok = token.Token{Type: token.LT, Literal: string(l.ch)}

	case []byte(token.GT)[0]:
		tok = token.Token{Type: token.GT, Literal: string(l.ch)}

	case nullChar: // NULL
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isValidIdentChar(l.ch) {
			ident := l.readIdentifier()
			return token.Token{Type: token.IdentType(ident), Literal: ident}
		} else if isDigit(l.ch) {
			return token.Token{Type: token.INT, Literal: l.readInt()}
		}
		tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// reads char until end of integer
func (l *Lexer) readInt() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	wsChars := []byte{' ', '\t', '\n', '\r'}
	for inByteArray(l.ch, wsChars) {
		l.readChar()
	}
}

func inByteArray(ch byte, chs []byte) bool {
	for _, c := range chs {
		if ch == c {
			return true
		}
	}
	return false
}

// reads char until end of identifier
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isValidIdentChar(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// can ch be part of identifier?
func isValidIdentChar(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = nullChar
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
