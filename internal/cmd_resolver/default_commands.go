package cmd_resolver

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/sirzag/templater/internal/utils"
)

func OpenTemplateDir() error {
	tempDir := utils.GetTemplateDir()

	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", tempDir).Start()
	case "linux":
		return exec.Command("xdg-open", tempDir).Start()
	case "windows":
		return exec.Command("explorer", tempDir).Start()
	default:
		return fmt.Errorf("Unsupported OS: %v", runtime.GOOS)
	}
}
