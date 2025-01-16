package main

import (
	"context"
	"fmt"
)

func usersHandler(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("no arguments needed for %s command", cmd.Name)
	}

	usersSlice, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch users from database: %w", err)
	}

	for _, user := range usersSlice {
		name := user.Name
		if s.cfg.CurrentUserName == name {
			fmt.Printf("- %s (current)\n", name)
		} else {
			fmt.Printf("- %s\n", name)
		}
	}

	return nil
}
