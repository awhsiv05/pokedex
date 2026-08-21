package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		cmd := cleanInput[0]
		if _, ok := config.registry[cmd]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := config.registry[cmd].callback(config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func help(cnf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for name, command := range cnf.registry {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandExit(cnf *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}
