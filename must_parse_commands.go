package main

import (
	"os"
)

func mustParseCommand(args []string) command {
	cmd, err := parseCommand(args)
	if err != nil {
		os.Exit(1)
	}
	return cmd
}
