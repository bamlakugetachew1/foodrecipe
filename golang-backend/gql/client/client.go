package client

import (
	"bytes"
	"encoding/json"
	"food-recipe-backend/config"
	"net/http"
)

// SendRequest sends a request to the Hasura GraphQL endpoint
func SendRequest(reqBody map[string]interface{}) (map[string]interface{}, error) {
	config.LoadConfig()
	jsonReq, _ := json.Marshal(reqBody)
	httpReq, err := http.NewRequest("POST", config.HasuraEndpoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-hasura-admin-secret", config.HasuraAdminSecret) // Replace with your admin secret

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)
	return responseBody, nil
}
