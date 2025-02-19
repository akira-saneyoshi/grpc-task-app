-- name: FindUserByID :one
SELECT id, name, email, password, is_active, created_at, updated_at
FROM users
WHERE id = ?
LIMIT 1;

-- name: FindUserByEmail :one
SELECT id, name, email, password, is_active, created_at, updated_at
FROM users
WHERE email = ?
LIMIT 1;
