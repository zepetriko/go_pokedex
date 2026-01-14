package main 

import (
	"fmt"
	//"strings"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("command 'pokedex' does not takes arguments")
	}

	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("You have not caught any Pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}