package template_config

import (
	"bytes"
	"os"
	"regexp"
)

func replaceComments(data []byte) []byte {
	// Remove single-line comments
	re := regexp.MustCompile(`//.*`)
	noComments := re.ReplaceAllString(string(data), "")

	// Remove multi-line comments (this is simplistic and may not work for all cases)
	re = regexp.MustCompile(`(?s)/\*.*?\*/`)
	noComments = re.ReplaceAllString(noComments, "")

	return []byte(noComments)
}

func replaceVariables(data []byte, params map[string]string) []byte {
	wd, _ := os.Getwd()
	data = bytes.Replace(data, []byte("{WORKING_DIR}"), []byte(wd), -1)
	// data = bytes.Replace(data, []byte("{FILE_NAME}"), []byte(params["FILE_NAME"]), -1)
	return data
}
