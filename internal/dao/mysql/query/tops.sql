-- name: CreateTop :exec
INSERT INTO tops (id, post_id)
VALUES (?, ?);

-- name: GetTopByPostID :one
SELECT *
FROM tops
where post_id = ?;

-- name: DeleteTopByPostID :exec
DELETE
FROM tops
WHERE post_id = ?;
/*
置顶:
    增加:
    删除:
*/
