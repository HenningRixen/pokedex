package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
func (c *Client) GetPokemon(pokemon *string) (error, Pokemon) {
	url := "https://pokeapi.co/api/v2/pokemon/" + *pokemon
	req, err := http.NewRequest("GET", url , nil)
	if err != nil {
		return err, Pokemon{}
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err, Pokemon{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)

	}
	if err != nil {
		return err, Pokemon{}
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return err, Pokemon{}
	}
	return nil, pokemonResponse
}

func (c *Client) GetLocationAreaBodyResponse(urlLocationArea *string) MapBodyResponse {
	url := "https://pokeapi.co/api/v2/location-area/"
	if urlLocationArea != nil {
		url = *urlLocationArea
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := MapBodyResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			log.Fatal(err)
		}

		return locationsResp
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	mapBodyResponse := MapBodyResponse{}
	err = json.Unmarshal(body, &mapBodyResponse)
	if err != nil {
		log.Fatal(err)
	}
	c.cache.Add(url, body)
	return mapBodyResponse
}

func (c *Client) GetLoactionPokemonEncounterBodyResponse(location *string) (error, ExploreBodyResponse) {

	url := "https://pokeapi.co/api/v2/location-area/" + *location

	if val, ok := c.cache.Get(url); ok {
		exploreResp := ExploreBodyResponse{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return err, ExploreBodyResponse{}
		}

		return nil, exploreResp
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, ExploreBodyResponse{}
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err, ExploreBodyResponse{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)

	}
	if err != nil {
		return err, ExploreBodyResponse{}
	}

	exploreBodyResponse := ExploreBodyResponse{}
	err = json.Unmarshal(body, &exploreBodyResponse)
	if err != nil {
		return err, ExploreBodyResponse{}
	}
	c.cache.Add(url, body)
	return nil, exploreBodyResponse
}
