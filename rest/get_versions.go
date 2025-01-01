package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Gets versions
//	@Description	Gets available versions for a resource.
//	@Tags			Resources
//	@Accept			json
//	@Produce		json
//	@Param			projectId	path		string			true	"Project ID"
//	@Param			languageId	path		string			true	"language tag"
//	@Success		200			{Array}		string			"Versions"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/versions/{projectId}/{languageId} [get]
//
// Gets the last resource.
func initGetVersions(engine *gin.Engine) {
	engine.GET("/versions/:projectId/:languageId", getVersions)
}

func getVersions(c *gin.Context) {
	project := c.Param("projectId")
	language := c.Param("languageId")

	resource, err := resources.FindVersions(server.GinLogFields(c), project, language)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, resource)
}
