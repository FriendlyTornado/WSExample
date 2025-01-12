package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PageInfo struct {
	Name  string
	Email string
	Title string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	info := PageInfo{
		Name:  "Jen",
		Title: "Main Page",
	}
	tplPath := filepath.Join("templates", "home.gohtml")
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "Error parsing page", 500)
	}

	err = t.Execute(w, info)
	if err != nil {
		http.Error(w, "Error getting page", 500)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	info := PageInfo{
		Name:  "Jen",
		Email: "ctrl.vee@gmail.com",
		Title: "Contact",
	}

	tplPath := filepath.Join("templates", "contact.gohtml")
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "Error parsing page", 500)
	}

	err = t.Execute(w, info)
	if err != nil {
		http.Error(w, "Error getting page", 500)
	}

}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		mainHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", 400)
	}
}

func main() {

	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}
