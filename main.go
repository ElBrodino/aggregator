package main

import (
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := mustReadConfig()
	dbQueries := mustOpenDB(cfg.DBURL)
	cmds := newCommands()
	cmd := mustParseCommand(os.Args)

	appState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	mustRun(cmds, appState, cmd)
}
