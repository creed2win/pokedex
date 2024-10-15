package main

import "os"

func commandExit(cfg *config) error {
	os.Exit(2)

	return nil
}
