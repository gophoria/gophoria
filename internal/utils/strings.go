package utils

import "fmt"

func Capitalize(str string) string {
	return fmt.Sprintf("%c%s", str[0]-32, str[1:])
}
