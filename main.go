package main

import (
	"fmt"
	"insight-api/config"
	"insight-api/handlers"
	"insight-api/middleware/loggers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	env := os.Getenv("ENV")
	if err := config.LoadEnvVars(env); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("ENV=%s", env)
	}

	router := mux.NewRouter()
	router.HandleFunc("/initialize_repository", loggers.HTTPLog(handlers.InitializeRepository)).Methods("POST")
	router.HandleFunc("/query_repository", loggers.HTTPLog(handlers.QueryRepository)).Methods("GET")
	router.HandleFunc("/reinitialize_repository", loggers.HTTPLog(handlers.ReinitializeRepository)).Methods("PUT")
	router.HandleFunc("/uninitialize_repository", loggers.HTTPLog(handlers.UninitializeRepository)).Methods("DELETE")
	router.HandleFunc("/validate_repository_id", loggers.HTTPLog(handlers.ValidateRepositoryId)).Methods("POST")

	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	log.Printf("Server running on %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
