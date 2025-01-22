package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bmamha/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	feedFetchedParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), feedFetchedParams)
	if err != nil {
		return fmt.Errorf("unable to mark feed: %w", err)
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
		pubDate, err := parseTime(item.PubDate)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
		}
		postParams := database.CreatePostsParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: pubDate,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			FeedID: feed.ID,
		}

		post, err := s.db.CreatePosts(context.Background(), postParams)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				log.Printf("Error saving item: %v", err)
			}
		}
		fmt.Printf("%s with %s ID has been saved to database\n", post.Title, post.ID)
	}

	return nil
}

func parseTime(dateStr string) (time.Time, error) {
	layouts := []string{
		time.Layout,
		"Mon, 02 Jan 2006 15:04:05 -0700",
		time.RFC822Z,
		time.RFC1123,
		time.RFC822,
		time.RFC1123Z,
	}
	var lastErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, dateStr)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, fmt.Errorf("could not parse time %s with any layout: %w", dateStr, lastErr)
}
