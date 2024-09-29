package main

import (
	"context"
	"fmt"
)

func handleReset(state *Florida, _ command) error {
	error := state.db.DeleteAllUsers(context.Background())
	if error != nil {
		return error
	}
	fmt.Println("deleted all users")
	return nil
}
