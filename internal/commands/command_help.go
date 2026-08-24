package commands

import (
	"fmt"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandHelp(cfg *structs.Config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n\n")

	for _, command := range cfg.Commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}
