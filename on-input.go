package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type LiveRequest struct {
	Text   string `json:"text"`
	Banner string `json:"banner"`
}

// Live typing handler
func asciiLiveHandler(w http.ResponseWriter, r *http.Request) {
	var req LiveRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result, err := generateASCIIArt(req.Text, req.Banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
