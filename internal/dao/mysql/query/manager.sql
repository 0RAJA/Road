-- name: CreateManager :exec
INSERT INTO manager (username, password, avatar_url)
VALUES (?, ?, ?);

-- name: UpdateManagerPassword :exec
UPDATE manager
SET password =?
WHERE username = ?;

-- name: UpdateManagerAvatar :exec
UPDATE manager
SET avatar_url =?
WHERE username = ?;

-- name: GetManagerByUsername :one
SELECT *
FROM manager
WHERE username = ?
LIMIT 1;

-- name: ListManager :many
SELECT username, avatar_url
FROM manager
ORDER BY username
LIMIT ?,?;

-- name: DeleteManager :exec
DELETE
FROM manager
WHERE username = ?;
