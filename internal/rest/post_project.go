package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/internal/projects"
	"github.com/nmarsollier/resourcesgo/internal/rest/server"
)

//	@Summary		Create a project
//	@Description	Create a new project tag.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateProjectRequest	true	"Project to add"
//	@Success		200		{object}	IDResult				"Project ID"
//	@Failure		400		{object}	errs.Validation			"Bad Request"
//	@Failure		404		{object}	errs.Custom				"Not Found"
//	@Failure		500		{object}	errs.Custom				"Internal Server Error"
//	@Router			/projects [post]
//
// Create a new project tab
func initPostProjects(engine *gin.Engine) {
	engine.POST("/projects", saveProject)
}

func saveProject(c *gin.Context) {
	body := CreateProjectRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	id, err := projects.Create(server.GinLogCtx(c), body.ID, body.Name)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &IDResult{id})
}

type CreateProjectRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
