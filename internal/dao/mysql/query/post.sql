-- name: CreatePost :exec
INSERT INTO post (id, cover, title, abstract, content, public)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdatePostByPostID :exec
UPDATE post
SET cover    = ?,
    title    = ?,
    abstract = ?,
    content  = ?,
    public   = ?
WHERE id = ?;

-- name: GetPostByPostID :one
SELECT *
FROM post
WHERE id = ?
LIMIT 1;

-- name: DeletePostByPostID :exec
DELETE
FROM post
WHERE id = ?
  and deleted = true;

-- name: ModifyPostDeletedByID :exec
UPDATE post
SET deleted = ?
WHERE id = ?;

-- name: ModifyPostPublicByID :exec
UPDATE post
SET public = ?
WHERE id = ?;

-- name: ListPostPublic :many
SELECT id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       create_time,
       modify_time
FROM post
where public = true
  and deleted = false
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostPrivate :many
SELECT id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       create_time,
       modify_time
FROM post
where public = false
  and deleted = false
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostDeleted :many
SELECT id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       create_time,
       modify_time
FROM post
where deleted = true
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostTopping :many
SELECT p.id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       p.create_time,
       modify_time,
       t.id
FROM post p,
     tops t
where p.id = t.post_id
ORDER BY t.id
LIMIT ?,?;

-- name: ListPostBySearchKey :many
SELECT p.id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       p.create_time,
       modify_time
FROM post p
where (title like ?
    or abstract like ?)
  and deleted = false
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostByStartTime :many
SELECT p.id,
       cover,
       title,
       abstract,
       star_num,
       visited_num,
       p.create_time,
       modify_time
FROM post p
where deleted = false
  and p.create_time between ? and ?
ORDER BY create_time Desc
LIMIT ?,?;

/*
增加帖子:ok
修改贴子:
    修改内容ok
    设置删除ok
    设置公开ok
查询帖子:
    ID 单个全部信息ok
    公开的ok
    私密的ok
    删除的ok
    时间范围ok
    访问数排序ok
    点赞数排序ok
    关键字查询
删除帖子:ok
*/
