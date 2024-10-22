package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, params paramSlice) error {
	baseUrl := "https://pokeapi.co/api/v2/location/"
	fmt.Println("calling explore test")

	client := http.Client{}
	location := params.params[0]
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading data from response body", err)
		return err
	}

	fmt.Println(dat)

	return nil
}
