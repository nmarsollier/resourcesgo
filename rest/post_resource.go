package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/engine"
)

/**
 * @api {post} /v1/user Registrar Usuario
 * @apiName Registrar Usuario
 * @apiGroup Seguridad
 *
 * @apiDescription Registra un nuevo usuario en el sistema.
 *
 * @apiExample {json} Body
 *    {
 *      "name": "{Nombre de Usuario}",
 *      "login": "{Login de usuario}",
 *      "password": "{Contraseña}"
 *    }
 *
 * @apiUse TokenResponse
 *
 * @apiUse ParamValidationErrors
 * @apiUse OtherErrors
 */

func init() {
	engine.Router().POST(
		"/resources",
		validateCreateBody,
		saveResource,
	)
}

func saveResource(c *gin.Context) {
	body := c.MustGet("data").(resources.CreateResourceRequest)

	resource, err := resources.Create(&body)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"id": resource.ID.Hex(),
	})
}

func validateCreateBody(c *gin.Context) {
	body := resources.CreateResourceRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.Set("data", body)
	c.Next()
}
