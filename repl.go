package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kevdbowers/pokedexclient/internal/pokeapi"
)

type configuration struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startREPL(config *configuration) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		var input = reader.Text()
		if len(input) == 0 {
			continue
		}
		var cleanedInput = cleanInput(input)

		command, exists := checkInput()[cleanedInput]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command\n\n")
		}
	}
}

func cleanInput(input string) string {
	var output = strings.ToLower(input)
	return output
}

type command struct {
	name        string
	description string
	callback    func(*configuration) error
}

func checkInput() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "exits the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "displays the names of the 20 previous location areas in the pokemon world",
			callback:    commandMapBack,
		},
	}
}
