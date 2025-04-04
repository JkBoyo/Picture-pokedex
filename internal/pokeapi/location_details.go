package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(areaName *string) (RespDetailLocations, error) {
	url := baseUrl + "/location-area/" + *areaName
	cachedResp, exists := c.cache.Get(url)
	if exists {
		var resp RespDetailLocations
		err := json.Unmarshal(cachedResp, &resp)
		if err != nil {
			return RespDetailLocations{}, err
		}
		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailLocations{}, err
	}

	c.cache.Add(url, dat)

	locationDetailsResp := RespDetailLocations{}

	err = json.Unmarshal(dat, &locationDetailsResp)
	if err != nil {
		return RespDetailLocations{}, err
	}

	return locationDetailsResp, nil
}
