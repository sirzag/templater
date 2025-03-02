package cmd_resolver

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/labstack/gommon/log"
	"github.com/sirzag/templater/internal/prompter"
	"github.com/sirzag/templater/internal/template_builder"
	"github.com/sirzag/templater/internal/template_config"
	"github.com/sirzag/templater/internal/utils"
	"github.com/urfave/cli/v3"
)

var (
	configPattern = regexp.MustCompile(`.*\.config\.(json|jsonc)$`)
)

func BuildCommands() ([]*cli.Command, error) {
	commands := []*cli.Command{
		{
			Name:  "locate",
			Usage: "The command will open the template directory",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				return OpenTemplateDir()
			},
		},
	}

	filepath.Walk(utils.GetTemplateDir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Errorf("Error walking path: %v", err)
			return nil
		}

		if info.IsDir() || !configPattern.MatchString(path) {
			return nil
		}

		conf, err := template_config.Parse(path)
		if err != nil {
			log.Errorf("Encountered invalid config: %v. Path: %v", err, path)
			return nil
		}

		commands = append(commands, &cli.Command{
			Name:  conf.Cmd,
			Usage: conf.Description,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				conf.ConfigureTarget()

				templateInput, err := CollectOptions(conf)
				if err != nil {
					return fmt.Errorf("Option collection failed: %w", err)
				}
				templateInput["FILE_NAME"] = conf.GetFileName()
				return template_builder.Build(context.WithValue(ctx, "template_input", templateInput), conf)
			},
		})

		return nil
	})

	return commands, nil
}

func CollectOptions(conf *template_config.TemplateConfig) (map[string]any, error) {
	templateOptions := map[string]any{}

	for _, opt := range conf.Options {
		switch opt.Type {
		case "string":
			templateOptions[opt.Name] = prompter.PromptString(opt.ToPromptConf())
		case "number":
			templateOptions[opt.Name] = prompter.PromptNumber(opt.ToPromptConf())
		case "boolean":
			templateOptions[opt.Name] = prompter.PromptBoolean(opt.ToPromptConf())
		case "enum":
			templateOptions[opt.Name] = prompter.PromptEnum(opt.ToPromptConf())
		default:
			return templateOptions, fmt.Errorf("Unsupported option type in config: %v\n", opt.Type)
		}
	}

	return templateOptions, nil
}
