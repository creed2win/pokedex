package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "pokedex > ")
		text, _ = reader.ReadString('\n')
		if text != "" {
			break
		}
	}
	fmt.Printf("printing what you just wrote: " + text)
}
