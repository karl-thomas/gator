package main

import (
	"context"
	"fmt"
)

func handleUsers(state *Florida, _ command) error {
	users, error := state.db.GetUsers(context.Background())
	if error != nil {
		return error
	}
	currentUser := state.Laws.Username
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
