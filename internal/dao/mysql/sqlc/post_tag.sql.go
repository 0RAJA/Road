// Code generated by sqlc. DO NOT EDIT.
// source: post_tag.sql

package db

import (
	"context"
	"time"
)

const createPost_Tag = `-- name: CreatePost_Tag :exec
INSERT INTO post_tag (post_id, tag_id)
VALUES (?, ?)
`

type CreatePost_TagParams struct {
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}

func (q *Queries) CreatePost_Tag(ctx context.Context, arg CreatePost_TagParams) error {
	_, err := q.db.ExecContext(ctx, createPost_Tag, arg.PostID, arg.TagID)
	return err
}

const deletePost_Tag = `-- name: DeletePost_Tag :exec
DELETE
FROM post_tag
WHERE post_id = ?
  and tag_id = ?
`

type DeletePost_TagParams struct {
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}

func (q *Queries) DeletePost_Tag(ctx context.Context, arg DeletePost_TagParams) error {
	_, err := q.db.ExecContext(ctx, deletePost_Tag, arg.PostID, arg.TagID)
	return err
}

const getPost_Tag = `-- name: GetPost_Tag :one
SELECT id, post_id, tag_id
FROM post_tag
where post_id = ?
  and tag_id = ?
`

type GetPost_TagParams struct {
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}

func (q *Queries) GetPost_Tag(ctx context.Context, arg GetPost_TagParams) (PostTag, error) {
	row := q.db.QueryRowContext(ctx, getPost_Tag, arg.PostID, arg.TagID)
	var i PostTag
	err := row.Scan(&i.ID, &i.PostID, &i.TagID)
	return i, err
}

const listPostByTagID = `-- name: ListPostByTagID :many
SELECT p.id,
       p.cover,
       p.title,
       p.abstract,
       p.public,
       p.deleted,
       pn.star_num,
       pn.visited_num,
       p.create_time,
       p.modify_time
FROM post_tag pt
         join post p
         join tags t
         join post_num pn on (pt.tag_id = ? and pt.post_id = p.id and pt.tag_id = t.id and pt.post_id = pn.post_id)
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
	Public     bool      `json:"public"`
	Deleted    bool      `json:"deleted"`
	StarNum    int64     `json:"star_num"`
	VisitedNum int64     `json:"visited_num"`
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
			&i.Public,
			&i.Deleted,
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
