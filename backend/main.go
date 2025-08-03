package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/atbu/slate/backend/auth"
	"github.com/atbu/slate/backend/db"
	"github.com/atbu/slate/backend/handlers"
	"github.com/atbu/slate/backend/middleware"
	"github.com/atbu/slate/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, using environment variables")
	}

	requiredVars := []string{"DATABASE_URL", "JWT_SECRET"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Required environment variable %s is not set", v)
		}
	}
}

func main() {
	loadEnv()

	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := mux.NewRouter()

	userRepo := models.NewUserRepository(database)
	refreshTokenRepo := models.NewRefreshTokenRepository(database)

	authService := auth.NewAuthService(userRepo, refreshTokenRepo, os.Getenv("JWT_SECRET"), 15*time.Minute)

	authHandler := handlers.NewAuthHandler(authService)
	//userHandler := handlers.NewUserHandler(userRepo)

	r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/auth/refresh", authHandler.RefreshToken).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(authService))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
