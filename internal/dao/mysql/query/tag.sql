-- name: CreateTag :exec
INSERT INTO tags (id, tag_name)
VALUES (?, ?);

-- name: UpdateTag :exec
UPDATE tags
SET tag_name = ?
WHERE id = ?;

-- name: GetTagById :one
SELECT *
FROM tags
WHERE id = ?
LIMIT 1;

-- name: ListTag :many
SELECT *
FROM tags
ORDER BY create_time Desc
LIMIT ?,?;

-- name: DeleteTagByTagID :exec
DELETE
FROM tags
WHERE id = ?;
/*
标签：
    增加
    删除
    修改
    查询 通过id查询
*/
