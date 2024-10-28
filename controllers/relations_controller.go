package controllers

import (
	"fmt"
	"groupie/models"
)

// FetchRelations fetches the relations data from the API for a specific artist by ID.
func FetchRelations(id int) (models.Relation, error) {
	var relationsResponse models.Relation

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	if err := FetchData(url, &relationsResponse); err != nil {
		return models.Relation{}, err
	}

	return relationsResponse, nil
}