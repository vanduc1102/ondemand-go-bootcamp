package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	CsvFile            string `mapstructure:"CSV_FILE"`
	PokemonApiEndpoint string `mapstructure:"POKEMON_API_ENDPOINT"`
}

func LoadConfig(path string, name string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
