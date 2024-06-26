package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func pokemap(currentState *currentState) error {
	if len(currentState.pokemap.current) == 0 {
		res, err := http.Get("https://pokeapi.co/api/v2/location/")
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
		fmt.Printf("%s", body)
	} else {
		for loc := range currentState.pokemap.current {
			fmt.Println(loc)
		}
	}
	return nil
}

func pokemapNext(currentState *currentState) error {
	res, err := http.Get(currentState.pokemap.next)
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
	fmt.Printf("%s", body)

	return nil
}

func pokemapPrev(currentState *currentState) error {
	res, err := http.Get(currentState.pokemap.prev)
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
	fmt.Printf("%s", body)
	return nil
}

type pokeMap struct {
	current []string
	prev    string
	next    string
}
