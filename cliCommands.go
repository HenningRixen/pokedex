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

func commandCatch(config *config, pokedexmap *map[string]pokeapi.PokemonExtended) error {
	pokemon := config.pokemon
	println("Throwing Pokeball at" + *pokemon + "...")
	err, pokemonResponse := config.pokeApiClient.GetPokemon(pokemon)
	if err != nil {
		fmt.Println("something wrong with request")
	}
	catchChance := rand.Intn(pokemonResponse.BaseExperience)
	if catchChance < 80 {
		println("Pokemon caught, adding it to pokedex")
		(*pokedexmap)[pokemonResponse.Name] = pokeapi.PokemonExtended{
			Pokemon:            pokemonResponse,
			PokemonLearntMoves: nil,
		}
	} else {
		println("didnt catch it bitch try again")
	}

	return nil
}

func commandPokedex(pokedexmap *map[string]pokeapi.PokemonExtended) error {
	for key, value := range *pokedexmap {
		println(key)
		println(value.Pokemon.BaseExperience)
	}

	return nil
}

func commandInspect(config *config, pokedexmap *map[string]pokeapi.PokemonExtended) error {
	pokemon := config.pokemonInspect
	pokemonInMap, exists := (*pokedexmap)[*pokemon]
	if exists {
		println("Name:" + pokemonInMap.Pokemon.Name)
		println(fmt.Sprintf("Height: %d", pokemonInMap.Pokemon.Height))
		println(fmt.Sprintf("Weight: %d", pokemonInMap.Pokemon.Weight))
		for _, stat := range pokemonInMap.Pokemon.Stats {
			println(fmt.Sprintf("-%s: %d", stat.Stat.Name, stat.BaseStat))
		}
		println("Types:")
		for _, typ := range pokemonInMap.Pokemon.Types {
			println("-" + typ.Type.Name)
		}
		println("Learned Moves:")
		for _, move := range pokemonInMap.PokemonLearntMoves {
			fmt.Println("Pokemon learnt moves:", move.Name)
		}
	} else {
		println("pokemon not caught yet")
	}

	return nil
}

func commandMoves(config *config, pokedexmap *map[string]pokeapi.PokemonExtended) error {
	pokemon := config.pokemonMoves
	values, exists := (*pokedexmap)[*pokemon]
	if exists {
		println("Moves:")
		for _, value := range values.Pokemon.Moves {
			println(value.Move.Name)
		}
	} else {
		println("pokemon not caught yet")
	}

	return nil
}

func commandLearnMove(pokemon *string, move *string, pokedexmap *map[string]pokeapi.PokemonExtended, config *config) error {
	existingPokemon, exists := (*pokedexmap)[*pokemon]

	if exists {
		err, move := config.pokeApiClient.GetMove(move)

		if err != nil {
			fmt.Println("something went wrong in the request")
		}

		existingPokemon.PokemonLearntMoves = append(existingPokemon.PokemonLearntMoves, move)
		(*pokedexmap)[*pokemon] = existingPokemon
		for _, move := range existingPokemon.PokemonLearntMoves {
			fmt.Println("Pokemon learnt moves:", move.Name)
		}
	}
	return nil
}
