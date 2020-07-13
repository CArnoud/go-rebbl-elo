package config

import (
	"github.com/spf13/viper"
)

// Config is a wrapper to get configs from file or environment variables.
type Config struct {
	SpikeAPIKey           string
	SpikeAPIHost          string
	SpikeCompetitionsPath string
	SpikeMatchesPath      string
}

// NewConfig returns an initialized Config instance.
func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		SpikeAPIKey:           viper.GetString("SPIKE_API_KEY"),
		SpikeAPIHost:          viper.GetString("SPIKE_HOST"),
		SpikeCompetitionsPath: viper.GetString("SPIKE_COMPETITIONS_PATH"),
		SpikeMatchesPath:      viper.GetString("SPIKE_MATCHES_PATH"),
	}, nil
}
