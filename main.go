package main

import (
	"time"

	"github.com/meowmonsters/pokedexcli/internal"
)

const (
	DefaultLocationAPI     string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	DefaultLocationAreaAPI string = "https://pokeapi.co/api/v2/location-area/"
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
}

func main() {
	configs := &config{
		registry:  getCommand(),
		prev:      "",
		next:      DefaultLocationAPI,
		cache:     pokecache.NewCache(5 * time.Second),
		arguments: make([]string, 0),
	}
	repl(configs)
}
func getCommand() map[string]cliCommand {
	registry := map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit this program",
			callback:    commandExit,
		},
		"help": cliCommand{
			name:        "help",
			description: "Display this help message",
			callback:    help,
		},
		"map": cliCommand{
			name:        "map",
			description: "Display map locations",
			callback:    getMap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Display map locations backwards",
			callback:    getMapb,
		},
		"explore": cliCommand{
			name:        "explore",
			description: "Displays Pokemon in given location explore ",
			callback:    Explore,
		},
	}
	return registry
}
