package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.catched) == 0 {
		fmt.Println("You have not caught any Pokemons yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.catched {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
