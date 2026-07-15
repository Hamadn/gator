package main

import (
	"context"
	"fmt"
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
		fmt.Println(item.Title)
	}
	return nil
}
