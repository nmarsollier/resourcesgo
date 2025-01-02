package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/internal/languages"
	"github.com/nmarsollier/resourcesgo/internal/rest/server"
)

//	@Summary		Create a language
//	@Description	Create a new language tag.
//	@Tags			Languages
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateLanguageRequest	true	"Language to add"
//	@Success		200		{object}	IDResult				"Language ID"
//	@Failure		400		{object}	errs.Validation			"Bad Request"
//	@Failure		404		{object}	errs.Custom				"Not Found"
//	@Failure		500		{object}	errs.Custom				"Internal Server Error"
//	@Router			/languages [post]
//
// Create a new project tab
func initPostLanguage(engine *gin.Engine) {
	engine.POST("/languages", saveLanguage)
}

func saveLanguage(c *gin.Context) {
	body := CreateLanguageRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	id, err := languages.Create(server.GinLogFields(c), body.ID, body.Name)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &IDResult{id})
}

type CreateLanguageRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
