package main

import (
	"log"
	"net/http"
	"os"

	"github.com/argo-agorshechnikov/golang-restApi/internal/handlers"
	"github.com/argo-agorshechnikov/golang-restApi/internal/repository"
	"github.com/argo-agorshechnikov/golang-restApi/internal/service"
)

func main() {

	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR env var is req")
	}

	userRepo, err := repository.NewUserRep(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.CreateUserHand)
	mux.HandleFunc("/user", userHandler.GetUserByIdHand)

	addr := ":8080"
	log.Printf("Server running")
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
