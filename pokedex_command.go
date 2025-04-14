package main

import (
	"fmt"
)

func pokedexCommand(c *config, args []string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("No pokemon caught. Try using the catch command to catch some :)")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
