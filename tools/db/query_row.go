package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func QueryRow[T any](logenv logx.Fields, query string, args ...interface{}) (*T, error) {
	conn, err := getDBConn(logenv)
	if err != nil {
		logx.Error(logenv, err)
		return nil, err
	}

	row := conn.QueryRow(context.Background(), query, args...)

	result := new(T)
	if err := row.Scan(fieldAddrs(result)...); err != nil {
		checkConnectionError(err)
		logx.Error(logenv, err)
		return nil, err
	}

	return result, nil
}
