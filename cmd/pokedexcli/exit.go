package main

import "os"

func commandExit(_ *currentState, _ []string) error {
	os.Exit(0)
	return nil
}
