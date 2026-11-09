package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (*ListLocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	return c.getRequest[ListLocationAreas](fullURL)
}

func (c *Client) GetLocationArea(locationName string) (*LocationArea, error) {
	endpoint := "/location-area" + "/" + locationName
	fullURL := baseURL + endpoint
	return c.getRequest[LocationArea](fullURL)
}

func (c *Client) GetPokemon(pokemonName string) (*Pokemon, error) {
	endpoint := "/pokemon" + "/" + pokemonName
	fullURL := baseURL + endpoint
	return c.getRequest[Pokemon](fullURL)
}

func (c *Client) getRequest[T any](fullURL string) (*T, error) {

	val, ok := c.cache.Get(fullURL)
	if ok {
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
	return unmarshal[T](data)
}

func unmarshal[T any](data []byte) (*T, error) {

	var locations T
	if err := json.Unmarshal(data, &locations); err != nil {
		return &locations, err
	}
	return &locations, nil
}
