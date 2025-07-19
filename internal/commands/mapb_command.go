package commands

import (
	"fmt"

	"github.com/frontendninja10/pokedexcli/internal/config"
)

func mapbCommand(args []string, cfg *config.Config) error {
	if cfg.NextLocationsUrl == "" && cfg.PreviousLocationsUrl == "" {
		return fmt.Errorf("you cannot go back")
	}

	if cfg.PreviousLocationsUrl == "" {
		return fmt.Errorf("you are on first page")
	} else {
		url := cfg.PreviousLocationsUrl

		res, err := getLocationAreas(url, cfg)
		if err != nil {
			return err
		}

		cfg.NextLocationsUrl = res.Next
		cfg.PreviousLocationsUrl = res.Previous

		for _, area := range res.Results {
			fmt.Println(area.Name)
		}
	}
	return nil
}