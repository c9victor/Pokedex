package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/c9victor/Pokedex/internal/structs"
)

type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func CommandExplore(cfg *structs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Argument location must be provided")
	}

	reqUrl := "https://pokeapi.co/api/v2/location-area/" + args[0]
	res, err := http.Get(reqUrl)
	if err != nil {
		fmt.Printf("Error getting location data for %v: %v\n", args[0], err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error getting data: %v\n", err)
		return err
	}
	defer res.Body.Close()

	var locationInformation locationArea
	err = json.Unmarshal(body, &locationInformation)
	if err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")
	pokemonString := ""
	for _, pokemon := range locationInformation.PokemonEncounters {
		pokemonString += "- " + pokemon.Pokemon.Name + "\n"
	}
	strings.Trim(pokemonString, " ")
	fmt.Printf("%v", pokemonString)

	return nil
}
