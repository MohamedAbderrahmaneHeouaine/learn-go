package main

import (
		"net/http"
		"encoding/json"
		"fmt"
		"time"
		"github.com/MohamedAbderrahmaneHeouain/learn-go/internal/database"
		"github.com/google/uuid"
		)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`	
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url:params.URL,
		UserID: user.ID,
	})
	if err!= nil {
        respondWithError(w, 400, err.Error())
    }	
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
        respondWithError(w, 403, err.Error())
        return
    }	
	
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}