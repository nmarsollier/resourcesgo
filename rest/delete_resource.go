package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/engine"
)

func init() {
	engine.Router().DELETE(
		"/resources/:project/:language/:semver",
		deleteResource,
	)
}

func deleteResource(c *gin.Context) {
	project := c.Param("project")
	language := c.Param("language")
	semver := c.Param("semver")

	resources.Delete(project, language, semver)
}
