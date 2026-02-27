package main

import "errors"

func parseCommand(args []string) (command, error) {
	if len(args) < 2 {
		return command{}, errors.New("not enough arguments")
	}
	return command{
		Name: args[1],
		Args: args[2:],
	}, nil
}
