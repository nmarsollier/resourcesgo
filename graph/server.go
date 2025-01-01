package graph

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nmarsollier/resourcesgo/graph/model"
	"github.com/nmarsollier/resourcesgo/graph/schema"
	"github.com/nmarsollier/resourcesgo/tools/env"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Start() {
	logfld := logx.NewFields()

	port := env.Get().GqlPort
	srv := handler.NewDefaultServer(model.NewExecutableSchema(model.Config{Resolvers: &schema.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logx.Info(logfld, "GraphQL playground in port : "+strconv.Itoa(port))
	logx.Error(logfld, http.ListenAndServe(fmt.Sprintf(":%d", env.Get().GqlPort), nil))
}
