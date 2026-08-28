package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokiRes struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []Areas `json:"results"`
}
type Areas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func mapPrint(data []byte) (string, string, error) {
	var PokRes PokiRes
	if err := json.Unmarshal(data, &PokRes); err != nil {
		return "", "", err
	}
	for _, a := range PokRes.Results {
		fmt.Println(a.Name)
	}
	return PokRes.Next, PokRes.Previous, nil
}

func maker(cnf *config, isNext bool) error {
	var url string
	if isNext {
		if cnf.next == "" {
			return errors.New("invalid request")
		}
		url = cnf.next
		cnf.prev = cnf.next
	} else {
		if cnf.prev == "" {
			return errors.New("invalid request")
		}
		url = cnf.prev
		cnf.next = cnf.prev
	}
	var data []byte
	var ok bool
	var next, prev string

	if data, ok = cnf.cache.Get(url); ok {
		var err error
		if next, prev, err = mapPrint(data); err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)

		if err != nil {
			return err
		}

		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Printf("error closing response body: %v", err)
			}
		}()
		if res.StatusCode != http.StatusOK {
			return errors.New(res.Status)
		}
		if data, err = io.ReadAll(res.Body); err != nil {
			return err
		}
		cnf.cache.Add(url, data)
		if next, prev, err = mapPrint(data); err != nil {
			return err
		}

	}

	if isNext {
		cnf.next = next
	} else {
		cnf.prev = prev
	}
	return nil
}

func getMap(cnf *config) error {

	if err := maker(cnf, true); err != nil {
		return err
	}
	return nil
}

func getMapb(cnf *config) error {
	if err := maker(cnf, false); err != nil {
		return err
	}
	return nil
}

// explore
type pokemonlist struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func printPokiList(data []byte) error {
	var pokilist pokemonlist
	if err := json.Unmarshal(data, &pokilist); err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")
	for _, p := range pokilist.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func Explore(cnf *config) error {
	if len(cnf.arguments) == 0 {
		return errors.New("invalid request , give a location name")
	}
	url := DefaultLocationAreaAPI + cnf.arguments[0]
	if data, ok := cnf.cache.Get(url); ok {
		fmt.Printf("Exploring %s...\n", cnf.arguments[0])
		if err := printPokiList(data); err != nil {
			return err
		}
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	cnf.cache.Add(url, data)
	fmt.Printf("Exploring %s...\n", cnf.arguments[0])
	if err := printPokiList(data); err != nil {
		return err
	}
	return nil
}
