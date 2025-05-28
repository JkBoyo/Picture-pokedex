package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jkboyo/pokedex/internal/pokeapi"
	"github.com/jkboyo/pokedex/internal/pokepng"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: client,
		pokedex:       pokedex,
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

func commandTestPng(c *config, args []string) error {
	if len(args) < 1 {
		return errors.New("No args presented")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	picturePath := fmt.Sprintf("%s/Pictures/%s.png", homeDir, args[0])
	png, err := os.ReadFile(picturePath)
	if err != nil {
		return err
	}
	string, err := pokepng.ConvertPNG(png)
	if err != nil {
		return err
	}
	fmt.Printf(string)
	return nil
}
