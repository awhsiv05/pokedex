package main

import (
	"time"

	"github.com/meowmonsters/pokedexcli/internal"
)

const (
	DefaultLocationAPI     string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	DefaultLocationAreaAPI string = "https://pokeapi.co/api/v2/location-area/"
	DefaultPokemonAPi      string = "https://pokeapi.co/api/v2/pokemon/"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cnf *config) error
}
type config struct {
	registry  map[string]cliCommand
	prev      string
	next      string
	cache     *pokecache.Cache
	arguments []string
	Pokemons  map[string]Pokemon
}

func main() {
	configs := &config{
		registry:  getCommand(),
		prev:      "",
		next:      DefaultLocationAPI,
		cache:     pokecache.NewCache(5 * time.Second),
		arguments: make([]string, 0),
		Pokemons:  make(map[string]Pokemon),
	}
	repl(configs)
}
func getCommand() map[string]cliCommand {
	registry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit this program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    help,
		},
		"map": {
			name:        "map",
			description: "Display map locations",
			callback:    getMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display map locations backwards",
			callback:    getMapb,
		},
		"explore": {
			"explore",
			"Displays Pokemon in given location explore ",
			Explore,
		},
		"catch": {
			"catch",
			"Tries to catch Pokemon",
			Catch,
		},
		"inspect": {
			"inspect",
			"describes the pokemon",
			Inspect,
		},
		"pokedex": {
			"pokedex",
			"Displays all pokemons of the user",
			Pokedex,
		},
	}
	return registry
}
