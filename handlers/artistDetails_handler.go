package handlers

import (
	"groupie/controllers"
	"groupie/models"
	"net/http"
	"strconv"
)

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/artists/"):] // Extract artist ID from URL
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	artist, err := controllers.FetchArtist(id) // Fetch artist details by ID
	if err != nil   {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dates, err := controllers.FetchDates(id) // Fetch concert dates by artist ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := controllers.FetchLocations(id) // Fetch locations by artist ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relations, err := controllers.FetchRelations(id) // Fetch relations by artist ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a struct to hold all the data for the template
	data := struct {
		Artist    models.Artist
		Dates     []models.Date
		Locations []models.Location
		Relations models.Relation
	}{
		Artist:    artist, // Assuming FetchArtist returns a slice
		Dates:     dates,
		Locations: locations,
		Relations: relations,
	}

	renderTemplate(w, "artist.html", data)
}
