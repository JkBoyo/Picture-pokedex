package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(c *config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	mapReq, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	mapRes, err := client.Do(mapReq)
	if err != nil {
		return err
	}

}
