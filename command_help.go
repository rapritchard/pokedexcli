package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome too the Pokedex!")
	fmt.Print("Usage:\n\n")

	availableCommands := getCLICommands()

	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}
