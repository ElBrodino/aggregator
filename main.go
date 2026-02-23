package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("Lane")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading u[dated config: %v", err)
	}
	fmt.Printf("%+v\n", updatedCfg)
}
