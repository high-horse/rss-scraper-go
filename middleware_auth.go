package main

import (
	"fmt"
	"net/http"
	"rss-scraper/auth"
	"rss-scraper/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondeWithError(w, 403, fmt.Sprintf("auth error : %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), api_key)
		if err != nil {
			respondeWithError(w, 400, fmt.Sprintf("could not get user : %v", err))
			return
		}
		// fmt.Println("userdata %s", user.ApiKey)
		// panic("")

		handler(w, r, user)
	}
}
