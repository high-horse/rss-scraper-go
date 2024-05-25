package main

import (
	"rss-scraper/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey 	string `json:"api_key"`
}

func dbuserToUser (dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name      string	`json:"name"`
	Url       string	`json:"url"`
	UserID    uuid.UUID	`json:"user_id"`
}

func dbfeedToFeed (dbFeed database.Feed) Feed {
	return Feed{
		ID  :      dbFeed.ID ,
		CreatedAt : dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt	,
		Name    :  dbFeed.Name	,
		Url       :dbFeed.Url	,
		UserID    :dbFeed.UserID	,
	}
}

func dbfeedsToFeeds (dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbfeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	UserID    uuid.UUID	`json:"user_id"`
	FeedID    uuid.UUID	`json:"feed_id"`
}

func dbFeedFollowToFeedFollow (dbFeedFolllow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFolllow.FeedID,
		CreatedAt: dbFeedFolllow.CreatedAt,
		UpdatedAt: dbFeedFolllow.UpdatedAt,
		UserID: dbFeedFolllow.UserID,
		FeedID: dbFeedFolllow.FeedID,
	}
}

func dbFeedFollowsToFeedFollows (dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeed := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(dbFeed))
	}
	return feedFollows
}