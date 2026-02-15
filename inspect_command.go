package main

import (
	"errors"
	"fmt"
)

var typeColors = map[string]struct {
	r int
	g int
	b int
}{
	"normal":   {r: 168, g: 167, b: 122},
	"fire":     {r: 238, g: 129, b: 48},
	"water":    {r: 99, g: 144, b: 240},
	"electric": {r: 247, g: 208, b: 44},
	"grass":    {r: 122, g: 199, b: 76},
	"ice":      {r: 150, g: 217, b: 214},
	"fighting": {r: 194, g: 46, b: 40},
	"poison":   {r: 163, g: 62, b: 161},
	"ground":   {r: 226, g: 191, b: 101},
	"flying":   {r: 169, g: 143, b: 243},
	"psychic":  {r: 249, g: 85, b: 135},
	"bug":      {r: 166, g: 185, b: 26},
	"rock":     {r: 182, g: 161, b: 54},
	"ghost":    {r: 115, g: 87, b: 151},
	"dragon":   {r: 111, g: 53, b: 252},
	"dark":     {r: 112, g: 87, b: 70},
	"steel":    {r: 183, g: 183, b: 206},
	"fairy":    {r: 214, g: 133, b: 173},
}

func inspectCommand(c *config, args []string) error {
	if len(args) < 1 {
		return errors.New("No Pokemon provided")
	}
	pokemon := &args[0]
	pokemonDat, exists := c.pokedex[*pokemon]
	if exists {
		var pokeInfo string
		pokeInfo += fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n",
			pokemonDat.Name,
			pokemonDat.Height,
			pokemonDat.Weight,
		)
		for _, stat := range pokemonDat.Stats {
			statName := stat.Stat.Name
			baseStat := stat.BaseStat
			pokeInfo += fmt.Sprintf("  -%s: %d\n",
				statName,
				baseStat,
			)
		}
		pokeInfo += "Types:\n"
		for _, pokemonType := range pokemonDat.Types {
			color := typeColors[pokemonType.Type.Name]
			colorStr := fmt.Sprintf("\033[38;2;%v;%v;%vm",
				color.r,
				color.g,
				color.b,
			)
			pokeInfo += fmt.Sprintf("  - %s%s\033[0m\n",
				colorStr,
				pokemonType.Type.Name,
			)
		}
		PrintPokePage(pokeInfo, pokemonDat.Sprites.FrontDefault)
		return nil
	}
	fmt.Printf("%s has yet to be caught.\n", *pokemon)
	return nil
}
