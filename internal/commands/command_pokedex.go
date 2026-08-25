package commands

import (
	"fmt"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandPokedex(cfg *structs.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.UserPokedex {
		fmt.Printf("\t- %v\n", pokemon.Name)
	}
	return nil
}
