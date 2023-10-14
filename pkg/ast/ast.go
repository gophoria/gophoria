package ast

type Enum struct {
	Items []*AssignItem
}

type Model struct {
}

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
