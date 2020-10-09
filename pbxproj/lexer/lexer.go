package lexer

import (
	"github.com/josa42/go-xcode-project/pbxproj/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	assign       bool
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	for l.skipWhitespace() || l.skipLineComment() || l.skipComment() {
	}

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		l.assign = true
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.assign = false
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		l.assign = false
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		l.assign = false
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.assign = true
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.assign = false
	case '"':
		if l.assign {
			tok.Type = token.STRING
		} else {
			tok.Type = token.KEY
		}
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetterOrDigit(l.ch) {

			tok.Literal = l.readIdentifier()
			if l.assign {
				tok.Type = token.STRING
			} else {
				tok.Type = token.KEY
			}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() bool {
	found := false
	for isWhitespace(l.ch) {
		l.readChar()
		found = true
	}
	return found
}

func (l *Lexer) skipLineComment() bool {
	found := false
	if l.ch == '/' && l.peekChar() == '/' {
		for l.ch != '\n' {
			l.readChar()
			found = true
		}
	}
	return found
}

func (l *Lexer) skipComment() bool {
	found := false
	if l.ch == '/' && l.peekChar() == '*' {
		found = true
		for l.ch != '*' || l.peekChar() != '/' {
			l.readChar()
		}

		// skip: */
		l.readChar()
		l.readChar()
	}

	return found
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetterOrDigit(l.peekChar()) {
		l.readChar()
	}
	return l.inputRange(position, l.position+1)
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.peekChar()) {
		l.readChar()
	}
	return l.inputRange(position, l.position+1)
}

func (l *Lexer) readString() string {
	out := ""
	for l.ch != 0 {
		l.readChar()

		if l.ch == '\\' {
			l.readChar()
		} else if l.ch == '"' {
			break
		}

		out += string(l.ch)
	}

	return out
}

func (l *Lexer) inputRange(begin, end int) string {
	ilen := len(l.input)

	if end == 0 {
		return ""
	}

	if end >= ilen {
		end = ilen - 1
	}

	return l.input[begin:end]
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetterOrDigit(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
