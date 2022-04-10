-- name: CreateUser_Star :exec
INSERT INTO user_star (username, post_id)
VALUES (?, ?);

-- name: GetUser_StarByUserNameAndPostId :one
SELECT id
FROM user_star
WHERE username = ?
  and post_id = ?;

-- name: DeleteUser_StarByUserNameAndPostID :exec
DELETE
FROM user_star
WHERE username = ?
  and post_id = ?;
/*
点赞关系:
    增加: username post_id
    删除
        通过username postID删除
    查询
        通过username和post_id查
*/
