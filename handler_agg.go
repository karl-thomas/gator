package main

import (
	"context"
	"fmt"
	"time"

	"github.com/karl-thomas/gator/internal/database"
	"github.com/karl-thomas/gator/rss"
)

func handleAgg(state *Florida, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("need to provide a time between requests")
	}

	duration, error := time.ParseDuration(cmd.Args[0])
	if error != nil {
		return error
	}
	fmt.Printf("Collecting feeds every %s\n", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err := scrapeFeeds(state)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(state *Florida) error {
	feed, error := state.db.GetNextFeedToFetch(context.Background())
	if error != nil {
		return error
	}

	_, error = state.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
	})

	rssFeed, error := rss.FetchFeed(context.Background(), feed.Url)
	if error != nil {
		return error
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}

	return nil
}
