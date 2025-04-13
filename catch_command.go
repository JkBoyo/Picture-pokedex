package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {

	pokemonName := &args[0]
	_, exists := c.pokedex[*pokemonName]
	if exists {
		fmt.Printf("%s was already caught.\n", *pokemonName)
		return nil
	}

	fmt.Printf("Throwing a pokeball at %s...\n", *pokemonName)

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
		c.pokedex[*pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", *pokemonName)
		return nil
	}

	fmt.Printf("%s escaped!\n", *pokemonName)
	return nil
}
