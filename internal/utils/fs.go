package utils

import "os"

func CreateDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
