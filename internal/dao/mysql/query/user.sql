-- name: CreateUser :exec
INSERT INTO user (username, avatar_url, depository_url, address)
VALUES (?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE user
SET avatar_url     = ?,
    depository_url = ?,
    address        = ?
WHERE username = ?;

-- name: GetUserByUsername :one
SELECT *
FROM user
WHERE username = ?
LIMIT 1;

-- name: ListUser :many
SELECT *
FROM user
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListUserByCreateTime :many
SELECT *
FROM user
where create_time between ? and ?
ORDER BY create_time Desc
LIMIT ?,?;
