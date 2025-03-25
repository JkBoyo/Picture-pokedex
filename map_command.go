package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args []string) error {
	locationsListResp, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationsListResp.Next
	c.Previous = locationsListResp.Previous

	for _, loc := range locationsListResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config, args []string) error {
	if c.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationsListResp, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationsListResp.Next
	c.Previous = locationsListResp.Previous

	for _, loc := range locationsListResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
