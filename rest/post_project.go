package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/projects"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Create a project
//	@Description	Create a new project tag.
//	@Tags			Resources
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
	engine.POST(
		"/projects",
		saveProject,
	)
}

func saveProject(c *gin.Context) {
	body := CreateProjectRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	resource, err := projects.Create(server.GinLogFields(c), body.Id, body.Name)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &IDResult{resource.ID})
}

type CreateProjectRequest struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type IDResult struct {
	ID string `json:"id"`
}
