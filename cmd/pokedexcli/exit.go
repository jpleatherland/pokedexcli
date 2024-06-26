package main

import "os"

func commandExit(_ *currentState) error {
	os.Exit(0)
	return nil
}
