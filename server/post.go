package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

type posts []Post

// Post represents a blog post
type Post struct {
	ID    string `json:"-"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Body  string `json:"-"`
}

func (p *Post) load(fs embed.FS) {
	filename := filepath.Join("blog", "posts", p.ID+".html")
	data, _ := fs.ReadFile(filename)
	p.Body = string(data)
}

func loadExistingPosts(postIndex []byte) posts {
	config := postConfig{}
	if err := json.Unmarshal(postIndex, &config); err != nil {
		log.Fatalf("failed to unmarshal postConfig: %v", err)
	}

	return config.Posts
}

func findPost(path string) (Post, error) {
	postID := idFromPath(path)
	for _, post := range existingPosts {
		if post.ID == postID {
			return post, nil
		}
	}

	return Post{}, fmt.Errorf("post: %v not found", path)
}
