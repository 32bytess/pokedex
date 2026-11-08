package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationsAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationsAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationsAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationsAreaResp{}, fmt.Errorf("bad status code : %v", res.StatusCode)
	}

	var locations LocationsAreaResp

	if err := (json.NewDecoder(res.Body)).Decode(&locations); err != nil {
		return LocationsAreaResp{}, err
	}
	return locations, nil
}
