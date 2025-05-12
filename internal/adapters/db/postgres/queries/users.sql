-- name: GetUserByID :one
select * from users WHERE id = $1;

-- name: GetUserByEmail :one
select * from users WHERE email = $1;

-- name: GetUserByUsername :one
select * from users WHERE username = $1;

-- name: CreateUser :exec
INSERT INTO users (id, email, username, password, status, created_at, created_by_id, updated_at)
VALUES (
           gen_random_uuid(), $1, $2, $3, $4,
           now(), $5, now()
       ) ON CONFLICT (email) DO NOTHING;

-- name: CreateUserWithID :exec
INSERT INTO users (id, email, username, password, status, created_at, created_by_id, updated_at)
VALUES (
           $1, $2, $3, $4, $5,
           now(), $6, now()
       ) ON CONFLICT (email) DO NOTHING;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $2,
    updated_at = now(),
    updated_by_id = $3
WHERE id = $1;

-- name: UpdateUserUsername :exec
UPDATE users
SET username = $2,
    updated_at = now(),
    updated_by_id = $3
WHERE id = $1;

-- name: UpdateUserStatus :exec
UPDATE users
SET status = $2,
    updated_at = now(),
    updated_by_id = $3
WHERE id = $1;
