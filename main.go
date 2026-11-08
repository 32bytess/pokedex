package main

import (
	"github.com/32bytess/pokedex/internals/pokeapi"
)

type config struct {
	pokiClient           pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
	commands             []string
}

func main() {

	cfg := config{
		pokiClient:           pokeapi.NewClient(),
		nextLocationArea:     nil,
		previousLocationArea: nil,
		commands:             []string{},
	}
	startRepl(&cfg)
}
