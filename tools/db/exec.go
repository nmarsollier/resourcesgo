package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/logx"
)

func Exec(logenv logx.Fields, query string, args ...any) (err error) {
	conn, err := getDBConn(logenv)

	if err != nil {
		logx.Error(logenv, err)
		return
	}

	_, err = conn.Exec(context.Background(), query, args...)
	if err != nil {
		checkConnectionError(err)
		logx.Error(logenv, err)
	}

	return
}
