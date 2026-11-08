package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.commands) != 2 {
		fmt.Println("Explore command need Name as an input")
		return errors.New("explore command need name as an input")
	}
	name := cfg.commands[1]
	pokemons, err := cfg.pokiClient.GetPokemonsInArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", name)
	for _, pokemons := range pokemons.PokemonEncounters {
		fmt.Printf("- %s\n", pokemons.Pokemon.Name)
	}
	return nil
}
