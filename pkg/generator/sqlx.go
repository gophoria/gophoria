package generator

import (
	"fmt"
	"io"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/gophoria/gophoria/pkg/ast"
	"github.com/gophoria/gophoria/pkg/code"
)

type SqlxGenerator struct {
	ast    *ast.Ast
	writer io.Writer
}

func NewSqlxGenerator() *SqlxGenerator {
	g := SqlxGenerator{}

	return &g
}

func (g *SqlxGenerator) GenerateAll(ast *ast.Ast, writer io.Writer) error {
	g.ast = ast
	g.writer = writer

	for _, enum := range ast.Enums {
		err := g.generateEnum(enum)
		if err != nil {
			return err
		}
	}

	for _, model := range ast.Models {
		err := g.generateModel(model)
		if err != nil {
			return err
		}
		err = g.generateStorage(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *SqlxGenerator) Generate(ast *ast.Ast, name string, writer io.Writer) error {
	g.ast = ast
	g.writer = writer

	isExist := false

	if name == "DateTime" {
		return g.generateDateTime(ast, writer)
	}

	for _, enum := range ast.Enums {
		if enum.Name.Identifier == name {
			isExist = true
			err := g.generateEnum(enum)
			if err != nil {
				return err
			}
		}
	}

	for _, model := range ast.Models {
		if model.Name.Identifier == name {
			isExist = true
			err := g.generateModel(model)
			if err != nil {
				return err
			}
			err = g.generateStorage(model)
			if err != nil {
				return err
			}
		}
	}

	if !isExist {
		return fmt.Errorf("enum or model %s not found", name)
	}

	return nil
}

func (g *SqlxGenerator) generateModel(model *ast.Model) error {
	g.writer.Write([]byte(`import "github.com/jmoiron/sqlx"
	`))

	g.writer.Write([]byte("type "))
	g.writer.Write([]byte(model.Name.Identifier))
	g.writer.Write([]byte(" struct {\n"))

	for _, item := range model.Items {
		err := g.generateModelItem(item)
		if err != nil {
			return err
		}
	}

	g.writer.Write([]byte("}\n\n"))

	return nil
}

func (g *SqlxGenerator) generateModelItem(item *ast.Declaration) error {
	g.writer.Write([]byte("  "))
	g.writer.Write([]byte(utils.Capitalize(item.Identifier.Identifier)))
	g.writer.Write([]byte(" "))

	switch item.DeclarationType.Type {
	case ast.VariableTypeInt:
		g.writer.Write([]byte("int"))
	case ast.VariableTypeReal:
		g.writer.Write([]byte("float64"))
	case ast.VariableTypeBool:
		g.writer.Write([]byte("bool"))
	case ast.VariableTypeString:
		g.writer.Write([]byte("string"))
	case ast.VariableTypeDateTime:
		g.writer.Write([]byte("DateTime"))
	case ast.VariableTypeObject:
		if g.isTypeModel(item.DeclarationType) {
			if item.DeclarationType.IsArray {
				g.writer.Write([]byte("[]"))
			}
			g.writer.Write([]byte("*"))
			g.writer.Write([]byte(item.DeclarationType.Name))
		} else if g.isTypeEnum(item.DeclarationType) {
			g.writer.Write([]byte(item.DeclarationType.Name))
		} else {
			return fmt.Errorf("not supported type (%s) for item %s", item.DeclarationType.Name, item.Identifier.Identifier)
		}
	default:
		return fmt.Errorf("not supported type (%s) for item %s", item.DeclarationType.Name, item.Identifier.Identifier)
	}
	g.writer.Write([]byte("\n"))

	return nil
}

func (g *SqlxGenerator) generateEnum(enum *ast.Enum) error {
	if len(enum.Items) == 0 {
		return fmt.Errorf("enum %s is empty", enum.Name.Identifier)
	}

	g.writer.Write([]byte("type "))
	g.writer.Write([]byte(enum.Name.Identifier))

	valueType := enum.Items[0].Value.Type
	switch valueType {
	case ast.ValueTypeInt:
		g.writer.Write([]byte(" int"))
	case ast.ValueTypeString:
		g.writer.Write([]byte(" string"))
	default:
		return fmt.Errorf("enum %s contains not supported type", enum.Name.Identifier)
	}
	g.writer.Write([]byte("\n\n"))

	g.writer.Write([]byte("const (\n"))
	for idx, item := range enum.Items {
		g.writer.Write([]byte("  "))
		g.writer.Write([]byte(enum.Name.Identifier))
		g.writer.Write([]byte(utils.Capitalize(item.Identifier.Identifier)))
		if idx == 0 {
			g.writer.Write([]byte(" "))
			g.writer.Write([]byte(enum.Name.Identifier))
		}
		g.writer.Write([]byte(" = "))
		if valueType == ast.ValueTypeString {
			g.writer.Write([]byte("\""))
		}
		g.writer.Write([]byte(item.Value.Value))
		if valueType == ast.ValueTypeString {
			g.writer.Write([]byte("\""))
		}
		g.writer.Write([]byte("\n"))
	}
	g.writer.Write([]byte(")\n\n"))

	return nil
}

func (g *SqlxGenerator) generateDateTime(ast *ast.Ast, writer io.Writer) error {
	writer.Write(code.DateTime)

	return nil
}

func (g *SqlxGenerator) isTypeModel(decType *ast.DeclarationType) bool {
	name := decType.Name

	for _, model := range g.ast.Models {
		if model.Name.Identifier == name {
			return true
		}
	}

	return false
}

func (g *SqlxGenerator) isTypeEnum(decType *ast.DeclarationType) bool {
	name := decType.Name

	for _, enum := range g.ast.Enums {
		if enum.Name.Identifier == name {
			return true
		}
	}

	return false
}

func (g *SqlxGenerator) generateStorage(model *ast.Model) error {
	g.writer.Write(code.GenerateStorageBytes(model.Name.Identifier))
	g.writer.Write(code.GenerateNewStorageBytes(model.Name.Identifier))
	return nil
}

func (g *SqlxGenerator) generateCRUD(ast *ast.Ast, writer io.Writer) error {

	return nil
}
