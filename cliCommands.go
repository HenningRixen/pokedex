package main

import (
	"fmt"
	"os"
)

func commandHelp(config *config) error {
	commandMap := commandsMapCreate()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Print("Pokedex: ")
	return nil
}

func commandExit(config *config) error {
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
	fmt.Print("Pokedex: ")

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
		fmt.Print("Pokedex: ")
	}

	return nil
}
