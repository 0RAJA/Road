// Code generated by sqlc. DO NOT EDIT.
// source: user_star.sql

package db

import (
	"context"
)

const createUser_Star = `-- name: CreateUser_Star :exec
INSERT INTO user_star (username, post_id)
VALUES (?, ?)
`

type CreateUser_StarParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) CreateUser_Star(ctx context.Context, arg CreateUser_StarParams) error {
	_, err := q.db.ExecContext(ctx, createUser_Star, arg.Username, arg.PostID)
	return err
}

const deleteUser_StarByUserNameAndPostID = `-- name: DeleteUser_StarByUserNameAndPostID :exec
DELETE
FROM user_star
WHERE username = ?
  and post_id = ?
`

type DeleteUser_StarByUserNameAndPostIDParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) DeleteUser_StarByUserNameAndPostID(ctx context.Context, arg DeleteUser_StarByUserNameAndPostIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser_StarByUserNameAndPostID, arg.Username, arg.PostID)
	return err
}

const getUser_StarByUserNameAndPostId = `-- name: GetUser_StarByUserNameAndPostId :one
SELECT id
FROM user_star
WHERE username = ?
  and post_id = ?
`

type GetUser_StarByUserNameAndPostIdParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) GetUser_StarByUserNameAndPostId(ctx context.Context, arg GetUser_StarByUserNameAndPostIdParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getUser_StarByUserNameAndPostId, arg.Username, arg.PostID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
