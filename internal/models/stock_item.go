package models

import (
	"database/sql"
)

type StockItem struct {
	ID         string
	ClinicID   sql.NullString
	Quantity   sql.NullInt32
	Name       sql.NullString
	Price      sql.NullString
	LotNumber  sql.NullString
	ExpiryDate sql.NullTime
	ProductID  sql.NullString
}
