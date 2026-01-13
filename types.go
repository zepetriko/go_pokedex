package main

import (
	"github.com/zepetriko/go_pokedex/internal/pokecache"
)

type pokemon struct {
	Name			string 	`json:"name"`
	BaseExperience	int 	`json:"base_experience"`
}

type config struct {
	Next 		string
	Previous 	string
	Cache		*pokecache.Cache
	Pokedex		map[string]pokemon
}

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(cfg *config, args []string) error
}

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
