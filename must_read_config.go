package main

import (
	"gator/internal/config"
	"log"
)

func mustReadConfig() config.Config {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	return cfg
}
