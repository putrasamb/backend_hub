package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// NewViper returns viper instance
func NewViper(filename, filetype string, paths ...string) (*viper.Viper, error) {
	config := viper.New()

	// Set config file name and type
	config.SetConfigName(filename)
	config.SetConfigType(filetype)

	// Add provided paths to search for the config file
	for _, path := range paths {
		config.AddConfigPath(path)
	}

	// Enable reading from environment variables
	config.AutomaticEnv()

	// Try to read config file
	if err := config.ReadInConfig(); err != nil {
		log.Warnf("Failed to load config file '%s': %v", filename, err)
	} else {
		log.Infof("Successfully loaded config file: %s", config.ConfigFileUsed())
	}

	return config, nil
}
