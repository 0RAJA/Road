// Code generated by sqlc. DO NOT EDIT.
// source: manager.sql

package db

import (
	"context"
)

const createManager = `-- name: CreateManager :exec
INSERT INTO manager (username, password)
VALUES (?, ?)
`

type CreateManagerParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateManager(ctx context.Context, arg CreateManagerParams) error {
	_, err := q.db.ExecContext(ctx, createManager, arg.Username, arg.Password)
	return err
}

const getManagerByUsername = `-- name: GetManagerByUsername :one
SELECT username, password
FROM manager
WHERE username = ?
LIMIT 1
`

func (q *Queries) GetManagerByUsername(ctx context.Context, username string) (Manager, error) {
	row := q.db.QueryRowContext(ctx, getManagerByUsername, username)
	var i Manager
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}

const listManager = `-- name: ListManager :many
SELECT username, password
FROM manager
ORDER BY username
LIMIT ?,?
`

type ListManagerParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListManager(ctx context.Context, arg ListManagerParams) ([]Manager, error) {
	rows, err := q.db.QueryContext(ctx, listManager, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Manager{}
	for rows.Next() {
		var i Manager
		if err := rows.Scan(&i.Username, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateManager = `-- name: UpdateManager :exec
UPDATE manager
SET password =?
WHERE username = ?
`

type UpdateManagerParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func (q *Queries) UpdateManager(ctx context.Context, arg UpdateManagerParams) error {
	_, err := q.db.ExecContext(ctx, updateManager, arg.Password, arg.Username)
	return err
}
