package handlers

import (
	"encoding/json"
	"fmt" // Import the fmt package
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// UploadFilesHandler handles multiple file uploads and CORS preflight requests
func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle POST request for file upload
	if r.Method == http.MethodPost {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		// Parse the multipart form, with a maximum memory size
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit for form data
		if err != nil {
			log.Printf("Error parsing form: %v", err)
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the uploaded files
		files := r.MultipartForm.File["files"]
		if files == nil {
			log.Println("No files were uploaded")
			http.Error(w, "No files uploaded", http.StatusBadRequest)
			return
		}

		uploadDir := "uploads/"
		fileNames := []string{}

		// Create upload directory if it doesn't exist
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.MkdirAll(uploadDir, os.ModePerm)
		}

		// Loop through each file and save it
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				log.Printf("Error opening file: %v", err)
				http.Error(w, "Unable to upload file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Create a unique filename with a timestamp
			fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(fileHeader.Filename))
			filePath := filepath.Join(uploadDir, fileName)

			// Create the destination file
			destFile, err := os.Create(filePath)
			if err != nil {
				log.Printf("Error creating file: %v", err)
				http.Error(w, "Unable to save file", http.StatusInternalServerError)
				return
			}
			defer destFile.Close()

			// Copy the uploaded file's data to the destination file
			if _, err := io.Copy(destFile, file); err != nil {
				log.Printf("Error saving file: %v", err)
				http.Error(w, "Unable to save file", http.StatusInternalServerError)
				return
			}

			// Add filename to the response list
			fileNames = append(fileNames, fileName)
		}

		// Set the response header and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Create the response map and encode it to JSON
		response := map[string]interface{}{
			"message":   "Files uploaded successfully",
			"fileNames": fileNames,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		}
	}
}
