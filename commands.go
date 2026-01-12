package main

func registerCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["help"] = cliCommand {
			name: 			"help",
			description: 	"Displays a help message",
			callback: 		commandHelp(commands),
	}

	commands["exit"] = cliCommand {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
	}
	
	commands["map"] = cliCommand {
			name: 			"map",
			description: 	"Displays the next 20 location areas",
			callback: 		commandMap,
	}

	commands["mapb"] = cliCommand {
			name: 			"mapb",
			description: 	"Displays the previous 20 location areas",
			callback: 		commandMapBack,
	}
	return commands
}