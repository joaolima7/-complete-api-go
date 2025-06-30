-- name: CreateUser :execresult
INSERT INTO users (id, name, email, password) 
VALUES(?, ?, ?, ?);

-- name: GetUserById :one
SELECT id, name, email, password, created_at, updated_at 
FROM users 
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at 
FROM users 
WHERE email = ?;

-- name: GetAllUsers :many
SELECT id, name, email, password, created_at, updated_at 
FROM users;

-- name: UpdateUser :execresult
UPDATE users 
SET name = ?, email = ?, password = ?
WHERE id = ?;

-- name: UpdateUserPassword :execresult
UPDATE users 
SET password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = ?;

-- name: CheckEmailExists :one
SELECT COUNT(*) as count
FROM users 
WHERE email = ? AND id != ?;

-- name: GetUserByEmailForAuth :one
SELECT id, name, email, password 
FROM users 
WHERE email = ?;