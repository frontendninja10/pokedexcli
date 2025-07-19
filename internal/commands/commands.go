package commands

import "github.com/frontendninja10/pokedexcli/internal/config"

type Command struct {
 Name string
 Description string
 Callback func(args []string, c *config.Config) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"map": {
			Name: "map",
			Description: "Displays names of 20 location areas",
			Callback: mapCommand,
		},
		"mapb": {
			Name: "mapb",
			Description: "Displays the previous 20 location areas",
			Callback: mapbCommand,
		},
		"explore": {
			Name: "explore <location_name>",
			Description: "Displays the names of the Pokemon located in the inputed area",
			Callback: exploreCommand,
		},
		"catch": {
			Name: "catch <pokemon_name>",
			Description: "Tries to catch a Pokemon",
			Callback: catchCommand,
		},
		"inspect": {
			Name: "inspect <pokemon_name>",
			Description: "Displays Pokemon name, height, weight, stats, and type(s)",
			Callback: inspectCommand,
		},
		"pokedex": {
			Name: "pokedex",
			Description: "Prints a list of all the names of the Pokemon the user has caught",
			Callback: pokedex,
		},
	}
}