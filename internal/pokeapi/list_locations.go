package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespListLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
		cachedResp, exists := c.cache.Get(url)
		if exists {
			var resp RespListLocations
			err := json.Unmarshal(cachedResp, &resp)
			if err != nil {
				return RespListLocations{}, err
			}

			return resp, nil
		}

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespListLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespListLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespListLocations{}, err
	}

	c.cache.Add(url, dat)

	locationsResp := RespListLocations{}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespListLocations{}, err
	}

	return locationsResp, nil

}
