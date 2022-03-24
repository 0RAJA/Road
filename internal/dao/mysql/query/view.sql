-- name: CreateView :exec
INSERT INTO views (views_num)
VALUES (?);

-- name: ListViewByCreateTime :many
SELECT *
FROM views
where create_time between ? and ?
ORDER BY create_time Desc
LIMIT ?,?;
