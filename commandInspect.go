package main

import "fmt"

func commandInspect(cfg *config, paramSlice paramSlice) error {

	pokemon := paramSlice.params[0]
	pokedex := cfg.pokedex

	for ownedPokemon := range pokedex {
		if pokemon == ownedPokemon {
			fmt.Println("Name: ", pokedex[pokemon].Name)
			fmt.Println("Height: ", pokedex[pokemon].Height)
			fmt.Println("Weight: ", pokedex[pokemon].Weight)
		} else {
			fmt.Println("No pokemon named: " + pokemon + " in your pokedex.")
		}
	}
	return nil
}
