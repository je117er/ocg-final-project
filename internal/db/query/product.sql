-- name: GetProducts :many
SELECT BIN_TO_UUID(id) id, name, price, vendor, vaccine_type, authorized_ages, dose, antigen_nature, route_of_administration, storage_requirements, available_formats, diluent, adjuvant, alternate_name, minimum_interval, immunization_schedule, authorized_interval, extended_interval, background, regulatory_actions, safety_status, authorization_status, trials, distribution, funding, slug, image, lot_number, expiry_date
FROM product;

-- name: GetProduct :one
SELECT *
FROM product
WHERE slug = ? LIMIT 1;
