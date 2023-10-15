// Package config is responsible for initiating client and server configuration.
// config consists of Config struct and NewConfig function to create Config instance.
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// NewConfig creates new Config instance.
func NewConfig(configPath string) (*Config, error) {
	conf, err := readConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error in readConfig: %w", err)
	}
	return conf, nil
}

// Config collects client and server configurations.
type Config struct {
	Address     string `json:"address"`      // server address
	DatabaseDSN string `json:"database_dsn"` // database address
}

// readConfig reads configuration file and returns Config instance.
func readConfig(configPath string) (*Config, error) {
	// read configuration file
	f, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	// extract configuration
	fileConf, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading configurations: %w", err)
	}

	// unmarshall json
	var conf Config
	err = json.Unmarshal(fileConf, &conf)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling to Config: %w", err)
	}

	return &conf, nil
}
