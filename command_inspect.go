package main 

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	pokemonName := strings.ToLower(args[0])
	pokemon, ok := cfg.Pokedex[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}