package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/c9victor/Pokedex/internal/structs"
)

func CommandMapB(cfg *structs.Config, args ...string) error {
	var reqUrl string
	if cfg.PrevUrl != nil {
		reqUrl = *cfg.PrevUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	cachedValue, exists := cfg.Cache.Get(reqUrl)
	var locations locations
	if exists {
		err := json.Unmarshal(cachedValue, &locations)
		if err != nil {
			fmt.Printf("Error unmarshaling data: %v\n", err)
		}
	} else {
		res, err := http.Get(reqUrl)
		if err != nil {
			fmt.Println("Error retrieving locations: %w", err)
			return err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading data: %w", err)
			return err
		}
		defer res.Body.Close()

		err = json.Unmarshal(body, &locations)
		if err != nil {
			fmt.Println("Error unmarshaling data: %w", err)
			return err
		}
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
