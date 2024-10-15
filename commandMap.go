package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func commandMap() error {
	request, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location/?offset=20&limit=20", nil)
	client := http.Client{}
	if err != nil {
		fmt.Println("error making request: ", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error getting respose: ", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading resposne body: ", err)
	}

	decoder := json.NewDecoder(strings.NewReader(string(bodyBytes)))

	var location []Location

	var stringData string

	err = decoder.Decode(&stringData)

	if err != nil {
		return err
	}

	fmt.Println("printing data: ", location)

	return nil
}

//TODO - need to create a data structure to store locations
//then I can unmarshal that data into it.
