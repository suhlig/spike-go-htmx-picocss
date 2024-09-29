package main

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"
)

// main is the entry point for the program. It sets up and executes the HTTP server.
func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":80", nil)
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

	err := page.ExecuteTemplate(w, templateName, map[string]int{
		"counter": counter,
		"next":    counter + 1,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
