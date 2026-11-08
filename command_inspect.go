package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if len(cfg.commands) != 2 {
		fmt.Println("inspect command need pokemon name as an input")
		return errors.New("inspect command need name as an input")
	}
	pokemon, ok := cfg.catched[cfg.commands[1]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	pokemon.Inspect()
	return nil
}
