package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rss-scraper/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	feed, err := urlToFeed("https://wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
	
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variable :", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found")
		os.Exit(1)
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database url not found")
		os.Exit(1)
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cound not connect to the database")
		os.Exit(1)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReady)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/user", apiCfg.handlerCreateuser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed", apiCfg.handlerGetFeed)

	v1Router.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))


	


	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	println("server starting on port :", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("coundnot start server :", err)
	}
}
