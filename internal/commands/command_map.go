package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/c9victor/Pokedex/internal/structs"
)

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(cfg *structs.Config, args ...string) error {
	reqUrl := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextUrl != nil {
		reqUrl = *cfg.NextUrl
	}

	cachedValue, exists := cfg.Cache.Get(reqUrl)
	var locations locations
	if exists {
		err := json.Unmarshal(cachedValue, &locations)
		if err != nil {
			fmt.Printf("Error unmarshaling data: %v\n", err)
			return err
		}
	} else {
		res, err := http.Get(reqUrl)
		if err != nil {
			fmt.Printf("Error retrieving locations: %v\n", err)
			return err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			return err
		}
		defer res.Body.Close()

		err = json.Unmarshal(body, &locations)
		if err != nil {
			fmt.Printf("Error unmarshaling data: %v\n", err)
			return err
		}
		cfg.Cache.Add(reqUrl, body)
	}

	cfg.PrevUrl = locations.Previous
	cfg.NextUrl = locations.Next

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	if exists {
		fmt.Println("Value retrieved from cache")
	}

	return nil
}
