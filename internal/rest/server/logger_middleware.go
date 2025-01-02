package server

import (
	"cmp"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

const ginLogFieldsKey = "ginlogfields"

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := logx.NewFields().
			Add(logx.CONTROLLER, "Rest").
			Add(logx.HTTP_METHOD, c.Request.Method).
			Add(logx.HTTP_PATH, c.Request.URL.Path).
			Add(logx.CORRELATION_ID, getCorrelationId(c))

		c.Set(ginLogFieldsKey, fields)

		c.Next()

		if c.Request.Method != "OPTIONS" {
			fields.Add(logx.HTTP_STATUS, strconv.Itoa(c.Writer.Status()))
			logx.Info(GinLogCtx(c), "Completed")
		}
	}
}

func GinLogCtx(c *gin.Context) context.Context {
	value := c.MustGet(ginLogFieldsKey).(logx.Fields)
	return logx.CtxWithFields(c, value)
}

func getCorrelationId(c *gin.Context) string {
	return cmp.Or(c.GetHeader(logx.CORRELATION_ID), uuid.New().String())
}
