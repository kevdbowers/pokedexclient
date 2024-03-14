package main

import (
	"time"

	"github.com/kevdbowers/pokedexclient/internal/pokeapi"
)

func main() {
	pokedexClient := pokeapi.NewClient(5 * time.Second)
	config := &configuration{
		pokeapiClient: pokedexClient,
	}
	startREPL(config)
}
