package main

import "fmt"

func commandHelp(_ *currentState) error {
	fmt.Print("\n")
	for _, contents := range getCommands() {
		fmt.Printf("%s: %s\n", contents.name, contents.description)
	}
	return nil
}
