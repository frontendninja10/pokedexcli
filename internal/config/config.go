package config

import (
	"github.com/frontendninja10/pokedexcli/internal/models"
	"github.com/frontendninja10/pokedexcli/internal/pokecache"
)

type Config struct {
	PreviousLocationsUrl string
	NextLocationsUrl string
	Cache *pokecache.Cache	
	CaughtPokemon map[string]models.Pokemon
}