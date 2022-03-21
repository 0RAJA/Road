// Code generated by sqlc. DO NOT EDIT.
// source: tops.sql

package db

import (
	"context"
)

const createTop = `-- name: CreateTop :exec
INSERT INTO tops (id, post_id)
VALUES (?, ?)
`

type CreateTopParams struct {
	ID     int64 `json:"id"`
	PostID int64 `json:"post_id"`
}

func (q *Queries) CreateTop(ctx context.Context, arg CreateTopParams) error {
	_, err := q.db.ExecContext(ctx, createTop, arg.ID, arg.PostID)
	return err
}

const deleteTopByTopID = `-- name: DeleteTopByTopID :exec
DELETE
FROM tops
WHERE id = ?
`

func (q *Queries) DeleteTopByTopID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTopByTopID, id)
	return err
}

const getTopByTopID = `-- name: GetTopByTopID :one
SELECT id, post_id
FROM tops
where id = ?
`

func (q *Queries) GetTopByTopID(ctx context.Context, id int64) (Top, error) {
	row := q.db.QueryRowContext(ctx, getTopByTopID, id)
	var i Top
	err := row.Scan(&i.ID, &i.PostID)
	return i, err
}