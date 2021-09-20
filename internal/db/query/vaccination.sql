-- name: GetVaccination :one
SELECT doses_completed, vaccine_name, date_registered, clinic_name, date_booked
FROM booking
WHERE customer_id = ? limit 1;
