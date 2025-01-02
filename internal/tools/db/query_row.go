package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

// QueryRow executes a query for a single row result.
//
// Type Parameters:
//
//	T - The type of the struct to scan the result into.
//
// Parameters:
//
//	ctx - The context to use for the database operation.
//	query - The SQL query to execute.
//	args - The arguments for the SQL query.
//
// Returns:
//
//	*T - A pointer to the struct containing the result of the query.
//	error - An error if the query fails or no rows are found.
//
// If no rows are found, it returns errs.NotFound.
// If an error occurs during the query or scanning, it logs the error and returns it.
func QueryRow[T any](ctx context.Context, query string, args ...interface{}) (*T, error) {
	conn, err := getDBConn(ctx)
	if err != nil {
		logx.Error(ctx, err)
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		checkConnectionError(err)
		logx.Error(ctx, err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errs.NotFound
	}

	result := new(T)
	if err := rows.Scan(structFieldPointers(columnNames(rows), result)...); err != nil {
		checkConnectionError(err)
		logx.Error(ctx, err)
		return nil, err
	}

	return result, nil
}
