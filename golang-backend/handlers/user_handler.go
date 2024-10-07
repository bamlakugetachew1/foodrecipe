package handlers

import (
	"encoding/json"
	"fmt"
	"food-recipe-backend/config"
	"food-recipe-backend/gql/client"
	"food-recipe-backend/gql/hasura/mutations"
	"food-recipe-backend/gql/hasura/queries"
	"food-recipe-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"` // Incoming plain password
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
}

type RegisterRequest struct {
	Input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"input"`
}

type LoginRequest struct {
	Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"input"`
}

// SMTP configuration for sending emails
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// Email content structure
type EmailContent struct {
	From    string
	To      string
	Subject string
	Body    string
}

// RegisterHandler registers a new user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the incoming request body into the RegisterRequest struct
	var regReq RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid input"})
		return
	}

	// Extract name, email, and password from the input
	name := regReq.Input.Name
	email := regReq.Input.Email
	password := regReq.Input.Password

	fmt.Printf("Received: Name: %s, Email: %s, Password: %s\n", name, email, password)

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Server error: could not hash password"})
		return
	}

	// Create request body for Hasura GraphQL API
	reqBody := map[string]interface{}{
		"query": mutations.RegisterUserMutation, // Ensure this is your actual GraphQL mutation string
		"variables": map[string]interface{}{
			"email":    email,
			"name":     name,
			"password": string(hashedPassword),
		},
	}

	// Send request to Hasura
	resp, err := client.SendRequest(reqBody) // client.SendRequest should be a function to send HTTP requests
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "User registration failed: " + err.Error()})
		return
	}

	// Check if there are any errors in the GraphQL response
	if resp["errors"] != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": fmt.Sprintf("GraphQL error: %+v", resp["errors"])})
		return
	}

	// Respond to the client with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User registered successfully, please log in."})
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	// Decode the request body into the LoginRequest struct
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create request body for Hasura
	reqBody := map[string]interface{}{
		"query": queries.LoginUserQuery,
		"variables": map[string]interface{}{
			"email": loginReq.Input.Email,
		},
	}

	// Send request to Hasura
	resp, err := client.SendRequest(reqBody)
	if err != nil {
		http.Error(w, "Error contacting server", http.StatusInternalServerError)
		return
	}

	// Check for errors in the Hasura response
	if resp["errors"] != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Validate user and password from Hasura response
	data, ok := resp["data"].(map[string]interface{})
	if !ok || data["users"] == nil {
		http.Error(w, "Invalid response from server", http.StatusInternalServerError)
		return
	}

	users, ok := data["users"].([]interface{})
	if !ok || len(users) == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Extract user data
	user, ok := users[0].(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid user data", http.StatusInternalServerError)
		return
	}

	// Ensure the password field exists and validate it
	passwordHash, ok := user["password"].(string)
	if !ok || passwordHash == "" {
		http.Error(w, "Invalid user password", http.StatusUnauthorized)
		return
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginReq.Input.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Extract user details (id, name, email, role)
	userID, ok := user["id"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	userName, ok := user["name"].(string)
	if !ok || userName == "" {
		http.Error(w, "Invalid user name", http.StatusInternalServerError)
		return
	}

	userEmail, ok := user["email"].(string)
	if !ok || userEmail == "" {
		http.Error(w, "Invalid user email", http.StatusInternalServerError)
		return
	}

	userRole, ok := user["role"].(string)
	if !ok || userRole == "" {
		http.Error(w, "Invalid user role", http.StatusInternalServerError)
		return
	}

	// Generate a JWT token for the user
	token, err := utils.GenerateJWT(int(userID), userName, userEmail, userRole)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := map[string]interface{}{
		"id":    int(userID),
		"token": token,
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Function to configure SMTP settings
func configureSMTP() *SMTPConfig {
	config.LoadConfig()
	return &SMTPConfig{
		Host:     config.EMAIL_HOST,     // Replace with your SMTP host
		Port:     config.EMAIL_PORT,     // Replace with your SMTP port
		Username: config.FROM,           // Your email
		Password: config.EMAIL_PASSWORD, // Your email password
	}
}

// Function to create static email content when a new user is added
func createUserEmailContent() *EmailContent {
	return &EmailContent{
		From:    config.FROM, // Your email address
		To:      config.FROM, // Admin or recipient email address
		Subject: "New User Added!",
		Body:    "A new user has been added to the platform.",
	}
}

// Function to send email
func sendEmail(smtpConfig *SMTPConfig, emailContent *EmailContent) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailContent.From)
	m.SetHeader("To", emailContent.To)
	m.SetHeader("Subject", emailContent.Subject)
	m.SetBody("text/plain", emailContent.Body)

	// Create a new dialer with the configured SMTP settings
	d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Username, smtpConfig.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Println("Email sent successfully!")
	return nil
}

// Webhook handler that triggers email notification
func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	// No need to parse event data, we're sending a static email
	emailContent := createUserEmailContent()
	smtpConfig := configureSMTP()

	// Send the email
	err := sendEmail(smtpConfig, emailContent)
	if err != nil {
		log.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Respond to Hasura with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email notification sent"))
}
