package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (*LocationsAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	return c.getRequest[LocationsAreaResp](fullURL)
}

func (c *Client) GetPokemonsInArea(name string) (*PokemonsResp, error) {
	endpoint := "/location-area" + "/" + name
	fullURL := baseURL + endpoint
	return c.getRequest[PokemonsResp](fullURL)
}

func (c *Client) getRequest[T any](fullURL string) (*T, error) {

	val, ok := c.cache.Get(fullURL)
	if ok {
		log.Print("from cache")
		return unmarshal[T](val)
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("bad status code : %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(fullURL, data)
	log.Print("from http")
	return unmarshal[T](data)
}

func unmarshal[T any](data []byte) (*T, error) {

	var locations T
	if err := json.Unmarshal(data, &locations); err != nil {
		return &locations, err
	}
	return &locations, nil
}
