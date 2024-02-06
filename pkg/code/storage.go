package code

import "fmt"

func GenerateStore(name string) []byte {
	return []byte(
		fmt.Sprintf(`type %sStore struct {
	conn	*sqlx.DB
}

`, name))
}

func GenerateNewStore(name string) []byte {
	return []byte(
		fmt.Sprintf(`func New%[1]sStore(conn *sqlx.DB) *%[1]sStore {
	return &%[1]sStore{conn: conn}
}

`, name))
}

func generateInsertQuery(name string, items []string) string {
	tmp := ""
	tmpVar := ""
	for i, item := range items {
		if i == 0 {
			tmp = item
			tmpVar = fmt.Sprintf(":%s", item)
		} else {
			tmp = fmt.Sprintf("%s,\n\t%s", tmp, item)
			tmpVar = fmt.Sprintf("%s,\n\t:%s", tmpVar, item)
		}
	}
	query := fmt.Sprintf(`INSERT INTO %s (
	%s)
	VALUES (
	%s)`, name, tmp, tmpVar)
	return query
}

func GenerateStoreInsertMethod(name string, items []string) []byte {
	query := "`" + generateInsertQuery(name, items) + "`"
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) Insert(p %[1]s) error {
	_, err := s.conn.NamedExec(%s, p)
	
	if err != nil {
		return err
	}

	return  nil
}

`, name, query))
}

func generateUpdateQuery(name string, id string, items []string) string {
	tmp := ""
	for i, item := range items {
		if i == 0 {
			tmp = fmt.Sprintf("%[1]s=:%[1]s", item)
		} else {
			tmp = fmt.Sprintf("%s,\n\t%[2]s=:%[2]s", tmp, item)
		}
	}
	query := fmt.Sprintf("UPDATE %[1]s SET %s\n	WHERE %[3]s=:%[3]s", name, tmp, id)
	return query
}

func GenerateStoreUpdateMethod(name string, items []string) []byte {
	query := "`" + generateUpdateQuery(name, items[0], items[1:]) + "`"
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) Update(p %[1]s) error {
	_, err := s.conn.NamedExec(%s, p)

	if err != nil {
		return err
	}

	return  nil
}

`, name, query, items[0]))
}

func GenerateStorSaveMethod(name string, items []string) []byte {
	query := "`" + generateUpdateQuery(name, items[0], items[1:]) + "`"
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) Save(p %[1]s) error {
	if song.Id == "" {
		return s.Insert(song)
	}

	return s.Update(song)
}

`, name, query, items[0]))
}

func GenerateStoreDeleteMethod(name string, items []string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) Delete(p %[1]s) error {
	query := "DELETE FROM %[1]s WHERE %[2]s=:%[2]s"
	_, err := s.conn.NamedExec(query, p)
  return err
}

	`, name, items[0]))
}

func GenerateStoreGetAllMethod(name string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) GetAll() ([]*%[1]s, error) {
	var result []*%[1]s
	query := "SELECT * FROM %[1]s"
	err := s.conn.Select(&result, query)
	return  result, err
}

	`, name))
}

func GenerateStoreGetByIdMethod(name string, items []string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStore) GetById(%[2]s int) (*%[1]s, error) {
	var result %[1]s
	query := "SELECT * FROM %[1]s WHERE %[2]s=:%[2]s"
	err := s.conn.Get(&result, query, p.%[2]s)
	return  &result, err
}

	`, name, items[0]))
}
