package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/errs"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func QueryRow[T any](fields logx.Fields, query string, args ...interface{}) (*T, error) {
	conn, err := getDBConn(fields)
	if err != nil {
		logx.Error(fields, err)
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		checkConnectionError(err)
		logx.Error(fields, err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errs.NotFound
	}

	result := new(T)
	if err := rows.Scan(fieldAddrs(columnNames(rows), result)...); err != nil {
		checkConnectionError(err)
		logx.Error(fields, err)
		return nil, err
	}

	return result, nil
}
