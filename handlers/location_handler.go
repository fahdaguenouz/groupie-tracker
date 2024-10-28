package handlers

import (
	"net/http"
	"strconv"
)

// LocationsHandler renders the locations for a specific artist
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/locations/"):] // Extract ID from URL
	id, err := strconv.Atoi(idStr)           // Convert to int
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/artists/"+strconv.Itoa(id), http.StatusFound)
}
