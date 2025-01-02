package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	handleErrorIfNeeded(c)
}

func AbortWithError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func handleErrorIfNeeded(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

	switch value := err.Err.(type) {
	case *errs.Custom:
		handleCustom(c, value)
	case *errs.Validation:
		c.JSON(400, err)
	case validator.ValidationErrors:
		handleValidationError(c, value)
	case error:
		c.JSON(500, gin.H{
			"error": value.Error(),
		})
	default:
		handleCustom(c, errs.Internal)
	}
}

func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	err := errs.NewValidation()

	for _, e := range validationErrors {
		err.Add(strings.ToLower(e.Field()), e.Tag())
	}

	c.JSON(400, err)
}

func handleCustom(c *gin.Context, err *errs.Custom) {
	c.JSON(err.Status(), err)
}
