package template_config

import "github.com/sirzag/templater/internal/prompter"

type OptionConfig struct {
	Name        string      `json:"name" validate:"nonzero"`
	Type        string      `json:"type" validate:"nonzero,regexp=^(string|number|boolean|enum)$"`
	Required    bool        `json:"required,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty" validate:"nonzero"`
	Values      []string    `json:"values,omitempty"` // If type is enum
}

func (optConf *OptionConfig) isValid() bool {
	return true
}

func (optConf *OptionConfig) ToPromptConf() *prompter.PromptConfig {
	return &prompter.PromptConfig{
		Required:    optConf.Required,
		Default:     optConf.Default,
		Description: optConf.Description,
		Items:       optConf.Values,
	}
}
