package main

import (
	"fmt"
	"net/http"

	groupie "groupie/ressources"
)

func main() {
	http.HandleFunc("/", groupie.HandleHome)
	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
