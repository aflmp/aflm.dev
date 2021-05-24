package server

import (
	"regexp"
	"strings"
)

func idFromTitle(title string) string {
	validURLChars := regexp.MustCompile("[^a-zA-Z0-9-_ ]")
	dashes := regexp.MustCompile("-[-]*")

	id := validURLChars.ReplaceAllString(title, "")
	id = strings.ReplaceAll(strings.ToLower(id), " ", "-")
	return dashes.ReplaceAllString(id, "-")
}

func idFromPath(path string) string {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	return strings.Replace(path, "/posts/", "", 1)
}
