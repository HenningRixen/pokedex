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
	pokemon             *string
	pokemonInspect      *string
}

func startLoop(config *config) {
	fmt.Print("Pokedex: ")
	scanner := bufio.NewScanner(os.Stdin)
	pokedexmap := map[string]pokeapi.Pokemon{}
	pokedexmapPoint := &pokedexmap
	commandmap := commandsMapCreate(pokedexmapPoint)
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
					fmt.Println("Pokedex: explore needs two inputs (command and location) seperated by withespace")
				} else {
					config.location = &cleanInputCommand[1]
					cmd.callback(config)
				}
			}
		}
		if cleanInputCommand[0] == "catch" {
			if cmd, exits := commandmap["catch"]; exits {
				if len(cleanInputCommand) == 1 {
					fmt.Println("Pokedex: catch needs two inputs (command and pokemon) seperated by withespace")
				} else {
					config.pokemon = &cleanInputCommand[1]
					cmd.callback(config)
				}
			}
		}
		if cleanInputCommand[0] == "pokedex" {
			if cmd, exits := commandmap["pokedex"]; exits {
				cmd.callback(config)
			} 
		}
		if cleanInputCommand[0] == "inspect" {
			if cmd, exits := commandmap["inspect"]; exits {
				if len(cleanInputCommand) == 1 {
					fmt.Println("Pokedex: inspect needs two inputs (command and pokemon) seperated by withespace")
				} else {
					config.pokemonInspect = &cleanInputCommand[1]
					cmd.callback(config)
				}
			}
		}

		if _, exits := commandmap[cleanInputCommand[0]]; !exits {
			fmt.Println("Unkown Command", cleanInputCommand)
		}
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

func commandsMapCreate(pokedexmap *map[string]pokeapi.Pokemon) map[string]cliCommand {
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
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon with a propability and Save it to your bag",
			callback:    func(c *config) error {
				return commandCatch(c, pokedexmap)
			},

		},
		"pokedex": {
			name:        "pokedex",
			description: "Look at the Pokemon in the Pokedex",
			callback:    func(c *config) error {
				return commandPokedex(pokedexmap)
			},
		},
		"inspect": {
			name:"inspect",
			description: "Inspect caught Pokemon",
			callback: func(c *config) error {
				return commandInspect(c, pokedexmap)
			},
		},
		
	}
}
