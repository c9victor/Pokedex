package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandCatch(cfg *structs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("One pokemon can be caught at a time")
	}

	pokemonToCatch := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonToCatch)

	reqUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemonToCatch
	res, err := http.Get(reqUrl)
	if err != nil {
		fmt.Printf("Error retrieving Pokemon: %v - %v\n", pokemonToCatch, err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return err
	}
	defer res.Body.Close()

	var pokemon structs.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return err
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	val := rng.Intn(pokemon.BaseExperience)

	if val > int(math.Floor(float64(pokemon.BaseExperience)*0.4)) {
		cfg.UserPokedex[pokemonToCatch] = pokemon

		// FOR TESTING PURPOSES
		// for _, p := range cfg.UserPokedex {
		// 	fmt.Printf("Have caught %v...\n", p.Name)
		// }
		fmt.Printf("%v was caught!\n", pokemonToCatch)
	} else {
		fmt.Printf("%v escaped!\n", pokemonToCatch)
	}

	return nil
}
