package main

import (
	"github.com/zepetriko/go_pokedex/internal/pokecache"
)

type config struct {
	Next 		string
	Previous 	string
	Cache		*pokecache.Cache
}

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(cfg *config) error
}

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}