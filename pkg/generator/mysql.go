package generator

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gophoria/gophoria/pkg/ast"
)

var mysqlTypes = map[ast.VariableType][]byte{
	ast.VariableTypeInt:      []byte("INT"),
	ast.VariableTypeReal:     []byte("DECIMAL"),
	ast.VariableTypeBool:     []byte("BOOL"),
	ast.VariableTypeString:   []byte("VARCHAR(255"),
	ast.VariableTypeDateTime: []byte("DATETIME"),
}

func init() {
	RegisterGenerator("mysql", NewSqlite3Generator())
}

type MysqlGenerator struct {
	ast    *ast.Ast
	writer io.Writer
	cfg    *GeneratorConfig
}

func NewMysqlGenerator() *MysqlGenerator {
	g := MysqlGenerator{}

	return &g
}

func (g *MysqlGenerator) GenerateAll(ast *ast.Ast, cfg *GeneratorConfig) error {
	g.ast = ast
	g.cfg = cfg

	err := os.MkdirAll(path.Join(g.cfg.WorkingDir, "migrations"), 0755)
	if err != nil {
		return err
	}

	for idx, model := range ast.Models {
		err := g.generateModel(model, idx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *MysqlGenerator) Generate(ast *ast.Ast, cfg *GeneratorConfig, name string) error {
	g.ast = ast
	g.cfg = cfg

	err := os.MkdirAll(path.Join(g.cfg.WorkingDir, "migrations"), 0755)
	if err != nil {
		return err
	}

	isExist := false

	for idx, model := range ast.Models {
		if model.Name.Identifier == name {
			isExist = true

			err := g.generateModel(model, idx)
			if err != nil {
				return err
			}
		}
	}

	if !isExist {
		return fmt.Errorf("model %s not found", name)
	}

	return nil
}

func (g *MysqlGenerator) generateModel(model *ast.Model, idx int) error {
	f, err := os.Create(path.Join(g.cfg.WorkingDir, "migrations", fmt.Sprintf("%d_%s.sql", idx+1, model.Name.Identifier)))
	if err != nil {
		return err
	}
	defer f.Close()

	g.writer = f

	g.writer.Write([]byte("CREATE TABLE IF NOT EXISTS "))
	g.writer.Write([]byte(model.Name.Token.Literal))
	g.writer.Write([]byte(" (\n"))

	for idx, item := range model.Items {
		err := g.generateItem(item, idx == len(model.Items)-1)
		if err != nil {
			return err
		}
	}

	g.writer.Write([]byte(");\n\n"))

	return nil
}

func (g *MysqlGenerator) generateItem(item *ast.Declaration, isLast bool) error {
	decType, ok := g.typeToSqliteType(item.DeclarationType)
	if !ok {
		if g.isTypeModel(item.DeclarationType) {
			return nil
		} else {
			return fmt.Errorf("invalid type %s", item.DeclarationType.Name)
		}
	}

	g.writer.Write([]byte("  "))
	g.writer.Write([]byte(item.Identifier.Identifier))
	g.writer.Write([]byte(" "))
	g.writer.Write(decType)

	if _, ok := g.getDecorator("id", item.Decorators); ok {
		g.writer.Write([]byte(" PRIMARY KEY"))
	}

	if !isLast {
		g.writer.Write([]byte(","))
	}
	g.writer.Write([]byte("\n"))

	return nil
}

func (g *MysqlGenerator) typeToSqliteType(decType *ast.DeclarationType) ([]byte, bool) {
	sqlType, ok := mysqlTypes[decType.Type]
	if ok {
		return sqlType, true
	}

	name := decType.Name

	for _, enum := range g.ast.Enums {
		if enum.Name.Identifier == name {
			if len(enum.Items) > 0 {
				enumType := enum.Items[0].Value.Type
				if enumType == ast.ValueTypeString {
					return []byte("VARCHAR(255)"), true
				}
				if enumType == ast.ValueTypeInt {
					return []byte("INT"), true
				}
			}
		}
	}

	return []byte{}, false
}

func (g *MysqlGenerator) isTypeModel(decType *ast.DeclarationType) bool {
	name := decType.Name

	for _, model := range g.ast.Models {
		if model.Name.Identifier == name {
			return true
		}
	}

	return false
}

func (g *MysqlGenerator) getDecorator(name string, decorators []*ast.Decorator) (*ast.Decorator, bool) {
	for _, dec := range decorators {
		if dec.Name.Identifier == name {
			return dec, true
		}
	}

	return nil, false
}