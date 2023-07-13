package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func get(url string) []uint8 {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonListResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

func GetPokemonList(limit, offset int) (PokemonListResponse, error) {
	url := fmt.Sprintf("%spokemon?limit=%d&offset=%d", viper.GetString("POKEMON_API_ENDPOINT"), limit, offset)
	body := get(url)

	var result PokemonListResponse
	err := json.Unmarshal(body, &result)
	return result, err
}
