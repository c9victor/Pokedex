package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/c9victor/Pokedex/internal/commands"
	"github.com/c9victor/Pokedex/internal/structs"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := structs.NewCache(5 * time.Minute)
	cfg := &structs.Config{
		Commands:    getCommands(),
		Cache:       cache,
		UserPokedex: make(map[string]structs.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		commandName := words[0]
		args := words[1:]
		cmd, exists := cfg.Commands[commandName]

		if exists {
			err := cmd.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	trimmedText := strings.Trim(text, " ")
	loweredText := strings.ToLower(trimmedText)
	words := strings.Split(loweredText, " ")
	return words
}

func getCommands() map[string]structs.CliCommand {
	cmds := make(map[string]structs.CliCommand)
	cmds["exit"] = structs.CliCommand{
		Name:        "exit",
		Description: "Exits the Pokedex",
		Callback:    commands.CommandExit,
	}
	cmds["help"] = structs.CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commands.CommandHelp,
	}
	cmds["map"] = structs.CliCommand{
		Name:        "map",
		Description: "Displays the next 20 locations in the Pokemon world",
		Callback:    commands.CommandMap,
	}
	cmds["mapb"] = structs.CliCommand{
		Name:        "mapb",
		Description: "Displays the previous 20 locations in the Pokemon world",
		Callback:    commands.CommandMapB,
	}
	cmds["explore"] = structs.CliCommand{
		Name:        "explore <LOCATION>",
		Description: "Lists all pokemon in a given area. Argument LOCATION required",
		Callback:    commands.CommandExplore,
	}
	cmds["catch"] = structs.CliCommand{
		Name:        "catch <POKEMON>",
		Description: "Attempt to catch a Pokemon! Argument POKEMON required",
		Callback:    commands.CommandCatch,
	}
	cmds["inspect"] = structs.CliCommand{
		Name:        "inspect <POKEMON>",
		Description: "Take a look at the pokemon you have caught! Argument POKEMON required",
		Callback:    commands.CommandInspect,
	}
	cmds["pokedex"] = structs.CliCommand{
		Name:        "pokedex",
		Description: "Shows all the Pokemon you have caught!",
		Callback:    commands.CommandPokedex,
	}
	return cmds
}
