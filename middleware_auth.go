package main

import (
		"github.com/MohamedAbderrahmaneHeouain/learn-go/internal/database"
		"github.com/MohamedAbderrahmaneHeouain/learn-go/internal/auth"
		"fmt"
		"net/http"
		)
type authedHandler func(http.ResponseWriter, *http.Request, database.User)
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err!= nil {
            respondWithError(w, 403, fmt.Sprintf("Auth error: %v",err))
            return
        }
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)	
		if err!= nil {
            respondWithError(w, 403, fmt.Sprintf("couldn't get user: %v",err))
            return
        }
		handler(w, r, user)
	}
}