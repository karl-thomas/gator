package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karl-thomas/gator/internal/database"
)

func handleFollow(state *Florida, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("need to provide a url")
	}
	feed, err := state.db.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}
	feedFollow, err := state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
		ID:     uuid.New(),
	})

	fmt.Printf("%+v\n", feedFollow)

	return nil
}
