package rest

import (
	"fmt"

	"github.com/nmarsollier/resourcesgo/rest/engine"
	"github.com/nmarsollier/resourcesgo/tools/env"
)

// Start this server
func Start() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}
