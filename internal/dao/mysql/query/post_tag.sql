-- name: CreatePost_Tag :exec
INSERT INTO post_tag (post_id, tag_id)
VALUES (?, ?);

-- name: GetPost_Tag :one
SELECT *
FROM post_tag
where post_id = ?
  and tag_id = ?;

-- name: DeletePost_Tag :exec
DELETE
FROM post_tag
WHERE post_id = ?
  and tag_id = ?;

-- name: ListPostByTagID :many
SELECT p.id,
       p.cover,
       p.title,
       p.abstract,
       p.public,
       p.deleted,
       pn.star_num,
       pn.visited_num,
       p.create_time,
       p.modify_time
FROM post_tag pt
         join post p
         join tags t
         join post_num pn on (pt.tag_id = ? and pt.post_id = p.id and pt.tag_id = t.id and pt.post_id = pn.post_id)
ORDER BY create_time Desc
LIMIT ?,?;

-- name: ListTagByPostID :many
SELECT t.*
FROM post_tag pt
         join tags t
         join post p on pt.post_id = ? and pt.tag_id = t.id and pt.post_id = p.id
ORDER BY create_time Desc
LIMIT ?,?;

/*
标签-帖子：
    增加 id,post_id,tag_id
    删除 id
    查询
        通过post_id查全部tag
        通过tag_id查全部post
*/
