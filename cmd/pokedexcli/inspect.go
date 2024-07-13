package main

import (
	"errors"
	"fmt"
)

func inspect(currentState *currentState, args []string) error {
	fmt.Print("\n")
	if len(args) <= 0 {
		return errors.New("need a pokemon to inspect")
	}
	pokemon, exists := currentState.caughtPokemon[args[0]]
	if !exists {
		return errors.New("you haven't caught that pokemon yet")
	}
	pokemon.print()
	return nil
}
