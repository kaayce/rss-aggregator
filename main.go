package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("No port found in the environment")
	}

	//  main router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get(Healthz, handlerReadiness)
	v1Router.Get(Err, handlerError)

	router.Mount(V1, v1Router)

	//  Not found route
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		log.Printf("Route does not exist: %s %s", r.Host, r.URL.Path)
		w.Write([]byte("Route does not exist")) // cast string to byte slice
	})

	//  Method not allowed
	router.MethodNotAllowed(func(w http.ResponseWriter, r * http.Request) {
		w.WriteHeader(405)
		log.Printf("Method is not valid: %s", r.Method)
		w.Write([]byte("Method is not valid"))
	})

	serve := &http.Server{
		Handler: router,
		Addr: ":" + PORT,
	}

	log.Printf("Server statrting on: %v port", PORT)
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}