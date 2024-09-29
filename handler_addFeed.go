package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karl-thomas/gator/internal/database"
)

func handleAddFeed(state *Florida, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("need to provide a feed name and url")
	}

	feed, error := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
		ID:     uuid.New(),
	})
	if error != nil {
		return error
	}
	fmt.Printf("%+v\n", feed)

	_, error = state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
		ID:     uuid.New(),
	})
	if error != nil {
		return error
	}

	return nil
}
