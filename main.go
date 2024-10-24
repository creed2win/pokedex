package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/creed2win/pokedex/internal/pokecache"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	cache := pokecache.NewCache(time.Minute * 5)
	cfg := &config{
		client:           http.Client{},
		cache:            cache,
		nextLocationsURL: "https://pokeapi.co/api/v2/location-area/",
		prevLocationsURL: "https://pokeapi.co/api/v2/location-area/",
		locationAreaUrl:  "https://pokeapi.co/api/v2/location-area/",
		pokedex:          Pokedex{},
	}

	for {
		fmt.Fprint(os.Stderr, "pokedex > ")
		scanner.Scan()
		//user input - command + parametr
		input := scanner.Text()
		inputsMap := strings.Fields(input)

		command := inputsMap[0]
		paramsSlice := paramSlice{
			params: inputsMap[1:],
		}
		if commands[command].name == command {
			commands[command].callback(cfg, paramsSlice)
		}
	}
}
