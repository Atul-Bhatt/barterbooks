package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.InfoLevel)

		// generate a correlation ID
		correlationID := uuid.NewString()
		c.Set("correlationID", correlationID)

		log.WithFields(log.Fields{
			"correlationID": correlationID,
			"Method":        c.Request.Method,
			"Path":          c.Request.URL.Path,
			"Status":        c.Writer.Status,
		}).Info()
	}
}
