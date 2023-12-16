package web

import (
	"net/http"
	"path/filepath"
)

// Page represents a webpage on the blog
type Page struct {
	ID    string
	Title string
	Body  string
	Posts []Post
}

func (p *Page) loadPage() {
	filename := filepath.Join("pages", p.ID+".html")
	p.Body = string(readFromFile(filename))
}

func (p *Page) renderTemplate(w http.ResponseWriter, template string) {
	p.loadPage()
	if err := templates.ExecuteTemplate(w, template, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	page := Page{Title: "alamp.dev", ID: "home", Posts: existingPosts}
	page.renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "About me", ID: "about"}
	page.renderTemplate(w, "about.html")
}
