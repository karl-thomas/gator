package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karl-thomas/gator/internal/config"
	"github.com/karl-thomas/gator/internal/database"
)

func handleRegister(state *Florida, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("need to provide a username")
	}
	user, error := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: cmd.Args[0],
	})
	if error != nil {
		return error
	}
	error = config.SetUser(user.Name)
	if error != nil {
		return error
	}

	fmt.Printf("%+v\n", user)
	return nil
}
