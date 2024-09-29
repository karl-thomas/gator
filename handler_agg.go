package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		publishedAt, error := time.Parse(time.RFC1123Z, item.PubDate)
		if error != nil {
			return error
		}
		post, error := state.db.AddPost(context.Background(), database.AddPostParams{
			FeedID:      feed.ID,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			ID:          uuid.New(),
		})

		if error.Error() != "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
			fmt.Printf("Error adding post: %s\n", error)
		}

		if error == nil {
			fmt.Printf("Added post %s\n", post.Title)
		}
	}

	return nil
}
