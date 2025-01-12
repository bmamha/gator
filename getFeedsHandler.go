package main

import (
	"context"
	"fmt"
)

func getFeedsHandler(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("no arguments needed for %s command", cmd.Name)
	}

	feedsSlice, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch feeds from database: %w", err)
	}

	for _, feed := range feedsSlice {
		user, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			fmt.Println("Could not get user name")
		}
		fmt.Println("---")
		fmt.Println(feed.Name.String)
		fmt.Println(feed.Url.String)
		fmt.Println(user.Name.String)
		fmt.Println("---")

	}

	return nil
}
