package models

type Pokemon struct {
	Name string `json:"name"`
	Weight int `json:"weight"`
	Height int `json:"height"`
	BaseExperience int `json:"base_experience"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type LocationAreaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string 	`json:"url"`
		} `json:"pokemon"`
		
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance 	int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL string 	`json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method struct {
					Name string `json:"name"`
					URL string 	`json:"url"`
				}
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version struct {
				Name string `json:"name"`
				URL string 	`json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationAreas struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}