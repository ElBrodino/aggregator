package main

import (
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Register should only have 1 argument")
	}

	err := s.config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error on setting user")
	}

	return nil
}
