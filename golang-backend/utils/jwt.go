package utils

import (
	"food-recipe-backend/config"
	"github.com/dgrijalva/jwt-go"
	"strconv" // Import the strconv package for string conversion
	"time"
)

// GenerateJWT generates a JWT token for the user
func GenerateJWT(userID int, name, email, role string) (string, error) {
	config.LoadConfig()
	var jwtSecret = []byte(config.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    userID,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []interface{}{"user", "admin"}, // Allowed roles for Hasura
			"x-hasura-default-role":  role,                           // Default role for the user
			"x-hasura-user-id":       strconv.Itoa(userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}
