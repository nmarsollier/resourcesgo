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
//	@Param			project		path		string			true	"Project ID"
//	@Param			language	path		string			true	"language tag"
//	@Success		200			{Array}		string			"Versions"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/versions/{project}/{language} [get]
//
// Gets the last resource.
func initGetVersions(engine *gin.Engine) {
	engine.GET(
		"/versions/:project/:language",
		getVersions,
	)
}

func getVersions(c *gin.Context) {
	project := c.Param("project")
	language := c.Param("language")

	resource, err := resources.FindVersions(server.GinLogFields(c), project, language)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, resource)
}
