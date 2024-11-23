package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MapBodyResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ExploreBodyResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func (c *Client) GetLoactionPokemonEncounterBodyResponse() ExploreBodyResponse {
	url := "https://pokeapi.co/api/v2/location-area/canalave-city-area"

	if val, ok := c.cache.Get(url); ok {
		exploreResp := ExploreBodyResponse{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			log.Fatal(err)
		}

		return exploreResp
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

	exploreBodyResponse := ExploreBodyResponse{}
	err = json.Unmarshal(body, &exploreBodyResponse)
	if err != nil {
		log.Fatal(err)
	}
	c.cache.Add(url, body)
	return exploreBodyResponse
}
