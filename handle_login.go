package main

import (
	"context"
	"fmt"

	"github.com/karl-thomas/gator/internal/config"
)

func handleLogin(state *Florida, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("need to provide a username")
	}
	user, error := state.db.GetUser(context.Background(), cmd.Args[0])
	if error != nil {
		return fmt.Errorf("user not found with name %s", cmd.Args[0])
	}
	error = config.SetUser(user.Name)
	if error != nil {
		return error
	}
	fmt.Println("logged in as", cmd.Args[0])
	return nil
}
