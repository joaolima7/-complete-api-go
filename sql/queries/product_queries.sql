-- name: CreateProduct :one
INSERT INTO products (name, price, mark_id) VALUES(?, ?, ?) RETURNING id, name, price, mark_id, created_at, updated_at;

-- name: GetAllProducts :many
SELECT id, name, price, mark_id, created_at, updated_at FROM products;

-- name: GetProductById :one
SELECT id, name, price, mark_id, created_at, updated_at FROM products WHERE id = ?;

-- name: UpdateProduct :one
UPDATE products SET name = ?, price = ?, mark_id = ? WHERE id = ? RETURNING id, name, price, mark_id, created_at, updated_at;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?;

-- name: GetProductsByMarkId :many
SELECT id, name, price, mark_id, created_at, updated_at FROM products WHERE mark_id = ?;