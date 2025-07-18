package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ErrorCode *string     `json:"errorCode"`
}

// Réponse succès
func RespondJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := APIResponse{
		Status:    "success",
		Message:   message,
		Data:      data,
		ErrorCode: nil,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Failed to encode JSON response:", err)
	}
}

// Réponse erreur
func RespondError(w http.ResponseWriter, code int, message string, errorCode error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var codeStr *string
	if errorCode != nil {
		s := errorCode.Error()
		codeStr = &s
	}

	response := APIResponse{
		Status:    "error",
		Message:   message,
		Data:      nil,
		ErrorCode: codeStr,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Failed to encode JSON response:", err)
	}
}
