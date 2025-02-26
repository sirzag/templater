package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getTemplateDir() string {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		// Fall back to current directory if can't get home
		return "./templates"
	}

	// Create a single consistent location
	templateDir := filepath.Join(home, ".yourapp", "templates")

	// Create directory if it doesn't exist
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		os.MkdirAll(templateDir, 0755)
		fmt.Printf("Created template directory at: %s\n", templateDir)
	}

	return templateDir
}

func checkTemplates(dir string) bool {
	// Read directory contents
	files, err := os.ReadDir(dir)
	if err != nil || len(files) == 0 {
		fmt.Printf("Warning: No templates found in %s\n", dir)
		return false
	}
	return true
}

func main() {
	p, _ := os.Getwd()
	fmt.Println("Work dir:", p)

	templateDir := getTemplateDir()
	hasTemplates := checkTemplates(templateDir)

	if !hasTemplates {
		fmt.Printf("Please add template files to: %s\n", templateDir)
		// Exit or continue with defaults
	}
}
