package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jkboyo/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for cmd, reg := range getRegister() {
		commandHelpText := cmd + ": " + reg.description + "\n"
		fmt.Printf("%v", commandHelpText)
	}
	return nil
}
