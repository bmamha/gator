package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bmamha/gator/internal/database"
	"github.com/google/uuid"
)

func addfeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("%s command expects only two arguments", cmd.Name)
	}
	name := cmd.Args[0]

	url := cmd.Args[1]
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
	feedFollowParams := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollows(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("unable to create follow for feed: %w", err)
	}
	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)
	fmt.Printf("%s is automatically followed by %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}
