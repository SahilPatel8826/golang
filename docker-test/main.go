package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hi there %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8080", nil)
}
