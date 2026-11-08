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
	callback    func() error
}

func commands() map[string]cliCommand {
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if s := scanner.Scan(); s {
			if err := scanner.Err(); err != nil {
				fmt.Printf("error scanning")
				return
			}
			text := scanner.Text()
			words := cleanInput(text)
			if len(words) == 0 {
				continue
			}
			if c, ok := commands()[words[0]]; ok {
				c.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
