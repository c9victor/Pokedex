package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Print(words)
		switch words[0] {
		case "exit":
			os.Exit(0)
		default:
			continue
		}
		// func commandExist() error {
		// 	fmt.Print("Closing the Pokedex... Goodbye!")
		// 	os.Exit(0)
		// }
	}
}
