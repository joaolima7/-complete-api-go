-- name: CreateProduct :execresult
INSERT INTO products (id, name, price, mark_id) 
VALUES(?, ?, ?, ?);

-- name: GetProductById :one
SELECT id, name, price, mark_id, created_at, updated_at 
FROM products 
WHERE id = ?;

-- name: GetAllProducts :many
SELECT id, name, price, mark_id, created_at, updated_at 
FROM products;

-- name: UpdateProduct :execresult
UPDATE products 
SET name = ?, price = ?, mark_id = ? 
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products 
WHERE id = ?;

-- name: GetProductsByMarkId :many
SELECT id, name, price, mark_id, created_at, updated_at 
FROM products 
WHERE mark_id = ?;