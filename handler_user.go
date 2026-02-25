package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login should only have 1 argument")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error on setting user")
	}

	return nil
}
