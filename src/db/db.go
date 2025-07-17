package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"interview-teamex-v1/src/config"
	"time"
)

func Init(ctx context.Context) (*pgx.Conn, error) {
	//
	contextTTL, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()

	//
	conn, err := pgx.Connect(contextTTL, config.Env.DbPostgresUrl)
	if err != nil {
		return nil, err
	}

	// migrate some tables, no versioning here... just lazy check
	if err = migrate(ctx, conn); err != nil {
		// wrapping error is fun, not sure if I'll unwrap it later on...
		return nil, fmt.Errorf("db migration: %w", err)
	}

	// seeding here, not advanced seed checking... just populate
	if err = seed(ctx, conn); err != nil {
		return nil, fmt.Errorf("db seed: %w", err)
	}

	return conn, nil
}
