package ast

type Ast struct {
	Config []*Config
	Enums  []*Enum
	Models []*Model
}

func NewAst() *Ast {
	ast := Ast{
		Config: []*Config{},
		Enums:  []*Enum{},
		Models: []*Model{},
	}

	return &ast
}

func (a *Ast) AddConfig(config *Config) {
	a.Config = append(a.Config, config)
}
