package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command expects a single argument- the username", cmd.Name)
	}

	name := sql.NullString{
		String: cmd.Args[0],
		Valid:  true,
	}

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User: %s has been set", cmd.Args[0])

	return nil
}
