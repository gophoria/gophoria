package ast

import (
	"strings"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type VariableType int

const (
	VariableTypeInt VariableType = iota
	VariableTypeReal
	VariableTypeBool
	VariableTypeString
	VariableTypeDateTime
	VariableTypeObject
)

var VariableTypeMap = map[lexer.TokenType]VariableType{
	lexer.TokenTypeTInt:      VariableTypeInt,
	lexer.TokenTypeTReal:     VariableTypeReal,
	lexer.TokenTypeTBool:     VariableTypeBool,
	lexer.TokenTypeTString:   VariableTypeString,
	lexer.TokenTypeTDateTime: VariableTypeDateTime,
	lexer.TokenTypeIdent:     VariableTypeObject,
}

type Declaration struct {
	Identifier      *Identifier
	DeclarationType *DeclarationType
}

type DeclarationType struct {
	Token   *lexer.Token
	Type    VariableType
	Name    string
	IsArray bool
}

func NewDeclaration(ident *Identifier, declType *DeclarationType) *Declaration {
	d := Declaration{
		Identifier:      ident,
		DeclarationType: declType,
	}

	return &d
}

func (d *Declaration) String() string {
	var sb strings.Builder

	sb.WriteString(d.Identifier.String())
	sb.WriteString(" ")
	sb.WriteString(d.DeclarationType.String())

	return sb.String()
}

func NewDeclarationType(token *lexer.Token, isArray bool) *DeclarationType {
	vType := VariableTypeMap[token.Type]

	v := DeclarationType{
		Token:   token,
		Type:    vType,
		Name:    token.Literal,
		IsArray: isArray,
	}

	return &v
}

func (v *DeclarationType) String() string {
	var sb strings.Builder

	sb.WriteString(v.Name)
	if v.IsArray {
		sb.WriteString("[]")
	}

	return sb.String()
}
