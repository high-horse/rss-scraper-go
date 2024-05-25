package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss-scraper/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string 	`json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url :      params.URL,
		UserID :   user.ID,
		})
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	respondeWithJSON(w, 201, dbfeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	
	respondeWithJSON(w, 201, dbfeedsToFeeds(feeds))
}
