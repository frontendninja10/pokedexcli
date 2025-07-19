package commands

import (
	"fmt"

	"github.com/frontendninja10/pokedexcli/internal/config"
)


func helpCommand(args []string, c *config.Config) error {
	fmt.Println()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")

	for k, v := range GetCommands() {
		fmt.Printf("%s: %s\n", k, v.Description)
	}

	fmt.Println()

	return nil
}