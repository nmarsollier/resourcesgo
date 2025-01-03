package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/internal/projects"
	"github.com/nmarsollier/resourcesgo/internal/rest/server"
)

//	@Summary		Gets a project detail
//	@Description	Gets projects details.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			projectId	path		string			true	"Project ID"
//	@Success		200			{object}	Project			"Project"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/projects/{projectId} [get]
//
// Gets a project details.
func initGetProject(engine *gin.Engine) {
	engine.GET("/projects/:projectId", getProject)
}

func getProject(c *gin.Context) {
	proj, err := projects.FindByID(server.GinLogCtx(c), c.Param("projectId"))

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &Project{
		ID:   proj.ID,
		Name: proj.Name,
	})
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
