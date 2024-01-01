package main

import (
	"fmt"
	"insight-api/config"
	"insight-api/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	if err := config.LoadEnvVars(env); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("ENV=%s", env)
	}

	router := gin.Default()

	router.POST("/initialize_repository", handlers.InitializeRepository)
	router.GET("/query_repository", handlers.QueryRepository)
	router.PUT("/reinitialize_repository", handlers.ReinitializeRepository)
	router.DELETE("/uninitialize_repository", handlers.UninitializeRepository)
	router.POST("/validate_repository_id", handlers.ValidateRepositoryId)

	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	log.Printf("Server running on %s", address)
	log.Fatal(router.Run(address))
}
