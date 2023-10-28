package ast

import (
	"strings"
)

type ArgumentType int

const (
	ArgumentTypeValue ArgumentType = iota
	ArgumentTypeCallable
)

type Argument struct {
	Name     *Identifier
	Type     ArgumentType
	Value    *Value
	Callable *Callable
}

type Callable struct {
	Identifier *Identifier
	Arguments  []*Argument
}

func NewCallable(identifier *Identifier) *Callable {
	c := Callable{
		Identifier: identifier,
	}

	return &c
}

func (c *Callable) String() string {
	var sb strings.Builder
	sb.WriteString(c.Identifier.Identifier)
	sb.WriteString("(")
	for i, arg := range c.Arguments {
		sb.WriteString(arg.String())
		if i < len(c.Arguments)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")

	return sb.String()
}

func (c *Callable) AddArgument(argument *Argument) {
	c.Arguments = append(c.Arguments, argument)
}

func NewArgument(name *Identifier, value *Value, callable *Callable) *Argument {
	argType := ArgumentTypeValue
	if callable != nil {
		argType = ArgumentTypeCallable
	}

	a := Argument{
		Type:     argType,
		Name:     name,
		Value:    value,
		Callable: callable,
	}

	return &a
}

func (a *Argument) String() string {
	var sb strings.Builder

	if a.Type == ArgumentTypeCallable {
		sb.WriteString(a.Callable.String())
		return sb.String()
	}

	if a.Name != nil {
		sb.WriteString(a.Name.Identifier)
		sb.WriteString(": ")
	}

	sb.WriteString(a.Value.Value)

	return sb.String()
}
