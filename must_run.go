package main

import (
	"fmt"
	"os"
)

func mustRun(cmds commands, appState *state, cmd command) {
	err := cmds.run(appState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
