package commands

import (
	"errors"
	"fmt"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandInspect(cfg *structs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon specified in call to inspect")
	}

	pokemon, ok := cfg.UserPokedex[args[0]]
	if !ok {
		fmt.Println("You have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("\t-%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("\t- %v\n", t.Type.Name)
		}
	}

	return nil
}
