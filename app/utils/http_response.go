package utils

import (
	"encoding/json"
	"net/http"

	"product_search_and_vectorize_service/app/Structures/Response"
)

// SendJSON sends a JSON response
func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// SendError sends an error JSON response
func SendError(w http.ResponseWriter, status int, message string) {
	SendJSON(w, status, Response.ErrorResponse{
		Success: false,
		Error:   message,
	})
}

// SendSuccess sends a success JSON response with data
func SendSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	SendJSON(w, status, map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	})
}
