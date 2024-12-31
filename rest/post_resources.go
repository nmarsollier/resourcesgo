package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Create a resource
//	@Description	Create a new resource.
//	@Tags			Resources
//	@Accept			json
//	@Produce		json
//	@Param			body	body		resources.CreateResourceRequest	true	"Project to add"
//	@Success		200		{object}	IDResult						"Project ID"
//	@Failure		400		{object}	errs.Validation					"Bad Request"
//	@Failure		404		{object}	errs.Custom						"Not Found"
//	@Failure		500		{object}	errs.Custom						"Internal Server Error"
//	@Router			/resources [post]
//
// Create a new resource version
func initPostResources(engine *gin.Engine) {
	engine.POST(
		"/resources",
		validateCreateBody,
		saveResource,
	)
}

func saveResource(c *gin.Context) {
	body := c.MustGet("data").(resources.CreateResourceRequest)

	resource, err := resources.Create(server.GinLogFields(c), &body)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &IDResult{resource.ID})
}

func validateCreateBody(c *gin.Context) {
	body := resources.CreateResourceRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.Set("data", body)
	c.Next()
}
