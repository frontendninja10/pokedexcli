package commands

import (
	"fmt"
	"os"

	"github.com/frontendninja10/pokedexcli/internal/config"
)


func exitCommand(args []string, c *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}