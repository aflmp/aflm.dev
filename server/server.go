package server

import (
	"embed"
	"log"
	"net"
	"net/http"
	"text/template"
)

var existingPosts posts

type Server struct {
	mux      *http.ServeMux
	config   Config
	template *template.Template
}

type renderable interface {
	load(embed.FS)
}

func New(config Config) *Server {
	templates := template.Must(template.ParseFS(config.Blog, "blog/templates/*.html"))
	return &Server{
		mux:      http.NewServeMux(),
		config:   config,
		template: templates,
	}
}

func (s *Server) routes() {
	s.mux.Handle("/assets/", http.FileServer(http.FS(s.config.Assets)))
	s.mux.HandleFunc("/", s.home)
	s.mux.HandleFunc("/about", s.about)
	s.mux.HandleFunc("/posts/", s.post)
	s.mux.HandleFunc("/posts", s.posts)
}

// Run does blah
func (s *Server) Run() error {
	s.routes()
	existingPosts = loadExistingPosts(s.config.PostList)
	log.Printf("server listening on port: %v", s.config.Port)
	return http.ListenAndServe(net.JoinHostPort("0.0.0.0", s.config.Port), s.mux)
}

func (s *Server) render(w http.ResponseWriter, r renderable, template string) {
	r.load(s.config.Blog)
	if err := s.template.ExecuteTemplate(w, template, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
