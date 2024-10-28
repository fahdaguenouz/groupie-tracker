package controllers

import (
	"fmt"
	"groupie/models"
)

// FetchDates fetches the concert dates from the API for a specific artist by ID.
func FetchDates(id int) ([]models.Date, error) {
	
		var dates []models.Date 
	

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	if err := FetchData(url, &dates); err != nil {
		return nil, err
	}

	return dates, nil
}