package main

import (
	"context"
	"fmt"

	"github.com/karl-thomas/gator/internal/database"
)

func handleFollowing(state *Florida, cmd command, user database.User) error {
	follows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Follows for %s\n", state.Laws.Username)
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
