package lexer

import (
	"c3lang/token"
	"strings"
)

const (
	SINGLE_LINE_COMMENT           = "//"
	MULTI_LINE_COMMENT_OPEN_WITH  = "/*"
	MULTI_LINE_COMMENT_CLOSE_WITH = "*/"
)

type Lexer struct {
	source_code string
	rowNumber   int
	colNumber   int
	position    int
}

func New(input string) *Lexer {
	l := &Lexer{source_code: input, position: -1, rowNumber: 1, colNumber: 1}
	return l
}

func (l *Lexer) nextChar() byte {
	if l.endOfFile() {
		return 0
	}

	return l.source_code[l.position+1]
}

func (l *Lexer) peekChar() byte {
	l.position += 1
	l.colNumber += 1
	return l.source_code[l.position]
}

func (l *Lexer) endOfFile() bool {
	return l.position+1 >= len(l.source_code)
}

func (l *Lexer) startsWith(part string) bool {
	if l.endOfFile() {
		return false
	}

	return strings.HasPrefix(l.source_code[l.position+1:], part)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	l.skipComments()
	l.skipWhitespace()

	tok.RowNumber = l.rowNumber
	tok.ColNumber = l.colNumber

	if l.endOfFile() {
		tok.Literal = ""
		tok.Type = token.EOF
		return tok
	}

	if l.isIdentifierOrKeyword() {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	}

	if l.isNumber() {
		tok.Literal = l.readNumber()
		tok.Type = token.NUMBER
		return tok
	}

	if l.isString() {
		tok.Literal = l.readString()
		tok.Type = token.STRING
		return tok
	}

	if l.isOperator() {
		tok.Literal = l.readOperator()
		tok.Type = token.OPERATOR
		return tok
	}

	if l.isSymbol() {
		tok.Literal = string(l.peekChar())
		tok.Type = token.LookupSymbol(tok.Literal)
		return tok
	}

	tok.Literal = ""
	tok.Type = token.ILLEGAL
	return tok
}

func (l *Lexer) skipWhitespace() {
	for _isWhiteSpace(l.nextChar()) {
		if l.nextChar() == '\n' {
			l.rowNumber += 1
			l.colNumber = 1
		}
		l.peekChar()
	}
}

func (l *Lexer) skipComments() {
	if l.startsWith(SINGLE_LINE_COMMENT) {
		for l.nextChar() != '\n' && l.nextChar() != 0 {
			l.peekChar()
		}
	}

	if l.startsWith(MULTI_LINE_COMMENT_OPEN_WITH) {
		for !l.startsWith(MULTI_LINE_COMMENT_CLOSE_WITH) && !l.endOfFile() {
			l.peekChar()
		}
		if l.startsWith(MULTI_LINE_COMMENT_CLOSE_WITH) {
			l.peekChar()
			l.peekChar()
		}
	}
}

func (l *Lexer) isIdentifierOrKeyword() bool {
	return _isLetter(l.nextChar()) || l.nextChar() == '_'
}

func (l *Lexer) readIdentifier() string {
	result := ""

	for _isValidIdentifierChar(l.nextChar()) {
		result += string(l.peekChar())
	}

	return result
}

func (l *Lexer) isNumber() bool {
	return _isDigit(l.nextChar())
}

func (l *Lexer) readNumber() string {
	result := ""
	for _isValidNumberChar(l.nextChar()) {
		result += string(l.peekChar())
	}
	return result
}

func (l *Lexer) isString() bool {
	return l.nextChar() == '"'
}

func (l *Lexer) readString() string {
	result := ""

	result += string(l.peekChar())
	for l.nextChar() != '"' {
		result += string(l.peekChar())
	}
	result += string(l.peekChar())

	return result
}

func (l *Lexer) isOperator() bool {
	return token.IsOperatorSymbol(string(l.nextChar()))
}

func (l *Lexer) readOperator() string {
	result := ""

	for token.IsOperatorSymbol(string(l.nextChar())) {
		result += string(l.peekChar())
	}

	return result
}

func (l *Lexer) isSymbol() bool {
	return token.IsSymbol(string(l.nextChar()))
}

func _isValidIdentifierChar(ch byte) bool {
	return _isLetter(ch) || _isDigit(ch) || ch == '_' || ch == '?' || ch == '!'
}

func _isValidNumberChar(ch byte) bool {
	return _isDigit(ch) || ch == '_' || ch == '.'
}

func _isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func _isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func _isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
