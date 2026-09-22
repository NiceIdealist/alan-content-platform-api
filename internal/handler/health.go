package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health reports basic liveness of the service.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
