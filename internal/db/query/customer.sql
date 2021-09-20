-- name: GetAllCustomers :many
select *
from customer;

-- name: GetCustomer :one
SELECT *
FROM customer
where id = ? limit 1;


-- name: UpdateCustomer :exec
UPDATE customer
SET
    name = ?,
    address = ?,
    gender = ?,
    dob = ?,
    insurance_number = ?,
    district = ?,
    city = ?,
    commune = ?,
    ethnicity = ?,
    nationality = ?
WHERE id = ?;


