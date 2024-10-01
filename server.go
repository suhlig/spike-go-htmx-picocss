package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type server struct {
	templates map[string]*template.Template
}

//go:embed *.html.tmpl
var templateFS embed.FS

func NewServer() (*server, error) {
	master, err := template.New("master.html.tmpl").ParseFS(templateFS, "master.html.tmpl")

	if err != nil {
		return nil, err
	}

	templateFiles, err := templateFS.ReadDir(".")

	if err != nil {
		return nil, err
	}

	s := server{
		templates: make(map[string]*template.Template),
	}

	for _, t := range templateFiles {
		clone, err := master.Clone()

		if err != nil {
			return nil, err
		}

		p, err := clone.ParseFS(templateFS, t.Name())

		if err != nil {
			return nil, err
		}

		s.templates[t.Name()] = p
	}

	return &s, nil
}

func (s *server) handleCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var (
		err   error
		count int
	)

	count, _ = strconv.Atoi(r.URL.Query().Get("count"))
	data := map[string]int{
		"count": count,
		"next":  count + 1,
	}

	blockName := r.URL.Query().Get("block")

	if blockName == "" {
		err = s.render(w, "counter.html.tmpl", data)
	} else {
		err = s.renderBlock(w, "counter.html.tmpl", blockName, data)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := s.render(w, "foo.html.tmpl", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleBar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := s.render(w, "bar.html.tmpl", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) render(w http.ResponseWriter, templateName string, data any) error {
	template, found := s.templates[templateName]

	if !found {
		return fmt.Errorf("undefined template %s", templateName)
	}

	return template.Execute(w, data)
}

func (s *server) renderBlock(w http.ResponseWriter, templateName, block string, data any) error {
	template, found := s.templates[templateName]

	if !found {
		return fmt.Errorf("undefined template %s", templateName)
	}

	return template.ExecuteTemplate(w, block, data)
}
