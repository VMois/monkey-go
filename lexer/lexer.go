package lexer

import "monkey/token"

type Lexer struct {
	input        string
	current_char byte
	position     int
	nextPosition int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.current_char = 0
	} else {
		l.current_char = l.input[l.nextPosition]
	}

	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.current_char {
	case '=':
		tok = newToken(token.ASSIGN, l.current_char)
	case ';':
		tok = newToken(token.SEMICOLON, l.current_char)
	case '(':
		tok = newToken(token.LPAREN, l.current_char)
	case ')':
		tok = newToken(token.RPAREN, l.current_char)
	case ',':
		tok = newToken(token.COMMA, l.current_char)
	case '+':
		tok = newToken(token.PLUS, l.current_char)
	case '{':
		tok = newToken(token.LBRACE, l.current_char)
	case '}':
		tok = newToken(token.RBRACE, l.current_char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.current_char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.current_char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.current_char)
			return tok
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.current_char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.current_char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.current_char == ' ' || l.current_char == '\t' || l.current_char == '\n' || l.current_char == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
