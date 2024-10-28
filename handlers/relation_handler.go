package handlers

import (
	"net/http"
	"strconv"
)

// RelationsHandler renders the relations for a specific artist
func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/relations/"):] // Extract ID from URL
	id, err := strconv.Atoi(idStr)           // Convert to int
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/artists/"+strconv.Itoa(id), http.StatusFound)
}
