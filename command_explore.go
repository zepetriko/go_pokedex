package main 

import (
	"fmt"
	"encoding/json"
	"strings"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-area>")
	}

	area := strings.ToLower(args[0])
	url := "https://pokeapi.co/api/v2/location-area/" + area

	data, err := getFromCacheOrFetch(cfg, url)

	if err != nil {
		return err
	}

	var location locationAreaResponse
	json.Unmarshal(data, &location)

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, encounter := range location.PokemonEncounter {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}