-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
select 
  f.name as feed_name, 
  u.name as user_name, 
  new_feed_follow.* 
from new_feed_follow 
inner join feeds f on new_feed_follow.feed_id = f.id 
inner join users u on new_feed_follow.user_id = u.id;

-- name: GetFeedFollowsForUser :many
select 
  f.name as feed_name, 
  u.name as user_name, 
  ff.*
from feed_follows ff
inner join feeds f on ff.feed_id = f.id
inner join users u on ff.user_id = u.id
where ff.user_id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2;
