package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) < 1 {
		return errors.New("No args passed")
	}
	for _, area := range args {
		locationDetailResp, err := c.pokeapiClient.LocationDetails(&area)
		if err != nil {
			return err
		}

		fmt.Println("Exploring " + area + "...")
		fmt.Println("Found Pokemon:")
		for _, mon := range locationDetailResp.PokemonEncounters {
			fmt.Println("	-" + mon.Pokemon.Name)
		}
	}
	return nil
}
