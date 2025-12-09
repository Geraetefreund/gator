package main

import (
	"context"
	"fmt"

	"github.com/Geraetefreund/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	dbParams := database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), dbParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Printf("\n==============================\n")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:        %s\n", feed.ID)
	fmt.Printf("* Created:   %s\n", feed.CreatedAt)
	fmt.Printf("* Updated:   %s\n", feed.UpdatedAt)
	fmt.Printf("* Name:      %s\n", feed.Name)
	fmt.Printf("* URL:       %s\n", feed.Url)
	fmt.Printf("* UserID:    %s\n", feed.UserID)
}

func handlerGetFeeds(s *state, cmd command) error {
	// feeds is []Feed
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		userName, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get username from database: %w", err)
		}
		fmt.Printf("* Name: %v\n", feed.Name)
		fmt.Printf("* URL:  %v\n", feed.Url)
		fmt.Printf("* User: %v\n", userName)
		fmt.Printf("==========================\n")

	}
	return nil
}
