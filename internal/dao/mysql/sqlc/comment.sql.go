// Code generated by sqlc. DO NOT EDIT.
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :exec
INSERT INTO comment (id, post_id, username, content, to_comment_id)
VALUES (?, ?, ?, ?, ?)
`

type CreateCommentParams struct {
	ID          int64  `json:"id"`
	PostID      int64  `json:"post_id"`
	Username    string `json:"username"`
	Content     string `json:"content"`
	ToCommentID int32  `json:"to_comment_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) error {
	_, err := q.db.ExecContext(ctx, createComment,
		arg.ID,
		arg.PostID,
		arg.Username,
		arg.Content,
		arg.ToCommentID,
	)
	return err
}

const deleteCommentByCommentID = `-- name: DeleteCommentByCommentID :exec
DELETE
FROM comment
WHERE id = ?
`

func (q *Queries) DeleteCommentByCommentID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCommentByCommentID, id)
	return err
}

const getCommentByCommentID = `-- name: GetCommentByCommentID :one
SELECT id, post_id, username, content, to_comment_id, create_time, modify_time
FROM comment
WHERE id = ?
`

func (q *Queries) GetCommentByCommentID(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getCommentByCommentID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.Username,
		&i.Content,
		&i.ToCommentID,
		&i.CreateTime,
		&i.ModifyTime,
	)
	return i, err
}

const listCommentByPostID = `-- name: ListCommentByPostID :many
SELECT id, post_id, username, content, to_comment_id, create_time, modify_time
FROM comment
where post_id = ?
ORDER BY create_time Desc
LIMIT ?,?
`

type ListCommentByPostIDParams struct {
	PostID int64 `json:"post_id"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListCommentByPostID(ctx context.Context, arg ListCommentByPostIDParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listCommentByPostID, arg.PostID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.Username,
			&i.Content,
			&i.ToCommentID,
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

const updateCommentByCommentID = `-- name: UpdateCommentByCommentID :exec
UPDATE comment
set content = ?
WHERE id = ?
`

type UpdateCommentByCommentIDParams struct {
	Content string `json:"content"`
	ID      int64  `json:"id"`
}

func (q *Queries) UpdateCommentByCommentID(ctx context.Context, arg UpdateCommentByCommentIDParams) error {
	_, err := q.db.ExecContext(ctx, updateCommentByCommentID, arg.Content, arg.ID)
	return err
}