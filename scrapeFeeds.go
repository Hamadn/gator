package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Hamadn/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		var pubDate sql.NullTime
		for _, layout := range []string{time.RFC1123Z, time.RFC1123} {
			if t, err := time.Parse(layout, item.PubDate); err == nil {
				pubDate = sql.NullTime{Time: t, Valid: true}
				break
			}
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			fmt.Printf("error creating post %q: %v\n", item.Title, err)
			continue
		}
	}
	return nil
}
