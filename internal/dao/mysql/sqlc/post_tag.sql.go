// Code generated by sqlc. DO NOT EDIT.
// source: post_tag.sql

package db

import (
	"context"
	"time"
)

const createPost_Tag = `-- name: CreatePost_Tag :exec
INSERT INTO post_tag (id, post_id, tag_id)
VALUES (?, ?, ?)
`

type CreatePost_TagParams struct {
	ID     int64 `json:"id"`
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}

func (q *Queries) CreatePost_Tag(ctx context.Context, arg CreatePost_TagParams) error {
	_, err := q.db.ExecContext(ctx, createPost_Tag, arg.ID, arg.PostID, arg.TagID)
	return err
}

const deletePost_TagByID = `-- name: DeletePost_TagByID :exec
DELETE
FROM post_tag
WHERE id = ?
`

func (q *Queries) DeletePost_TagByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost_TagByID, id)
	return err
}

const listPostByTagID = `-- name: ListPostByTagID :many
SELECT p.id,
       p.cover,
       p.title,
       p.abstract,
       p.star_num,
       p.visited_num,
       p.create_time,
       p.modify_time
FROM post_tag pt
         join post p
         join tags t on (pt.tag_id = ? and pt.post_id = p.id and pt.tag_id = t.id)
ORDER BY create_time Desc
LIMIT ?,?
`

type ListPostByTagIDParams struct {
	TagID  int64 `json:"tag_id"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type ListPostByTagIDRow struct {
	ID         int64     `json:"id"`
	Cover      string    `json:"cover"`
	Title      string    `json:"title"`
	Abstract   string    `json:"abstract"`
	StarNum    int32     `json:"star_num"`
	VisitedNum int32     `json:"visited_num"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}

func (q *Queries) ListPostByTagID(ctx context.Context, arg ListPostByTagIDParams) ([]ListPostByTagIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listPostByTagID, arg.TagID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPostByTagIDRow{}
	for rows.Next() {
		var i ListPostByTagIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Cover,
			&i.Title,
			&i.Abstract,
			&i.StarNum,
			&i.VisitedNum,
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

const listTagByPostID = `-- name: ListTagByPostID :many
SELECT t.id, t.tag_name, t.create_time
FROM post_tag pt
         join tags t
         join post p on pt.post_id = ? and pt.tag_id = t.id and pt.post_id = p.id
ORDER BY create_time Desc
LIMIT ?,?
`

type ListTagByPostIDParams struct {
	PostID int64 `json:"post_id"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListTagByPostID(ctx context.Context, arg ListTagByPostIDParams) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTagByPostID, arg.PostID, arg.Offset, arg.Limit)
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
