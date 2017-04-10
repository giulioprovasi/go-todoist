package main

import (
	"net/http"
	"todo/server"
	"github.com/GeertJohan/go.rice"
	"log"
	"path/filepath"
	"os"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	log.Println(filepath.Dir(os.Args[0]))
	http.ListenAndServe(":8080", nil)
}