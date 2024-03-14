package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationsResponse := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return ResponseLocations{}, err
	}

	return locationsResponse, nil
}
