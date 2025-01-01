package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/resourcesgo/languages"
	"github.com/nmarsollier/resourcesgo/rest/server"
)

//	@Summary		Gets a language detail
//	@Description	Gets language details.
//	@Tags			Languages
//	@Accept			json
//	@Produce		json
//	@Param			languageId	path		string			true	"Language ID"
//	@Success		200			{object}	Language			"Project"
//	@Failure		400			{object}	errs.Validation	"Bad Request"
//	@Failure		404			{object}	errs.Custom		"Not Found"
//	@Failure		500			{object}	errs.Custom		"Internal Server Error"
//	@Router			/languages/:languageId [get]
//
// Gets a project details.
func initGetLanguage(engine *gin.Engine) {
	engine.GET("/languages/:languageId", getLanguage)
}

func getLanguage(c *gin.Context) {
	project := c.Param("languageId")

	proj, err := languages.FindByID(server.GinLogFields(c), project)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, &Language{
		ID:   proj.ID,
		Name: proj.Name,
	})
}

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
