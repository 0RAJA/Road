-- name: CreatePost_Num :exec
INSERT INTO post_num (post_id)
VALUES (?);

-- name: UpdatePost_Num_Star :exec
UPDATE post_num
SET star_num = (star_num + ?)
WHERE post_id = ?;

-- name: UpdatePost_Num_Visited :exec
UPDATE post_num
SET visited_num = (visited_num + ?)
WHERE post_id = ?;
