package main

import (
	"github.com/32bytess/pokedex/internals/pokeapi"
)

type config struct {
	pokiClient               pokeapi.Client
	nextListLocationArea     *string
	previousListLocationArea *string
	catchedPokemons          map[string]pokeapi.Pokemon
	currentCommand           []string
}

func main() {
	cfg := config{
		pokiClient:               pokeapi.NewClient(),
		nextListLocationArea:     nil,
		previousListLocationArea: nil,
		catchedPokemons:          map[string]pokeapi.Pokemon{},
		currentCommand:           []string{},
	}
	startRepl(&cfg)
}
