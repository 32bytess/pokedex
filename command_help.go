package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, val := range getCommands() {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
