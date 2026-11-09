package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.currentCommand) != 2 {
		fmt.Println("Explore command need Name as an input")
		return errors.New("explore command need name as an input")
	}
	locationName := cfg.currentCommand[1]
	locationArea, err := cfg.pokiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", locationName)
	for _, pokemons := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemons.Pokemon.Name)
	}
	return nil
}
