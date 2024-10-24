package main

import "fmt"

func commandPokedex(cfg *config, paramSlice paramSlice) error {

	pokedex := cfg.pokedex
	fmt.Println("Listing all pokemon in you pokedex:")

	for pokemon := range pokedex {
		fmt.Println(" - ", pokemon)
	}
	return nil
}
