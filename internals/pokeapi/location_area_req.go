package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationsAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	val, ok := c.cache.Get(fullURL)
	if ok {
		log.Print("from cache")
		return unmarshal(val)
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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsAreaResp{}, err
	}
	c.cache.Add(fullURL, data)
	log.Print("from http")
	return unmarshal(data)
}

func unmarshal(data []byte) (LocationsAreaResp, error) {

	var locations LocationsAreaResp
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationsAreaResp{}, err
	}
	return locations, nil
}
