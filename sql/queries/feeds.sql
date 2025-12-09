-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  GEN_RANDOM_UUID(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
    GEN_RANDOM_UUID(),
    NOW(),
    NOW(),
    $1,
    $2
    )
  RETURNING *
)
SELECT 
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN inserted_feed_follow.user_id ON users.id = inserted_feed_follow.user_id
INNER JOIN inserted_feed_follow.feed_id ON feeds.id = inserted_feed_follow.feed_id ;
