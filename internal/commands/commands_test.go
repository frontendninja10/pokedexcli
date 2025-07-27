package commands

import (
	"testing"

	"github.com/frontendninja10/pokedexcli/internal/config"
	"github.com/frontendninja10/pokedexcli/internal/models"
)

func TestCatchCommand_NoArg(t *testing.T) {
	cfg := &config.Config{
		CaughtPokemon: make(map[string]models.Pokemon),
	}

	err := catchCommand([]string{"catch"}, cfg)

	if err == nil {
		t.Error("Expected error message for missing pokemon name")
	}
}

func TestCatchCommand_AlreadyCaught(t *testing.T) {
	cfg := &config.Config{
		CaughtPokemon: map[string]models.Pokemon{
			"pikachu": {Name: "pikachu"},
		},
	}

	err := catchCommand([]string{"catch", "pikachu"}, cfg)

	if err == nil {
		t.Error("Expected error message for already caught pokemon")
	}
}