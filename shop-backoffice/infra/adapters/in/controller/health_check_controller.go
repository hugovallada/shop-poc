package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func (hc HealthCheckController) HealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
