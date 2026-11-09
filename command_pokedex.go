package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.catchedPokemons) == 0 {
		fmt.Println("You have not caught any Pokemons yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.catchedPokemons {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
