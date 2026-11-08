package main

import (
	"github.com/32bytess/pokedex/internals/pokeapi"
)

type config struct {
	pokiClient           pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
	catched              map[string]pokeapi.PokemonResp
	commands             []string
}

func main() {

	cfg := config{
		pokiClient:           pokeapi.NewClient(),
		nextLocationArea:     nil,
		previousLocationArea: nil,
		catched:              map[string]pokeapi.PokemonResp{},
		commands:             []string{},
	}
	startRepl(&cfg)
}
