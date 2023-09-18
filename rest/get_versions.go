package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/engine"
)

func init() {
	engine.Router().GET(
		"/versions/:project/:language",
		getVersions,
	)
}

func getVersions(c *gin.Context) {
	project := c.Param("project")
	language := c.Param("language")

	resource, err := resources.GetVersions(project, language)

	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, resource)
}
