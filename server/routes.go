package server

import (
	"fmt"
	"net/http"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	page := &Page{Title: "aflm.dev", ID: "home", Posts: existingPosts}
	s.render(w, page, "home.html")
}

func (s *Server) about(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "About me", ID: "about"}
	s.render(w, page, "about.html")
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	post, err := findPost(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.render(w, &post, "post.html")
}

func (s *Server) posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list posts")
}
