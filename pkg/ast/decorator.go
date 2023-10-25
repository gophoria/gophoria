package ast

import (
	"strings"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type DecoratorType int

const (
	DecoratorTypeNormal DecoratorType = iota
	DecoratorTypeCallable
)

type Decorator struct {
	Token *lexer.Token
	Name  *Identifier
}

func NewDecorator(token *lexer.Token, name *lexer.Token) *Decorator {
	d := Decorator{
		Token: token,
		Name:  NewIdentifier(name),
	}

	return &d
}

func (d *Decorator) String() string {
	var sb strings.Builder

	sb.WriteString(d.Token.Literal)
	sb.WriteString(d.Name.String())

	return sb.String()
}
