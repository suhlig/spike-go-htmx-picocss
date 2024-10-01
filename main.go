package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := mainE()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
}

func mainE() error {
	s, err := NewServer()

	if err != nil {
		return err
	}

	http.HandleFunc("/", s.handleRoot)
	http.HandleFunc("/foo", s.handleFoo)
	http.HandleFunc("/bar", s.handleBar)

	fmt.Fprintln(os.Stderr, "Starting up")
	return http.ListenAndServe("localhost:8080", nil)
}
