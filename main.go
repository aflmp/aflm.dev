package main

import (
	"embed"
	"log"
	"syscall"

	server "github.com/aflmp/aflm.dev/server"
)

//go:embed posts.json
var postList []byte

//go:embed assets/*
var assets embed.FS

//go:embed blog/*
var blog embed.FS

func main() {
	port, found := syscall.Getenv("PORT")
	if !found {
		port = "3000"
	}

	config := server.Config{
		Port:     port,
		Blog:     blog,
		PostList: postList,
		Assets:   assets,
	}

	srv := server.New(config)
	err := srv.Run()
	log.Fatal(err)
}
