package main

import (
	"net/http"

	"github.com/creed2win/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, paramSlice) error
}

type config struct {
	cache            pokecache.Cache
	client           http.Client
	pokedex          Pokedex
	nextLocationsURL string
	prevLocationsURL string
	locationAreaUrl  string
}

type paramSlice struct {
	params []string
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
		"explore": {
			name:        "explore",
			description: "Exploring selected location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Command for catching pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspecting pokemon in you pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Listing all pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
