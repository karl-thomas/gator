package main

import (
	"context"
	"fmt"

	"github.com/karl-thomas/gator/internal/database"
)

func handleUnfollow(state *Florida, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("unfollow requires a feed url")
	}

	url := cmd.Args[0]
	feed, err := state.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = state.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Unfollowed %s\n", url)
	return nil
}
