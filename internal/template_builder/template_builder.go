package template_builder

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/sirzag/templater/internal/template_config"
	"github.com/sirzag/templater/internal/utils"
)

func Build(ctx context.Context, config *template_config.TemplateConfig) error {
	templateInput := ctx.Value("template_input").(map[string]any)

	t, err := template.ParseFiles(config.Template)
	if err != nil {
		return fmt.Errorf("Cannot parse template: %w", err)
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, templateInput)
	if err != nil {
		return fmt.Errorf("Cannot execute template: %w", err)
	}

	processedBytes := bytes.ReplaceAll(buf.Bytes(), []byte("\n\n"), []byte("\n"))
	err = utils.CreateFileWithDir(config.GetOutFile(), processedBytes)
	return err
}
