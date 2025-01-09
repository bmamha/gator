package main

import (
	"context"
	"fmt"
)

const url = "https://www.wagslane.dev/index.xml"

func aggCommand(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("no arguments needed for %s command", cmd.Name)
	}

	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to fetch rss: %w", err)
	}

	fmt.Println(rss)

	return nil
}
