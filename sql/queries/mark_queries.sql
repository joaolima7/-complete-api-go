-- name: CreateMark :execresult
INSERT INTO marks (id, name) 
VALUES(?, ?);

-- name: GetMarkById :one
SELECT id, name, created_at, updated_at 
FROM marks 
WHERE id = ?;

-- name: GetAllMarks :many
SELECT id, name, created_at, updated_at 
FROM marks;

-- name: UpdateMark :execresult
UPDATE marks 
SET name = ?
WHERE id = ?;   

-- name: DeleteMark :exec
DELETE FROM marks 
WHERE id = ?;