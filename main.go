package main

import (
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := mustReadConfig()
	dbQueries, db := mustOpenDB(cfg.DBURL)
	defer db.Close()
	cmds := newCommands()
	cmd := mustParseCommand(os.Args)

	appState := &state{
		config: &cfg,
		db:     dbQueries,
	}

	mustRun(cmds, appState, cmd)
}
