package main

import (
	"errors"
	"gator/internal/config"
)

func readConfig() (config.Config, error) {
	cfg, err := config.Read()
	if err != nil {
		return config.Config{}, errors.New("can't read config")
	}
	return cfg, nil
}
