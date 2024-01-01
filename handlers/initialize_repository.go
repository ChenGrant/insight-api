package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRepository(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "id1",
	})
}
