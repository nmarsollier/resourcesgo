package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/resources"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Marks a resource as deleted.
//	@Description	Delete a resource.
//	@Tags			Resources
//	@Accept			json
//	@Produce		json
//	@Param			projectId	path	string	true	"Project ID"
//	@Param			languageId	path	string	true	"language tag"
//	@Param			semver		path	string	true	"Sem version, you can not use wildcards"
//	@Success		200			"No Content"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/resources/{projectId}/{languageId}/{semver} [delete]
//
// Marks a resource as deleted.
func initDeleteResource(engine *gin.Engine) {
	engine.DELETE(
		"/resources/:projectId/:languageId/:semver",
		deleteResource,
	)
}

func deleteResource(c *gin.Context) {
	project := c.Param("projectId")
	language := c.Param("languageId")
	semver := c.Param("semver")

	resources.Delete(server.GinLogFields(c), project, language, semver)
}
