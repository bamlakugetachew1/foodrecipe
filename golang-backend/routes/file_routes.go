package routes

import (
	"food-recipe-backend/handlers"
	"food-recipe-backend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterFilesRoutes(router *mux.Router) {
	recipesRouter := router.PathPrefix("/files").Subrouter()
	recipesRouter.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		middleware.VerifyJWT(http.HandlerFunc(handlers.UploadFilesHandler)).ServeHTTP(w, r)
	}).Methods("POST", "OPTIONS")

}
