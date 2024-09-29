package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karl-thomas/gator/internal/database"
)

func handleAddFeed(state *Florida, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("need to provide a feed name and url")
	}
	currentUser, error := state.db.GetUser(context.Background(), state.Laws.Username)
	if error != nil {
		return error
	}
	feed, error := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: currentUser.ID,
		ID:     uuid.New(),
	})
	if error != nil {
		return error
	}
	fmt.Printf("%+v\n", feed)

	return nil
}
