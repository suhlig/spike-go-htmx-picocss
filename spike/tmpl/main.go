package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	err := mainE()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
}

const (
	master = `
# {{ block "title" . }}default title{{ end }}

Here are all the names:
{{ block "list" . }}
	{{ range . }}
		{{ println "-" . }}
	{{ end }}
{{ end }}
That's it!
`

	overlay1Content = `
{{ define "title" }}This is Overlay One {{ end }}
{{ define "list" }} one, two, three {{ end }}
	`
	overlay2Content = `
{{ define "title" }}O2{{ end }}
{{ define "list" }} four, five, six {{ end }}
`
)

func mainE() error {
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)

	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)

	if err != nil {
		return err
	}

	overlay1Tmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay1Content)

	if err != nil {
		return err
	}

	if err := overlay1Tmpl.Execute(os.Stdout, guardians); err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "-------------------")

	overlay2Tmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay2Content)

	if err != nil {
		return err
	}

	if err := overlay2Tmpl.Execute(os.Stdout, guardians); err != nil {
		return err
	}

	return nil
}
