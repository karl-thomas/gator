package main

import (
	"context"

	"github.com/karl-thomas/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *Florida, cmd command, user database.User) error) func(*Florida, command) error {
	return func(state *Florida, cmd command) error {
		user, err := state.db.GetUser(context.Background(), state.Laws.Username)
		if err != nil {
			return err
		}
		return handler(state, cmd, user)
	}
}
