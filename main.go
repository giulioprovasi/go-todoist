package main

import (
	"net/http"
	"github.com/kitensei/go-todoist/server"
	"os"
	//"github.com/GeertJohan/go.rice"
)

var boxPrefix = getenv("BOXPATH", "")

func main() {
	server.RegisterHandlers()
	// go.rice causes issues on Heroku, but when editing on Vagrant must use this
	//http.Handle("/", http.FileServer(rice.MustFindBox(boxPrefix + "static").HTTPBox()))
	http.Handle("/", http.FileServer(http.Dir(boxPrefix+"static")))

	// added for Heroku, let him chose the port used, or defaults to 8080
	port := getenv("PORT", "8080")
	http.ListenAndServe(":" + port, nil)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}