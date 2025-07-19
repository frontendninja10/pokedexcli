package commands

import (
	"fmt"

	"github.com/frontendninja10/pokedexcli/internal/config"
)

func inspectCommand(args []string, cfg *config.Config) error {
	if len(args) == 1 {
		return fmt.Errorf("input Pokemon to inspect")
	}

	pokemon, exists := cfg.CaughtPokemon[args[1]]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeObject := range pokemon.Types {
		fmt.Printf("- %s\n", typeObject.Type.Name)
	}
	return nil
}