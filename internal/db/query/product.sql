-- name: GetProduct :one
SELECT *
FROM product
WHERE id = ? LIMIT 1;
