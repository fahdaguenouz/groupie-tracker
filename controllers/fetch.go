package controllers


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FetchData is a generic function to fetch data from a given URL.
func FetchData(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}
