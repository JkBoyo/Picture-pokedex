package main

import (
	"fmt"
	"math/rand"

	"github.com/jkboyo/pokedex/internal/pokepng"
)

func commandCatch(c *config, args []string) error {

	pokemonName := &args[0]
	_, exists := c.pokedex[*pokemonName]
	if exists {
		fmt.Printf("%s was already caught.\n", *pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemonName)

	pokemon, err := c.pokeapiClient.PokemonDetails(pokemonName)
	if err != nil {
		fmt.Printf("%s was gone before the pokeball could hit him\n", *pokemonName)
		return err
	}

	baseExp := pokemon.BaseExperience

	denom := baseExp / 20

	numer := rand.Intn(denom + 1)

	result := (float64(numer) / float64(denom))

	if result > 0.5 {
		spriteUrl := pokemon.Sprites.FrontDefault
		pokePngDat, err := c.pokeapiClient.PokemonPicture(spriteUrl)
		if err != nil {
			return err
		}

		pokeAscii, err := pokepng.ConvertPNG(pokePngDat)
		if err != nil {
			return err
		}

		pokemon.Sprites.FrontDefault = pokeAscii

		c.pokedex[*pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", *pokemonName)
		return nil
	}

	fmt.Printf("%s escaped!\n", *pokemonName)
	return nil
}
