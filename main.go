package main

import (
	"time"

	"github.com/rapritchard/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	pokedex             map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
