package main

import (
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Register should only have 1 argument")
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error on setting user: %w", err)
	}

	fmt.Printf("User '%s' has been created and logged in.\n", name)
	return nil
}
