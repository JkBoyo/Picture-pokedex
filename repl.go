package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jkboyo/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Previous      *string
	Next          *string
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]

		if cmd, ok := getRegister()[command]; ok {
			err := cmd.callback(c, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getRegister() map[string]cliCommand {
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
			description: "Display list of 20 locations at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous list of 20 locations at a time",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list pokemon found in areas passed in with spaces",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch the pokemon passed as your argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "list out details of previously caught pokemon",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show all caught pokemon in a list",
			callback:    pokedexCommand,
		},
	}
}
