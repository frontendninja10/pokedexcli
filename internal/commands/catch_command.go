package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/frontendninja10/pokedexcli/internal/config"
	"github.com/frontendninja10/pokedexcli/internal/models"
)



func catchPokemon(url, pokemonName string, cfg *config.Config) (models.Pokemon, error) {
	pokemon, exists := cfg.CaughtPokemon[pokemonName]
	if exists {
		return pokemon, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return models.Pokemon{}, fmt.Errorf("possible network error")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.Pokemon{}, fmt.Errorf("unsuccessful request: %d, try again", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Pokemon{}, err
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return models.Pokemon{}, err
	}
	return pokemon, nil
}

func catchCommand(args []string, cfg *config.Config) error {
	if len(args) == 1 {
		return fmt.Errorf("you must provide a Pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[1])

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[1])

	pokemon, err := catchPokemon(url, args[1], cfg)
	if err != nil {
		return err
	}
	baseExperience := pokemon.BaseExperience
	randomValue := rand.Intn(baseExperience)

	fmt.Println(randomValue)

	if randomValue < 20 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}