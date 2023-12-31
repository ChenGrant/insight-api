package handlers

import (
	"encoding/json"
	"net/http"
)

func InitializeRepository(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]string{
		"repository_id": "id1",
	})
}
