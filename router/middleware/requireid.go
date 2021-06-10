package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		c.Set("X-Request-Id", requestId)

		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
