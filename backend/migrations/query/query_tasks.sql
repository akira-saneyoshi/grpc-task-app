-- query_tasks.sql

-- name: FindTaskByID :one
SELECT id, user_id, title, description, status, due_date, created_at, updated_at
FROM tasks
WHERE id = ?
LIMIT 1;

-- name: FindTasksByUserID :many
SELECT id, user_id, title, description, status, due_date, created_at, updated_at
FROM tasks
WHERE user_id = ?
ORDER BY updated_at DESC;

-- name: CreateTask :exec
INSERT INTO tasks(id, user_id, title, description, status, due_date, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: GetLastInsertID :one
SELECT LAST_INSERT_ID();

-- name: UpdateTask :exec
UPDATE tasks
SET title = ?, description = ?, status = ?, due_date = ?, updated_at = NOW()
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
