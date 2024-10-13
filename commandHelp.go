package main

import "fmt"

func commandHelp() error {
	fmt.Println(getCommands()["help"].description)
	for cmd := range getCommands() {
		fmt.Println("- " + cmd)
	}
	return nil
}
