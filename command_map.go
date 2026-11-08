package main

import (
	"fmt"
)

func commandMap(cfg *config, next bool) error {
	var fullURL *string
	if next {
		fullURL = cfg.nextLocationArea
	} else {
		if cfg.previousLocationArea == nil {
			fmt.Println("you're on the first page")
			return nil
		}
		fullURL = cfg.previousLocationArea
	}

	locations, err := cfg.pokiClient.ListLocationAreas(fullURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationArea = locations.Next

	cfg.previousLocationArea = locations.Previous

	return nil
}
