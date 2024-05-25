package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss-scraper/internal/database"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	respondeWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	respondeWithJSON(w, 200, dbFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "feedFollowID")
	fmt.Printf("feed follow %s \n", feedFollowIDstr)

	feedFollowID, err := uuid.Parse(feedFollowIDstr)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("coulnot parse feed follow id : %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("coulnot delete feed follow : %v", err))
		return
	}
	respondeWithJSON(w, 200, struct{}{})
}
