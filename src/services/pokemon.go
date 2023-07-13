package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/clients"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/models"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/utils"
)

func FindAll() ([]models.Pokemon, error) {
	var pokemonList, error = LoadPokemonFromCSV()
	if error != nil {
		return nil, error
	}
	return pokemonList, nil
}

func FindById(id int) (*models.Pokemon, error) {
	var pokemonList, error = LoadPokemonFromCSV()
	if error != nil {
		return nil, error
	}
	for _, pokemon := range pokemonList {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}
	return nil, fmt.Errorf("pokemon id=%d does not exist", id)
}

func LoadPokemonFromCSV() ([]models.Pokemon, error) {
	csvFilePath := viper.GetString("CSV_FILE")
	data, err := utils.Read(csvFilePath)
	if err != nil {
		return nil, err
	}
	return createPokemonList(data)
}

func createPokemonList(data [][]string) ([]models.Pokemon, error) {
	var pokemonList []models.Pokemon

	for _, line := range data {
		var pokemon models.Pokemon
		for j, field := range line {
			if j == 0 {
				id, error := strconv.Atoi(field)
				if error != nil {
					return nil, fmt.Errorf("id=%s is not a valid integer type", field)
				}
				pokemon.Id = id
			} else if j == 1 {
				pokemon.Name = field
			}
		}
		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList, nil
}

func Import(limit, offset int) (*clients.PokemonListResponse, error) {
	csvFilePath := viper.GetString("CSV_FILE")
	body, error := clients.GetPokemonList(10, 10)
	if error != nil {
		return nil, error
	}
	records := [][]string{}
	for _, record := range body.Results {
		row := []string{getId(record.Url), record.Name}
		records = append(records, row)
	}
	utils.Write(csvFilePath, records)
	return &body, error
}

func getId(url string) string {
	parts := strings.Split(strings.Trim(url, "/"), "/")
	return parts[len(parts)-1]
}
