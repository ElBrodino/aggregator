package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg := config.Config{}
	cfg.SetUser("Lane")
	fmt.Println(config.Read())
}
