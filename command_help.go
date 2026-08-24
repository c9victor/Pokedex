package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n\n")

	for _, command := range cfg.commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
