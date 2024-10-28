package controllers

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
	"strconv"
)

func FetchArtist(id int) (models.Artist, error) {
	var artist models.Artist
	url := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return artist, fmt.Errorf("error fetching artist: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		return artist, err
	}

	return artist, nil
}
