package web

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var (
	port          string
	existingPosts posts
	templates     *template.Template
)

func init() {
	port = os.Getenv("PORT")
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
