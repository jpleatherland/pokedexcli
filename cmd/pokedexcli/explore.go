package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func explore(currentState *currentState, args []string) error {
	if len(args) <= 0 {
		return errors.New("require a location to explore")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + args[0])
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		errText := fmt.Sprintf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		return errors.New(errText)
	}
	var response exploredArea
	json.Unmarshal(body, &response)
	fmt.Println("Found pokemon:")
	currentState.pokemonInArea = nil
	for encounter := range response.PokemonEncounters {
		currentState.pokemonInArea = append(currentState.pokemonInArea, response.PokemonEncounters[encounter].Pokemon.Name)
		fmt.Println(response.PokemonEncounters[encounter].Pokemon.Name)
	}
	return nil
}

type exploredArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
