package generator

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gophoria/gophoria/internal/code"
	"github.com/gophoria/gophoria/internal/utils"
	"github.com/gophoria/gophoria/pkg/ast"
)

type SqlxGenerator struct {
	ast    *ast.Ast
	writer io.Writer
	cfg    *GeneratorConfig
}

func init() {
	RegisterGenerator("sqlx", NewSqlxGenerator())
}

func NewSqlxGenerator() *SqlxGenerator {
	g := SqlxGenerator{}

	return &g
}

func (g *SqlxGenerator) GenerateAll(ast *ast.Ast, cfg *GeneratorConfig) error {
	g.ast = ast
	g.cfg = cfg

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
		err = g.generateStore(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *SqlxGenerator) Generate(ast *ast.Ast, cfg *GeneratorConfig, name string) error {
	g.ast = ast
	g.cfg = cfg

	isExist := false

	if name == "DateTime" {
		return g.generateDateTime(ast, g.writer)
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
			err = g.generateStore(model)
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
	f, err := os.Create(path.Join(g.cfg.WorkingDir, "db", fmt.Sprintf("%s.go", model.Name.Identifier)))
	if err != nil {
		return err
	}
	defer f.Close()
	g.writer = f

	g.writer.Write([]byte("import (\n\"github.com/jmoiron/sqlx\"\n)\n\n"))

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
	g.writer.Write([]byte(" `db:\""))
	g.writer.Write([]byte(item.Identifier.Identifier))
	g.writer.Write([]byte("\"`"))
	g.writer.Write([]byte("\n"))

	return nil
}

func (g *SqlxGenerator) generateEnum(enum *ast.Enum) error {
	if len(enum.Items) == 0 {
		return fmt.Errorf("enum %s is empty", enum.Name.Identifier)
	}

	f, err := os.Create(path.Join(g.cfg.WorkingDir, "db", fmt.Sprintf("%s.go", enum.Name.Identifier)))
	if err != nil {
		return err
	}
	defer f.Close()
	g.writer = f

	g.writer.Write([]byte("package db\n\n"))

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
	for _, item := range enum.Items {
		g.writer.Write([]byte("  "))
		g.writer.Write([]byte(enum.Name.Identifier))
		g.writer.Write([]byte(utils.Capitalize(item.Identifier.Identifier)))
		g.writer.Write([]byte(" "))
		g.writer.Write([]byte(enum.Name.Identifier))
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

func (g *SqlxGenerator) generateStore(model *ast.Model) error {
	var storeItems []string
	for _, item := range model.Items {
		if item.DeclarationType.Type == ast.VariableTypeObject {
			if !g.isTypeEnum(item.DeclarationType) {
				continue
			}
		}
		storeItems = append(storeItems, string(item.Identifier.Identifier))
	}

	g.writer.Write(code.GenerateStore(model.Name.Identifier))
	g.writer.Write(code.GenerateNewStore(model.Name.Identifier))
	g.writer.Write(code.GenerateStoreInsertMethod(model.Name.Identifier, storeItems))
	g.writer.Write(code.GenerateStoreUpdateMethod(model.Name.Identifier, storeItems))
	g.writer.Write(code.GenerateStoreDeleteMethod(model.Name.Identifier, storeItems))
	g.writer.Write(code.GenerateStoreGetAllMethod(model.Name.Identifier))
	g.writer.Write(code.GenerateStoreGetByIdMethod(model.Name.Identifier, storeItems))
	return nil
}

func (g *SqlxGenerator) generateStoreNewMethod(model *ast.Model) error {
	code := fmt.Sprintf(`func New%[1]sStore(conn *sqlx.DB) *%[1]sStore {
	return &%[1]sStore{conn: conn}
}
`, model.Name.Identifier)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreInsertMethod(model *ast.Model) error {
	query := ""
	queryVar := ""

	for _, item := range model.Items {
		if item.DeclarationType.Type == ast.VariableTypeObject {
			if !g.isTypeEnum(item.DeclarationType) {
				continue
			}
		}

		if query != "" {
			query += ",\n"
			queryVar += ",\n:"
		}

		query += item.Identifier.Identifier
		queryVar += ":" + item.Identifier.Identifier
	}

	code := fmt.Sprintf(`func (s *%[1]sStore) Insert(p *%[1]s) error {
  if p.Id == "" {
    p.Id = uuid.NewString()
  }

	_, err := s.conn.NamedExec("INSERT INTO %[1]s (
    %s
  ) VALUES (
    %s
  )", p)
	
	if err != nil {
		return err
	}

	return  nil
}
`, model.Name.Identifier, query, queryVar)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreUpdateMethod(model *ast.Model) error {
	query := ""
	queryVar := ""

	for _, item := range model.Items {
		if item.DeclarationType.Type == ast.VariableTypeObject {
			if !g.isTypeEnum(item.DeclarationType) {
				continue
			}
		}

		if query != "" {
			query += ",\n"
		}

		query += item.Identifier.Identifier + "=" + item.Identifier.Identifier
	}

	code := fmt.Sprintf(`func (s *%[1]sStore) Update(p *%[1]s) error {
  if p.Id == "" {
    p.Id = uuid.NewString()
  }

	_, err := s.conn.NamedExec("UPDATE %[1]s SET
    %s
  WHERE id=:id", p)
	
	if err != nil {
		return err
	}

	return  nil
}
`, model.Name.Identifier, query, queryVar)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreSaveMethod(model *ast.Model) error {
	code := fmt.Sprintf(`func (s *%[1]sStore) Save(p *%[1]s) error {
  if p.Id == "" {
    return p.Insert()
  } else {
    return p.Update()
  }
}
`, model.Name.Identifier)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreDeleteMethod(model *ast.Model) error {
	code := fmt.Sprintf(`func (s *%[1]sStore) Delete(p *%[1]s) error {
	query := "DELETE FROM %[1]s WHERE id=:id"
	_, err := s.conn.NamedExec(query, p)
  return err
}
`, model.Name.Identifier)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreGetAllMethod(model *ast.Model) error {
	code := fmt.Sprintf(`func (s *%[1]sStore) GetAll() ([]*%[1]s, error) {
	var result []*%[1]s
	query := "SELECT * FROM %[1]s"
	err := s.conn.Select(&result, query)
	return result, err
}
`, model.Name.Identifier)

	g.writer.Write([]byte(code))
	return nil
}

func (g *SqlxGenerator) generateStoreGetbyIdMethod(model *ast.Model) error {
	code := fmt.Sprintf(`func (s *%[1]sStore) GetById(id string) (*%[1]s, error) {
	var result %[1]s
	query := "SELECT * FROM %[1]s WHERE id=?"
	err := s.conn.Get(&result, query, id)
	return  &result, err
}
`, model.Name.Identifier)

	g.writer.Write([]byte(code))
	return nil
}
