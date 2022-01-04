package main

import (
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	DBHost   string
	DBPort   int
	DBUser   string
	DBPass   string
	HTTPPort string
}

func getEnvConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
