// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO user (username, avatar_url, depository_url, address)
VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	Username      string         `json:"username"`
	AvatarUrl     string         `json:"avatar_url"`
	DepositoryUrl string         `json:"depository_url"`
	Address       sql.NullString `json:"address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.AvatarUrl,
		arg.DepositoryUrl,
		arg.Address,
	)
	return err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT username, avatar_url, depository_url, address, create_time, modify_time
FROM user
WHERE username = ?
LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.AvatarUrl,
		&i.DepositoryUrl,
		&i.Address,
		&i.CreateTime,
		&i.ModifyTime,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT username, avatar_url, depository_url, address, create_time, modify_time
FROM user
ORDER BY create_time Desc
LIMIT ?,?
`

type ListUserParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.AvatarUrl,
			&i.DepositoryUrl,
			&i.Address,
			&i.CreateTime,
			&i.ModifyTime,
		); err != nil {
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

const listUserByCreateTime = `-- name: ListUserByCreateTime :many
SELECT username, avatar_url, depository_url, address, create_time, modify_time
FROM user
where create_time between ? and ?
ORDER BY create_time Desc
LIMIT ?,?
`

type ListUserByCreateTimeParams struct {
	CreateTime   time.Time `json:"create_time"`
	CreateTime_2 time.Time `json:"create_time_2"`
	Offset       int32     `json:"offset"`
	Limit        int32     `json:"limit"`
}

func (q *Queries) ListUserByCreateTime(ctx context.Context, arg ListUserByCreateTimeParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUserByCreateTime,
		arg.CreateTime,
		arg.CreateTime_2,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.AvatarUrl,
			&i.DepositoryUrl,
			&i.Address,
			&i.CreateTime,
			&i.ModifyTime,
		); err != nil {
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

const updateUser = `-- name: UpdateUser :exec
UPDATE user
SET avatar_url     = ?,
    depository_url = ?,
    address        = ?
WHERE username = ?
`

type UpdateUserParams struct {
	AvatarUrl     string         `json:"avatar_url"`
	DepositoryUrl string         `json:"depository_url"`
	Address       sql.NullString `json:"address"`
	Username      string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.AvatarUrl,
		arg.DepositoryUrl,
		arg.Address,
		arg.Username,
	)
	return err
}
