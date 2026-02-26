package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	dbURL := cfg.DbURL // "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.Ping()

	dbQueries := database.New(db)
	appState := &state{
		config: &cfg,
		db:     dbQueries,
	}

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	// user input checks
	if len(os.Args) < 2 {
		fmt.Printf("Missing arguments, got: %d\n", len(os.Args))
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(appState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//appState.db = dbQueries
}
