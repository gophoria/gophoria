package lexer

type TokenType int

const (
	TokenTypeIllegal TokenType = iota
	TokenTypeEof

	// identifiers + literals
	TokenTypeIdent
	TokenTypeInt
	TokenTypeString

	// operators
	TokenTypeAssign

	// delimeters
	TokenTypeLParen
	TokenTypeRParen

	TokenTypeLBrace
	TokenTypeRBrace

	TokenTypeLSquareBrace
	TokenTypeRSquareBrace

	TokenTypeColon
	TokenTypeComma

	TokenTypeDecorator

	// keywords
	TokenTypeEnum
	TokenTypeModel
	TokenTypeDb
	TokenTypeTrue
	TokenTypeFalse

	// types
	TokenTypeTInt
	TokenTypeTReal
	TokenTypeTString
	TokenTypeTBool
	TokenTypeTDateTime
)

type Token struct {
	Type    TokenType
	Literal string

	Row int
	Col int
}

func NewToken(tokenType TokenType, literal string, row int, col int) *Token {
	token := Token{
		Type:    tokenType,
		Literal: literal,
		Row:     row,
		Col:     col,
	}

	return &token
}
