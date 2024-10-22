package main

import "os"

func commandExit(cfg *config, params paramSlice) error {
	os.Exit(2)

	return nil
}
