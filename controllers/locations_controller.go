package controllers

import (
	"fmt"
	"groupie/models"
)

// FetchLocations fetches the location data from the API for a specific artist by ID.
func FetchLocations(id int) ([]models.Location, error) {
	var locationsResponse struct {
		Locations []models.Location `json:"locations"`
	}

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	if err := FetchData(url, &locationsResponse); err != nil {
		return nil, err
	}

	return locationsResponse.Locations, nil
}