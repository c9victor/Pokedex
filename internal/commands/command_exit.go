package commands

import (
	"fmt"
	"os"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandExit(cfg *structs.Config, args ...string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
