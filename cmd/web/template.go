package web

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func loadTemplates(templateDir string) []string {
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatalf("failed to load templates from %v: %v", templateDir, err)
	}

	result := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") && !file.IsDir() {
			result = append(result, filepath.Join(templateDir, file.Name()))
		}
	}

	return result
}
