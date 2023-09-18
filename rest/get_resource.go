package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/engine"
)

func init() {
	engine.Router().GET(
		"/resources/:project/:language/:semver",
		getResource,
	)
}

func getResource(c *gin.Context) {
	project := c.Param("project")
	language := c.Param("language")
	semver := c.Param("semver")

	resource, err := resources.GetLastResource(project, language, semver)

	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"id":       resource.ID.Hex(),
		"project":  resource.Project,
		"language": resource.Language,
		"semver":   resource.SemVer,
		"values":   resource.Values,
	})
}
