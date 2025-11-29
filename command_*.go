package main

import (
	"fmt"
	"os"
	"math/rand"

	"github.com/kblasti/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *config, client *pokeapi.Client, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, client *pokeapi.Client, args []string) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, j := range commands {
		fmt.Printf("%v: %v\n", j.name, j.description)
	}
	return nil
}

func commandMap(cfg *config, client *pokeapi.Client, args []string) error {
	url := ""
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	} else {
		url = *cfg.Next
	}

	resp, err := client.GetLocationAreas(url)
	if err != nil {
    	return err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, r := range resp.Results {
    	fmt.Println(r.Name)
	}
	return nil
}

func commandMapb(cfg *config, client *pokeapi.Client, args []string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.Previous
	resp, err := client.GetLocationAreas(url)
	if err != nil {
    	return err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, r := range resp.Results {
    	fmt.Println(r.Name)
	}
	return nil
}

func commandExplore(cfg *config, client *pokeapi.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Must pass location name")
		return nil
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	url := "https://pokeapi.co/api/v2/location-area/" + areaName
	loc, err := client.GetLocation(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, client *pokeapi.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Must pass Pokemon name")
		return nil
	}

	pokemonName := args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pok, err := client.GetPokemon(url)
	if err != nil {
		return err
	}

	baseEXP := pok.BaseExperience
	if baseEXP < 50 {
		baseEXP = 50
	}
	if baseEXP > 300 {
		baseEXP = 300
	}

	position := (baseEXP - 50) / 250
	prob := 0.9 - float64(position) * 0.8

	random :=  float64(rand.Intn(100))
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if random <= (prob * 100) {
		fmt.Printf("%s was caught!\n", pokemonName)
		(*cfg.caughtPokemon)[pokemonName] = pok
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	
	return nil
}

func commandInspect(cfg *config, client *pokeapi.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Must pass Pokemon name")
		return nil
	}
	
	pokemonName := args[0]

	pok, ok := (*cfg.caughtPokemon)[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pok.Height)
	fmt.Printf("Weight: %v\n", pok.Weight)
	fmt.Println("Stats:")
	for _, s := range pok.Stats {
		fmt.Printf(" -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pok.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, client *pokeapi.Client, args []string) error {
	if len(*cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}
	
	fmt.Println("Your Pokedex")
	for _, n := range (*cfg.caughtPokemon) {
		fmt.Printf(" - %s\n", n.Name)
	}
	return nil
}