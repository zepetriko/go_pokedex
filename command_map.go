package main 

import "fmt"

func commandMap(cfg *config, args []string) error {
	if cfg.Next == "" {
		fmt.Println("No more locations.")
		return nil
	}

	return fetchAndPrintLocations(cfg, cfg.Next)
}

func commandMapBack(cfg *config, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	return fetchAndPrintLocations(cfg, cfg.Previous)
}