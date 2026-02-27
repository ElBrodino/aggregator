package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handlerFunc, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("command not found")
	}
	return handlerFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
