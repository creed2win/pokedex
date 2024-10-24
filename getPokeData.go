package main

import (
	"fmt"
	"io"
	"net/http"
)

func getPokeData(cfg *config, urlType string) []byte {
	var data []byte
	var url string

	switch urlType {
	case "nextLocationUrl":
		url = cfg.nextLocationsURL
	case "prevLocationUrl":
		url = cfg.prevLocationsURL
	case "locationAreaUrl":
		url = cfg.locationAreaUrl
	}

	if url == "" {
		fmt.Println("no prev or next URL. Try other commands.")

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
	}

	resp, err := cfg.client.Do(req)
	if err != nil {
		fmt.Println("error getting response: ", err)
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading data from response", err)
	}
	return data
}
