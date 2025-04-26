package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

func (c *Client) PokemonPicture(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	pokePng, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, errors.New("PNG not fetched")
	}
	defer pokePng.Body.Close()

	pokePngDat, err := io.ReadAll(pokePng.Body)
	if err != nil {
		return []byte{}, err
	}

	return pokePngDat, nil

}
