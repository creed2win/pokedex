package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "pokedex > ")
		scanner.Scan()
		command := scanner.Text()

		if commands[command].name == command {
			commands[command].callback()
		}
	}
}
