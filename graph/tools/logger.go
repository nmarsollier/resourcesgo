package tools

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func LoggerFields(ctx context.Context) logx.Fields {
	operationContext := graphql.GetOperationContext(ctx)

	return logx.NewFields().
		Add(logx.CONTROLLER, "GraphQL").
		Add(logx.HTTP_METHOD, operationContext.OperationName).
		Add(logx.HTTP_PATH, operationContext.OperationName).
		Add(logx.CORRELATION_ID, getCorrelationId(ctx))
}

func getCorrelationId(ctx context.Context) string {
	operationContext := graphql.GetOperationContext(ctx)
	value := operationContext.Headers.Get("Authorization")

	if len(value) == 0 {
		value = uuid.New().String()
	}

	return value
}
