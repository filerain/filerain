package main

import (
	"embed"
	"fmt"
	"github.com/a-h/templ"
	"io/fs"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	var staticFS = fs.FS(staticFiles)
	fs := http.FileServer(http.FS(staticFS))

	// Serve static files
	http.Handle("/", fs)

	http.HandleFunc("GET /auth/signup", signUpHandler)
	http.HandleFunc("POST /auth/signup", signUpHandlerPost)
	http.Handle("/hello", templ.Handler(hello("Title")))

	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":80", nil)
	print("Serving at port 80")
}
