package main

import (
	"fmt"
	"os"
)

func main() {
	startRepl()
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for cmd, reg := range getRegister() {
		commandHelpText := cmd + ": " + reg.description + "\n"
		fmt.Printf("%v", commandHelpText)
	}
	return nil
}
