package middleware

import (
	"context"
	"food-recipe-backend/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

type contextKey string

const userIDKey contextKey = "userID"

// VerifyJWT verifies the token, extracts the user ID, and calls the next handler if valid
func VerifyJWT(next http.Handler) http.Handler {
	config.LoadConfig()
	var jwtSecret = []byte(config.JWT_SECRET_KEY)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Extract the token string (without the "Bearer " prefix)
		tokenString = tokenString[len("Bearer "):]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return jwtSecret, nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			// Extract the user ID from the token claims
			userID, ok := (*claims)["id"].(float64) // Assuming `id` is stored as a float64
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Store the user ID in the request context
			ctx := context.WithValue(r.Context(), userIDKey, int(userID))

			// Pass the request to the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}

// GetUserIDFromContext extracts the user ID from the context
func GetUserIDFromContext(r *http.Request) (int, bool) {
	userID, ok := r.Context().Value(userIDKey).(int)
	return userID, ok
}
