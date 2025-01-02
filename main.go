package main

import (
	"github.com/nmarsollier/resourcesgo/internal/graph"
	routes "github.com/nmarsollier/resourcesgo/internal/rest"
)

func main() {
	go graph.Start()
	routes.Start()
}
