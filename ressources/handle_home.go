package groupie

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound)
		return
	}

	var artists []Artists
	fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
