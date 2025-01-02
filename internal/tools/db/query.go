package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

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

		if err = rows.Scan(fieldAddrs(columns, result)...); err != nil {
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
