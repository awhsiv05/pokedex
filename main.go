package main

type cliCommand struct {
	name        string
	description string
	callback    func(cnf *config) error
}
type config struct {
	registry map[string]cliCommand
}

func main() {
	configs := &config{
		registry: getCommand(),
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
	}
	return registry
}
