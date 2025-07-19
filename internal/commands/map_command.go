package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/frontendninja10/pokedexcli/internal/config"
	"github.com/frontendninja10/pokedexcli/internal/models"
)

func getLocationAreas(url string, c *config.Config) (models.LocationAreas, error) {
	cache, exists := c.Cache.Get(url)
	if exists {
		var apiRes models.LocationAreas
		if err := json.Unmarshal(cache, &apiRes); err != nil {
			return apiRes, err
		}
		return apiRes, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return models.LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.LocationAreas{}, err
	}

	var apiRes models.LocationAreas
	if err := json.Unmarshal(data, &apiRes); err != nil {
		return models.LocationAreas{}, err
	}

	c.Cache.Add(url, data)
	return  apiRes, nil
}

func mapCommand(args []string, cfg *config.Config) error {
	var url string

	if cfg.NextLocationsUrl == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	} else {
		url = cfg.NextLocationsUrl
	}

	res, err := getLocationAreas(url, cfg)
	if err != nil {
		return err
	}

	cfg.PreviousLocationsUrl = res.Previous
	cfg.NextLocationsUrl = res.Next

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}
