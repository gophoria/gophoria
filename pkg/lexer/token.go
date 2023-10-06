package lexer

type TokenType int

const (
	TokenTypeInvalid TokenType = iota
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

	// decorators
	TokenTypeDecId
	TokenTypeDecNullable
	TokenTypeDecDefault
	TokenTypeDecRelation

	// keywords
	TokenTypeEnum
	TokenTypeModel
	TokenTypeDb
	TokenTypeTrue
	TokenTypeFalse

	// types
	TokenTypeTInt
	TokenTypeTString
	TokenTypeTBool
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