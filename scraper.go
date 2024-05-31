package main

import (
	"context"
	"database/sql"
	"log"
	"rss-scraper/internal/database"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries, 
	concurrency int, 
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(), 
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds :", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched :", err )
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed :", err )
		return
	}

	for _, item := range rssFeed.Channel.Item {
		// log.Println("Post found ", item.Title, " on feed ", feed.Name)
		desc :=  sql.NullString{}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldnot parse date %v with err %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(), 
			database.CreatePostParams{
				ID :uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Title: item.Title,
				Description: desc,
				PublishedAt: pubAt,
				Url: item.Link,
				FeedID: feed.ID,
			})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post :", err)
		}
	}
	log.Printf("feed %s collected, %v post found", feed.Name, len(rssFeed.Channel.Item))
}
