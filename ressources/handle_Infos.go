package groupie

import (
	"html/template"
	"net/http"
)

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	var artist Artists
	var locations Locations
	var dates Dates
	var relations Relations

	fetch("https://groupietrackers.herokuapp.com/api/", "artists/"+id, &artist)
	tmpl, err := template.ParseFiles("templates/infos.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	fetch("https://groupietrackers.herokuapp.com/api/", "locations/"+id, &locations)
	_, err = template.ParseFiles("templates/infos.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	fetch("https://groupietrackers.herokuapp.com/api/", "dates/"+id, &dates)
	_, err = template.ParseFiles("templates/infos.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	fetch("https://groupietrackers.herokuapp.com/api/", "relation/"+id, &relations)
	_, err = template.ParseFiles("templates/infos.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"artist":    artist,
		"locations": locations.Locations,
		"dates":     dates.Dates,
		"DatesLocations": relations.DatesLocations,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
