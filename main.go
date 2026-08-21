package main

type cliCommand struct {
	name        string
	description string
	callback    func(cnf *config) error
}
type config struct {
	registry map[string]cliCommand
	prev     string
	next     string
}

func main() {
	defaultLocationAPI := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	configs := &config{
		registry: getCommand(),
		prev:     "",
		next:     defaultLocationAPI,
	}
	repl(configs)
}
func getCommand() map[string]cliCommand {
	var registry = map[string]cliCommand{
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
			callback:    getmap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Display map locations backwards",
			callback:    getmapb,
		},
	}
	return registry
}
