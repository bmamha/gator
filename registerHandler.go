package main

import (
	"fmt"
	"time"
	"github.com/bmamha/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"database/sql"
	"context"
)
func registerHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command expects a single argument- the username", cmd.Name)
  }

	name := cmd.Args[0]

 params := database.CreateUserParams{
	 ID: uuid.New(),
	 CreatedAt: time.Now(),
	 UpdatedAt: time.Now(),
	 Name: sql.NullString{
    String: cmd.Args[0],    
    Valid: true,         
},
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
    os.Exit(1)
	}


	err = s.cfg.SetUser(name)
	if err != nil {
		
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("User: %s has been registered\n", name)

	fmt.Println(user)

	return nil 

}
