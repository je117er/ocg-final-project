package models

import (
	"database/sql"
)

type StockItem struct {
	ID                   string
	ClinicID             sql.NullString
	Quantity             sql.NullInt32
	Name                 sql.NullString
	Price                sql.NullString
	LotNumber            sql.NullString
	ExpiryDate           sql.NullTime
	ProductID            sql.NullString
	StockLeft            sql.NullInt32
	ImmunizationSchedule sql.NullInt64
	AuthorizedInterval   sql.NullInt64
}
