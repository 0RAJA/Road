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

-- name: GetPostInfoByPostID :one
SELECT id,
       cover,
       title,
       abstract,
       public,
       deleted,
       create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post,
     post_num pn
WHERE id = ?
  and id = post_id
LIMIT 1;

-- name: GetPostByPostID :one
SELECT p.*, pn.star_num, pn.visited_num
FROM post p,
     post_num pn
WHERE id = ?
  and id = pn.post_id
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
       public,
       deleted,
       create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post,
     post_num pn
where public = true
  and deleted = false
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostPrivate :many
SELECT id,
       cover,
       title,
       abstract,
       public,
       deleted,
       create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post,
     post_num pn
where public = false
  and deleted = false
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostDeleted :many
SELECT id,
       cover,
       title,
       abstract,
       public,
       deleted,
       create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post,
     post_num pn
where deleted = true
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostTopping :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     tops t,
     post_num pn
where p.id = t.post_id
  and p.id = pn.post_id
ORDER BY t.id Desc
LIMIT ?,?;

-- name: ListPostBySearchKey :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     post_num pn
where (title like ?
    or abstract like ?)
  and deleted = false
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostByStartTime :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     post_num pn
where deleted = false
  and p.create_time between ? and ?
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostOrderByCreatedTime :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     post_num pn
where deleted = false
  and id = post_id
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListPostOrderByStarNum :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     post_num pn
where deleted = false
  and id = post_id
ORDER BY pn.star_num Desc
LIMIT ?,?;

-- name: ListPostOrderByVisitedNum :many
SELECT p.id,
       cover,
       title,
       abstract,
       public,
       deleted,
       p.create_time,
       modify_time,
       pn.star_num,
       pn.visited_num
FROM post p,
     post_num pn
where deleted = false
  and id = post_id
ORDER BY pn.visited_num Desc
LIMIT ?,?;
/*
????????????:ok
    ??????star???
    ??????visited???
????????????:
    ????????????ok
    ????????????ok
    ????????????ok
????????????:
    ID ??????????????????ok
    ?????????ok
    ?????????ok
    ?????????ok
    ????????????ok
    ???????????????ok
    ???????????????ok
    ???????????????
????????????:ok
*/
