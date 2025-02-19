-- name: FindUserByID :one
SELECT id, name, email, password, created_at, updated_at
FROM users
WHERE id = ?
LIMIT 1;

-- name: FindUserByEmail :one
SELECT id, name, email, password, created_at, updated_at
FROM users
WHERE email = ?
LIMIT 1;
