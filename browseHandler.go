package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bmamha/gator/internal/database"
)

func browseHandler(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("format: %s <limit(optional)>", cmd.Name)
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to fetch user: %w", err)
	}

	var limit int32
	if len(cmd.Args) == 1 {
		value, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("wrong values used for limit. it should be in integer form")
		}
		limit = int32(value)
	} else {
		limit = 2
	}

	postsForUserParams := database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), postsForUserParams)
	if err != nil {
		return fmt.Errorf("unable to fetch posts: %w", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}

	return nil
}
