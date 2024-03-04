package main

import (
		"fmt"
		"os"
		"log"	
		"github.com/joho/godotenv"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		"net/http"
	)
func main(){
	fmt.Println("hello")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found the env")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
		MaxAge:           300,	
	}))
	v1 := chi.NewRouter()
	v1.Get("/ready", handleReadiness)
	v1.Get("/err", handlerErr)
	router.Mount("/v1",v1)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
        
	}
	log.Printf("serving on port %v",portString)
	err :=srv.ListenAndServe()
	if err!= nil {
        log.Fatal(err)
    }	
	fmt.Println(portString)
}