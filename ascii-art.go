package main

import (
	"html/template"
	"net/http"
)

func asciiHandler(w http.ResponseWriter, r *http.Request) { //implement POST method to handle form submission
	if r.Method != http.MethodPost {
		http.Error(w, "404 Page Not Found", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}
	result, err := generateASCIIArt(text, banner)
	if err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, map[string]string{
		"Result": result,
	})

}
