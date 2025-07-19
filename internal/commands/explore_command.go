package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/frontendninja10/pokedexcli/internal/config"
	"github.com/frontendninja10/pokedexcli/internal/models"
)

func exploreLocation(url string, c *config.Config) (models.LocationAreaDetails, error) {
	cache, exists := c.Cache.Get(url)
	if exists {
		var locaionAreaRes models.LocationAreaDetails
		if err := json.Unmarshal(cache, &locaionAreaRes); err != nil {
			return locaionAreaRes, err
		}
		return locaionAreaRes, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return models.LocationAreaDetails{}, fmt.Errorf("error creating request")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.LocationAreaDetails{}, fmt.Errorf("recieved non-200 status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.LocationAreaDetails{}, err
	}

	var locaionAreaDetails models.LocationAreaDetails

	if err := json.Unmarshal(data, &locaionAreaDetails); err != nil {
		return  models.LocationAreaDetails{}, err
	}

	c.Cache.Add(url, data)

	return locaionAreaDetails, nil
} 

func exploreCommand(args []string, c *config.Config) error {
	if len(args) == 1 {
		return fmt.Errorf("you must provide a location")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[1]

	res, err := exploreLocation(url, c)
	if err != nil {
		return err
	}

	pokemonEncounters := res.PokemonEncounters
	fmt.Printf("Exploring %s...\n", args[1])
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range pokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}