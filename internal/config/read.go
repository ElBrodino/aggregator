package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func read() Config {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}
	}

	return cfg
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullpath := filepath.Join(homeDir, configFileName)
	return fullpath, nil
}
