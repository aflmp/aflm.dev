package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

// postConfig represents posts.yaml
type postConfig struct {
	Posts posts `json:"posts"`
}

type posts []Post

// Post represents a blog post
type Post struct {
	ID    string `json:"-"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Body  string `json:"-"`
}

func postIDfromTitle(title string) string {
	validURLChars := regexp.MustCompile("[^a-zA-Z0-9-_ ]")
	dashes := regexp.MustCompile("-[-]*")

	id := validURLChars.ReplaceAllString(title, "")
	id = strings.ReplaceAll(strings.ToLower(id), " ", "-")
	return dashes.ReplaceAllString(id, "-")
}

func postIDfromPath(path string) string {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	return strings.Replace(path, "/posts/", "", 1)
}

func (p *postConfig) UnmarshalJSON(data []byte) error {
	var pc struct {
		Posts []struct {
			Title string `json:"title"`
			Date  string `json:"date"`
		} `json:"Posts"`
	}

	err := json.Unmarshal(data, &pc)
	if err != nil {
		return err
	}

	p.Posts = make(posts, len(pc.Posts))
	for i, post := range pc.Posts {
		p.Posts[i].Title = post.Title
		p.Posts[i].Date = post.Date
		p.Posts[i].ID = postIDfromTitle(post.Title)
	}

	return nil
}

func loadExistingPosts() posts {
	data := readFromFile("./posts.json") // TODO: load as an argument
	config := postConfig{}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("failed to unmarshal postConfig: %v", err)
	}

	return config.Posts
}

func findPost(path string) (Post, error) {
	postID := postIDfromPath(path)
	for _, post := range existingPosts {
		if post.ID == postID {
			return post, nil
		}
	}

	return Post{}, fmt.Errorf("post: %v not found", path)
}

func (p *Post) loadPost() {
	filename := filepath.Join("posts", p.ID+".html")
	p.Body = string(readFromFile(filename))
}

func (p *Post) renderTemplate(w http.ResponseWriter, template string) {
	p.loadPost()
	if err := templates.ExecuteTemplate(w, template, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	post, err := findPost(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post.renderTemplate(w, "post.html")
}
