package ast

import (
	"strings"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type Config struct {
	Token *lexer.Token
	Type  string
	Items []*AssignItem
}

func NewConfig(token *lexer.Token) *Config {
	c := Config{
		Token: token,
		Type:  token.Literal,
		Items: []*AssignItem{},
	}

	return &c
}

func (c *Config) String() string {
	var sb strings.Builder
	sb.WriteString(c.Type)
	sb.WriteString(" {\n")
	for _, item := range c.Items {
		sb.WriteString(item.String())
		sb.WriteString("\n")
	}
	sb.WriteString(" }\n")

	return sb.String()
}

func (c *Config) AddItem(item *AssignItem) {
	c.Items = append(c.Items, item)
}
