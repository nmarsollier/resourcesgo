package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

func Query[T any](fields logx.Fields, query string, args ...interface{}) (results []*T, err error) {
	conn, err := getDBConn(fields)
	if err != nil {
		logx.Error(fields, err)
		return
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		logx.Error(fields, err)
		return
	}
	defer rows.Close()
	columns := columnNames(rows)

	for rows.Next() {
		result := new(T)

		if err = rows.Scan(fieldAddrs(columns, result)...); err != nil {
			checkConnectionError(err)
			logx.Error(fields, err)
			return
		}

		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		logx.Error(fields, err)
		return
	}

	return
}
