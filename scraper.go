package main

import (
	"context"
	"log"
	"rss-scraper/internal/database"
	"sync"
	"time"
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
		log.Println("Post found ", item.Title, " on feed ", feed.Name)
	}
	log.Printf("feed %s collecte, %v post found", feed.Name, len(rssFeed.Channel.Item))
}