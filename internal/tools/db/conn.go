package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/resourcesgo/internal/tools/env"
	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

var instance *pgxpool.Pool

const ERR_EXIST = 23505
const ERR_FOREIGN_KEY = 23503

// getDBConn establishes a connection to the PostgreSQL database using pgxpool.
// It returns a connection pool instance or an error if the connection fails.
// If a connection pool instance already exists, it returns the existing instance.
//
// Parameters:
//
//	ctx - The context for managing the connection lifecycle.
//
// Returns:
//
//	*pgxpool.Pool - The connection pool instance.
//	error - An error if the connection fails.
func getDBConn(ctx context.Context) (*pgxpool.Pool, error) {
	if instance != nil {
		return instance, nil
	}

	config, err := pgxpool.ParseConfig(env.Get().PostgresURL)
	if err != nil {
		logx.Error(ctx, err)
		return nil, err
	}

	instance, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logx.Error(ctx, err)
		return nil, err
	}

	logx.Info(ctx, "Postgres Connected")

	return instance, nil
}
