package handlers

import (
	"encoding/json"
	"net/http"
)

func ValidateRepositoryId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{
		"repository_id_is_valid": true,
	})
}
