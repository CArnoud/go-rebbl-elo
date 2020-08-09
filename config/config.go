package config

import (
	"github.com/spf13/viper"
)

// Config is a wrapper to get configs from file or environment variables.
type Config struct {
	SpikeAPIKey           string
	SpikeAPIHost          string
	SpikeCompetitionsPath string
	SpikeContestsPath     string
	DatabaseHost          string
	DatabasePort          string
	DatabaseUser          string
	DatabasePassword      string
	DatabaseName          string
}

// NewConfig returns an initialized Config instance.
func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		SpikeAPIKey:           viper.GetString("SPIKE_API_KEY"),
		SpikeAPIHost:          viper.GetString("SPIKE_HOST"),
		SpikeCompetitionsPath: viper.GetString("SPIKE_COMPETITIONS_PATH"),
		SpikeContestsPath:     viper.GetString("SPIKE_CONTESTS_PATH"),
		DatabaseHost:          viper.GetString("DATABASE_HOST"),
		DatabasePort:          viper.GetString("DATABASE_PORT"),
		DatabaseUser:          viper.GetString("DATABASE_USER"),
		DatabasePassword:      viper.GetString("DATABASE_PASSWORD"),
		DatabaseName:          viper.GetString("DATABASE_NAME"),
	}, nil
}
