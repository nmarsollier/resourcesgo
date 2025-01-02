package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/internal/resources"
	"github.com/nmarsollier/resourcesgo/internal/rest/server"
)

//	@Summary		Create a resource
//	@Description	Create a new resource.
//	@Tags			Resources
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateResourceBody	true	"Project to add"
//	@Success		200		{object}	IDResult			"Project ID"
//	@Failure		400		{object}	errs.Validation		"Bad Request"
//	@Failure		404		{object}	errs.Custom			"Not Found"
//	@Failure		500		{object}	errs.Custom			"Internal Server Error"
//	@Router			/resources [post]
//
// Create a new resource version
func initPostResources(engine *gin.Engine) {
	engine.POST("/resources", saveResource)
}

func saveResource(c *gin.Context) {
	body := CreateResourceBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	newRes := resources.NewResource(
		body.ProjectID,
		body.LanguageID,
		body.SemVer,
		body.Values,
	)

	id, err := resources.Create(server.GinLogCtx(c), newRes)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &IDResult{id})
}

type CreateResourceBody struct {
	ProjectID  string            `json:"projectId" binding:"required"`
	LanguageID string            `json:"languageId" binding:"required"`
	SemVer     string            `json:"semver" binding:"required"`
	Values     map[string]string `json:"values" binding:"required"`
}
