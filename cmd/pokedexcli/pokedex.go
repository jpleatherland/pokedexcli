package main

import (
	"fmt"
)

func pokedex(currentState *currentState, _ []string) error {
	fmt.Print("\n")
	fmt.Println("Your caught pokemon:")
	for k := range currentState.caughtPokemon {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}
