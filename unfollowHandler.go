package main

import (
	"context"
	"fmt"

	"github.com/bmamha/gator/internal/database"
)

func unfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("unable to obtain feed from url: %w", err)
	}
	deleteParams := database.DeleteFollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFollowFeed(context.Background(), deleteParams)
	if err != nil {
		return fmt.Errorf("unable to delete follow of feed for user: %w", err)
	}
	fmt.Printf("%s has unfollowed %s feed\n", user.Name, feed.Name)

	return nil
}
