package utils

import (
	"io"

	"github.com/gophoria/gophoria/pkg/code"
)

func GenerateProject(writer io.Writer, provider string, url string, withExample bool) {
	writer.Write([]byte("db {\n"))
	writer.Write([]byte("  provider = \""))
	writer.Write([]byte(provider))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  url = \""))
	writer.Write([]byte(url))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  lib = \"sqlx\"\n"))
	writer.Write([]byte("}\n"))

	if withExample {
		writer.Write([]byte("\n"))
		writer.Write(code.ExampleProject)
	}
}
