package main

import (
	"fmt"
)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command expects a single argument- the username", cmd.Name)
  }

	name := cmd.Args[0]

	err := s.cfg.SetUser(name)

	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("User: %s has been set", name)

	return nil 

}
