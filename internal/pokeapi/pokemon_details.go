package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(pokemonName *string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + *pokemonName
	cachedResp, exists := c.cache.Get(url)
	if exists {
		var resp Pokemon
		err := json.Unmarshal(cachedResp, &resp)
		if err != nil {
			return Pokemon{}, err
		}
		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	pokemon := Pokemon{}

	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
