-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
 RETURNING *)
 SELECT inserted_feed_follow.*,
        feeds.name AS feed_name,
        users.name AS user_name
  FROM inserted_feed_follow
  INNER JOIN users ON user_id = users.id 
  INNER JOIN feeds ON feed_id = feeds.id;


-- name: GetFeedFollowsForUser :many
SELECT users.name AS user_name,
       feeds.name AS feed_name
       FROM feed_follows
       INNER JOIN users ON user_id = users.id
       INNER JOIN feeds ON feed_id = feeds.id
       WHERE feed_follows.user_id = $1;

-- name: DeleteFollowFeed :exec
DELETE FROM feed_follows
       WHERE user_id = $1 AND feed_id = $2;



      
