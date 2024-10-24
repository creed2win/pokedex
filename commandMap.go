package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// TODO - do I have to refactor to not rewrite so much code?
func commandMap(cfg *config, params paramSlice) error {
	url := cfg.nextLocationsURL

	if url == "" {
		fmt.Println("no prev or next URL. Try other commands.")
		return errors.New("empty URL")
	}

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

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading resp.Body: ", err)
	}

	var pokeLocations pokeLocations

	err = json.Unmarshal(dat, &pokeLocations)

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
