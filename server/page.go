package server

import (
	"embed"
	"path/filepath"
)

// Page represents a webpage on the blog
type Page struct {
	ID    string
	Title string
	Body  string
	Posts []Post
}

func (p *Page) load(fs embed.FS) {
	filename := filepath.Join("blog", "pages", p.ID+".html")
	data, _ := fs.ReadFile(filename)
	p.Body = string(data)
}
