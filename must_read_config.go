package main

import (
	"gator/internal/config"
	"os"
)

func mustReadConfig() config.Config {
	cfg, err := readConfig()
	if err != nil {
		os.Exit(1)
	}

	return cfg
}
