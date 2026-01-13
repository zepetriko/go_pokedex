package main

import (
	"github.com/zepetriko/go_pokedex/internal/pokecache"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Cache: pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]pokemon),
	}

	commands := registerCommands()
	
	for {
		fmt.Print("Pokedex > ")
		
		if !scanner.Scan(){
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		words := strings.Fields(input)
		commandName := strings.ToLower(words[0])
		args := words[1:]

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg, args); err != nil {
			fmt.Println(err)
		}

	}

}