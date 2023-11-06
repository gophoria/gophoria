package utils

import "io"

func GenerateProject(writer io.Writer, provider string, url string) {
	writer.Write([]byte("db {\n"))
	writer.Write([]byte("  provider = \""))
	writer.Write([]byte(provider))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  url = \""))
	writer.Write([]byte(url))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("}\n"))
}
