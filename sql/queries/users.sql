-- sql/queries/users.sql

-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE($2, username),
    email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
