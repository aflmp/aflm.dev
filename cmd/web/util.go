package web

import (
	"io/ioutil"
	"log"
	"strings"
)

func readFromFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %v: %v", filename, err)
	}

	return data
}

func contains(title string, posts []Post) bool {
	for _, post := range posts {
		if strings.HasPrefix(post.Title, title) {
			return true
		}
	}

	return false
}
