package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss-scraper/internal/database"
	"time"
	"rss-scraper/internal/auth"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateuser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string 	`json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondeWithError(w, 500, fmt.Sprintf("database error: %v", err))
		return
	}
	respondeWithJSON(w, 201, dbuserToUser(user))
}


func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondeWithError(w, 403, fmt.Sprintf("auth error : %v", err))
	}

	user , err := apiCfg.DB.GetUserByAPIKey(r.Context(), api_key)
	if err != nil {
		respondeWithError(w, 400, fmt.Sprintf("could not get user : %v", err))
		return
	}

	respondeWithJSON(w, 200, dbuserToUser(user))
}