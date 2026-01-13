package main

import "fmt"

func commandHelp(commands map[string]cliCommand) func(cfg *config, args []string) error {
	return func(cfg *config, args []string) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()

		for _, command := range commands {
			fmt.Printf(" %s: %s\n", command.name, command.description)
		}

		return nil
	}
}