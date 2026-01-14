package main 

import (
	"fmt"
	"encoding/json"
	"strings"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemonName := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	data, err := getFromCacheOrFetch(cfg, url)

	if err != nil {
		return err
	}

	var pokemon pokemon
	json.Unmarshal(data, &pokemon)

	catched := throwBall(pokemon.BaseExperience)

	if !catched {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	
	if _, ok := cfg.Pokedex[pokemonName]; ok {
		fmt.Printf("%s is already in your Pokedex\n", pokemon.Name)
		return nil
	}

	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the 'inspect' command.")
	return nil
}

func throwBall(base_experience int) bool {
	chance := max(20, 100 - base_experience)
	roll := rand.Intn(100)
	if roll < chance {
		return true
	}

	return false
}