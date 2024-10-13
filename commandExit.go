package main

import "os"

func commandExit() error {
	os.Exit(2)

	return nil
}
