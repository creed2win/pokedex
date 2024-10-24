package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

func commandCatch(cfg *config, params paramSlice) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"

	pokemonName := params.params[0]

	req, err := http.NewRequest("GET", (baseUrl + pokemonName), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// actual catching logic
	playerRoll := rand.IntN(2000)
	fmt.Println("Catching pokemon with name: ", pokemon.Name)
	fmt.Println("This pokemon has base experience of: ", pokemon.BaseExperience)
	fmt.Println("You rolled in range 1-2000:", playerRoll)
	if playerRoll < pokemon.BaseExperience {
		fmt.Println("Bad luck! You did not catch ", pokemonName, ".  Maybe next time!")
	} else {
		cfg.pokedex[pokemonName] = pokemon
		fmt.Println("Congatulation!! ", pokemonName, "in now in your pokedex.")
	}
	return nil
}
