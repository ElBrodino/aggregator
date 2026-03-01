package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Register should only have 1 argument")
	}
	newID := uuid.New()

	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        newID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("Could not create user: %w\n", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("User '%s' has been created but not logged in: %w\n", name, err)
	}

	fmt.Println("User createdsuccessfully!")
	fmt.Printf("* ID:      %v\n", user.ID)
	fmt.Printf("* Name:    %v\n", user.Name)
	return nil
}
