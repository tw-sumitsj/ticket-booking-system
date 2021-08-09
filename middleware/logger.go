package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func LoggingMiddleWare(c *gin.Context) {
	t := time.Now()
	reqDetails := map[string]string{
		"host":   c.Request.Host,
		"method": c.Request.Method,
		"uri":    c.Request.RequestURI,
	}
	c.Next()
	log.Info().Dur("Duration", time.Since(t)).Interface("reqDetails", reqDetails).Send()
}
