package main

import (
	"bufio"
	"fmt"
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
		cache:            cache,
		nextLocationsURL: "https://pokeapi.co/api/v2/location-area/",
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
		} //TODO - problem if statment is not evaluating correctly
		if commands[command].name == command {
			commands[command].callback(cfg, paramsSlice)
		}
	}
}
