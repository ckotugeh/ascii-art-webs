package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static")) // maps Get request Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/ascii-live", asciiLiveHandler)
	fmt.Println("server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}
