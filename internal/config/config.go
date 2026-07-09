package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func write(cfg Config) error {
	configFileName, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Unable to find config file")
		return err
	}
	configContent, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("Unable to marshal config")
		return err
	}
	return os.WriteFile(configFileName, configContent, 0600)
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}

func Read() (Config, error) {
	configFileName, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Unable to find config file")
		return Config{}, err
	}
	var content []byte

	content, err = os.ReadFile(configFileName)
	if err != nil {
		fmt.Println("Unable to read config file")
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Unable to unmarshal config file")
		return Config{}, err
	}

	return config, nil
}
