package rest

import (
	"fmt"

	"github.com/nmarsollier/resourcesgo/rest/server"
	"github.com/nmarsollier/resourcesgo/tools/env"
)

// Start this server
func Start() {
	engine := server.Router()
	initGetProject(engine)
	initGetResources(engine)
	initGetVersions(engine)
	initPostResources(engine)
	initDeleteResource(engine)
	initPostProjects(engine)
	engine.Run(fmt.Sprintf(":%d", env.Get().Port))
}
