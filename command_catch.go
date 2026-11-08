package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if len(cfg.commands) != 2 {
		fmt.Println("catch command need pokemon name as an input")
		return errors.New("catch command need name as an input")
	}

	name := cfg.commands[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokiClient.GetPokemonDetails(name)
	if err != nil {
		return err
	}
	r := rand.Intn(200)
	if r >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", name)
		cfg.catched[name] = *pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", name)
	return nil
}
