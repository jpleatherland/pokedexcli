package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Pokedex > ")
	userInput := bufio.NewScanner(os.Stdin)
	for userInput.Scan() {
		fmt.Println(eval(userInput))
		fmt.Print("Pokedex > ")
	}
}

func eval(command *bufio.Scanner) string {
	return fmt.Sprintf("Gotta parse 'em all: %v", command.Text())
}
