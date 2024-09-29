package main

import (
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

// init function sets up the template+fragment.  Most of the work is actually done here.
// In a larger program, this would likely be stored in a separate file, but this makes for a
// simple example.
func init() {

	page = template.New("main")

	page = template.Must(page.Parse(`<!DOCTYPE html>

	<html>
	<head>
		<script src="https://unpkg.com/htmx.org@1.8.0"></script>
		<link rel="stylesheet" href="https://unpkg.com/missing.css@1.1.1"/>
		<title>Template Fragment Example</title>
	</head>
	<body>
		<h1>Template Fragment Example</h1>

		<p>This page demonstrates how to create and serve
		<a href="https://htmx.org/essays/template-fragments/">template fragments</a>
		using the <a href="https://pkg.go.dev/text/template">built-in template package</a> in Go.</p>

		<p>This is accomplished by using the "block" action in the template, which lets you
		define and execute a sub-template in a single step.</p>

		<!-- Here's the fragment.  We can target it by executing the "buttonOnly" template. -->
		{{block "buttonOnly" .}}
			<button hx-get="/?counter={{.next}}&template=buttonOnly" hx-swap="outerHTML">
				This Button Has Been Clicked {{.counter}} Times
			</button>
		{{end}}

	</body>
	</html>`))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	counter, _ := strconv.Atoi(r.URL.Query().Get("counter"))

	templateName := r.URL.Query().Get("template")

	if templateName == "" {
		templateName = "main"
	}

	err := page.ExecuteTemplate(w, templateName, map[string]int{
		"counter": counter,
		"next":    counter + 1,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
