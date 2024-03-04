package main
import (
		"net/http"
	    "encoding/json"
		"log"
		)

func respondWithError(w http.ResponseWriter, code int, msg string)  {
	if(code > 499){
		log.Println("Responding with 5XX error:",msg)
	}
	type errResponse struct {
		Error string `json:"error"`	
	}
	respondWithJSON(w, code,errResponse{Error:msg} )
}
func respondWithJSON(w http.ResponseWriter, code int, pyload interface {}){
	data, err := json.Marshal(pyload)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }	
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(code)
	w.Write(data)
}