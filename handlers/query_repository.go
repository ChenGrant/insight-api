package handlers

import (
	"encoding/json"
	"net/http"
)

type FileMatch struct {
	Path      string `json:"path"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Content   string `json:"content"`
}

func QueryRepository(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode([]FileMatch{
		{
			Path:      "/server/src/config/database.js",
			StartLine: 3,
			EndLine:   15,
			Content:   "const connectToDatabase = async () => {...};",
		},
		{
			Path:      "/server/src/server.js",
			StartLine: 25,
			EndLine:   25,
			Content:   "await connectToDatabase(app);",
		},
	})
}
