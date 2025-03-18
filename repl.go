package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		command := cleanedInput[0]
		conf := config{}
		if cmd, ok := getRegister()[command]; ok {
			err := cmd.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

type config struct {
	Previous string
	Next     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
