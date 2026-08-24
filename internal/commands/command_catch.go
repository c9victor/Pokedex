package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandCatch(cfg *structs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("One pokemon can be caught at a time")
	}

	pokemon := args[0]
	fmt.Printf("Thowing a Pokeball at %v...\n", pokemon)
	rng := rand.New(rand.NewSource(0))

	return nil
}
