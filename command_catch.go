package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide a pokemon to catch")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	rand := rand.IntN(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)

	if rand > threshold {
		return fmt.Errorf("%s escaped!\n", pokemon.Name)
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Println("You may now inspect it with the inspect command")
	return nil
}
