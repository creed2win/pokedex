package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	url := cfg.nextLocationsURL

	if url == "" {
		fmt.Println("no prev or next URL. Try other commands.")
		return errors.New("empty URL")
	}

	request, err := http.NewRequest("GET", url, nil)
	client := http.Client{}
	if err != nil {
		fmt.Println("error making request: ", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error getting respose: ", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var pokeLocations pokeLocations

	err = decoder.Decode(&pokeLocations)

	if err != nil {
		fmt.Println("error decoding JSON: ", err)
	}

	fmt.Println(url)
	for _, location := range pokeLocations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationsURL = pokeLocations.Next
	cfg.prevLocationsURL = pokeLocations.Previous
	return nil
}
