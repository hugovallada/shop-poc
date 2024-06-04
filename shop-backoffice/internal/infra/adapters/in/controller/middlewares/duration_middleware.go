package middlewares

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func CallDuration(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	elapsedTime := time.Since(startTime)
	slog.Info(fmt.Sprintf("Call executed in %s", elapsedTime))
}
