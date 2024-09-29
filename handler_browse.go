package main

import (
	"context"
	"strconv"

	"github.com/karl-thomas/gator/internal/database"
)

func handleBrowse(state *Florida, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = i
	}

	posts, err := state.db.PostsForUser(context.Background(), database.PostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		println(post.Title)
	}

	return nil
}
