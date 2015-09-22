// Convert markdown to html
package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Convert markdown to html
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

// GenerateMarkdown ...
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {

	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
