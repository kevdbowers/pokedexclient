package main

import "os"

func commandExit(config *configuration) error {
	os.Exit(0)
	return nil
}
