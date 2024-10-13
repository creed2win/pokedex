package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	request, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location/?offset=20&limit=20", nil)
	client := http.Client{}
	if err != nil {
		return fmt.Errorf("error making request: ", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Errorf("error getting respose: ", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	//var data string

	//json.Unmarshal(bodyBytes, &data)

	fmt.Println("printing data: ", string(bodyBytes))

	return nil
}

//TODO - need to create a data structure to store locations
//then I can unmarshal that data into it.
