package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"time"

	"github.com/kblasti/pokedexcli/internal/pokeapi"
)

type repl struct {
	client pokeapi.Client
}

type config struct {
	Next	 *string
	Previous *string
	caughtPokemon *map[string]*pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokeapi.Client, []string) error
}

func startRepl(cfg *config) {
	timeout := 5 * time.Second
    cacheInterval := 5 * time.Second

    r := repl{
        client: pokeapi.NewClient(timeout, cacheInterval),
    }
	m := make(map[string]*pokeapi.Pokemon)
	cfg.caughtPokemon = &m
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := strings.Fields(strings.ToLower(input))
		command := cleaned[0]
		args := cleaned[1:]
		_, ok := commands[command]
		if ok {
			commands[command].callback(cfg, &r.client, args)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"map": {
			name: 		 "map",
			description: "Get the next page of locations",
			callback: 	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Get the previous page of locations",
			callback: 	 commandMapb,
		},
		"explore": {
			name: 		 "explore",
			description: "Show list of what pokemon are at a given location",
			callback: 	 commandExplore,
		},
		"catch": {
			name: 		 "catch",
			description: "Attempt to catch a pokemon",
			callback: 	 commandCatch,
		},
		"inspect": {
			name: 		 "inspect",
			description: "Look at pokemon info in your Pokedex",
			callback: 	 commandInspect,
		},
		"pokedex": {
			name: 		 "pokedex",
			description: "Look at list of caught pokemon",
			callback: 	 commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
