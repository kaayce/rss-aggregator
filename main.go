package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/kaayce/rss-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("No port found in the environment")
	}

	//  connect to db
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("No DB_URL found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
	defer conn.Close()

	// create queries object from our database package
	queries := database.New(conn)

	// init apiConfig struct
	config := apiConfig{DB: queries}

	startScraping(queries, 10, time.Minute)

	//  main router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get(Healthz, handlerReadiness)
	v1Router.Get(Err, handlerError)

	v1Router.Post(Users, config.handlerCreateUser)
	v1Router.Get(Users, config.middleware(config.handlerGetUser))

	v1Router.Post(Feeds, config.middleware(config.handlerCreateFeed))
	v1Router.Get(Feeds, config.handlerGetFeeds)

	v1Router.Post(FeedFollows, config.middleware(config.handlerCreateFeedFollow))
	v1Router.Get(FeedFollows, config.middleware(config.handlerGetFeedFollows))

	v1Router.Delete(SingleFeedFollow, config.middleware(config.handlerDeleteFeedFollow))

	v1Router.Get(UserFollowedPosts, config.middleware(config.handlerGetPostsForUser))

	router.Mount(V1, v1Router)

	//  Not found route
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		log.Printf("Route does not exist: %s %s", r.Host, r.URL.Path)
		w.Write([]byte("Route does not exist")) // cast string to byte slice
	})

	//  Method not allowed
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		log.Printf("Method is not valid: %s", r.Method)
		w.Write([]byte("Method is not valid"))
	})

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server starting on port: %v", PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
