package main

import (
	"context"
	"fmt"
	"food-recipe-backend/config"
	"food-recipe-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load the configuration
	config.LoadConfig()

	// Connect to the PostgreSQL database
	conn := config.ConnectDB()
	defer conn.Close(context.Background())

	// Create a new router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterUserRoutes(router)
	routes.RegisterFilesRoutes(router)

	// Serve static files from the "uploads" directory
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	// Set up a simple HTTP handler for the root path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Food Recipe API!")
	})

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	// Start the server on the port defined in .env
	fmt.Printf("Server is listening on port %s...\n", config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
