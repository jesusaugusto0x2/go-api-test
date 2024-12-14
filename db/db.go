package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"example.com/go-api-test/ent"
	"go.uber.org/fx"
)

type Config struct {
	DSN string
}

func NewEntClient(lc fx.Lifecycle, cfg Config) (*ent.Client, error) {
	client, err := ent.Open(dialect.Postgres, cfg.DSN)

	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
