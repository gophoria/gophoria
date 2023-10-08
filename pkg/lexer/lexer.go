package lexer

import "strings"

type Lexer struct {
	input        string
	position     int
	peekPosition int

	ch  byte
	row int
	col int
}

var keywordsMap = map[string]TokenType{
	"enum":     TokenTypeEnum,
	"model":    TokenTypeModel,
	"db":       TokenTypeDb,
	"true":     TokenTypeTrue,
	"false":    TokenTypeFalse,
	"int":      TokenTypeTInt,
	"real":     TokenTypeTReal,
	"string":   TokenTypeTString,
	"bool":     TokenTypeTBool,
	"DateTime": TokenTypeTDateTime,
}

func NewLexer(input string) *Lexer {
	lexer := Lexer{
		input:        input,
		position:     0,
		peekPosition: 0,

		row: 0,
		col: -1,
	}

	lexer.readChar()

	return &lexer
}

func (l *Lexer) Next() *Token {
	var tok *Token = nil

	l.skipWhitespaces()

	switch l.ch {
	case 0:
		tok = NewToken(TokenTypeEof, "", l.row, l.col)
		break
	case '=':
		tok = NewToken(TokenTypeAssign, "=", l.row, l.col)
		break
	case '(':
		tok = NewToken(TokenTypeLParen, "(", l.row, l.col)
		break
	case ')':
		tok = NewToken(TokenTypeRParen, ")", l.row, l.col)
		break
	case '{':
		tok = NewToken(TokenTypeLBrace, "{", l.row, l.col)
		break
	case '}':
		tok = NewToken(TokenTypeRBrace, "}", l.row, l.col)
		break
	case '[':
		tok = NewToken(TokenTypeLSquareBrace, "[", l.row, l.col)
		break
	case ']':
		tok = NewToken(TokenTypeRSquareBrace, "]", l.row, l.col)
		break
	case ':':
		tok = NewToken(TokenTypeColon, ":", l.row, l.col)
		break
	case ',':
		tok = NewToken(TokenTypeComma, ",", l.row, l.col)
		break
	case '@':
		tok = NewToken(TokenTypeDecorator, "@", l.row, l.col)
		break
	case '"':
		row := l.row
		col := l.col
		literal := l.readString()
		tok = NewToken(TokenTypeString, literal, row, col)
		break
	}

	if tok == nil {
		if l.isLetter(l.ch) {
			row := l.row
			col := l.col
			tokenType, literal := l.readIdentifier()
			tok = NewToken(tokenType, literal, row, col)
			return tok
		} else if l.isNumber(l.ch) {
			row := l.row
			col := l.col
			literal := l.readNumber()
			tok = NewToken(TokenTypeInt, literal, row, col)
			return tok
		}
	}

	if tok == nil {
		tok = NewToken(TokenTypeIllegal, string(l.ch), l.row, l.col)
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.peekPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.peekPosition]
	}

	l.position = l.peekPosition
	l.peekPosition++
	l.col++
}

func (l *Lexer) readString() string {
	var sb strings.Builder

	l.readChar()
	for l.ch != '"' {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	return sb.String()
}

func (l *Lexer) readIdentifier() (TokenType, string) {
	var sb strings.Builder

	for l.isLetter(l.ch) || l.isNumber(l.ch) || l.ch == '_' {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	literal := sb.String()
	tokenType, ok := keywordsMap[literal]
	if !ok {
		tokenType = TokenTypeIdent
	}

	return tokenType, literal
}

func (l *Lexer) readNumber() string {
	var sb strings.Builder

	for l.isNumber(l.ch) {
		sb.WriteByte(l.ch)
		l.readChar()
	}

	return sb.String()
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' {
		if l.ch == '\n' {
			l.row++
			l.col = 0
		}

		l.readChar()
	}
}

func (l *Lexer) isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
