package graph

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nmarsollier/resourcesgo/internal/graph/model"
	"github.com/nmarsollier/resourcesgo/internal/graph/schema"
	"github.com/nmarsollier/resourcesgo/internal/tools/env"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func Start() {
	ctx := logx.CtxWithFields(context.Background(), logx.NewFields())

	port := env.Get().GqlPort
	srv := handler.NewDefaultServer(model.NewExecutableSchema(model.Config{Resolvers: &schema.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logx.Info(ctx, "GraphQL playground in port : "+strconv.Itoa(port))
	logx.Error(ctx, http.ListenAndServe(fmt.Sprintf(":%d", env.Get().GqlPort), nil))
}
