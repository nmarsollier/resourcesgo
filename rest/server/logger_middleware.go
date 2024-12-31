package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := logx.NewFields().
			Add(logx.CONTROLLER, "Rest").
			Add(logx.HTTP_METHOD, c.Request.Method).
			Add(logx.HTTP_PATH, c.Request.URL.Path).
			Add(logx.CORRELATION_ID, getCorrelationId(c))

		c.Set("logfields", fields)

		c.Next()

		if c.Request.Method != "OPTIONS" {
			fields.Add(logx.HTTP_STATUS, strconv.Itoa(c.Writer.Status()))
			logx.Info(fields, "Completed")
		}
	}
}

func GinLogFields(c *gin.Context) logx.Fields {
	value, exists := c.Get("logfields")

	if !exists {
		return logx.NewFields()
	}

	return value.(logx.Fields)
}

func getCorrelationId(c *gin.Context) string {
	value := c.GetHeader(logx.CORRELATION_ID)

	if len(value) == 0 {
		value = uuid.New().String()
	}

	return value
}
