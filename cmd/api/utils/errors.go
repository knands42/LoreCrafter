package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
}

func WriteJSONValidationError(w http.ResponseWriter, details any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   "validation failed",
		Details: details,
	})
	if err != nil {
		log.Printf("failed to write validation error: %v", err)
		return err
	}

	return nil
}

func WriteJSONError(w http.ResponseWriter, statusCode int, err string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodingErr := json.NewEncoder(w).Encode(ErrorResponse{
		Error: err,
	})
	if encodingErr != nil {
		log.Printf("failed to write error: %v", err)
		return encodingErr
	}

	return nil
}
