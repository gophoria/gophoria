package code

import "fmt"

func GenerateStorage(name string) []byte {
	return []byte(
		fmt.Sprintf(`type %sStorage struct {
	conn	*sqlx.DB
}

`, name))
}

func GenerateNewStorage(name string) []byte {
	return []byte(
		fmt.Sprintf(`func New%[1]sStorage(conn *sqlx.DB) *%[1]sStorage {
	return &%[1]sStorage{conn: conn}
}

`, name))
}

func generateInsertQuery(name string, items []string) string {
	tmp := ""
	tmpVar := ""
	for i, item := range items {
		if i == 0 {
			tmp = fmt.Sprintf("%s", item)
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

func GenerateStorageCreateMethod(name string, items []string) []byte {
	query := "`" + generateInsertQuery(name, items) + "`"
	return []byte(fmt.Sprintf(`func (s *%[1]sStorage) CreateNew%[1]s(p %[1]s) error {
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

func GenerateStorageUpdateMethod(name string, items []string) []byte {
	query := "`" + generateUpdateQuery(name, items[0], items[1:]) + "`"
	return []byte(fmt.Sprintf(`func (s *%[1]sStorage) Update%[1]s(p %[1]s) error {
	_, err := s.conn.NamedExec(%s, p)

	if err != nil {
		return err
	}

	return  nil
}

`, name, query, items[0]))
}

func GenerateStorageDeleteMethod(name string, items []string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStorage) Delete%[1]s(p %[1]s) error {
	query := "DELETE FROM %[1]s WHERE %[2]s=:%[2]s"
	_, err := s.conn.NamedExec(query, p)
	
	if err != nil {
		return err
	}

	return  nil
}

	`, name, items[0]))
}

func GenerateStorageGetallMethod(name string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStorage) Getall%[1]ss() ([]*%[1]s, error) {
	var result []*%[1]s
	query := "SELECT * FROM %[1]s"
	err := s.conn.Select(&result, query)
	return  result, err
}

	`, name))
}

func GenerateStorageGetByIdMethod(name string, items []string) []byte {
	return []byte(fmt.Sprintf(`func (s *%[1]sStorage) Get%[1]sById(%[2]s int) (*%[1]s, error) {
	var result %[1]s
	query := "SELECT * FROM %[1]s WHERE %[2]s=:%[2]s"
	err := s.conn.Get(&result, query)
	return  &result, err
}

	`, name, items[0]))
}
