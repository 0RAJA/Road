// Code generated by sqlc. DO NOT EDIT.
// source: tag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :exec
INSERT INTO tags (id, tag_name)
VALUES (?, ?)
`

type CreateTagParams struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) error {
	_, err := q.db.ExecContext(ctx, createTag, arg.ID, arg.TagName)
	return err
}

const deleteTagByTagID = `-- name: DeleteTagByTagID :exec
DELETE
FROM tags
WHERE id = ?
`

func (q *Queries) DeleteTagByTagID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTagByTagID, id)
	return err
}

const getTagById = `-- name: GetTagById :one
SELECT id, tag_name, create_time
FROM tags
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetTagById(ctx context.Context, id int64) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagById, id)
	var i Tag
	err := row.Scan(&i.ID, &i.TagName, &i.CreateTime)
	return i, err
}

const getTagByName = `-- name: GetTagByName :one
select id, tag_name, create_time
from tags
where tag_name = ?
`

func (q *Queries) GetTagByName(ctx context.Context, tagName string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagByName, tagName)
	var i Tag
	err := row.Scan(&i.ID, &i.TagName, &i.CreateTime)
	return i, err
}

const listTag = `-- name: ListTag :many
SELECT id, tag_name, create_time
FROM tags
ORDER BY create_time Desc
LIMIT ?,?
`

type ListTagParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListTag(ctx context.Context, arg ListTagParams) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTag, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.TagName, &i.CreateTime); err != nil {
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

const updateTag = `-- name: UpdateTag :exec
UPDATE tags
SET tag_name = ?
WHERE id = ?
`

type UpdateTagParams struct {
	TagName string `json:"tag_name"`
	ID      int64  `json:"id"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) error {
	_, err := q.db.ExecContext(ctx, updateTag, arg.TagName, arg.ID)
	return err
}
