// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: bucket_query.sql

package database

import (
	"context"
	"time"
)

const countBuckets = `-- name: CountBuckets :one
select count(1) as count
from storage.buckets
`

func (q *Queries) CountBuckets(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countBuckets)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createBucket = `-- name: CreateBucket :one
insert into storage.buckets
(name, allowed_content_types, max_allowed_object_size, public, disabled)
values ($1,
        $2,
        $3,
        $4,
        $5) returning id
`

type CreateBucketParams struct {
	Name                 string
	AllowedContentTypes  []string
	MaxAllowedObjectSize *int64
	Public               bool
	Disabled             bool
}

func (q *Queries) CreateBucket(ctx context.Context, arg *CreateBucketParams) (string, error) {
	row := q.db.QueryRow(ctx, createBucket,
		arg.Name,
		arg.AllowedContentTypes,
		arg.MaxAllowedObjectSize,
		arg.Public,
		arg.Disabled,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteBucket = `-- name: DeleteBucket :exec
delete
from storage.buckets
where id = $1
`

func (q *Queries) DeleteBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteBucket, id)
	return err
}

const disableBucket = `-- name: DisableBucket :exec
update storage.buckets
set disabled = true
where id = $1
`

func (q *Queries) DisableBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, disableBucket, id)
	return err
}

const enableBucket = `-- name: EnableBucket :exec
update storage.buckets
set disabled = false
where id = $1
`

func (q *Queries) EnableBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, enableBucket, id)
	return err
}

const getBucketById = `-- name: GetBucketById :one
select id,
       version,
       name,
       allowed_content_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where id = $1
limit 1
`

func (q *Queries) GetBucketById(ctx context.Context, id string) (*StorageBucket, error) {
	row := q.db.QueryRow(ctx, getBucketById, id)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.AllowedContentTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getBucketByName = `-- name: GetBucketByName :one
select id,
       version,
       name,
       allowed_content_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name = $1
limit 1
`

func (q *Queries) GetBucketByName(ctx context.Context, name string) (*StorageBucket, error) {
	row := q.db.QueryRow(ctx, getBucketByName, name)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.AllowedContentTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getBucketObjectCountById = `-- name: GetBucketObjectCountById :one
select o.bucket_id as id, count(1) as count
from storage.objects as o
where o.bucket_id = $1
group by o.bucket_id
`

type GetBucketObjectCountByIdRow struct {
	ID    string
	Count int64
}

func (q *Queries) GetBucketObjectCountById(ctx context.Context, id string) (*GetBucketObjectCountByIdRow, error) {
	row := q.db.QueryRow(ctx, getBucketObjectCountById, id)
	var i GetBucketObjectCountByIdRow
	err := row.Scan(&i.ID, &i.Count)
	return &i, err
}

const getBucketSizeById = `-- name: GetBucketSizeById :one
select o.bucket_id as id, b.name as name, SUM(o.size) as size
from storage.objects as o
         join storage.buckets as b on o.bucket_id = b.id
where o.bucket_id = $1
group by o.bucket_id, b.name
`

type GetBucketSizeByIdRow struct {
	ID   string
	Name string
	Size int64
}

func (q *Queries) GetBucketSizeById(ctx context.Context, id string) (*GetBucketSizeByIdRow, error) {
	row := q.db.QueryRow(ctx, getBucketSizeById, id)
	var i GetBucketSizeByIdRow
	err := row.Scan(&i.ID, &i.Name, &i.Size)
	return &i, err
}

const listAllBuckets = `-- name: ListAllBuckets :many
select id,
       version,
       name,
       allowed_content_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
`

func (q *Queries) ListAllBuckets(ctx context.Context) ([]*StorageBucket, error) {
	rows, err := q.db.Query(ctx, listAllBuckets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.AllowedContentTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBucketsPaginated = `-- name: ListBucketsPaginated :many
select id,
       name,
       allowed_content_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where id >= $1
limit $2
`

type ListBucketsPaginatedParams struct {
	Cursor string
	Limit  int32
}

type ListBucketsPaginatedRow struct {
	ID                   string
	Name                 string
	AllowedContentTypes  []string
	MaxAllowedObjectSize *int64
	Public               bool
	Disabled             bool
	Locked               bool
	LockReason           *string
	LockedAt             *time.Time
	CreatedAt            time.Time
	UpdatedAt            *time.Time
}

func (q *Queries) ListBucketsPaginated(ctx context.Context, arg *ListBucketsPaginatedParams) ([]*ListBucketsPaginatedRow, error) {
	rows, err := q.db.Query(ctx, listBucketsPaginated, arg.Cursor, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListBucketsPaginatedRow
	for rows.Next() {
		var i ListBucketsPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AllowedContentTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const lockBucket = `-- name: LockBucket :exec
update storage.buckets
set locked      = true,
    lock_reason = $1::text,
    locked_at   = now()
where id = $2
`

type LockBucketParams struct {
	LockReason string
	ID         string
}

func (q *Queries) LockBucket(ctx context.Context, arg *LockBucketParams) error {
	_, err := q.db.Exec(ctx, lockBucket, arg.LockReason, arg.ID)
	return err
}

const makeBucketPrivate = `-- name: MakeBucketPrivate :exec
update storage.buckets
set public = false
where id = $1
`

func (q *Queries) MakeBucketPrivate(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeBucketPrivate, id)
	return err
}

const makeBucketPublic = `-- name: MakeBucketPublic :exec
update storage.buckets
set public = true
where id = $1
`

func (q *Queries) MakeBucketPublic(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeBucketPublic, id)
	return err
}

const searchBucketsPaginated = `-- name: SearchBucketsPaginated :many
select id,
       name,
       allowed_content_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name ilike $1
limit $3 offset $2
`

type SearchBucketsPaginatedParams struct {
	Name   *string
	Offset *int32
	Limit  *int32
}

type SearchBucketsPaginatedRow struct {
	ID                   string
	Name                 string
	AllowedContentTypes  []string
	MaxAllowedObjectSize *int64
	Public               bool
	Disabled             bool
	Locked               bool
	LockReason           *string
	LockedAt             *time.Time
	CreatedAt            time.Time
	UpdatedAt            *time.Time
}

func (q *Queries) SearchBucketsPaginated(ctx context.Context, arg *SearchBucketsPaginatedParams) ([]*SearchBucketsPaginatedRow, error) {
	rows, err := q.db.Query(ctx, searchBucketsPaginated, arg.Name, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*SearchBucketsPaginatedRow
	for rows.Next() {
		var i SearchBucketsPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AllowedContentTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unlockBucket = `-- name: UnlockBucket :exec
update storage.buckets
set locked      = false,
    lock_reason = null,
    locked_at   = null
where id = $1
`

func (q *Queries) UnlockBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, unlockBucket, id)
	return err
}

const updateBucket = `-- name: UpdateBucket :exec
update storage.buckets
set max_allowed_object_size = coalesce($1, max_allowed_object_size),
    public                  = coalesce($2, public),
    allowed_content_types   = coalesce($3, allowed_content_types)
where id = $4
`

type UpdateBucketParams struct {
	MaxAllowedObjectSize *int64
	Public               *bool
	AllowedContentTypes  []string
	ID                   string
}

func (q *Queries) UpdateBucket(ctx context.Context, arg *UpdateBucketParams) error {
	_, err := q.db.Exec(ctx, updateBucket,
		arg.MaxAllowedObjectSize,
		arg.Public,
		arg.AllowedContentTypes,
		arg.ID,
	)
	return err
}
