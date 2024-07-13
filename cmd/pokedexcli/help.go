package main

import (
	"fmt"
	"sort"
)

func commandHelp(_ *currentState, _ []string) error {
	fmt.Print("\n")
	commands := getCommands()
	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, commands[key].description)
	}
	return nil
}
