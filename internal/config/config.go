package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBurl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func write(cfg Config) error {
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return filepath.Join(homeDir, configFileName), nil
}
