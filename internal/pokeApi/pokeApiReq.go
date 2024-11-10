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

func (c *Client) GetLocationAreaBodyResponse(urlLocationArea *string) MapBodyResponse {
	url := "https://pokeapi.co/api/v2/location-area/"
	if urlLocationArea != nil {
		url = *urlLocationArea
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
	return mapBodyResponse
}
