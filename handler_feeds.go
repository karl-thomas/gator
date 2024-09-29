package main

import (
	"context"
	"fmt"
)

func handleFeeds(state *Florida, _ command) error {
	feeds, error := state.db.GetFeedsWithUser(context.Background())
	if error != nil {
		return error
	}
	for _, feed := range feeds {
		fmt.Printf("* %+v\n", feed)
	}
	return nil
}
