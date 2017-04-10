package main

import (
	"net/http"
	"github.com/kitensei/go-todoist/server"
	"github.com/GeertJohan/go.rice"
	"os"
	"log"
	"fmt"
)

var boxPrefix = getenv("BOXPATH", "")

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DIRECTORY CWD: " + dir)
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(rice.MustFindBox(boxPrefix + "static").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
