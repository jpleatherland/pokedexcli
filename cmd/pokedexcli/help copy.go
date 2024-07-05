package main

import "fmt"

func explore(_ *currentState, _ []string) error {
	fmt.Print("\n")
	for _, contents := range getCommands() {
		fmt.Printf("%s: %s\n", contents.name, contents.description)
	}
	return nil
}
