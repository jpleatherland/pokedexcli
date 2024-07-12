package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
)

func catch(currentState *currentState, args []string) error {
	if len(args) <= 0 {
		if len(currentState.pokemonInArea) > 0 {
			fmt.Println("Pokemon in your area:")
			for pokemon := range currentState.pokemonInArea {
				fmt.Println(currentState.pokemonInArea[pokemon])
			}
		}
		return errors.New("catch requires a pokemon name")
	}
	if !slices.Contains(currentState.pokemonInArea, args[0]) {
		fmt.Println("Pokemon in your area:")
		for pokemon := range currentState.pokemonInArea {
			fmt.Println(currentState.pokemonInArea[pokemon])
		}
		return errors.New(args[0] + " is not in your most recently explored area")
	}
	fmt.Printf("Threw a pokeball at %s...\n", args[0])
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + args[0])
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
	var tmpPokemon pokemon
	json.Unmarshal(body, &tmpPokemon)
	// check pokemon base experience
	// if < 100 rand int over 25 will catch
	// if < 200 rand int over 50 will catch
	// if > 200 rand int over 75 will catch
	currentState.caughtPokemon = append(currentState.caughtPokemon, tmpPokemon)
	fmt.Println("caught pokemon: " + currentState.caughtPokemon[0].Name)
	return nil
}

type pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
