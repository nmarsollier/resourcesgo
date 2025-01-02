package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

// Exec executes a mutation SQL.
//
// Parameters:
//   - ctx: The context to use for the database operation.
//   - query: The SQL query to execute.
//   - args: The arguments for the SQL query.
//
// Returns:
//   - err: An error if the operation fails, otherwise nil.
func Exec(ctx context.Context, query string, args ...any) (err error) {
	conn, err := getDBConn(ctx)

	if err != nil {
		logx.Error(ctx, err)
		return
	}

	_, err = conn.Exec(context.Background(), query, args...)
	if err != nil {
		checkConnectionError(err)
		logx.Error(ctx, err)
	}

	return
}
