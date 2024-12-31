package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Gets a resource
//	@Description	Gets the latest resource given a semver.
//	@Tags			Resources
//	@Accept			json
//	@Produce		json
//	@Param			project			path		string				true	"Project ID"
//	@Param			language		path		string				true	"language tag"
//	@Param			semver			path		string				true	"Sem version, you can use wildcards + or *"
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	resources.Resource	"Resource"
//	@Failure		400				{object}	errs.Validation		"Bad Request"
//	@Failure		404				{object}	errs.Custom			"Not Found"
//	@Failure		500				{object}	errs.Custom			"Internal Server Error"
//	@Router			/resources/:project/:language/:semver [get]
//
// Gets the last resource.
func initGetResources(engine *gin.Engine) {
	engine.GET(
		"/resources/:project/:language/:semver",
		getResource,
	)
}

func getResource(c *gin.Context) {
	project := c.Param("project")
	language := c.Param("language")
	semver := c.Param("semver")

	resource, err := resources.GetLastResource(server.GinLogFields(c), project, language, semver)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"id":       resource.ID,
		"project":  resource.Project,
		"language": resource.Language,
		"semver":   resource.SemVer,
		"values":   resource.Values,
	})
}
