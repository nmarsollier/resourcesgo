package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/resourcesgo/tools/env"
	"github.com/nmarsollier/resourcesgo/tools/logx"
	"github.com/nmarsollier/resourcesgo/tools/strs"
)

var instance *pgxpool.Pool

const ERR_EXIST = 23505
const ERR_FOREIGN_KEY = 23503

func getDBConn(fields logx.Fields) (*pgxpool.Pool, error) {
	if instance != nil {
		return instance, nil
	}

	config, err := pgxpool.ParseConfig(env.Get().PostgresURL)
	if err != nil {
		logx.Error(fields, err)
		return nil, err
	}

	instance, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logx.Error(fields, err)
		return nil, err
	}

	logx.Info(fields, "Postgres Connected")

	return instance, nil
}

func checkConnectionError(err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "08000", "08003", "08006", "08001", "08004", "08007", "08P01":
			instance = nil
		}
	}

	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		instance = nil
	}
}

func ErrorCode(err error) int {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return strs.AtoiZero(pgErr.Code)
	}
	return 0
}
