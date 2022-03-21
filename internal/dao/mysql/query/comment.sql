-- name: CreateComment :exec
INSERT INTO comment (id, post_id, username, content, to_comment_id)
VALUES (?, ?, ?, ?, ?);

-- name: GetCommentByCommentID :one
SELECT *
FROM comment
WHERE id = ?;

-- name: UpdateCommentByCommentID :exec
UPDATE comment
set content = ?
WHERE id = ?;

-- name: ListCommentByPostID :many
SELECT *
FROM comment
where post_id = ?
ORDER BY create_time Desc
LIMIT ?,?;

-- name: DeleteCommentByCommentID :exec
DELETE
FROM comment
WHERE id = ?;
/*
评论:
    增加
    删除
        通过ID删除
    修改
        修改内容
    查询
        通过post_id查
        通过id查 //测试
*/
