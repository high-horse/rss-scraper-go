package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss-scraper/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateuser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	respondeWithJSON(w, 201, dbuserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondeWithJSON(w, 200, dbuserToUser(user))
}


func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("Couldnot get posts %v", err))
		return
	}

	
	respondeWithJSON(w, 200, dbPostsToPosts(posts))
}
