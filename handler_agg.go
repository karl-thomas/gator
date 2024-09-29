package main

import (
	"context"
	"fmt"

	"github.com/karl-thomas/gator/rss"
)

func handleAgg(state *Florida, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
