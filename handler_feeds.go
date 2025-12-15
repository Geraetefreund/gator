package main

import (
	"context"
	"fmt"

	"github.com/Geraetefreund/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed name> <url>", cmd.Name)
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

	dbParamsFollow := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), dbParamsFollow)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
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

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed not found: %w", err)
	}

	dbParams := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), dbParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow: %w", err)
	}
	fmt.Printf("Successfully subscirbed to feed:\n")
	fmt.Printf("* Name:      %v\n", feed_follow.FeedName)
	fmt.Printf("* User Name: %v\n", feed_follow.UserName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get follows for user: %w", err)
	}

	fmt.Printf("User: %v\n", user.Name)
	fmt.Printf("Following:\n")

	for _, feed := range following {
		fmt.Printf("* %v\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed not found: %w", err)
	}

	dbParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), dbParams)
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	return nil
}
