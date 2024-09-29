-- name: AddPost :one
INSERT INTO posts (id, title, url, description, published_at, feed_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: PostsForUser :many
SELECT * FROM posts JOIN feeds ON posts.feed_id = feeds.id WHERE feeds.user_id = $1 ORDER BY published_at DESC limit $2;
