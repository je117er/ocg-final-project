package models

import "database/sql"

type PriceRule struct {
	ID              sql.NullInt32
	ProductID       string
	Value           sql.NullString
	ValueType       sql.NullInt32
	OncePerCustomer sql.NullBool
	UsageLimit      sql.NullInt32
	StartsAt        sql.NullTime
	EndsAt          sql.NullTime
	Title           sql.NullString
}
