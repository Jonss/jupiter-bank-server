package db

import (
	"context"
	"time"
)

type Config struct {
	ID        int64
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetConfigByKeyParams struct {
	Key string
}

const getConfigByKeyQuery = `
SELECT id, key, value
FROM configs
WHERE key = $1
`

func (q *Queries) GetConfigByKey(ctx context.Context, arg GetConfigByKeyParams) (Config, error) {
	row := q.db.QueryRowContext(ctx, getConfigByKeyQuery, arg.Key)
	var c Config
	err := row.Scan(
		&c.ID,
		&c.Key,
		&c.Value,
	)
	return c, err
}
