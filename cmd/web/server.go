package web

import (
	"log"
	"net/http"
	"syscall"
	"text/template"
)

var (
	port          string
	existingPosts posts
	templates     *template.Template
)

func init() {
	var found bool

	port, found = syscall.Getenv("PORT")
	if !found {
		port = "3000"
	}

	existingPosts = loadExistingPosts()
	files := loadTemplates("templates")
	templates = template.Must(template.ParseFiles(files...))
}

// Main contains the http server
func Main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/about/", about)
	mux.HandleFunc("/posts/", post)

	log.Printf("server listening on port: %v", port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
