-- name: CreateBucket :exec
insert into storage.buckets
(id, name, allowed_content_types, max_allowed_object_size, public, disabled)
values (sqlc.arg('id'),
        sqlc.arg('name'),
        sqlc.narg('allowed_content_types'),
        sqlc.narg('max_allowed_object_size'),
        sqlc.arg('public'),
        sqlc.arg('disabled'));

-- name: UpdateBucket :exec
update storage.buckets
set max_allowed_object_size = coalesce(sqlc.narg('max_allowed_object_size'), max_allowed_object_size),
    public                  = coalesce(sqlc.narg('public'), public)
where id = sqlc.arg('id');

-- name: UpdateBucketAllowedContentTypes :exec
update storage.buckets
set allowed_content_types = sqlc.arg('allowed_content_types')
where id = sqlc.arg('id');

-- name: DisableBucket :exec
update storage.buckets
set disabled = true
where id = sqlc.arg('id');

-- name: EnableBucket :exec
update storage.buckets
set disabled = false
where id = sqlc.arg('id');

-- name: MakeBucketPublic :exec
update storage.buckets
set public = true
where id = sqlc.arg('id');

-- name: MakeBucketPrivate :exec
update storage.buckets
set public = false
where id = sqlc.arg('id');

-- name: LockBucket :exec
update storage.buckets
set locked      = true,
    lock_reason = sqlc.arg('lock_reason')::text,
    locked_at   = now()
where id = sqlc.arg('id');

-- name: UnlockBucket :exec
update storage.buckets
set locked      = false,
    lock_reason = null,
    locked_at   = null
where id = sqlc.arg('id');

-- name: DeleteBucket :exec
delete
from storage.buckets
where id = sqlc.arg('id');

-- name: GetBucketById :one
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
where id = sqlc.arg('id')
limit 1;

-- name: GetBucketByName :one
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
where name = sqlc.arg('name')
limit 1;

-- name: ListAllBuckets :many
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
from storage.buckets;

-- name: ListBucketsPaginated :many
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
where id >= sqlc.arg('cursor')
limit sqlc.arg('limit');

-- name: SearchBucketsPaginated :many
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
where name ilike sqlc.narg('name')
limit sqlc.narg('limit') offset sqlc.narg('offset');

-- name: CountBuckets :one
select count(1) as count
from storage.buckets;

-- name: GetBucketSizeById :one
select id, name, sum(size) as size
from storage.objects
where bucket_id = sqlc.arg('id');

-- name: GetBucketObjectCountById :one
select count(1) as count
from storage.objects
where bucket_id = sqlc.arg('id');