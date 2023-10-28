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
	Token    *lexer.Token
	Type     DecoratorType
	Name     *Identifier
	Callable *Callable
}

func NewDecorator(token *lexer.Token, name *Identifier) *Decorator {
	d := Decorator{
		Token: token,
		Name:  name,
		Type:  DecoratorTypeNormal,
	}

	return &d
}

func (d *Decorator) String() string {
	var sb strings.Builder

	sb.WriteString(d.Token.Literal)

	if d.Type == DecoratorTypeCallable {
		sb.WriteString(d.Callable.String())
	} else {
		sb.WriteString(d.Name.String())
	}

	return sb.String()
}

func (d *Decorator) SetCallable(callable *Callable) {
	d.Type = DecoratorTypeCallable
	d.Callable = callable
}
