package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if len(cfg.currentCommand) != 2 {
		fmt.Println("catch command need pokemon name as an input")
		return errors.New("catch command need name as an input")
	}

	pokemonName := cfg.currentCommand[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	r := rand.Intn(200)
	if r >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Printf("yuu may now inspect it with the inspect command\n")
		cfg.catchedPokemons[pokemonName] = *pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
