package code

import "fmt"

func GenerateStorageBytes(name string) []byte {
	return []byte(fmt.Sprintf(`
	 type %sStorage struct {
		conn	*sqlx.DB
	 }
	
	`, name))
}

func GenerateNewStorageBytes(name string) []byte {
	return []byte(fmt.Sprintf(`
	func New%[1]sStorage(conn *sqlx.DB) *%[1]sStorage {
		return &%[1]sStorage{conn: conn}
	}
	`, name))
}
