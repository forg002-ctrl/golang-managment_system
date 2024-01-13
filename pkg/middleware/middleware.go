package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next();
		endTime := time.Now();
		
		requestDuration := endTime.Sub(startTime)
		fmt.Printf("[%s] %s %s %s %v\n",
			endTime.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			requestDuration,
		)
	}
} 