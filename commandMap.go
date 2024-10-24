package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *config, params paramSlice) error {
	url := cfg.nextLocationsURL

	if val, ok := cfg.cache.Get(url); ok {
		var pokeLocations pokeLocations
		err := json.Unmarshal(val, &pokeLocations)
		if err != nil {
			return err
		}
		fmt.Println(url)
		for _, location := range pokeLocations.Results {
			fmt.Println(location.Name)
		}
		cfg.nextLocationsURL = pokeLocations.Next
		cfg.prevLocationsURL = pokeLocations.Previous
		return nil
	}

	dat := getPokeData(cfg, "nextLocationUrl")

	var pokeLocations pokeLocations

	err := json.Unmarshal(dat, &pokeLocations)

	if err != nil {

		fmt.Println("error Unmarshaling 'dat': ", err)
	}

	fmt.Println(url)
	for _, location := range pokeLocations.Results {
		fmt.Println(location.Name)
	}

	cfg.cache.Add(url, dat)

	cfg.nextLocationsURL = pokeLocations.Next
	cfg.prevLocationsURL = pokeLocations.Previous
	return nil
}
