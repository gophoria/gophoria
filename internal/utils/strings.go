package utils

import (
	"fmt"
	"io"
)

func Capitalize(str string) string {
	return fmt.Sprintf("%c%s", str[0]-32, str[1:])
}

func WriteString(writer io.Writer, str string) (int, error) {
	return writer.Write([]byte(str))
}
