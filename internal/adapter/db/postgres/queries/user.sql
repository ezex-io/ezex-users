-- name: GetSecurityImageByEmail :one
SELECT security_image, security_phrase
FROM users
WHERE email = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id, firebase_uuid, email)
VALUES ($1, $2, $3);

-- name: SaveSecurityImage :exec
UPDATE users
SET security_image = $2,
    security_phrase = $3
WHERE email = $1;
