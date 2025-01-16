package main

import (
	"context"
	"fmt"

	"github.com/bmamha/gator/internal/database"
)

func followingHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s ", cmd.Name)
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to find feeds for user: %w", err)
	}
	if len(feedFollows) == 0 {
		fmt.Printf("No feed currently followed by %s\n", user.Name)
	} else {
		for _, item := range feedFollows {
			fmt.Println(item.FeedName)
		}
	}
	return nil
}
