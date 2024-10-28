package handlers

import (
	"net/http"
	"strconv"
)

// DatesHandler renders the concert dates for a specific artist
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/dates/"):] // Extract ID from URL
	id, err := strconv.Atoi(idStr)       // Convert to int
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/artists/"+strconv.Itoa(id), http.StatusFound)
}
