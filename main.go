package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	appState := state{}
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState.config = &cfg

	//fmt.Printf("%+v\n", updatedCfg)

	cmds := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	// user input checks
	if len(os.Args) < 2 {
		fmt.Printf("Missing arguments, got: %d\n", len(os.Args))
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&appState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
