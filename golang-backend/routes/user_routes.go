package routes

import (
	"food-recipe-backend/handlers"
	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers the user-related routes
func RegisterUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter() // Create a subrouter for /user
	userRouter.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	userRouter.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	userRouter.HandleFunc("/sendemail", handlers.NotifyHandler).Methods("POST")

}
