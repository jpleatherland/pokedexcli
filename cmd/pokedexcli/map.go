package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func pokemap(currentState *currentState, _ []string) error {
	if len(currentState.pokemap.Results) == 0 {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
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
		json.Unmarshal(body, &currentState.pokemap)
		for loc := range currentState.pokemap.Results {
			fmt.Println(currentState.pokemap.Results[loc].Name)
		}
	} else {
		for loc := range currentState.pokemap.Results {
			fmt.Println(currentState.pokemap.Results[loc].Name)
		}
	}
	return nil
}

func pokemapNext(currentState *currentState, _ []string) error {
	val, exists := currentState.pokecache.Get(currentState.pokemap.Next)
	if exists {
		json.Unmarshal(val, &currentState.pokemap)
	} else {
		res, err := http.Get(currentState.pokemap.Next)
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
		currentState.pokecache.Add(currentState.pokemap.Next, body)
		json.Unmarshal(body, &currentState.pokemap)
	}
	for loc := range currentState.pokemap.Results {
		fmt.Println(currentState.pokemap.Results[loc].Name)
	}
	return nil
}

func pokemapPrev(currentState *currentState, _ []string) error {
	val, exists := currentState.pokecache.Get(currentState.pokemap.Previous)
	if exists {
		json.Unmarshal(val, &currentState.pokemap)
	} else {
		res, err := http.Get(currentState.pokemap.Previous)
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
		currentState.pokecache.Add(currentState.pokemap.Previous, body)
		json.Unmarshal(body, &currentState.pokemap)
	}
	for loc := range currentState.pokemap.Results {
		fmt.Println(currentState.pokemap.Results[loc].Name)
	}
	return nil
}

type pokeMap struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}
