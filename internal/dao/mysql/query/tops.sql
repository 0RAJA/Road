-- name: CreateTop :exec
INSERT INTO tops (id, post_id)
VALUES (?, ?);

-- name: GetTopByTopID :one
SELECT *
FROM tops
where id = ?;

-- name: DeleteTopByTopID :exec
DELETE
FROM tops
WHERE id = ?;
/*
置顶:
    增加:
    删除:
*/
