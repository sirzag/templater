package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func GetTemplateDir() string {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		// Fall back to current directory if can't get home
		return "./templates"
	}

	// Create a single consistent location
	templateDir := filepath.Join(home, ".templater", "templates")

	// Create directory if it doesn't exist
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		os.MkdirAll(templateDir, 0755)
		fmt.Printf("Created template directory at: %s\n", templateDir)
	}

	return templateDir
}

func CheckTemplates(dir string) bool {
	// Read directory contents
	files, err := os.ReadDir(dir)
	if err != nil || len(files) == 0 {
		fmt.Printf("Warning: No templates found in %s\n", dir)
		return false
	}
	return true
}

func ApplyExtension(file string, ext string) string {
	if strings.HasSuffix(file, ext) {
		return file
	}
	return file + "." + strings.ReplaceAll(ext, ".", "")
}

// should transform snakecase and/or spacings to camelcase
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	parts := strings.Split(s, "/")

	for i, part := range parts {
		if i > 0 {
			result.WriteRune('/')
		}

		if part == "" {
			continue
		}

		startsWithUnderscore := strings.HasPrefix(part, "_")
		if startsWithUnderscore {
			result.WriteRune('_')
			part = part[1:]
		}

		isAllCaps := true
		hasLetter := false
		for _, c := range part {
			if unicode.IsLetter(c) {
				hasLetter = true
				if !unicode.IsUpper(c) {
					isAllCaps = false
					break
				}
			}
		}
		isAllCaps = isAllCaps && hasLetter

		var wordBuilder strings.Builder
		var capitalizeNext bool
		for j, c := range part {
			if c == '_' || c == ' ' {
				capitalizeNext = true
				continue
			}

			if j == 0 {
				if isAllCaps {
					wordBuilder.WriteRune(unicode.ToLower(c))
				} else {
					wordBuilder.WriteRune(c)
				}
			} else if capitalizeNext {
				wordBuilder.WriteRune(unicode.ToUpper(c))
				capitalizeNext = false
			} else if unicode.IsUpper(c) && !isAllCaps {
				wordBuilder.WriteRune(c)
			} else {
				wordBuilder.WriteRune(unicode.ToLower(c))
			}
		}

		result.WriteString(wordBuilder.String())
	}

	return result.String()
}

// Should replace spaces with _ and transform camelcase to snakecase
func ToSnakeCase(s string) string {
	// Return empty string if input is empty
	if len(s) == 0 {
		return ""
	}

	// Trim leading and trailing whitespace
	s = strings.TrimSpace(s)

	// Replace spaces with underscores
	s = strings.ReplaceAll(s, " ", "_")

	// Handle consecutive uppercase letters (like "JSONParser" -> "json_parser")
	re := regexp.MustCompile(`([A-Z])([A-Z][a-z])`)
	s = re.ReplaceAllString(s, "${1}_${2}")

	// Handle camelCase (like "fileName" -> "file_name")
	re = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = re.ReplaceAllString(s, "${1}_${2}")

	// Replace consecutive underscores with a single underscore
	re = regexp.MustCompile(`_+`)
	s = re.ReplaceAllString(s, "_")

	// Convert to lowercase
	return strings.ToLower(s)
}


func CreateFileWithDir(pathInput string, content []byte) error {
	// Split the path into directory and filename components
	dir, fileName := filepath.Split(pathInput)

	// If dir is empty, it means the user only provided a filename
	// Use current directory in that case (or you could use another default)
	if dir == "" {
		// Using current directory as default
		dir = "."
	}

	// Ensure the directory exists
	if dir != "." {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// Directory doesn't exist, create it
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return fmt.Errorf("error creating directory: %w", err)
			}
			fmt.Printf("Created directory: %s\n", dir)
		}
	}

	// Create or open the file
	file, err := os.OpenFile(filepath.Join(dir, fileName), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	// Write content to the file if provided
	if len(content) > 0 {
		_, err = file.Write(content)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	fmt.Printf("Successfully created file: %s\n", filepath.Join(dir, fileName))
	return nil
}
