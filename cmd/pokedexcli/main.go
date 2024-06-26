package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userInput := bufio.NewScanner(os.Stdin)
	currentState := new(currentState)
	currentState.pokemap = new(pokeMap)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()
		err := executeCommand(userInput.Text(), *currentState)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("\n")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*currentState) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints a list of available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Prints the current locations",
			callback:    pokemap,
		},
		"mapn": {
			name:        "mapn",
			description: "Prints the next set of locations and sets the current map",
			callback:    pokemapNext,
		},
		"mapp": {
			name:        "mapp",
			description: "Prints the previous set of locations and sets the current map",
			callback:    pokemapPrev,
		},
	}
}

func executeCommand(userInput string, currentState currentState) error {
	commands := getCommands()
	command, exists := commands[userInput]
	if !exists {
		return fmt.Errorf("unknown command")
	}
	err := command.callback(&currentState)
	if err != nil {
		return err
	}
	return nil
}

type currentState struct {
	pokemap *pokeMap
}
