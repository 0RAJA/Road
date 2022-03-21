-- name: CreateManager :exec
INSERT INTO manager (username, password)
VALUES (?, ?);

-- name: UpdateManager :exec
UPDATE manager
SET password =?
WHERE username = ?;

-- name: GetManagerByUsername :one
SELECT *
FROM manager
WHERE username = ?
LIMIT 1;

-- name: ListManager :many
SELECT *
FROM manager
ORDER BY username
LIMIT ?,?;
