package utils_test

import (
	"testing"

	"github.com/sirzag/templater/internal/utils"
)

func TestApplyExtension(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		extension string
		expected  string
	}{
		{"Test with .txt extension", "file", ".txt", "file.txt"},
		{"Test with .jpg extension", "image", ".jpg", "image.jpg"},
		{"Test with existing extension", "document.pdf", ".pdf", "document.pdf"},
		{"Test with empty input", "", ".txt", ".txt"},
		{"Test with empty extension", "file", "", "file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ApplyExtension(tt.input, tt.extension)
			if result != tt.expected {
				t.Errorf("ApplyExtension(%q, %q) = %q; want %q", tt.input, tt.extension, result, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test with snake_case", "file_name", "fileName"},
		{"Test with spaces", "file name", "fileName"},
		{"Test with camelCase", "fileName", "fileName"},
		{"Test with CamelCase", "FileName", "FileName"},
		{"Test with empty input", "", ""},
		{"Test first letter stays as is", "nested/path name/File name_test.js", "nested/pathName/FileNameTest.js"},
		{"Test with multiple underscores", "user_first_name", "userFirstName"},
		{"Test with multiple spaces", "user first name", "userFirstName"},
		{"Test with mixed delimiters", "user_first name", "userFirstName"},
		{"Test with numbers", "page_1_details", "page1Details"},
		{"Test with leading underscore", "_hidden_file", "_hiddenFile"},
		{"Test with trailing underscore", "temp_file_", "tempFile"},
		{"Test with double underscores", "base__name", "baseName"},
		{"Test with all caps", "API_KEY", "apiKey"},
		{"Test with mixed case", "Mixed_Case_String", "MixedCaseString"},
		{"Test with period", "file.name", "file.name"},
		{"Test with complex path", "/root/nested_dir/file_name.js", "/root/nestedDir/fileName.js"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ToCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test with CamelCase", "FileName", "file_name"},
		{"Test with spaces", "file name", "file_name"},
		{"Test with snake_case", "file_name", "file_name"},
		{"Test with snake_case and spaces", "file_name ", "file_name"},
		{"Test with empty input", "", ""},

		{"Test with camelCase", "fileName", "file_name"},
		{"Test with multiple CamelCase words", "FileNameTest", "file_name_test"},
		{"Test with abbreviations", "APIKey", "api_key"},
		{"Test with consecutive uppercase", "JSONParser", "json_parser"},
		{"Test with numbers", "Page1Details", "page1_details"},
		{"Test with mixed case and spaces", "File Name Test", "file_name_test"},
		{"Test with mixed case and underscores", "File_Name_Test", "file_name_test"},
		{"Test with leading spaces", "  FileName", "file_name"},
		{"Test with trailing spaces", "FileName  ", "file_name"},
		{"Test with mixed delimiters", "File_Name Test", "file_name_test"},
		{"Test with leading underscore", "_hiddenFile", "_hidden_file"},
		{"Test with trailing underscore", "tempFile_", "temp_file_"},
		{"Test with special characters", "user.name", "user.name"},
		{"Test with mixed case and special characters", "User.Name", "user.name"},
		{"Test with path separator", "path/to/FileName", "path/to/file_name"},
		{"Test with double uppercase", "UserID", "user_id"},
		{"Test with all uppercase", "CONSTANT", "constant"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
