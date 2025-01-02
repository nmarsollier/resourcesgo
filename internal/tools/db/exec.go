package db

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

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
