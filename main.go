package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	fmt.Print("Pokedex: ")
	scanner := bufio.NewScanner(os.Stdin)
	commandmap := commandsMapCreate()
	for scanner.Scan() {
		inputCommand := scanner.Text()
		cleanInputCommand := cleanInput(inputCommand)
		switch cleanInputCommand {
		case "help":
			if cmd, exits := commandmap["help"]; exits {
				cmd.callback()
			}
		case "exit":
			if cmd, exits := commandmap["exit"]; exits {
				cmd.callback()
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
	}
}

func commandHelp() error {
	commandMap := commandsMapCreate()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
