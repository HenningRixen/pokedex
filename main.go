package main

import (
	"time"

	pokeapi "github.com/HenningRixen/pokedex/internal/pokeApi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &config{
		pokeApiClient: pokeClient,
	}
	startLoop(config)
}
