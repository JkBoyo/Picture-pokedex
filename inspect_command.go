package main

import (
	"fmt"
)

func inspectCommand(c *config, args []string) error {
	pokemon := &args[0]
	pokemonDat, exists := c.pokedex[*pokemon]
	if exists {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n",
			pokemonDat.Name,
			pokemonDat.Height,
			pokemonDat.Weight,
		)
		for _, stat := range pokemonDat.Stats {
			statName := stat.Stat.Name
			baseStat := stat.BaseStat
			fmt.Printf("  -%s: %d\n",
				statName,
				baseStat,
			)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemonDat.Types {
			fmt.Printf("  - %s\n",
				pokemonType.Type.Name,
			)
		}
		fmt.Println(pokemonDat.Sprites.FrontDefault)
		return nil
	}
	fmt.Printf("%s has yet to be caught.", *pokemon)
	return nil
}
