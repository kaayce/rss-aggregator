package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaayce/rss-aggregator/internal/auth"
	"github.com/kaayce/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authApiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication key error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), authApiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

// Pass user to context
func (apiCfg *apiConfig) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authApiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication key error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), authApiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't retrieve user: %v", err))
			return
		}

		// Pass user to context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
