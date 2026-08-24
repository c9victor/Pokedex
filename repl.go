package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	prevUrl  string
	nextUrl  string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		commands: getCommands(),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		commandName := words[0]
		cmd, exists := cfg.commands[commandName]
		if exists {
			err := cmd.callback(cfg)
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

func getCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 locations in the Pokemon world",
		callback:    commandMap,
	}
	return commands
}
