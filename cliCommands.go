package main

import (
	"fmt"
	"math/rand"
	"os"

	pokeapi "github.com/HenningRixen/pokedex/internal/pokeApi"
)

func commandHelp(config *config) error {
	commandMap := commandsMapCreate(nil)
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(config *config) error {
	println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	locationAreaStruct := config.pokeApiClient.GetLocationAreaBodyResponse(config.nextLocationUrl)
	config.nextLocationUrl = locationAreaStruct.Next
	config.previousLocationUrl = locationAreaStruct.Previous

	for _, result := range locationAreaStruct.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	locationAreaStruct := config.pokeApiClient.GetLocationAreaBodyResponse(config.previousLocationUrl)
	config.nextLocationUrl = locationAreaStruct.Next
	if config.previousLocationUrl == nil {
		fmt.Print("Now Previous Locations\nPokedex: ")
	} else {
		config.previousLocationUrl = locationAreaStruct.Previous

		for _, result := range locationAreaStruct.Results {
			fmt.Println(result.Name)
		}
	}


	return nil
}

func commandExplore(config *config) error {
	location := config.location
	err, exploreLocationAreaStruct := config.pokeApiClient.GetLoactionPokemonEncounterBodyResponse(location)
	if err != nil {
		fmt.Println("something went wrong in the request")
	}
	for _, result := range exploreLocationAreaStruct.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *config, pokedexmap *map[string]pokeapi.Pokemon) error {
	pokemon := config.pokemon
	println("Throwing Pokeball at" + *pokemon + "...")
	err, pokemonResponse := config.pokeApiClient.GetPokemon(pokemon)
	if err != nil {
		fmt.Println("something wrong with request")
	}
	catchChance := rand.Intn(pokemonResponse.BaseExperience)
	if (catchChance < 80) {
		println("Pokemon caught, adding it to pokedex")
		// pointer here?
		(*pokedexmap)[pokemonResponse.Name] = pokemonResponse
	} else {
		println("didnt catch it bitch try again")
	}

	return nil
}

func commandPokedex(pokedexmap *map[string]pokeapi.Pokemon) error {
	for key := range *pokedexmap {
		println(key)
	}

	return nil
}
