package main

import (
		"net/http"
		"encoding/json"
		"fmt"
		"time"
		"github.com/MohamedAbderrahmaneHeouain/learn-go/internal/database"
		"internal/auth"
		"github.com/google/uuid"
		)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"` 
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!= nil {
        respondWithError(w, 400, err.Error())
    }	
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetAPIKey(r.Header)
	if err!= nil {
        respondWithError(w, 403, err.Error())
        return
    }	
	user,err := apiCfg.DB.GetUserByAPIKey(r.Content(), apiKey)
	if err!= nil {
        respondWithError(w, 400, err.Error())
        return
    }		
	respondWithJSON(w, 200, databaseUserToUser(user))
}