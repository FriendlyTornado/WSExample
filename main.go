package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<B>Example page</B><br><p><a href=\"contact\">Contact</a>")
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<B>Contact</B><p>You can contact me at "+
		"<a href=\"mailto:ctrl.vee@gmail.com\">ctrl.vee@gmail.com</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlerFunc(w, r)
	case "/contact":
		contactFunc(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}
