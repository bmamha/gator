package main

import (
	"fmt"
	"os"
	"context"
)

func resetHandler(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("No arguments needed for %s command", cmd.Name)	
	}
  
	err := s.db.Reset(context.Background())

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("All users have been deleted")

	return nil 
	
}
