-- name: GetClinic :one
SELECT *
FROM clinic
WHERE id = ? LIMIT 1;
