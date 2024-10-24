package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, params paramSlice) error {
	baseUrl := cfg.locationAreaUrl
	if len(params.params) < 1 {
		fmt.Println("No parameters provided to 'explore'")
		return fmt.Errorf("no params")
	}
	location := params.params[0]

	dat, ok := cfg.cache.Get(baseUrl + location)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Exploring " + locationArea.Name + "...")
		fmt.Println("Found pokemon: ")
		for _, pokemon := range locationArea.PokemonEncounters {
			fmt.Println("- " + pokemon.Pokemon.Name)
		}
	}

	client := http.Client{}

	request, err := http.NewRequest("GET", (baseUrl + location), nil)
	if err != nil {
		fmt.Println("error making a request", err)
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error getting response", err)
		return err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading data from response body", err)
		return err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Exploring " + locationArea.Name + "...")
	fmt.Println("Found pokemon: ")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	cfg.cache.Add((baseUrl + location), dat)
	return nil
}
