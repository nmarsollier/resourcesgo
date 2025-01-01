package main

import (
	"github.com/nmarsollier/resourcesgo/graph"
	routes "github.com/nmarsollier/resourcesgo/rest"
)

func main() {
	go graph.Start()
	routes.Start()
}
