package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bmamha/gator/internal/database"
	"github.com/google/uuid"
)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
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
		return fmt.Errorf("unable to create a follow: %w", err)
	}

	fmt.Println(feedFollow.FeedName)
	fmt.Println(feedFollow.UserName)

	return nil
}
