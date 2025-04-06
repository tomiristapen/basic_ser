package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		log.Println("[LOGGING] -------------------------")
		log.Printf(" METHOD: %s", c.Request.Method)
		log.Printf(" PATH:   %s", c.Request.URL.Path)
		log.Printf(" IP:     %s", c.ClientIP())
		log.Printf(" STATUS: %d", c.Writer.Status())
		log.Printf(" TIME:   %s", duration)
		log.Println("-------------------------------------")
	}
}
