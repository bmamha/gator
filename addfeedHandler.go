package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bmamha/gator/internal/database"
	"github.com/google/uuid"
)

func addfeedHandler(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("%s command expects two arguments", cmd.Name)
	}
	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(),
		sql.NullString{
			String: userName,
			Valid:  true,
		})
	if err != nil {
		return fmt.Errorf("unable to fetch current user: %w", err)
	}

	name := sql.NullString{
		String: cmd.Args[0],
		Valid:  true,
	}

	url := sql.NullString{
		String: cmd.Args[1],
		Valid:  true,
	}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("unable to create feed: %w", err)
	}

	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)

	return nil
}
