package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, statusCode int, data interface{}, isMessage bool, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	switch isMessage {
	case true:
		resp := map[string]string{"message": message}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("could not encode response: %v", err)
			return
		}
		return
	default:
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("could not encode response: %v", err)
			return
		}
		return
	}
}
