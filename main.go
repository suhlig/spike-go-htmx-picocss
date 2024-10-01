package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// main is the entry point for the program. It sets up and executes the HTTP server.
func main() {
	http.HandleFunc("/", handleRequest)

	fmt.Fprintln(os.Stderr, "Starting up")

	http.ListenAndServe("localhost:8080", nil)
}

var page *template.Template

//go:embed *.html.tmpl
var content embed.FS

func init() {
	page = template.Must(template.New("").ParseFS(content, "*.html.tmpl"))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	counter, _ := strconv.Atoi(r.URL.Query().Get("counter"))

	templateName := r.URL.Query().Get("template")

	if templateName == "" {
		templateName = "default.html.tmpl"
	}

	w.Header().Set("Content-Type", "text/html")

	err := page.ExecuteTemplate(w, templateName, map[string]int{
		"counter": counter,
		"next":    counter + 1,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
