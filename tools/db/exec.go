package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Exec(fields logx.Fields, query string, args ...any) (err error) {
	conn, err := getDBConn(fields)

	if err != nil {
		logx.Error(fields, err)
		return
	}

	_, err = conn.Exec(context.Background(), query, args...)
	if err != nil {
		checkConnectionError(err)
		logx.Error(fields, err)
	}

	return
}
