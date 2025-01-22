package main

import (
	"fmt"
	"time"
)

const url = "https://www.wagslane.dev/index.xml"

func aggHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("format: %s <url>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("unable to parse time specified: %w", err)
	}
	fmt.Printf("Collecting feeds every %s", cmd.Args[0])

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
