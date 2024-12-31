package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Query[T any](logenv logx.Fields, query string, args ...interface{}) (results []*T, err error) {
	conn, err := getDBConn(logenv)
	if err != nil {
		logx.Error(logenv, err)
		return
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		logx.Error(logenv, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		result := new(T)

		if err = rows.Scan(fieldAddrs(result)...); err != nil {
			checkConnectionError(err)
			logx.Error(logenv, err)
			return
		}

		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		logx.Error(logenv, err)
		return
	}

	return
}
