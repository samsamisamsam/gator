package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	db_url            string
	current_user_name string
}

func Read() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config: %w", err)
	}
	defer file.Close()

	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json config: %w", err)
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.current_user_name = username
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading home path: %w", err)
	}
	path = path + "/.gatorconfig.json"
	return path, nil
}
