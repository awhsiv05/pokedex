package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		fmt.Println(fmt.Sprintf("Your command was: %s", cleanInput[0]))

	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}
