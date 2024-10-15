package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationsURL string
	prevLocationsURL string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Help command should guide users in need.  \nTo use pokedex try one of these commands:",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit command to quit application",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Loading next page of locations from Pokedex API",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Loading previous page of locations from Pokedex API",
			callback:    commandMapb,
		},
	}
}
