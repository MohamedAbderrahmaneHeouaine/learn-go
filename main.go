package main

import (
		"fmt"
		"os"
		"log"	
		"github.com/joho/godotenv"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
		"net/http"
		_"github.com/lib/pq"
		"github.com/MohamedAbderrahmaneHeouain/learn-go/internal/database"
		"database/sql"
		
	)
	type apiConfig struct {
		DB *database.Queries
	}
func main(){
	fmt.Println("hello")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found the env")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("db is not connect")
	}
	connection, err := sql.Open("postgres",dbURL)
	if err!= nil {
        log.Fatal(err)
    }
	dbQueries := database.New(connection)
			
	apiCfg := apiConfig{
		DB:dbQueries,
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
	v1.Post("/users", apiCfg.handlerCreateUser)		
	v1.Get("/users",apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1.Post("/feeds",apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1.Get("/feeds",apiCfg.handlerGetFeeds)
	router.Mount("/v1",v1)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
        
	}
	log.Printf("serving on port %v",portString)
	
	fmt.Println(portString)
	log.Fatal(srv.ListenAndServe())
}