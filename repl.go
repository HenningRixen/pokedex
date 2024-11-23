package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/HenningRixen/pokedex/internal/pokeApi"
)

type config struct {
	pokeApiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
	location            *string
}

func startLoop(config *config) {
	fmt.Print("Pokedex: ")
	scanner := bufio.NewScanner(os.Stdin)
	commandmap := commandsMapCreate()
	for scanner.Scan() {
		inputCommand := scanner.Text()
		cleanInputCommand := cleanInput(inputCommand)
		if cleanInputCommand[0] == "help" {
			if cmd, exits := commandmap["help"]; exits {
				cmd.callback(config)
			}
		}
		if cleanInputCommand[0] == "exit" {
			if cmd, exits := commandmap["exit"]; exits {
				cmd.callback(config)
			}
		}
		if cleanInputCommand[0] == "map" {
			if cmd, exits := commandmap["map"]; exits {
				cmd.callback(config)
			}
		}
		if cleanInputCommand[0] == "mapb" {
			if cmd, exits := commandmap["mapb"]; exits {
				cmd.callback(config)
			}
		}
		if cleanInputCommand[0] == "explore" {
			if cmd, exits := commandmap["explore"]; exits {
				if len(cleanInputCommand) == 1 {
					fmt.Print("Pokedex: explore needs two inputs seperated by withespace")
					fmt.Print("Pokedex: ")
				} else {
					config.location = &cleanInputCommand[1]
					cmd.callback(config)
				}
			}
		}
		fmt.Println("Unkown Command", cleanInputCommand)
		fmt.Print("Pokedex: ")
	}
}

func cleanInput(inputText string) []string {
	lower := strings.ToLower(inputText)
	words := strings.Split(lower, " ")
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandsMapCreate() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get 20 Locations in Pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get Previous 20 Locations in Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get Pokemon you can encounter in this Location",
			callback:    commandExplore,
		},
	}
}
