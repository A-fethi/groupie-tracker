package main

import (
	"fmt"
	"net/http"

	groupie "groupie/ressources"
)

func main() {
	http.HandleFunc("/static/", groupie.HandleStatic)
	http.HandleFunc("/", groupie.HandleHome)
	http.HandleFunc("/informations/{id}", groupie.HandleInfos)
	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
