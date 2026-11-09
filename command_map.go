package main

import (
	"fmt"
)

func commandMap(cfg *config, next bool) error {
	var fullURL *string
	if next {
		fullURL = cfg.nextListLocationArea
	} else {
		if cfg.previousListLocationArea == nil {
			fmt.Println("you're on the first page")
			return nil
		}
		fullURL = cfg.previousListLocationArea
	}

	listLocationAreas, err := cfg.pokiClient.ListLocationAreas(fullURL)
	if err != nil {
		return err
	}

	for _, location := range listLocationAreas.Results {
		fmt.Println(location.Name)
	}
	cfg.nextListLocationArea = listLocationAreas.Next

	cfg.previousListLocationArea = listLocationAreas.Previous

	return nil
}
