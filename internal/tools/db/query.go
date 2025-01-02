package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

// Query executes a SQL query and returns the results.
//
// Type Parameters:
//
//	T: The type of the result objects.
//
// Parameters:
//
//	ctx: The context for the query execution.
//	query: The SQL query string to be executed.
//	args: The arguments for the SQL query.
//
// Returns:
//
//	results: A slice of pointers to the result objects of type T.
//	err: An error if the query execution or scanning fails.
func Query[T any](ctx context.Context, query string, args ...interface{}) (results []*T, err error) {
	conn, err := getDBConn(ctx)
	if err != nil {
		logx.Error(ctx, err)
		return
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		logx.Error(ctx, err)
		return
	}
	defer rows.Close()
	columns := columnNames(rows)

	for rows.Next() {
		result := new(T)

		if err = rows.Scan(structFieldPointers(columns, result)...); err != nil {
			checkConnectionError(err)
			logx.Error(ctx, err)
			return
		}

		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		logx.Error(ctx, err)
		return
	}

	return
}
