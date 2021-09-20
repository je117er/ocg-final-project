-- name: UpdateStockItem :exec
UPDATE stock_item
SET
    quantity = ?,
    price = ?,
    lot_number = ?,
    expiry_date = ?
WHERE id = ?;