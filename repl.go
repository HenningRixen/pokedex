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
}

func startLoop(config *config) {
	fmt.Print("Pokedex: ")
	scanner := bufio.NewScanner(os.Stdin)
	commandmap := commandsMapCreate()
	for scanner.Scan() {
		inputCommand := scanner.Text()
		cleanInputCommand := cleanInput(inputCommand)
		switch cleanInputCommand {
		case "help":
			if cmd, exits := commandmap["help"]; exits {
				cmd.callback(config)
			}
		case "exit":
			if cmd, exits := commandmap["exit"]; exits {
				cmd.callback(config)
			}
		case "map":
			if cmd, exits := commandmap["map"]; exits {
				cmd.callback(config)
			}
		case "mapb":
			if cmd, exits := commandmap["mapb"]; exits {
				cmd.callback(config)
			}
		default:
			fmt.Println("Unkown Command", cleanInputCommand)
			fmt.Print("Pokedex: ")
		}
	}
}

func cleanInput(inputText string) string {
	return strings.ToLower(inputText)
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
			name:        "map",
			description: "Get Previous 20 Locations in Pokemon",
			callback:    commandMapb,
		},
	}
}
