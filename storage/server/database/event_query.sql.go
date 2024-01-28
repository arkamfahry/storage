// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: event_query.sql

package database

import (
	"context"
	"time"
)

const createEvent = `-- name: CreateEvent :exec
insert into storage.events
    (id, name, content, status, retries, expires_at, created_at)
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7)
`

type CreateEventParams struct {
	ID        string
	Name      string
	Content   []byte
	Status    string
	Retries   int32
	ExpiresAt *time.Time
	CreatedAt time.Time
}

func (q *Queries) CreateEvent(ctx context.Context, arg *CreateEventParams) error {
	_, err := q.db.Exec(ctx, createEvent,
		arg.ID,
		arg.Name,
		arg.Content,
		arg.Status,
		arg.Retries,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	return err
}
