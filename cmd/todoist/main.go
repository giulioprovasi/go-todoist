package todoist

import (
	"net/http"
	"todo/server"
	"github.com/GeertJohan/go.rice"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}