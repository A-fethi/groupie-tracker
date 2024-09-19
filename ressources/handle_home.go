package groupie

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var artists []Artist
	fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists)
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "500 Internal server error!", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "500 Internal server error!", http.StatusInternalServerError)
		return
	}
}
