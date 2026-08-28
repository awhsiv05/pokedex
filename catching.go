package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name           string  `json:"name"`
	URL            string  `json:"url"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Stat struct {
	Name string `json:"name"`
}
type Types struct {
	Type Type `json:"type"`
}
type Type struct {
	Name string `json:"name"`
}

func Catch(cnf *config) error {
	if len(cnf.arguments) == 0 {
		return errors.New("no arguments")
	}
	pokemonToCatch := cnf.arguments[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatch)
	url := DefaultPokemonAPi + pokemonToCatch
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	if float32(randNum) > (float32(pokemon.BaseExperience) * 0.75) {
		fmt.Printf("%s escaped!\n", pokemonToCatch)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonToCatch)
	cnf.Pokemons[pokemonToCatch] = pokemon
	return nil
}

// Inspect

func Inspect(cnf *config) error {
	if len(cnf.arguments) == 0 {
		return errors.New("no arguments")
	}
	pokemonToInspect := cnf.arguments[0]
	if info, ok := cnf.Pokemons[pokemonToInspect]; ok {
		fmt.Printf("Name: %s\n", info.Name)
		fmt.Printf("Height %d\n", info.Height)
		fmt.Printf("Weight %d\n", info.Weight)
		fmt.Println("Stats:")
		for _, stats := range info.Stats {
			fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat+stats.Effort)
		}
		fmt.Println("Types:")
		for _, types := range info.Types {
			fmt.Printf("  - %s\n", types.Type.Name)
		}

		return nil
	}

	fmt.Println("you have not caught this pokemon")
	return nil
}

// pokedex

func Pokedex(cnf *config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cnf.Pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
