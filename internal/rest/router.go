package rest

import (
	"fmt"

	"github.com/nmarsollier/resourcesgo/internal/rest/server"
	"github.com/nmarsollier/resourcesgo/internal/tools/env"
)

// Start this server
func Start() {
	engine := server.Router()
	initGetProject(engine)
	initGetLanguage(engine)
	initGetResources(engine)
	initGetVersions(engine)
	initPostResources(engine)
	initDeleteResource(engine)
	initPostProjects(engine)
	initPostLanguage(engine)
	engine.Run(fmt.Sprintf(":%d", env.Get().Port))
}
