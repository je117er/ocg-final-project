-- name: GetBooking :one
SELECT id, date_registered, total_bill, vaccine_name, lot_number
FROM booking
WHERE customer_id = ? limit 1;

-- name: UpdateBooking :exec
UPDATE booking
SET
    date_booked = ?,
    time_period = ?,
    discount_id = ?
WHERE customer_id = ?;
