package main

import (
	"net/http"
	"github.com/kitensei/go-todoist/server"
	"os"
	"log"
	"fmt"
	"path"
	"strconv"
)

var boxPrefix = getenv("BOXPATH", "")

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DIRECTORY CWD: " + dir)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)
	fmt.Println("EXECUTABLE PATH: " + exPath)
	exists, err := exists(boxPrefix + "static")
	if err != nil {
		panic(err)
	}
	fmt.Println("CHECK IF (" +boxPrefix + "static) EXISTS: " + strconv.FormatBool(exists))
	server.RegisterHandlers()
	//http.Handle("/", http.FileServer(rice.MustFindBox(boxPrefix + "static").HTTPBox()))
	http.Handle("/", http.FileServer(http.Dir(boxPrefix + "static")))
	http.ListenAndServe(":8080", nil)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
