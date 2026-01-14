package main

import (
	"github.com/zepetriko/go_pokedex/internal/pokecache"
)

type pokemon struct {
	Name			string 		`json:"name"`
	Height 			int			`json:"height"`
	Weight 			int			`json:"weight"`

	Stats 			[]struct {
		BaseStat 	int 		`json:"base_stat"`
		Stat		struct {
			Name 	string 		`json:"name"`
		} 						`json:"stat"`
	} 							`json:"stats"`

	Types 			[]struct {
		Type 		struct {
			Name	string 		`json:"name"`
		} 						`json:"type"`
	}							`json:"types"`

	BaseExperience	int 		`json:"base_experience"`
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
