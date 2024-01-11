package utils

import (
	"io"

	"github.com/gophoria/gophoria/pkg/code"
)

type GenerateConfig struct {
	DbProvider string
	DbUrl      string
	DbLib      string

	UiLib        string
	UiComponents string

	WithExample bool
}

func GenerateProject(writer io.Writer, cfg GenerateConfig) {
	generateProjectDb(writer, cfg)
	writer.Write([]byte("\n"))
	generateProjectUi(writer, cfg)

	if cfg.WithExample {
		writer.Write([]byte("\n"))
		writer.Write(code.ExampleProject)
	}
}

func generateProjectDb(writer io.Writer, cfg GenerateConfig) {
	writer.Write([]byte("db {\n"))
	writer.Write([]byte("  provider = \""))
	writer.Write([]byte(cfg.DbProvider))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  url = \""))
	writer.Write([]byte(cfg.DbUrl))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  lib = \""))
	writer.Write([]byte(cfg.DbLib))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("}\n"))
}

func generateProjectUi(writer io.Writer, cfg GenerateConfig) {
	writer.Write([]byte("ui {\n"))
	writer.Write([]byte("  lib = \""))
	writer.Write([]byte(cfg.UiLib))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("  components = \""))
	writer.Write([]byte(cfg.UiComponents))
	writer.Write([]byte("\"\n"))
	writer.Write([]byte("}\n"))
}
