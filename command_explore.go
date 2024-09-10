package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide an area to explore")
	}
	location := args[0]
	resp, err := cfg.pokeapiClient.GetLocationArea(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}
	return nil
}
