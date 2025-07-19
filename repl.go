package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/frontendninja10/pokedexcli/internal/commands"
	"github.com/frontendninja10/pokedexcli/internal/config"
	"github.com/frontendninja10/pokedexcli/internal/models"
	"github.com/frontendninja10/pokedexcli/internal/pokecache"
)



func startRepl() {
	var c config.Config

	newCacheVal := pokecache.NewCache(60 * time.Second)

	c.Cache = &newCacheVal
	c.CaughtPokemon = map[string]models.Pokemon{}

	
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		reader.Scan()

		input := reader.Text()

		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}


		command, exists := commands.GetCommands()[cleanedInput[0]]
		if exists {
			err := command.Callback(cleanedInput, &c)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)

	return  output
}







