package template_config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirzag/templater/internal/prompter"
	"github.com/sirzag/templater/internal/utils"
)

type TemplateConfig struct {
	OutDir      string         `json:"out,omitempty"`
	Cmd         string         `json:"cmd,omitempty"`
	Description string         `json:"description,omitempty"`
	Template    string         `json:"template,omitempty"`
	Ext         string         `json:"extension,omitempty"`
	Options     []OptionConfig `json:"options,omitempty"`

	configDir string
	target  string // pathLike
}

func Parse(path string) (*TemplateConfig, error) {
	fsInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot read template: %w", err)
	}
	if fsInfo.IsDir() {
		return nil, fmt.Errorf("Expected config file - got dir: %v", path)
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot read file: %w", err)
	}


	conf := &TemplateConfig{
		configDir: filepath.Dir(path),
	}

	fileData = replaceVariables(fileData, map[string]string{"FILE_NAME": conf.GetFileName()})
	fileData = replaceComments(fileData)

	err = json.Unmarshal(fileData, conf)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse template: %w", err)
	}

	err = conf.checkOutDir()
	if err != nil {
		return nil, err
	}
	err = conf.checkTemplate()
	if err != nil {
		return nil, err
	}

	return conf, err
}

func (tc *TemplateConfig) GetOutFile() string {
	path := tc.target
	return filepath.Join(tc.OutDir, utils.ApplyExtension(path, tc.Ext))
}

func (tc *TemplateConfig) GetFileName() string {
	ext := filepath.Ext(tc.target)
	return strings.TrimSuffix(filepath.Base(tc.target), ext)
}

func (tc *TemplateConfig) ConfigureTarget() {
	FILE_NAME := prompter.PromptString(&prompter.PromptConfig{
		Description: "Enter file name",
		Required:    true,
	})

	FILE_NAME = utils.ToCamelCase(FILE_NAME)
	tc.target = FILE_NAME
}

func (tc *TemplateConfig) checkOutDir() error {
	if strings.HasPrefix(tc.OutDir, ".") {
		tc.OutDir = filepath.Join(tc.configDir, tc.OutDir)
	}

	outDir, err := os.Stat(tc.OutDir)
	if err != nil || !outDir.IsDir() {
		return fmt.Errorf("Invalid out dir: %v", tc.OutDir)
	}

	return nil
}

func (tc *TemplateConfig) checkTemplate() error {
	if strings.HasPrefix(tc.Template, ".") {
		tc.Template = filepath.Join(tc.configDir, tc.Template)
	}

	templFile, err := os.Stat(tc.Template)
	if err != nil || templFile.IsDir() {
		return fmt.Errorf("Invalid template file: %v", tc.Template)
	}
	return nil
}
