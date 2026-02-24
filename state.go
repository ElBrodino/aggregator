package main

import (
	"fmt"
	"gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	return nil
}

func (c *commands) register(name string, f func(*state, command)) error {
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("Login should only have 1 argument")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error on setting user")
	}

	fmt.Printf("User set: %s", s.config.CurrentUserName)
	return nil
}
