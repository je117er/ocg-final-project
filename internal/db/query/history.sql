-- name: GetMedicalHistory :one
SELECT *
FROM medical_condition
where customer_id = ? limit 1;

-- name: UpdateMedicalHistory :exec
UPDATE medical_condition
set
    description = ?,
    condition_status = ?
where customer_id = ?;
