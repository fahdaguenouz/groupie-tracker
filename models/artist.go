package models

import (
	"encoding/json"
	"fmt"
)

type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    string              `json:"locations"`
	ConcertDates ConcertDatesField   `json:"concertDates"` // Ensure this is a slice of Date structs
	Relations    Relation `json:"relations"`    // Assuming relations is a map
}
type ConcertDatesField struct {
	Dates []string
}
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
func (c *ConcertDatesField) UnmarshalJSON(data []byte) error {
	// Attempt to unmarshal as a single date string
	var singleDate string
	if err := json.Unmarshal(data, &singleDate); err == nil {
		c.Dates = []string{singleDate}
		return nil
	}

	// Attempt to unmarshal as an array of strings
	var dateArray []string
	if err := json.Unmarshal(data, &dateArray); err == nil {
		c.Dates = dateArray
		return nil
	}

	// Attempt to unmarshal as an object with a "dates" key containing an array
	var dateObject struct {
		Dates []string `json:"dates"`
	}
	if err := json.Unmarshal(data, &dateObject); err == nil {
		c.Dates = dateObject.Dates
		return nil
	}

	// Return an error if none of the above worked
	return fmt.Errorf("concertDates is neither a string, an array of strings, nor an object with a dates key")
}
// Location represents a location with associated dates.
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Date represents a date associated with an event.
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// // Relation represents relationships between dates and locations.
// type Relation struct {
// 	ID             int                 `json:"id"`
// 	DatesLocations map[string][]string `json:"datesLocations"`
// }

func (r *Relation) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal into a map[string][]string
	var datesLocations map[string][]string
	if err := json.Unmarshal(data, &datesLocations); err == nil {
		r.DatesLocations = datesLocations
		return nil
	}

	// If that fails, try unmarshaling as a string (if sometimes relations is just a placeholder string)
	var singleString string
	if err := json.Unmarshal(data, &singleString); err == nil {
		r.DatesLocations = map[string][]string{"placeholder": {singleString}}
		return nil
	}

	// Return an error if neither method works
	return fmt.Errorf("relations is neither a map nor a string")
}