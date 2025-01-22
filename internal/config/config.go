package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	gatorConfigFile := fmt.Sprintf("%s/%s", home, configFileName)

	return gatorConfigFile, nil
}

func Read() (Config, error) {
	gatorConfigfile, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	fileContent, err := os.Open(gatorConfigfile)
	if err != nil {
		return Config{}, err
	}

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)

	var config Config

	err = json.Unmarshal(byteResult, &config)
	if err != nil {
		return Config{}, nil
	}
	return config, nil
}

func write(cfg Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (cfg Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	err := write(cfg)
	if err != nil {
		return err
	}

	return nil
}
