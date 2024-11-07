package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
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
	c.CurrentUserName = username
	return write(*c)
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading home path: %w", err)
	}
	path = path + configFileName
	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonConfig, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshaling the data: %w", err)
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return fmt.Errorf("error opening config: %w", err)
	}
	defer file.Close()

	_, err = file.Write(jsonConfig)
	if err != nil {
		return fmt.Errorf("error writing config: %w", err)
	}

	return nil
}
