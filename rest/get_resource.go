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
//	@Param			projectId	path		string			true	"Project ID"
//	@Param			languageId	path		string			true	"language tag"
//	@Param			semver		path		string			true	"Sem version, you can use wildcards + or *"
//	@Success		200			{object}	Resource		"Resource"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/resources/{projectId}/{languageId}/{semver} [get]
//
// Gets the last resource.
func initGetResources(engine *gin.Engine) {
	engine.GET("/resources/:projectId/:languageId/:semver", getResource)
}

func getResource(c *gin.Context) {
	project := c.Param("projectId")
	language := c.Param("languageId")
	semver := c.Param("semver")

	resource, err := resources.GetLastResource(server.GinLogFields(c), project, language, semver)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &Resource{
		ID:         resource.ID,
		ProjectID:  resource.ProjectID,
		LanguageID: resource.LanguageID,
		SemVer:     resource.SemVer,
		Values:     resource.Values,
	})
}

type Resource struct {
	ID         string            `json:"id"`
	ProjectID  string            `json:"projectId"`
	LanguageID string            `json:"languageId"`
	SemVer     string            `json:"semver"`
	Values     map[string]string `json:"values"`
}
