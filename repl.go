package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations in Pokemon",
			callback: func(cfg *config) error {
				return commandMap(cfg, true)
			},
		},
		"mapb": {
			name:        "map",
			description: "Displays previous 20 locations in Pokemon",
			callback: func(cfg *config) error {
				return commandMap(cfg, false)
			},
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of Pokemons in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch Pokemons",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of Pokemons you have caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if s := scanner.Scan(); s {
			if err := scanner.Err(); err != nil {
				fmt.Printf("error scanning")
				return
			}
			text := scanner.Text()
			cfg.commands = cleanInput(text)
			if len(cfg.commands) == 0 {
				continue
			}
			if c, ok := getCommands()[cfg.commands[0]]; ok {
				c.callback(cfg)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
