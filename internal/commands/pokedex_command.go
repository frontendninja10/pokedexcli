package commands

import (
	"fmt"

	"github.com/frontendninja10/pokedexcli/internal/config"
)

func pokedex(args []string, cfg *config.Config) error {
	fmt.Println("Your Pokedex:")
	for _, value := range cfg.CaughtPokemon {
		fmt.Printf("- %s\n", value.Name)
	}

	return nil
}