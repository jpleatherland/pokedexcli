package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jpleatherland/pokedexcli/internal/pokecache"
)

func main() {
	userInput := bufio.NewScanner(os.Stdin)
	currentState := new(currentState)
	currentState.pokemap = new(pokeMap)
	currentState.pokecache = pokecache.NewCache(time.Duration(5 * time.Minute))

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
	callback    func(*currentState, []string) error
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
			description: "Prints the current locations or gets locations if no current locations exist",
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
		"explore": {
			name:        "explore",
			description: "Pass in an area name to receive additional information about the area",
			callback:    explore,
		},
	}
}

func executeCommand(userInput string, currentState currentState) error {
	commands := getCommands()
	splitUserInput := strings.Split(userInput, " ")
	command, exists := commands[splitUserInput[0]]
	var userArgs []string
	if !exists {
		return fmt.Errorf("unknown command")
	}
	if len(splitUserInput) > 1 {
		userArgs = splitUserInput[1:]
	} else {
		userArgs = nil
	}
	err := command.callback(&currentState, userArgs)
	if err != nil {
		return err
	}
	return nil
}

type currentState struct {
	pokemap   *pokeMap
	pokecache *pokecache.Cache
}
