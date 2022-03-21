-- name: CreatePost_Tag :exec
INSERT INTO post_tag (id, post_id, tag_id)
VALUES (?, ?, ?);

-- name: DeletePost_TagByID :exec
DELETE
FROM post_tag
WHERE id = ?;

-- name: ListPostByTagID :many
SELECT p.id,
       p.cover,
       p.title,
       p.abstract,
       p.star_num,
       p.visited_num,
       p.create_time,
       p.modify_time
FROM post_tag pt
         join post p
         join tags t on (pt.tag_id = ? and pt.post_id = p.id and pt.tag_id = t.id)
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
