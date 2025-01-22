// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPosts = `-- name: CreatePosts :one
INSERT INTO posts (id, created_at, updated_at, published_at, title, url, description, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
RETURNING id, created_at, updated_at, published_at, title, url, description, feed_id
`

type CreatePostsParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	Title       string
	Url         string
	Description sql.NullString
	FeedID      uuid.UUID
}

func (q *Queries) CreatePosts(ctx context.Context, arg CreatePostsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPosts,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.PublishedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PublishedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.FeedID,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.id, posts.created_at, posts.updated_at, published_at, title, posts.url, description, feed_id, feeds.id, feeds.created_at, feeds.updated_at, last_fetched_at, name, feeds.url, user_id FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
WHERE feeds.user_id = $1
ORDER BY posts.created_at DESC LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

type GetPostsForUserRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   time.Time
	Title         string
	Url           string
	Description   sql.NullString
	FeedID        uuid.UUID
	ID_2          uuid.UUID
	CreatedAt_2   time.Time
	UpdatedAt_2   time.Time
	LastFetchedAt sql.NullTime
	Name          string
	Url_2         string
	UserID        uuid.UUID
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]GetPostsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsForUserRow
	for rows.Next() {
		var i GetPostsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PublishedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.FeedID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LastFetchedAt,
			&i.Name,
			&i.Url_2,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
