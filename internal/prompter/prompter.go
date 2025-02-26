package prompter

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/manifoldco/promptui"
)

type PromptConfig struct {
	Required    bool
	Default     interface{}
	Description string
	Items       []string
}

func PromptString(opt *PromptConfig) string {
	prompt := promptui.Prompt{
		Label: opt.Description,
		Validate: func(input string) error {
			if len(input) == 0 && opt.Required {
				return fmt.Errorf("The input is required")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed %w\n", err)
	}
	if len(result) == 0 && opt.Default != nil {
		result = opt.Default.(string)
	}

	return result
}

func PromptNumber(opt *PromptConfig) int64 {
	prompt := promptui.Prompt{
		Label: opt.Description,
		Validate: func(input string) error {
			if len(input) == 0 && opt.Required {
				return fmt.Errorf("The input is required")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed %w\n", err)
	}
	if len(result) == 0 && opt.Default != nil {
		result = opt.Default.(string)
	}

	num, _ := strconv.ParseInt(result, 10, 32)
	if err != nil {
		log.Fatal("Prompt failed %w\n", err)
	}
	return num
}

func PromptBoolean(opt *PromptConfig) bool {
	prompt := promptui.Select{
		Label:     opt.Description,
		Items:     []string{"Yes", "No"},
		CursorPos: map[bool]int{true: 0, false: 1}[opt.Default.(bool)],
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed %w\n", err)
	}
	return map[string]bool{"Yes": true, "No": false}[result]
}

func PromptEnum(opt *PromptConfig) string {
	prompt := promptui.Select{
		Label:     opt.Description,
		Items:     opt.Items,
		CursorPos: slices.Index(opt.Items, opt.Default.(string)),
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed %w\n", err)
	}
	return result
}
